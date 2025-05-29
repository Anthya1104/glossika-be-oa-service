package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/redis"
	"github.com/gin-gonic/gin"
)

func GetRecommendationListHandler(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	var req model.GetRecommendationListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := fmt.Errorf("failed to bind request body: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadRequestBody,
			RawErr:     err,
		})
		return
	}

	offset := (req.Page - 1) * req.PageSize

	productList := []db.Product{}
	// get cache
	cacheKey := fmt.Sprintf("user_recommendation:%v", userID)
	cacheVal, err := redis.Client.Get(c, cacheKey).Result()
	if err == nil && cacheVal != "" {
		_ = json.Unmarshal([]byte(cacheVal), &productList)
	} else {

		recommendationList, wrapErr := database.GetSqlDb().GetUserRecommendation(c, userID)
		if wrapErr.RawErr != nil {
			respondError(c, wrapErr)
			return
		}
		log.C(c).Debugf("GetRecommendationListHandler get all recommendationList from DB :%v\n", recommendationList)

		productIDList := []uint{}
		for _, v := range recommendationList {
			productIDList = append(productIDList, v.ProductID)
		}

		productList, wrapErr = database.GetSqlDb().BatchGetProducts(c, productIDList)
		if wrapErr.RawErr != nil {
			respondError(c, wrapErr)
			return
		}
		log.C(c).Debugf("GetRecommendationListHandler get all productList from DB :%v\n", productList)

		// set cache
		b, _ := json.Marshal(productList)
		redis.Client.Set(c, cacheKey, b, time.Minute*10)
	}

	// TODO: the sliced data would be 1.5x by the pageSize, need to check the limit log
	pagedrecommendationList, total := util.SliceDataByPaging(productList, offset, req.PageSize)
	log.C(c).Debugf("GetRecommendationListHandler sliced data: %v, total: %v\n", pagedrecommendationList, total)

	resp := model.GetRecommendationListResp{}

	for _, p := range pagedrecommendationList {
		resp.Data.RecommendationList = append(resp.Data.RecommendationList, model.RecommendationList{
			ProductID:   p.ID,
			ProductName: p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	// TODO: fixed the empty resp.Data.Total issue

	respondSuccess(c, http.StatusOK, &resp, false)
}
