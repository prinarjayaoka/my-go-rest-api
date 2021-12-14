package dao

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prinarjayaoka/my-go-rest-api/models"
)

type Models struct {
	models.ThreadStore
	models.PostStore
	models.CommentStore
}

func NewStore(dataSourceName string) (*Models, error) {
	db, errConn := sqlx.Open("mysql", dataSourceName)
	if errConn != nil {
		return nil, fmt.Errorf("error connecting to database: %w", errConn)
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", errPing)
	}

	return &Models{
		ThreadStore: NewThreadStore(db),
	}, nil
}
