package types

type Category string

const (
	CategoryInformation Category = "Information"
	CategoryModeration  Category = "Moderation"
	CategorySettings    Category = "Settings"
	CategoryUtilities   Category = "Utilities"
)

func (c Category) String() string {
	return string(c)
}
