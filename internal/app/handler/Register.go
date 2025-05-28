package handler

import (
	"fmt"
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
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

	// Check if the email already exists
	if count, err := database.GetSqlDb().CountUserByEmail(c, req.Email); err.RawErr != nil {
		respondError(c, err)
		return
	} else if count > 0 {
		err := fmt.Errorf("email already exists: %s", req.Email)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusConflict,
			ErrCode:    errcode.DBDuplicatedUser,
			RawErr:     err,
		})
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := db.User{
		Email:    req.Email,
		Password: string(hashed),
	}
	// TODO: refactor to database package as a method
	if err := database.GetSqlDb().Orm.Create(&user).Error; err != nil {
		err := fmt.Errorf("failed to create user: %w", err)
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBCreateUserFailed,
			RawErr:     err,
		})
		return
	}

	// TODO: gen email verify token and send email

	resp := model.CommonSuccessRes{}

	respondSuccess(c, http.StatusOK, &resp, false)

}
