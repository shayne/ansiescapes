package ansiescapes

import "strconv"

// ESC blah blah
const ESC = "\u001b["

// EraseEndLine blah blah
const EraseEndLine = ESC + "K"

// CursorLeft blah blah
const CursorLeft = ESC + "1000D"

// CursorUp blah blah
func CursorUp(count int) string {
	return ESC + strconv.Itoa(count) + "A"
}

// EraseLines blah blah
func EraseLines(count int) string {
	clear := ""
	for i := 0; i < count; i++ {
		clear += CursorLeft + EraseEndLine
		if i < count-1 {
			clear += CursorUp(1)
		}
	}
	return clear
}
