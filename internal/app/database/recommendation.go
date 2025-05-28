package database

import (
	"context"
	"net/http"

	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
)

func (db *SqlDb) GetUserRecommendation(ctx context.Context, userID uint) (products []dbModel.UserRecommendation, wrapErr errcode.WrapErr) {

	err := db.Orm.WithContext(ctx).
		Model(&dbModel.UserRecommendation{}).
		Where("user_id = ?", userID).
		Find(&products).Error

	if err != nil {
		log.C(ctx).Error("GetUserRecommendationProductsAll failed to get list: %w", err)
		wrapErr = errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBGetUserRecommendationFailed,
			RawErr:     err,
		}
	}

	return
}
