package queue

type Queue []string

func New() Queue {
	var queue Queue
	return queue
}

func (s *Queue) Push(val string) {
	*s = append(*s, val)
}

func (s *Queue) Pop() (string, bool) {
	if (*s).IsEmpty() {
		return "", false
	}

	element := (*s)[0]
	*s = (*s)[1:]

	return element, true
}

func (s *Queue) Top() string {
	if (*s).IsEmpty() {
		return ""
	}

	return (*s)[0]
}

func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}
