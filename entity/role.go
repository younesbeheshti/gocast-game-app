package entity

type Role uint8

const (
	UserRole Role = iota + 1
	AdminRole
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return "User"
	case AdminRole:
		return "Admin"
	}
	return ""
}
