package collection

import (
	"reflect"
	"testing"
)

func TestIntCollection_Insert(t *testing.T) {
	{
		a := NewIntCollection([]int{1, 2, 3})
		b, err := a.Insert(1, 10).ToInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int{1, 10, 2, 3}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewIntCollection([]int{1, 2, 3})
		b, err := a.Insert(0, 10).ToInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int{10, 1, 2, 3}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewIntCollection([]int{1, 2, 3})
		b, err := a.Insert(3, 10).ToInts()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []int{1, 2, 3, 10}) {
			t.Fatal("insert length error")
		}
	}
}

func TestIntCollection_Filter(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3})
	a, err := intColl.First().ToInt()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, 1) {
		t.Fatal("filter error")
	}
}

func TestIntCollection_Index(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3})
	iMix := intColl.Index(2)
	if iMix.Err() != nil {
		t.Fatal(iMix.Err())
	}

	i, err := iMix.ToInt()
	if err != nil {
		t.Fatal("index error")
	}

	if i != 3 {
		t.Fatal("not equal")
	}
}

func TestIntCollection_Remove(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3})
	r := intColl.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}

func TestIntCollection_Split(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6, 7, 8})
	ret := intColl.Split(3)

	if len(ret) != 3 {
		t.Fatal("split len not right")
	}

	// ret[0].DD()
	// ret[1].DD()
	// ret[2].DD()

	if ret[0].Count() != 3 || ret[2].Count() != 2 {
		t.Fatal("split not right")
	}

	int2Coll := NewIntCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	ret2 := int2Coll.Split(3)
	if len(ret2) != 3 {
		t.Fatal("split not right")
	}

	if ret2[2].Count() != 3 {
		t.Fatal("split not right")
	}

	// ret2[0].DD()
	// ret2[1].DD()
	// ret2[2].DD()
}
