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

func TestIntCollection_Filter(t *testing.T) {
	intColl := NewIntCollection([]int{1,2,3})
	a, err := intColl.First().ToInt()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(a, 1) {
		t.Error("filter error")
	}
}

func TestIntCollection_Index(t *testing.T) {
	intColl := NewIntCollection([]int{1,2,3})
	iMix := intColl.Index(2)
	if iMix.Err() != nil {
		t.Fatal(iMix.Err())
	}

	i,err := iMix.ToInt()
	if err != nil {
		t.Fatal("index error")
	}

	if i != 3 {
		t.Fatal("not equal")
	}
}

func TestIntCollection_Remove(t *testing.T) {
	intColl := NewIntCollection([]int{1,2,3})
	r := intColl.Remove(0)
	if r.Err() != nil{
		t.Fatal(r.Err())
	}
}