package parser

import (
	"unicode/utf8"

	"github.com/axetroy/fslint/internal/char_state"
)

func IsUppercaseWithUppercase(str string) bool {
	if !utf8.ValidString(str) {
		return false
	}

	state := char_state.NewState(str)

	for {
		char := state.Next()

		if char == nil {
			break
		}

		index := char.Index()
		prev := char.Prev()
		next := char.Next()

		switch true {
		case char.Is(char_state.CharTypeUpperCase):
			continue
		case char.Is(char_state.CharTypeNumber):
			if prev != nil && !(prev.Is(char_state.CharTypeKebab) || prev.Is(char_state.CharTypeNumber)) {
				return false
			}
			if next != nil && !(next.Is(char_state.CharTypeKebab) || next.Is(char_state.CharTypeNumber)) {
				return false
			}
			continue
		case char.Is(char_state.CharTypeKebab):
			if index == 0 || index == state.Len()-1 {
				return false
			}
			if prev != nil && prev.Is(char_state.CharTypeKebab) {
				return false
			}
		default:
			return false
		}
	}

	return true
}
