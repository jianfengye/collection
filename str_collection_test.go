package collection

import (
	"reflect"
	"testing"
)

func TestStrCollection_Insert(t *testing.T) {
	{
		a := NewStrCollection([]string{"1", "2", "3"})
		b, err := a.Insert(1, "10").ToStrings()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []string{"1", "10", "2", "3"}) {
			t.Fatal("insert error")
		}
	}
	{
		a := NewStrCollection([]string{"1", "2", "3"})
		b, err := a.Insert(0, "10").ToStrings()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []string{"10", "1", "2", "3"}) {
			t.Fatal("insert 0 error")
		}
	}

	{
		a := NewStrCollection([]string{"1", "2", "3"})
		b, err := a.Insert(3, "10").ToStrings()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(b, []string{"1", "2", "3", "10"}) {
			t.Fatal("insert length error")
		}
	}
}

func TestStrCollection_FromJson(t *testing.T) {
	data := `["aa", "bb"]`
	objColl := NewStrCollection([]string{})
	err := objColl.FromJson([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	objColl.DD()
}

func TestStrCollection_Remove(t *testing.T) {
	strColl := NewStrCollection([]string{"1", "2", "3"})
	r := strColl.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
}

func TestStrCollection_Diff(t *testing.T) {
	oldColl := NewStrCollection([]string{"1", "2"})
	newColl := NewStrCollection([]string{"2", "3"})
	o := oldColl.Diff(newColl)

	n, err := o.ToStrings()
	if err != nil {
		t.Fatal(err)
	}

	if len(n) != 1 {
		t.Fatal("diff error ")
	}
}

func TestUnion(t *testing.T) {
	oldColl := NewStrCollection([]string{"test1", "test2", "test3", "test4"})
	newColl := NewStrCollection([]string{"test3", "test4", "test5", "test6"})
	o := oldColl.Union(newColl)

	ret, err := o.ToStrings()
	if err != nil {
		t.Fatal(err)
	}

	if len(ret) != 6 {
		t.Fatal("union len error ")
	}

	if ret[5] != "test6" {
		t.Fatal("union val error")
	}
}

func TestIntersect(t *testing.T) {
	oldColl := NewStrCollection([]string{"test1", "test2", "test3", "test4"})
	newColl := NewStrCollection([]string{"test3", "test4", "test5", "test6"})
	o := oldColl.Intersect(newColl)
	o.DD()

	ret, err := o.ToStrings()
	if err != nil {
		t.Fatal(err)
	}

	if len(ret) != 2 {
		t.Fatal("intersect len error ")
	}

	if ret[1] != "test4" {
		t.Fatal("intersect val error")
	}
}
