package domain

type Post struct {
	ID     uint32
	UserID uint32
	Text   string

	User *User
}
