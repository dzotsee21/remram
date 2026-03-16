//go:build windows

package internal

var KeyToRobotGo = map[rune]string{
	// Lowercase Letters
	'a': "a", 'b': "b", 'c': "c", 'd': "d",
	'e': "e", 'f': "f", 'g': "g", 'h': "h",
	'i': "i", 'j': "j", 'k': "k", 'l': "l",
	'm': "m", 'n': "n", 'o': "o", 'p': "p",
	'q': "q", 'r': "r", 's': "s", 't': "t",
	'u': "u", 'v': "v", 'w': "w", 'x': "x",
	'y': "y", 'z': "z",

	// Numbers
	'1': "1", '2': "2", '3': "3", '4': "4",
	'5': "5", '6': "6", '7': "7", '8': "8",
	'9': "9", '0': "0",

	// Symbols
	'-':  "-",
	'=':  "=",
	'[':  "[",
	']':  "]",
	'\\': "\\",
	';':  ";",
	'\'': "'",
	',':  ",",
	'.':  ".",
	'/':  "/",
	'`':  "`",

	// Control & Whitespace
	' ':    "space",
	'\n':   "enter",
	'\r':   "enter",
	'\t':   "tab",
	'\b':   "backspace",
	'\x1b': "esc",
	'\x10': "shift",

	// Function Keys
	'\x01': "f1", '\x02': "f2", '\x03': "f3",
	'\x04': "f4", '\x05': "f5", '\x06': "f6",
	'\x07': "f7", '\x0b': "f11", '\x0c': "f12",
}
