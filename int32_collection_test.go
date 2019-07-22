package collection

import (
	"errors"
	"reflect"
	"testing"
)

func TestInt32Collection(t *testing.T) {
	arr := NewInt32Collection([]int32{1, 2, 3, 4, 5})

	arr.DD()

	max, err := arr.Max().ToInt32()
	if err != nil {
		t.Error(err)
	}

	if max != 5 {
		t.Error(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(int32)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Error(errors.New("filter error"))
	}

	out, err := arr2.ToInt32s()
	if err != nil || len(out) != 3 {
		t.Error("ToInt32s error")
	}

	json, err := arr2.ToJson()
	if err != nil {
		t.Error(err)
	}

	t.Log(string(json))

}

func TestInt32Collection_Insert(t *testing.T) {
	{
		a := NewInt32Collection([]int32{1,2,3})
		b, err := a.Insert(1, int32(10)).ToInt32s()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int32{1, 10, 2, 3}) {
			t.Error("insert error")
		}
	}
	{
		a := NewInt32Collection([]int32{1,2,3})
		b, err := a.Insert(0, int32(10)).ToInt32s()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int32{10, 1, 2, 3}) {
			t.Error("insert 0 error")
		}
	}

	{
		a := NewInt32Collection([]int32{1,2,3})
		b, err := a.Insert(3, int32(10)).ToInt32s()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int32{1, 2, 3, 10}) {
			t.Error("insert length error")
		}
	}
}