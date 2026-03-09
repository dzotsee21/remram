package internal

import "github.com/bendahl/uinput"

var KeyToUinput = map[rune]int{
    // Lowercase Letters
    'a': uinput.KeyA, 'b': uinput.KeyB, 'c': uinput.KeyC, 'd': uinput.KeyD,
    'e': uinput.KeyE, 'f': uinput.KeyF, 'g': uinput.KeyG, 'h': uinput.KeyH,
    'i': uinput.KeyI, 'j': uinput.KeyJ, 'k': uinput.KeyK, 'l': uinput.KeyL,
    'm': uinput.KeyM, 'n': uinput.KeyN, 'o': uinput.KeyO, 'p': uinput.KeyP,
    'q': uinput.KeyQ, 'r': uinput.KeyR, 's': uinput.KeyS, 't': uinput.KeyT,
    'u': uinput.KeyU, 'v': uinput.KeyV, 'w': uinput.KeyW, 'x': uinput.KeyX,
    'y': uinput.KeyY, 'z': uinput.KeyZ,

    // Numbers
    '1': uinput.Key1, '2': uinput.Key2, '3': uinput.Key3, '4': uinput.Key4,
    '5': uinput.Key5, '6': uinput.Key6, '7': uinput.Key7, '8': uinput.Key8,
    '9': uinput.Key9, '0': uinput.Key0,

    // Symbols (Base keys)
    '-':  uinput.KeyMinus,
    '=':  uinput.KeyEqual,
    '[':  uinput.KeyLeftbrace,
    ']':  uinput.KeyRightbrace,
    '\\': uinput.KeyBackslash,
    ';':  uinput.KeySemicolon,
    '\'': uinput.KeyApostrophe,
    ',':  uinput.KeyComma,
    '.':  uinput.KeyDot,
    '/':  uinput.KeySlash,
    '`':  uinput.KeyGrave,

    // Control Characters
    ' ':    uinput.KeySpace,
    '\n':   uinput.KeyEnter,
    '\r':   uinput.KeyEnter,
    '\t':   uinput.KeyTab,
    '\b':   uinput.KeyBackspace,
    '\x1b': uinput.KeyEsc,
}
