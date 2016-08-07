package maze

type backtrack struct {
}

func (b backtrack) Choose(length int) int {
	return length - 1
}

// Backtrack returns a backtrack algorithm implementation
func Backtrack() Chooser {
	return backtrack{}
}
