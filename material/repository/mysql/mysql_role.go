package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"example/e-learn/domain"
	"example/e-learn/material/repository"
)

type mysqlRoleRepository struct {
	Conn *sql.DB
}

// NewMysqlRoleRepository will create an object that represent the role.Repository interface
func NewMysqlRoleRepository(conn *sql.DB) domain.RoleRepository {
	return &mysqlRoleRepository{conn}
}

func (m *mysqlRoleRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Role, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Role, 0)
	for rows.Next() {
		t := domain.Role{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRoleRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Role, nextCursor string, err error) {
	query := `SELECT id, name, updated_at, created_at
  						FROM role WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}
func (m *mysqlRoleRepository) GetByID(ctx context.Context, id int64) (res domain.Role, err error) {
	query := `SELECT id, name, updated_at, created_at
  						FROM role WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Role{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlRoleRepository) Store(ctx context.Context, a *domain.Role) (err error) {
	query := `INSERT role SET name=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Name)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}

func (m *mysqlRoleRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM role WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysqlRoleRepository) Update(ctx context.Context, ar *domain.Role) (err error) {
	query := `UPDATE role set name=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Name, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
