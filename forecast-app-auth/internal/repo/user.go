package repo

import "github.com/jmoiron/sqlx"

var (
	ifUserExistsQuery    *sqlx.Stmt
	insertUserQuery      *sqlx.Stmt
	getUserPasswordQuery *sqlx.Stmt
)

func (r *Repository) CheckUser(username string) (bool, error) {
	var ok bool
	err := ifUserExistsQuery.QueryRowx(username).Scan(&ok)
	return ok, err
}

func (r *Repository) CreateUser(username, password string) error {
	_, err := insertUserQuery.Exec(username, password)
	return err
}

func (r *Repository) GetUserPassword(username string) (string, error) {
	var password string
	err := getUserPasswordQuery.QueryRowx(username).Scan(&password)
	return password, err
}
