package types

type Color int

const (
	ColorDefault Color = 0x2e79d5
	ColorSuccess Color = 0x54de3c
	ColorError   Color = 0xeb1010
	ColorWarn    Color = 0xd4bf33
)

func (c Color) Int() int {
	return int(c)
}
