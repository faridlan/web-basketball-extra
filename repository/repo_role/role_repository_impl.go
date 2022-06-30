package repo_role

import (
	"context"
	"database/sql"

	"github.com/faridlan/web-basketball-extra/helper"
	"github.com/faridlan/web-basketball-extra/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return RoleRepositoryImpl{}
}

func (repository RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) int64 {
	SQL := "insert into roles(name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, role.Name)
	helper.PanicIfError(err)

	r, err := result.RowsAffected()
	helper.PanicIfError(err)

	return r
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	panic("not implemented") // TODO: Implement
}

func (repository RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, role domain.Role) {
	panic("not implemented") // TODO: Implement
}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id,name from roles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	roles := []domain.Role{}

	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.Name)
		helper.PanicIfError(err)

		roles = append(roles, role)
	}

	return roles
}
