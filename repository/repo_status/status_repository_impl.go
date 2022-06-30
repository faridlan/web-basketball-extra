package repo_status

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type StatusRepositoryImpl struct {
}

func (repository StatusRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, status domain.Status) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository StatusRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, status domain.Status) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository StatusRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, statusId int) {
	panic("not implemented") // TODO: Implement
}

func (repository StatusRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, statusId int) domain.Status {
	panic("not implemented") // TODO: Implement
}

func (repository StatusRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Status {
	panic("not implemented") // TODO: Implement
}
