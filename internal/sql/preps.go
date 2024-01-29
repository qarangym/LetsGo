package sql

import (
	"database/sql"
	"log"
)

type exampleModelWithStmt struct {
	DB         *sql.DB
	insertStmt *sql.Stmt
}

func NewExampleModel(db *sql.DB) (*exampleModelWithStmt, error) {
	inStmt, err := db.Prepare("INSERT INTO...")
	if err != nil {
		return nil, err
	}

	return &exampleModelWithStmt{db, inStmt}, nil
}

func (m *exampleModelWithStmt) Insert(args ...interface{}) error {

	if _, err := m.insertStmt.Exec(args...); err != nil {
		log.Printf("prepared Insert statement failed: %v", err)
		return err
	}

	return nil
}

func examplMain() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	exampleModelWithStmt, err := NewExampleModel(db)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := exampleModelWithStmt.insertStmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
