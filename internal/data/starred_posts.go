package data

type StarredPosts interface {
	New() StarredPosts

	Create(starredPost StarredPost) (*StarredPost, error)
	Delete() error
	Get() (*StarredPost, error)

	FilterByUserID(int64) StarredPosts
	FilterByPostID(int64) StarredPosts

	Transaction(func() error) error
}

type StarredPost struct {
	ID     int64 `db:"id"`
	UserID int64 `db:"user_id"`
	PostID int64 `db:"post_id"`
}
