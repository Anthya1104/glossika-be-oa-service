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
	"golang.org/x/crypto/bcrypt"
)

func UserLoginHandler(c *gin.Context) {
	var req model.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadRequestBody,
			RawErr:     err,
		})
		return
	}

	// find the user from DB
	var user db.User
	if err := database.GetSqlDb().Orm.Where("user_email = ?", req.Email).First(&user).Error; err != nil {
		err := fmt.Errorf("failed to find user by email: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusUnauthorized,
			ErrCode:    errcode.DBUserNotFound,
			RawErr:     err,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		err := fmt.Errorf("password mismatch for user %s: %w", user.Email, err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusUnauthorized,
			ErrCode:    errcode.UserInvalidAuth,
			RawErr:     err,
		})
		return
	}

	if !user.IsActivated {
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusForbidden,
			ErrCode:    errcode.UserNotActivated,
			RawErr:     fmt.Errorf("user not activated"),
		})
		return
	}

	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		err := fmt.Errorf("failed to generate JWT token: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.JWTGenerateFailed,
			RawErr:     err,
		})
		return
	}

	resp := model.UserLoginResp{
		Data: model.UserLoginRespData{
			Token: token,
		},
	}

	respondSuccess(c, http.StatusOK, &resp, false)

}
