package repo_user

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) int64 {
	panic("not implemented") // TODO: Implement
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	panic("not implemented") // TODO: Implement
}

func (repository UserRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, userId int) domain.User {
	panic("not implemented") // TODO: Implement
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	panic("not implemented") // TODO: Implement
}
