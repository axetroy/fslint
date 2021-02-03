package parser

import (
	"unicode/utf8"

	"github.com/axetroy/fslint/internal/char_state"
)

func IsPascalCase(str string) bool {
	if !utf8.ValidString(str) {
		return false
	}

	state := char_state.NewState(str)

	for {
		char := state.Next()

		if char == nil {
			break
		}

		prev := char.Prev()

		switch true {
		case char.Is(char_state.CharTypeLowerCase):
			if prev != nil && prev.Is(char_state.CharTypeNumber) {
				return false
			}
			// if first char
			if char.Index() == 0 {
				return false
			}
		case char.Is(char_state.CharTypeUpperCase):
			continue
		case char.Is(char_state.CharTypeNumber):
			if char.Index() == 0 {
				return false
			}
		default:
			return false
		}
	}

	return true
}
