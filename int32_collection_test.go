package collection

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestInt32Collection(t *testing.T) {
	arr := NewInt32Collection([]int32{1, 2, 3, 4, 5})

	max, err := arr.Max().ToInt32()
	if err != nil {
		t.Fatal(err)
	}

	if max != 5 {
		t.Fatal(errors.New("max error"))
	}

	arr2 := arr.Filter(func(obj interface{}, index int) bool {
		val := obj.(int32)
		if val > 2 {
			return true
		}
		return false
	})
	if arr2.Count() != 3 {
		t.Fatal(errors.New("filter error"))
	}

	out, err := arr2.ToInt32s()
	if err != nil || len(out) != 3 {
		t.Fatal("ToInt32s error")
	}

	json, err := arr2.ToJson()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(json))

}

func TestInt32Collection_Insert(t *testing.T) {
	{
		a := NewInt32Collection([]int32{1, 2, 3})
		b, err := a.Insert(1, int32(10)).ToInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int32{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewInt32Collection([]int32{1, 2, 3})
		b, err := a.Insert(0, int32(10)).ToInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int32{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewInt32Collection([]int32{1, 2, 3})
		b, err := a.Insert(3, int32(10)).ToInt32s()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int32{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestInt32Collection_Remove(t *testing.T) {
	int32Coll := NewInt32Collection([]int32{1, 2, 3})
	r := int32Coll.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}

func TestInt32Collection_GroupBy(t *testing.T) {
	objColl := NewInt32Collection([]int32{1, 1, 20, 4})
	groupBy := objColl.GroupBy(func(item interface{}, i2 int) interface{} {
		foo := item.(int32)
		return foo
	})

	for k, collection := range groupBy {
		t.Log(k)
		collection.DD()
	}
}
