package stack

type Stack []string

func New() Stack {
	var stack Stack
	return stack
}

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (string, bool) {
	if (*s).IsEmpty() {
		return "", false
	}
	ind := len(*s) - 1
	element := (*s)[ind]
	*s = (*s)[:ind]

	return element, true
}

func (s *Stack) Top() string {
	if (*s).IsEmpty() {
		return ""
	}
	ind := len(*s) - 1

	return (*s)[ind]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
