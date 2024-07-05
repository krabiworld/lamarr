package types

type Color int

const (
	DEFAULT Color = 0x2e79d5
	SUCCESS Color = 0x54de3c
	ERROR   Color = 0xeb1010
	WARN    Color = 0xd4bf33
)

func (c Color) Int() int {
	return int(c)
}
