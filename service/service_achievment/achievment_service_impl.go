package service_achievment

import (
	"context"

	"github.com/faridlan/web-basketball-extra/model/web/web_achievment"
)

type AchievmentService interface {
	Create(ctx context.Context, request web_achievment.AchievmentCreate) int64
	Update(ctx context.Context, request web_achievment.AchievmentUpdate) int64
	Delete(ctx context.Context, achievId int)
	FindById(ctx context.Context, achievId int) web_achievment.AchievmentResponse
	FindAll(ctx context.Context) []web_achievment.AchievmentResponse
}
