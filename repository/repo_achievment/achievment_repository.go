package repo_achievment

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type AchievmentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, achievment domain.Achievment) int64
	Update(ctx context.Context, tx *sql.Tx, achievment domain.Achievment) int64
	Delete(ctx context.Context, tx *sql.Tx, achievId int)
	FindId(ctx context.Context, tx *sql.Tx, achievId int) domain.Achievment
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Achievment
}
