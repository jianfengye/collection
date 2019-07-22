package collection

import (
	"reflect"
	"testing"
)

func TestIntCollection_Insert(t *testing.T) {
	{
		a := NewIntCollection([]int{1,2,3})
		b, err := a.Insert(1, 10).ToInts()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int{1, 10, 2, 3}) {
			t.Error("insert error")
		}
	}
	{
		a := NewIntCollection([]int{1,2,3})
		b, err := a.Insert(0, 10).ToInts()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int{10, 1, 2, 3}) {
			t.Error("insert 0 error")
		}
	}

	{
		a := NewIntCollection([]int{1,2,3})
		b, err := a.Insert(3, 10).ToInts()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []int{1, 2, 3, 10}) {
			t.Error("insert length error")
		}
	}
}
