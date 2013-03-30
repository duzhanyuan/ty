package fun

import (
	"testing"
)

func TestMap(t *testing.T) {
	square := func(x int) int { return x * x }
	squares := Map(square, []int{1, 2, 3, 4, 5}).([]int)

	assertDeep(t, squares, []int{1, 4, 9, 16, 25})
	assertDeep(t, []int{}, Map(square, []int{}).([]int))
}

func TestFilter(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	evens := Filter(even, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).([]int)

	assertDeep(t, evens, []int{2, 4, 6, 8, 10})
	assertDeep(t, []int{}, Filter(even, []int{}).([]int))
}

func TestFoldl(t *testing.T) {
	// Use an operation that isn't associative so that we know we've got
	// the left/right folds done correctly.
	reducer := func(a, b int) int { return b % a }
	v := Foldl(reducer, 7, []int{4, 5, 6}).(int)

	assertDeep(t, v, 3)
	assertDeep(t, 0, Foldl(reducer, 0, []int{}).(int))
}

func TestFoldr(t *testing.T) {
	// Use an operation that isn't associative so that we know we've got
	// the left/right folds done correctly.
	reducer := func(a, b int) int { return b % a }
	v := Foldr(reducer, 7, []int{4, 5, 6}).(int)

	assertDeep(t, v, 1)
	assertDeep(t, 0, Foldr(reducer, 0, []int{}).(int))
}

func TestConcat(t *testing.T) {
	toflat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	flat := Concat(toflat).([]int)

	assertDeep(t, flat, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestPointers(t *testing.T) {
	type temp struct {
		val int
	}
	square := func(t *temp) *temp { return &temp{t.val * t.val} }
	squares := Map(square, []*temp{
		{1}, {2}, {3}, {4}, {5},
	})

	assertDeep(t, squares, []*temp{
		{1}, {4}, {9}, {16}, {25},
	})
}

func BenchmarkMap(b *testing.B) {
	if flagBuiltin {
		benchmarkMapBuiltin(b)
	} else {
		benchmarkMapReflect(b)
	}
}

func benchmarkMapReflect(b *testing.B) {
	b.StopTimer()
	list := randIntSlice(1000)
	square := func(a int) int {
		return a * a
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Map(square, list).([]int)
	}
}

func benchmarkMapBuiltin(b *testing.B) {
	b.StopTimer()
	list := randIntSlice(1000)
	square := func(a int) int {
		return a * a
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ret := make([]int, len(list))
		for i := 0; i < len(list); i++ {
			ret[i] = square(list[i])
		}
	}
}
