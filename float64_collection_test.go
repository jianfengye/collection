package collection

import (
	"github.com/pkg/errors"
	"testing"
)

func TestFloat64Collection(t *testing.T) {
	arr := NewFloat64Collection([]float64{1.0 ,2.0,3.0,4.0,5.0})

	arr.DD()

	max, err := arr.Max().ToFloat64()
	if err != nil {
		t.Error(err)
	}

	if max != 5 {
		t.Error(errors.New("max error"))
	}


	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(float64)
		if val > 2.0 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Error(errors.New("filter error"))
	}

	out, err := arr2.ToFloat64s()
	if err != nil || len(out) != 3 {
		t.Error(errors.New("to float64s error"))
	}
}
