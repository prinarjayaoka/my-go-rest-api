package dao

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prinarjayaoka/my-go-rest-api/models"
)

type DAO struct {
	models.ThreadStore
	models.PostStore
	models.CommentStore
}

func NewStore(dataSourceName string) (*DAO, error) {
	db, errConn := sqlx.Open("mysql", dataSourceName)
	if errConn != nil {
		return nil, fmt.Errorf("error connecting to database: %w", errConn)
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", errPing)
	}

	return &DAO{
		ThreadStore: NewThreadStore(db),
	}, nil
}
