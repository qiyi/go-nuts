package sets

type StrSet struct {
	m map[string]bool
}

func (set *StrSet) Add(s string) bool {
	contains := set.Contains(s)
	set.m[s] = true
	return contains
}

func (set *StrSet) Remove(s string) bool {
	contains := set.Contains(s)
	delete(set.m, s)
	return !contains 
}

func (set *StrSet) Size() int {
	return len(set.m)
}

func (set *StrSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *StrSet) Contains(s string) bool {
	_, contains := set.m[s]
	return contains
}


