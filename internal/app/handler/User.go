package handler

import (
	"fmt"
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/service"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetUserInfoAPI(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		err := fmt.Errorf("userId not found in query")
		respondError(c, errcode.WrapErr{
			HttpStatus: http.StatusBadRequest,
			ErrCode:    errcode.BadQuery,
			RawErr:     err,
		})
		return
	}

	// get user info from DB
	resp, _ := service.GetUserInfo(c, userId)

	respondSuccess(c, http.StatusOK, &resp, false)

	// c.String(http.StatusOK, "ok")
}
