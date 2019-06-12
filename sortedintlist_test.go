package sortedIntList_test

import (
	"math/rand"
	"sort"
	"testing"

	sortedIntList "github.com/arendtio/sortedintlist"
)

func check(t *testing.T, target []int, result sortedIntList.Interface) {
	if len(target) != result.Length() {
		t.Fatal("Length mismatch. Result:", result, "target:", target)
	}
	for i, v := range target {
		if result.At(i) != v {
			t.Fatal("Result:", result, "does not match target:", target)
			break
		}
	}
}

func TestSortedList_Basic(t *testing.T) {
	l := sortedIntList.NewSortedIntList()
	l.Insert(3)
	check(t, []int{3}, l)
	l.Insert(5)
	check(t, []int{3, 5}, l)
	l.Insert(8)
	check(t, []int{3, 5, 8}, l)
	l.Insert(1)
	check(t, []int{1, 3, 5, 8}, l)
	l.Insert(2)
	check(t, []int{1, 2, 3, 5, 8}, l)
	l.Insert(10)
	check(t, []int{1, 2, 3, 5, 8, 10}, l)
	l.Insert(-5)
	check(t, []int{-5, 1, 2, 3, 5, 8, 10}, l)
	l.Insert(0)
	check(t, []int{-5, 0, 1, 2, 3, 5, 8, 10}, l)
	l.Insert(2)
	check(t, []int{-5, 0, 1, 2, 2, 3, 5, 8, 10}, l)
	l.Remove(5)
	check(t, []int{-5, 0, 1, 2, 2, 3, 8, 10}, l)
	l.Remove(2)
	check(t, []int{-5, 0, 1, 2, 3, 8, 10}, l)
	l.Remove(2)
	check(t, []int{-5, 0, 1, 3, 8, 10}, l)
	l.Remove(-5)
	check(t, []int{0, 1, 3, 8, 10}, l)
	l.Remove(3)
	check(t, []int{0, 1, 8, 10}, l)
	l.Remove(10)
	check(t, []int{0, 1, 8}, l)
	l.Remove(1)
	check(t, []int{0, 8}, l)
	l.Remove(0)
	check(t, []int{8}, l)
	l.Remove(8)
	check(t, []int{}, l)
}

func TestSortedList_Random(t *testing.T) {
	source := rand.NewSource(42) // yes, we want the same random numbers every time we run it
	generator := rand.New(source)

	for i := 0; i < 1000; i++ {
		// preparing a random list of numbers
		target := make([]int, 10)
		for k := 0; k < 10; k++ {
			target[k] = generator.Intn(10)
		}

		// inserting the list of numbers into the sortedIntList
		result := sortedIntList.NewSortedIntList()
		for _, v := range target {
			result.Insert(v)
		}

		// sort the random list with the normal Go function
		sort.Sort(sort.IntSlice(target))

		// ckeck if both results are identical
		check(t, target, result)
	}
}
