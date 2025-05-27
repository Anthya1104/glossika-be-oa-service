package handler

import (
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func VersionHandler(c *gin.Context) {
	version := config.EnvVariable.Version
	c.String(http.StatusOK, version)
}

func respondSuccess(c *gin.Context, httpStatus int, resp model.CommonResInterface, logRespBody bool) {
	resp.SetVersion()

	if logRespBody {
		log.C(c).Infof("[Response][Success] %v %v respStatus=%v respBody=%v", c.Request.Method, c.Request.URL.String(), httpStatus, util.StructToJsonString(resp))
	} else {
		log.C(c).Infof("[Response][Success] %v %v respStatus=%v", c.Request.Method, c.Request.URL.String(), httpStatus)
	}

	c.JSON(httpStatus, resp)
}

func respondError(c *gin.Context, wrapErr errcode.WrapErr) {
	resp := model.CommonErrorRes{
		CommonRes: model.CommonRes{
			Version: config.EnvVariable.Version,
			Error:   string(wrapErr.ErrCode),
		},
		Msg: errcode.ErrCodeMsg[wrapErr.ErrCode],
	}

	log.C(c).Errorf("[Response][Error] %v %v respStatus=%v respBody=%v rawErr=%v", c.Request.Method, c.Request.URL.String(), wrapErr.HttpStatus, util.StructToJsonString(resp), wrapErr.RawErr)

	c.JSON(wrapErr.HttpStatus, resp)
}
