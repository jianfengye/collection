package collection

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestUInt32Collection(t *testing.T) {
	arr := NewUInt32Collection([]uint32{1, 2, 3, 4, 5})

	arr.DD()

	max, err := arr.Max().ToUInt32()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(uint32)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToUInt32s()
	if err != nil || len(out) != 3 {
		t.Fatal("ToUInt32s error")
	}

	json, err := arr2.ToJson()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(json))

}

func TestUInt32Collection_Insert(t *testing.T) {
	{
		a := NewUInt32Collection([]uint32{1, 2, 3})
		b, err := a.Insert(1, uint32(10)).ToUInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint32{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewUInt32Collection([]uint32{1, 2, 3})
		b, err := a.Insert(0, uint32(10)).ToUInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint32{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewUInt32Collection([]uint32{1, 2, 3})
		b, err := a.Insert(3, uint32(10)).ToUInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint32{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestUInt32Collection_Remove(t *testing.T) {
	uint32Coll := NewUInt32Collection([]uint32{1, 2, 3})
	r := uint32Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}

func TestUInt32Collection_GroupBy(t *testing.T) {
	objColl := NewUInt32Collection([]uint32{1, 1, 20, 4})
	groupBy := objColl.GroupBy(func(item interface{}, i2 int) interface{} {
		foo := item.(uint32)
		return foo
	})
	for k, collection := range groupBy {
		t.Log(k)
		collection.DD()
	}
}
