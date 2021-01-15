package char_state

import (
	"regexp"
	"unicode"
)

var (
	numRegExp = regexp.MustCompile(`\d`)
)

type Char struct {
	allString string
	val       string
	index     int
}

func (c *Char) Is(typeIota int) bool {
	r := []rune(c.Val())[0]

	switch typeIota {
	case CharTypeLowerCase:
		return unicode.IsLower(r)
	case CharTypeUpperCase:
		return unicode.IsUpper(r)
	case CharTypeNumber:
		return numRegExp.MatchString(c.Val())
	case CharTypeKebab:
		return c.Val() == "-"
	case CharTypeUnderscore:
		return c.Val() == "_"
	default:
		return false
	}
}

func (c *Char) Val() string {
	return c.val
}

func (c *Char) Index() int {
	return c.index
}

func (s *Char) Prev() *Char {
	if s.index == 0 {
		return nil
	}

	char := string(s.allString[s.index-1])

	return &Char{
		allString: s.allString,
		val:       char,
		index:     s.index - 1,
	}
}

func (s *Char) Next() *Char {
	if int(s.index) >= len(s.allString)-1 {
		return nil
	}

	char := string(s.allString[s.index+1])

	return &Char{
		allString: s.allString,
		val:       char,
		index:     s.index + 1,
	}
}
