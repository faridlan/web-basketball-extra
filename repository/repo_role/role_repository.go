package repo_role

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/model/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, role domain.Role) int64
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, role domain.Role)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}
