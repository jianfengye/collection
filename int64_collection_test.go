package collection

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestInt64Collection(t *testing.T) {
	arr := NewInt64Collection([]int64{1, 2, 3, 4, 5})

	arr.DD()

	max, err := arr.Max().ToInt64()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(int64)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToInt64s()
	if err != nil || len(out) != 3 {
		t.Fatal("ToInt64s error")
	}
}

func TestInt64Collection_Insert(t *testing.T) {
	{
		a := NewInt64Collection([]int64{1, 2, 3})
		b, err := a.Insert(1, int64(10)).ToInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int64{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewInt64Collection([]int64{1, 2, 3})
		b, err := a.Insert(0, int64(10)).ToInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int64{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewInt64Collection([]int64{1, 2, 3})
		b, err := a.Insert(3, int64(10)).ToInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int64{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestInt64Collection_Remove(t *testing.T) {
	int64Coll := NewInt64Collection([]int64{1, 2, 3})
	r := int64Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}
