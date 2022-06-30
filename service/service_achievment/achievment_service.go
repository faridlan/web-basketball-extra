package service_achievment

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/web/web_achievment"
	"github.com/faridlan/web-basketball-extra/repository/repo_achievment"
)

type AchievmentServiceImpl struct {
	RepoAchiev repo_achievment.AchievmentRepository
	DB         *sql.DB
}

func (service AchievmentServiceImpl) Create(ctx context.Context, request web_achievment.AchievmentCreate) int64 {
	panic("not implemented") // TODO: Implement
}

func (service AchievmentServiceImpl) Update(ctx context.Context, request web_achievment.AchievmentUpdate) int64 {
	panic("not implemented") // TODO: Implement
}

func (service AchievmentServiceImpl) Delete(ctx context.Context, achievId int) {
	panic("not implemented") // TODO: Implement
}

func (service AchievmentServiceImpl) FindById(ctx context.Context, achievId int) web_achievment.AchievmentResponse {
	panic("not implemented") // TODO: Implement
}

func (service AchievmentServiceImpl) FindAll(ctx context.Context) []web_achievment.AchievmentResponse {
	panic("not implemented") // TODO: Implement
}
