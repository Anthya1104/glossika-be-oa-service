package handler

import (
	"fmt"
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func UserActivateHandler(c *gin.Context) {
	token := c.Query("token")
	claims, err := util.ParseToken(token)
	if err != nil || claims["type"] != "email_verify" {
		err := fmt.Errorf("invalid or expired token: %v", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.UserInvalidAuth,
			RawErr:     err,
		})
		return
	}
	userID := uint(claims["user_id"].(float64))

	// set user as active in DB
	// TODO: make this as method in database package and refactor with create user function as upsert
	if err := database.GetSqlDb().Orm.Model(&db.User{}).
		Where("id = ?", userID).
		Update("is_activated", true).Error; err != nil {
		err = fmt.Errorf("failed to activate user: %v", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBUpdateUserFailed,
			RawErr:     err,
		})
		return
	}

	resp := model.CommonSuccessRes{}
	respondSuccess(c, http.StatusOK, &resp, false)
}
