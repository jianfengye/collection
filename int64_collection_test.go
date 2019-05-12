package collection

import (
	"errors"
	"testing"
)

func TestInt64Collection(t *testing.T) {
	arr := NewInt64Collection([]int64{1,2,3,4,5})

	arr.DD()

	max, err := arr.Max().ToInt64()
	if err != nil {
		t.Error(err)
	}

	if max != 5 {
		t.Error(errors.New("max error"))
	}


	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(int64)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Error(errors.New("filter error"))
	}

	out, err := arr2.ToInt64s()
	if err != nil || len(out) != 3{
		t.Error("ToInt64s error")
	}
}
