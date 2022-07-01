package repo_user

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) int64
	Update(ctx context.Context, tx *sql.Tx, user domain.User) int64
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindId(ctx context.Context, tx *sql.Tx, userId int) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
