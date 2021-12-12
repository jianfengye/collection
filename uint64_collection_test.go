package collection

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestUInt64Collection(t *testing.T) {
	arr := NewUInt64Collection([]uint64{1, 2, 3, 4, 5})

	arr.DD()

	max, err := arr.Max().ToUInt64()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(uint64)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToUInt64s()
	if err != nil || len(out) != 3 {
		t.Fatal("ToUInt64s error")
	}
}

func TestUInt64Collection_Insert(t *testing.T) {
	{
		a := NewUInt64Collection([]uint64{1, 2, 3})
		b, err := a.Insert(1, uint64(10)).ToUInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint64{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewUInt64Collection([]uint64{1, 2, 3})
		b, err := a.Insert(0, uint64(10)).ToUInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint64{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewUInt64Collection([]uint64{1, 2, 3})
		b, err := a.Insert(3, uint64(10)).ToUInt64s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint64{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestUInt64Collection_Remove(t *testing.T) {
	uint64Coll := NewUInt64Collection([]uint64{1, 2, 3})
	r := uint64Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}
