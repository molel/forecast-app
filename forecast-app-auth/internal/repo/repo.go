package repo

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	conn *sqlx.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Init(databaseUser, databasePassword, databaseAddress, databaseName string) error {
	var err error

	databaseSource := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		databaseUser,
		databasePassword,
		databaseAddress,
		databaseName,
	)

	r.conn, err = sqlx.Connect("pgx", databaseSource)
	if err != nil {
		return fmt.Errorf("cannot connect to database: %s", err)
	}

	var i int
	for i = 0; i < 10; i++ {
		log.Println("Pinging database...")

		if err = r.conn.Ping(); err == nil {
			break
		}
		log.Printf("Cannot ping database: %s, trying again", err)

		<-time.After(time.Second * 5)
	}
	if i == 10 {
		return fmt.Errorf("cannot ping database: %s", err)
	}

	if ifUserExistsQuery, err = r.conn.Preparex("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1);"); err != nil {
		return fmt.Errorf("cannot prepare ifUserExists query: %s", err)
	}
	if insertUserQuery, err = r.conn.Preparex("INSERT INTO users (username, password) VALUES ($1, $2);"); err != nil {
		return fmt.Errorf("cannot prepare ifUserExists query: %s", err)
	}
	if getUserPasswordQuery, err = r.conn.Preparex("SELECT password FROM users WHERE username=$1;"); err != nil {
		return fmt.Errorf("cannot prepare ifUserExists query: %s", err)
	}

	return nil
}
