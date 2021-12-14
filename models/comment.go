package models

type Comment struct {
	ID      string `db:"id"`
	PostID  string `db:"post_id"`
	Content string `db:"content"`
	Votes   int    `db:"votes"`
}

type CommentStore interface {
	Comment(id string) (Comment, error)
	CommentsByPost(postID string) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id string) error
}
