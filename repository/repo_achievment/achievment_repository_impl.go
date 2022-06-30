package repo_achievment

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type AchievmentRepositoryImpl struct {
}

func (repository AchievmentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, achievment domain.Achievment) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository AchievmentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, achievment domain.Achievment) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository AchievmentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, achievId int) {
	panic("not implemented") // TODO: Implement
}

func (repository AchievmentRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, achievId int) domain.Achievment {
	panic("not implemented") // TODO: Implement
}

func (repository AchievmentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Achievment {
	panic("not implemented") // TODO: Implement
}
