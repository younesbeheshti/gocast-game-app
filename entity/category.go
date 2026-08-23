package entity

type Category string

const (
	FootballCategory Category = "football"
)

func (c Category) IsValid() bool {
	switch c {
	case FootballCategory:
		return true
	}
	return false
}
