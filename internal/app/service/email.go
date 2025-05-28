package service

import (
	"context"

	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
)

func SendEmail(ctx context.Context) {
	log.C(ctx).Info("SendEmail function called")
}
