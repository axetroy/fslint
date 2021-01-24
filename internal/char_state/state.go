package char_state

type State struct {
	char  string
	index int
}

func (s *State) PreviewPrev() *Char {
	if s.index == 0 {
		return nil
	}

	char := string(s.char[s.index-1])

	return &Char{
		allString: s.char,
		val:       char,
		index:     s.index - 1,
	}
}

func (s *State) Prev() *Char {
	if s.index == 0 {
		return nil
	}

	char := string(s.char[s.index-1])

	s.index = s.index - 1

	return &Char{
		allString: s.char,
		val:       char,
		index:     s.index,
	}
}

func (s *State) Next() *Char {
	if int(s.index) >= len(s.char)-1 {
		return nil
	}

	char := string(s.char[s.index+1])

	s.index = s.index + 1

	return &Char{
		allString: s.char,
		val:       char,
		index:     s.index,
	}
}

func (s *State) PreviewNext() *Char {
	if int(s.index) >= len(s.char)-1 {
		return nil
	}

	char := string(s.char[s.index+1])

	return &Char{
		allString: s.char,
		val:       char,
		index:     s.index + 1,
	}
}

func (s *State) Index() int {
	return int(s.index)
}

func (s *State) Len() int {
	return len(s.char)
}

func NewState(char string) State {
	return State{
		char:  char,
		index: -1,
	}
}
