package models

type Post struct {
	ID       string `db:"id"`
	ThreadID string `db:"thread_id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Votes    int    `db:"votes"`
}

type PostStore interface {
	Post(id string) (Post, error)
	PostsByThread(threadID string) ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id string) error
}
