package parser

import (
	"unicode/utf8"

	"github.com/axetroy/fslint/internal/char_state"
)

func IsKebabCase(str string) bool {
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
		next := char.Next()

		switch true {
		case char.Is(char_state.CharTypeLowerCase):
			if prev != nil && prev.Is(char_state.CharTypeNumber) {
				return false
			}
		case char.Is(char_state.CharTypeUpperCase):
			if char.Index() == 0 {
				return false
			}
			if prev != nil && !prev.Is(char_state.CharTypeKebab) {
				return false
			}
		case char.Is(char_state.CharTypeNumber):
			if char.Index() == 0 {
				return false
			}
		case char.Is(char_state.CharTypeKebab):
			if prev != nil && prev.Is(char_state.CharTypeKebab) {
				return false
			}
			if next != nil {
				if next.Is(char_state.CharTypeKebab) {
					return false
				}

				if !next.Is(char_state.CharTypeLowerCase) && !next.Is(char_state.CharTypeNumber) {
					return false
				}
			}
		default:
			return false
		}
	}

	return true
}
