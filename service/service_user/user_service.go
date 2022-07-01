package service_user

import (
	"context"

	"github.com/faridlan/web-basketball-extra/model/web/web_user"
)

type UserService interface {
	Create(ctx context.Context, request web_user.UserCreateReq) int64
	Update(ctx context.Context, request web_user.UserUpdateReq) int64
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web_user.UserResponse
	FindAll(ctx context.Context) []web_user.UserResponse
}
