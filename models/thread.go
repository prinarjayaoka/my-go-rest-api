package models

type Thread struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `db:"description" json:"description"`
	IsActive    int    `db:"is_active" json:"isActive"`
}

type ThreadStore interface {
	Thread(id string) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) (int64, error)
	UpdateThread(t *Thread) (int64, error)
	DeleteThread(id string) (int64, error)
}
