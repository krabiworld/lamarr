package types

type Category string

const (
	INFORMATION Category = "INFORMATION"
	MODERATION  Category = "MODERATION"
	SETTINGS    Category = "SETTINGS"
	UTILITIES   Category = "UTILITIES"
)

func (c Category) String() string {
	return string(c)
}
