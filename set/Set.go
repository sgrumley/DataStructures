package set

import "fmt"

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func (s *Set) Has(key string) bool {
	_, exists := (*s)[key]
	return exists
}

func (s *Set) Add(key string) bool {
	(*s)[key] = struct{}{}

	return true
}

func (s *Set) List() {
	for key, _ := range *s {
		fmt.Println(key)
	}
}
