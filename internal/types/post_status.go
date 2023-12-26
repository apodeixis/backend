package types

type PostStatus string

const (
	NewPostStatus       PostStatus = "new"
	PendingPostStatus   PostStatus = "pending"
	ConfirmedPostStatus PostStatus = "confirmed"
	FailedPostStatus    PostStatus = "failed"
)
