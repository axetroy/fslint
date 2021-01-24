package parser

import (
	"unicode/utf8"

	"github.com/axetroy/fslint/internal/char_state"
)

func IsSnakeCase(str string, isBig bool) bool {
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
			// if first char
			if isBig && char.Index() == 0 {
				return false
			}
		case char.Is(char_state.CharTypeUpperCase):
			if !isBig && char.Index() == 0 {
				return false
			}
			if char.Index() != 0 {
				if !prev.Is(char_state.CharTypeUnderscore) {
					return false
				}
			}
		case char.Is(char_state.CharTypeNumber):
			if char.Index() == 0 {
				return false
			}
		case char.Is(char_state.CharTypeUnderscore):
			if char.Index() == 0 {
				return false
			}
			if next != nil {
				if next.Is(char_state.CharTypeUnderscore) {
					return false
				}
				if isBig && !next.Is(char_state.CharTypeUpperCase) {
					return false
				}
				if !isBig && !next.Is(char_state.CharTypeLowerCase) {
					return false
				}
			}
		default:
			return false
		}
	}

	return true
}
