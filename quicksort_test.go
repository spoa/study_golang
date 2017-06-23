package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	value := rand.Perm(10)
	result := quicksort(value)

	if !reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error("結果が正しくありません")
		t.Log("value : ", value)
		t.Log("result : ", result)
	}
}

func TestHasSameValue(t *testing.T) {
	value := Suffle([]int{9, 8, 5, 7, 6, 5, 4, 3, 5, 2, 1, 0})
	result := quicksort(value)

	if !reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9}) {
		t.Error("結果が正しくありません")
		t.Log("value : ", value)
		t.Log("result : ", result)
	}
}

func Suffle(a []int) (result []int) {
	for _, i := range rand.Perm(len(a)) {
		result = append(result, a[i])
	}
	return result
}
