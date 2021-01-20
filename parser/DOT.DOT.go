package parser

import (
	"unicode/utf8"

	"github.com/axetroy/fslint/char_state"
)

func IsUppercaseDotUppercase(str string) bool {
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

		switch true {
		case char.Is(char_state.CharTypeUpperCase):
			continue
		case char.Is(char_state.CharTypeDot):
			if index == 0 || index == state.Len()-1 {
				return false
			}
		default:
			return false
		}
	}

	return true
}
