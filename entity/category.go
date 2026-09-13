package entity

type Category string

const (
	FootballCategory Category = "football"
	HistoryCategory  Category = "history"
)

func (c Category) IsValid() bool {
	switch c {
	case FootballCategory:
		return true
	case HistoryCategory:
		return true
	}
	return false
}

func CategoryList() []Category {
	return []Category{
		FootballCategory,
		HistoryCategory,
	}
}
