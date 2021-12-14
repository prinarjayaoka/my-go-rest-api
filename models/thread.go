package models

type Thread struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ThreadStore interface {
	Thread(id string) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) (int64, error)
	UpdateThread(t *Thread) (int64, error)
	DeleteThread(id string) (int64, error)
}
