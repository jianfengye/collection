package collection

import (
	"testing"

	"github.com/pkg/errors"
)

func TestFloat32Collection(t *testing.T) {
	arr := NewFloat32Collection([]float32{1.0, 2.0, 3.0, 4.0, 5.0})

	max, err := arr.Max().ToFloat32()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(float32)
		if val > 2.0 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToFloat32s()
	if err != nil || len(out) != 3 {
		t.Fatal(errors.New("to float32s error"))
	}

	byt, err := arr2.ToJson()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(byt))
}

func TestFloat32Collection_Remove(t *testing.T) {
	float32Coll := NewFloat32Collection([]float32{1.0, 2.0, 3.0})
	r := float32Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}
