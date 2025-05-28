package handler

import (
	"fmt"
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
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

	recommendationList, wrapErr := database.GetSqlDb().GetUserRecommendation(c, userID)
	if wrapErr.RawErr != nil {
		respondError(c, wrapErr)
		return
	}
	log.C(c).Infof("GetRecommendationListHandler get all recommendationList from DB :%v\n", recommendationList)

	productIDList := []uint{}
	for _, v := range recommendationList {
		productIDList = append(productIDList, v.ProductID)
	}

	productList, wrapErr := database.GetSqlDb().BatchGetProducts(c, productIDList)
	if wrapErr.RawErr != nil {
		respondError(c, wrapErr)
		return
	}
	log.C(c).Infof("GetRecommendationListHandler get all productList from DB :%v\n", productList)

	pagedrecommendationList, total := util.SliceDataByPaging(productList, offset, req.PageSize)
	log.C(c).Infof("GetRecommendationListHandler sliced data: %v, total: %v\n", pagedrecommendationList, total)

	resp := model.GetRecommendationListResp{}

	for _, p := range pagedrecommendationList {
		resp.Data.RecommendationList = append(resp.Data.RecommendationList, model.RecommendationList{
			ProductID:   p.ID,
			ProductName: p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	respondSuccess(c, http.StatusOK, &resp, false)
}
