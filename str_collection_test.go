package collection

import (
	"reflect"
	"testing"
)

func TestStrCollection_Insert(t *testing.T) {
	{
		a := NewStrCollection([]string{"1","2","3"})
		b, err := a.Insert(1, "10").ToStrings()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []string{"1", "10", "2", "3"}) {
			t.Error("insert error")
		}
	}
	{
		a := NewStrCollection([]string{"1","2","3"})
		b, err := a.Insert(0, "10").ToStrings()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []string{"10", "1", "2", "3"}) {
			t.Error("insert 0 error")
		}
	}

	{
		a := NewStrCollection([]string{"1","2","3"})
		b, err := a.Insert(3, "10").ToStrings()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(b, []string{"1", "2", "3", "10"}) {
			t.Error("insert length error")
		}
	}
}

func TestStrCollection_FromJson(t *testing.T) {
	data := `["aa", "bb"]`
	objColl := NewStrCollection([]string{})
	err := objColl.FromJson([]byte(data))
	if err != nil {
		t.Error(err)
	}
	objColl.DD()
}

func TestStrCollection_Remove(t *testing.T) {
	strColl := NewStrCollection([]string{"1", "2", "3"})
	r := strColl.Remove(0)
	if r.Err() != nil{
		t.Fatal(r.Err())
	}
}