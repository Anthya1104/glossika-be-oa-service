package model

import "github.com/Anthya1104/gin-base-service/pkg/config"

type CommonResInterface interface {
	SetVersion()
}

type CommonRes struct {
	Version string `json:"version"`
	Error   string `json:"error"`
}

func (r *CommonRes) SetVersion() {
	r.Version = config.EnvVariable.Version
}

type CommonErrorRes struct {
	CommonRes
	Msg string `json:"msg"`
}

type CommonSuccessRes struct {
	CommonRes
	Data string `json:"data"`
}
