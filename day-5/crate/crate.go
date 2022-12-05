package crate

type (
	Crate  string
	Stack  []Crate
	Stacks []Stack
)

func (s Stacks) Move(n, from, to int) Stacks {
	if s == nil {
		return s
	}

	for i := 0; i < n; i++ {
		f := s[from]

		toMove := f[len(f)-1]
		s[from] = f[:len(f)-1]
		s[to] = append(s[to], toMove)
	}
	return s
}

func (s Stacks) GetTops() string {

	var result string
	for _, stack := range s {
		if len(stack) > 0 {
			result += string(stack[len(stack)-1])
		} else {
			result += " "
		}
	}
	return result
}

func (s Stacks) MoveMultiple(n, from, to int) Stacks {
	if s == nil {
		return s
	}

	f := s[from]
	toMove := f[len(f)-n:]
	s[from] = f[:len(f)-n]
	s[to] = append(s[to], toMove...)

	return s
}
