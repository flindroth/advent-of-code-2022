package util

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(elements ...T) {
	for _, v := range elements {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Has(element T) bool {
	if _, ok := s[element]; ok {
		return true
	}
	return false
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	ns := make(Set[T])
	for k := range s {
		if other.Has(k) {
			ns.Add(k)
		}
	}
	return ns
}

func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

func IntSet(start, end int) Set[int] {
	set := make(Set[int])
	for i := start; i <= end; i++ {
		set.Add(i)
	}
	return set
}
