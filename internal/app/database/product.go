package database

import (
	"context"
	"net/http"

	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
)

func (db *SqlDb) BatchGetProducts(ctx context.Context, productIDList []uint) (products []dbModel.Product, wrapErr errcode.WrapErr) {

	err := db.Orm.WithContext(ctx).
		Model(&dbModel.Product{}).
		Where("id IN ?", productIDList).
		Find(&products).Error

	if err != nil {
		log.C(ctx).Error("BatchGetProducts failed to get list: %w", err)
		wrapErr = errcode.WrapErr{
			HttpStatus: http.StatusInternalServerError,
			ErrCode:    errcode.DBGetProductFailed,
			RawErr:     err,
		}
	}

	return
}
