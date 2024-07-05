package types

type Category string

const (
	INFORMATION Category = "Information"
	MODERATION  Category = "Moderation"
	SETTINGS    Category = "Settings"
	UTILITIES   Category = "Utilities"
)

func (c Category) String() string {
	return string(c)
}
