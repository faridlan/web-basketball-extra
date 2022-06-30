package service_status

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/web/web_status"
	"github.com/faridlan/web-basketball-extra/repository/repo_status"
)

type StatusServiceimpl struct {
	RepoStatus repo_status.StatusRepository
	DB         *sql.DB
}

func (service StatusServiceimpl) Create(ctx context.Context, request web_status.StatusCreateReq) int64 {
	panic("not implemented") // TODO: Implement
}

func (service StatusServiceimpl) Update(ctx context.Context, request web_status.StatusUpdateReq) int64 {
	panic("not implemented") // TODO: Implement
}

func (service StatusServiceimpl) Delete(ctx context.Context, statusId int) {
	panic("not implemented") // TODO: Implement
}

func (service StatusServiceimpl) FindById(ctx context.Context, statusId int) web_status.StatusResponse {
	panic("not implemented") // TODO: Implement
}

func (service StatusServiceimpl) FindAll(ctx context.Context) []web_status.StatusResponse {
	panic("not implemented") // TODO: Implement
}
