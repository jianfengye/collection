package collection

import (
	"testing"

	"github.com/pkg/errors"
)

func TestFloat64Collection(t *testing.T) {
	arr := NewFloat64Collection([]float64{1.0, 2.0, 3.0, 4.0, 5.0})

	arr.DD()

	max, err := arr.Max().ToFloat64()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(float64)
		if val > 2.0 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToFloat64s()
	if err != nil || len(out) != 3 {
		t.Fatal(errors.New("to float64s error"))
	}
}

func TestFloat64Collection_Remove(t *testing.T) {
	float64Coll := NewFloat64Collection([]float64{1.0, 2.0, 3.0})
	r := float64Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}
