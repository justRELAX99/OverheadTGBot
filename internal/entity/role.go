package entity

const (
	AdminRole  Role = "admin"
	AuthorRole Role = "author"
)

type Role string

func (r Role) IsAdmin() bool {
	return r == AdminRole
}
