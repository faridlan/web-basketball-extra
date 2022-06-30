package repo_role

import (
	"context"
	"database/sql"
	"log"

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

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) int64 {
	SQL := "update roles set name = ? where id = ?"
	result, err := tx.ExecContext(ctx, SQL, role.Name, role.Id)
	helper.PanicIfError(err)

	r, err := result.RowsAffected()
	helper.PanicIfError(err)

	return r
}

func (repository RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, roleId int) {
	SQL := "delete from roles where id = ?"
	result, err := tx.ExecContext(ctx, SQL, roleId)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	helper.PanicIfError(err)
}

func (repository RoleRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, roleId int) domain.Role {
	SQL := "select id,name from roles where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)

	defer rows.Close()

	role := domain.Role{}
	for rows.Next() {
		err := rows.Scan(&role.Id, &role.Name)
		helper.PanicIfError(err)
		log.Print(role)
	}
	return role
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
