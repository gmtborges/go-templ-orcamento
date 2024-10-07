package types

type AutoCategoryType string

const (
	AutoCategoryTypeProduct AutoCategoryType = "PRODUCT"
	AutoCategoryTypeService AutoCategoryType = "SERVICE"
)

type AutoCategory struct {
	ID          int64
	Description string
	Type        AutoCategoryType
}
