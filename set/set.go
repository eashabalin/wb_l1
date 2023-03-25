package set

import (
	"fmt"
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](slice []T) Set[T] {
	m := make(map[T]struct{})
	for _, e := range slice {
		m[e] = struct{}{}
	}
	return Set[T]{m: m}
}

func (s *Set[T]) Print() {
	fmt.Printf("[ ")
	for k := range s.m {
		fmt.Printf("%v ", k)
	}
	fmt.Printf("]\n")
}

func (s *Set[T]) Add(elements ...T) {
	for _, el := range elements {
		s.m[el] = struct{}{}
	}
}

func (s *Set[T]) Delete(element T) {
	delete(s.m, element)
}

func (s *Set[T]) Slice() []T {
	slice := make([]T, len(s.m))
	i := 0
	for k, _ := range s.m {
		slice[i] = k
		i++
	}
	return slice
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Union(sets ...Set[T]) Set[T] {
	sets = append(sets, *s)
	m := make(map[T]struct{})
	for _, set := range sets {
		for k := range set.m {
			m[k] = struct{}{}
		}
	}
	return Set[T]{m}
}

func (s *Set[T]) Intersection(sets ...Set[T]) Set[T] {
	sets = append(sets, *s)
	m := make(map[T]int)
	for _, set := range sets {
		for k := range set.m {
			m[k]++
		}
	}
	length := len(sets)
	resMap := make(map[T]struct{})
	for k, v := range m {
		if v == length {
			resMap[k] = struct{}{}
		}
	}
	return Set[T]{m: resMap}
}

func (s *Set[T]) Difference(set Set[T]) Set[T] {
	m := make(map[T]int)
	for k := range s.m {
		m[k] = 0
	}
	for k := range set.m {
		m[k] = 0
	}
	for k := range set.m {
		m[k]--
	}
	resMap := make(map[T]struct{})
	for k, v := range m {
		if v == 0 {
			resMap[k] = struct{}{}
		}
	}
	return Set[T]{m: resMap}
}
