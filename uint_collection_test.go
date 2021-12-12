package collection

import (
	"reflect"
	"testing"
)

func TestUIntCollection_Insert(t *testing.T) {
	{
		a := NewUIntCollection([]uint{1, 2, 3})
		b, err := a.Insert(1, uint(10)).ToUInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewUIntCollection([]uint{1, 2, 3})
		b, err := a.Insert(0, uint(10)).ToUInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewUIntCollection([]uint{1, 2, 3})
		b, err := a.Insert(3, uint(10)).ToUInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []uint{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestUIntCollection_Filter(t *testing.T) {
	intColl := NewUIntCollection([]uint{1, 2, 3})
	a, err := intColl.First().ToUInt()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, uint(1)) {
		t.Fatal("filter error")
	}
}

func TestUIntCollection_Index(t *testing.T) {
	intColl := NewUIntCollection([]uint{1, 2, 3})
	iMix := intColl.Index(2)
	if iMix.Err() != nil {
		t.Fatal(iMix.Err())
	}

	i, err := iMix.ToUInt()
	if err != nil {
		t.Fatal("index error")
	}

	if i != 3 {
		t.Fatal("not equal")
	}
}
