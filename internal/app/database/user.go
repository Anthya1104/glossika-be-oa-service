package database

import (
	"context"
	"net/http"

	dbModel "github.com/Anthya1104/gin-base-service/internal/app/model/db"
	"github.com/Anthya1104/gin-base-service/pkg/errcode"
)

func (db *SqlDb) GetUserInfoByUserId(ctx context.Context, accountId string) (userInfo dbModel.UserInfo, wrapErr errcode.WrapErr) {
	err := db.Orm.WithContext(ctx).
		Model(&dbModel.UserInfo{}).
		Where("user_id = ? AND (expired_at IS NULL OR expired_at > ?)", accountId).
		Select("COALESCE(SUM(remaining_points), 0)").
		Row().Scan(&userInfo.Id)

	if err != nil {
		wrapErr = errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBGetUserInfoFailed,
			RawErr:     err,
		}
	}

	return
}
