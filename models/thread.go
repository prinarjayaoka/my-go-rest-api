package models

type Thread struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type ThreadStore interface {
	Thread(id string) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id string) error
}
