package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		
		got := Sum(numbers)
		want := 15
		
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		
		got := Sum(numbers)
		want := 6
		
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of two slices", func(t *testing.T) {		
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("sum of four slices", func(t *testing.T) {		
		got := SumAll([]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8})
		want := []int{3, 7, 11, 15}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("sum tails of collections", func(t *testing.T) {		
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}
		
		checkSums(t, got, want)
	})

	t.Run("sum tails of empty collection", func(t *testing.T) {		
		got := SumAllTails([]int{}, []int{3, 4})
		want := []int{0, 4}
		
		checkSums(t, got, want)
	})
}
