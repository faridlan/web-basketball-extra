package service_user

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/web/web_user"
	"github.com/faridlan/web-basketball-extra/repository/repo_user"
)

type UserServiceImpl struct {
	RepoUser repo_user.UserRepository
	DB       *sql.DB
}

func (service UserServiceImpl) Create(ctx context.Context, request web_user.UserCreateReq) int64 {
	panic("not implemented") // TODO: Implement
}

func (service UserServiceImpl) Update(ctx context.Context, request web_user.UserUpdateReq) int64 {
	panic("not implemented") // TODO: Implement
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) {
	panic("not implemented") // TODO: Implement
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) web_user.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service UserServiceImpl) FindAll(ctx context.Context) []web_user.UserResponse {
	panic("not implemented") // TODO: Implement
}
