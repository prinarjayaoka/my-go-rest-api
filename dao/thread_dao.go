package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/prinarjayaoka/my-go-rest-api/models"
)

type ThreadStore struct {
	*sqlx.DB
}

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{DB: db}
}

func (s *ThreadStore) Thread(id string) (models.Thread, error) {
	var t models.Thread
	err := s.Get(&t, `SELECT * FROM threads WHERE id = ?`, id)

	if err != nil {
		return models.Thread{}, fmt.Errorf("error retrieving thread: %w", err)
	}

	return t, nil
}

func (s *ThreadStore) Threads() ([]models.Thread, error) {
	var tt []models.Thread
	err := s.Select(&tt, `SELECT * FROM threads`)

	if err != nil {
		return []models.Thread{}, fmt.Errorf("error retrieving threads: %w", err)
	}

	return tt, nil
}

func (s *ThreadStore) CreateThread(t *models.Thread) (int64, error) {
	result, err := s.Exec(
		`INSERT INTO threads (id, title, description) VALUES (?,?,?)`,
		t.ID,
		t.Title,
		t.Description,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (s *ThreadStore) UpdateThread(t *models.Thread) (int64, error) {
	result, err := s.Exec(
		`UPDATE threads SET title=?, description=? WHERE id=?`,
		t.Title,
		t.Description,
		t.ID,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (s *ThreadStore) DeleteThread(id string) (int64, error) {
	result, err := s.Exec(`DELETE FROM threads WHERE id=?`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
