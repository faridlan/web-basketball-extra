package repo_status

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type StatusRepository interface {
	Save(ctx context.Context, tx *sql.Tx, status domain.Status) int64
	Update(ctx context.Context, tx *sql.Tx, status domain.Status) int64
	Delete(ctx context.Context, tx *sql.Tx, statusId int)
	FindId(ctx context.Context, tx *sql.Tx, statusId int) domain.Status
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Status
}
