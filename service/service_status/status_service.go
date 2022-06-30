package service_status

import (
	"context"

	"github.com/faridlan/web-basketball-extra/model/web/web_status"
)

type StatusService interface {
	Create(ctx context.Context, request web_status.StatusCreateReq) int64
	Update(ctx context.Context, request web_status.StatusUpdateReq) int64
	Delete(ctx context.Context, statusId int)
	FindById(ctx context.Context, statusId int) web_status.StatusResponse
	FindAll(ctx context.Context) []web_status.StatusResponse
}
