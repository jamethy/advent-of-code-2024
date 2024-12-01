package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(t ...T) {
	for _, o := range t {
		s[o] = struct{}{}
	}
}

func (s Set[T]) AddAll(other Set[T]) {
	for o := range other {
		s.Add(o)
	}
}

func (s Set[T]) RemoveAll(other Set[T]) {
	for o := range other {
		s.Remove(o)
	}
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) Remove(t ...T) int {
	count := 0
	for _, o := range t {
		if _, ok := s[o]; ok {
			count++
			delete(s, o)
		}
	}
	return count
}

func (s Set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Retain(t ...T) int {
	o := NewSet(t...)
	return s.RetainSet(o)
}

func (s Set[T]) RetainSet(o Set[T]) int {
	count := 0
	for k := range s {
		if !o.Has(k) {
			count++
			delete(s, k)
		}
	}
	return count
}

func (s Set[T]) Intersection(o Set[T]) Set[T] {
	i := s.Clone()
	i.RetainSet(o)
	return i
}

func (s Set[T]) Slice() []T {
	l := make([]T, 0, len(s))
	for s2 := range s {
		l = append(l, s2)
	}
	return l
}

func NewSet[T comparable](t ...T) Set[T] {
	s := make(Set[T], len(t))
	for _, o := range t {
		s[o] = struct{}{}
	}
	return s
}

func (s Set[T]) Clone() Set[T] {
	o := make(Set[T], len(s))
	for i := range s {
		o[i] = struct{}{}
	}
	return o
}

func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	shared := s1.Clone()
	shared.RetainSet(s2)
	return shared
}
