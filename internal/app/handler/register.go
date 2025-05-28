package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/service"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegisterHandler(c *gin.Context) {
	var req model.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := fmt.Errorf("failed to bind request body: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadRequestBody,
			RawErr:     err,
		})
		return
	}

	if !req.ValidateRegisterEmail() {
		err := fmt.Errorf("invalid email format: %s", req.Email)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadRequest,
			RawErr:     err,
		})
		return
	}

	if !req.ValidateRegisterPassword() {
		err := fmt.Errorf("invalid password format: %s", req.Password)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadRequest,
			RawErr:     err,
		})
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err := fmt.Errorf("failed to hash password: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.BcryptHashFailed,
			RawErr:     err,
		})
		return
	}

	// Create user
	user := db.User{
		Email:    req.Email,
		Password: string(hashed),
	}
	// TODO: refactor to database package as a method
	if err := database.GetSqlDb().Orm.Create(&user).Error; err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		err := fmt.Errorf("email already exists: %s", req.Email)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusConflict,
			ErrCode:    errcode.DBDuplicatedUser,
			RawErr:     err,
		})
		return

	} else if err != nil {
		err := fmt.Errorf("failed to create user: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBCreateUserFailed,
			RawErr:     err,
		})
		return
	}

	token, err := util.GenerateTokenWithType(user.ID, user.Email, "email_verify")
	if err != nil {
		err = fmt.Errorf("failed to generate token: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.JWTGenerateFailed,
			RawErr:     err,
		})
		return
	}
	// should be https in productive environment, here use http for testing purpose
	verifyLink := fmt.Sprintf("http://%s:%s/api/v1/users/verify?token=%s", config.EnvVariable.Host, config.EnvVariable.Port, token)

	// fake send email
	service.SendEmail(c)

	// response the verify mail for test since the email sending is not implemented yet
	// should not response in the productive environment
	resp := model.CommonSuccessRes{
		Data: verifyLink,
	}

	respondSuccess(c, http.StatusOK, &resp, false)

}
