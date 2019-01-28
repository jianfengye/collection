package collection

import (
	"testing"
	"reflect"
)

type Foo struct {
	A string
	b int
	C int
}

func TestObjArray(t *testing.T) {
	foo1 := Foo{
		A : "foo1_a",
		b : 1,
		C : 1,
	}
	foo2 := Foo{
		A : "foo2_a",
		b : 2,
		C : 2,
	}
	foos := []Foo{foo1, foo2}
	objArr := NewObjArray(reflect.ValueOf(foos))
	as := objArr.Column("A").ToString()
	if len(as) != 2 {
		t.Fatal("Column len error")
	}
	if as[0] != "foo1_a" {
		t.Fatal("column value error")
	}

	m := objArr.KeyBy("C")
	if m.Len() != 2 {
		t.Fatal("get map error")
	}
	if m.Get(1) == nil {
		t.Fatal("can not get map")
	}

	im := m.Get(1).ToInterface().(Foo)
	if im.A != "foo1_a" {
		t.Fatal("get map error")
	}

	max := objArr.Column("C").Max().ToInt()
	if max != 2 {
		t.Fatal("get max error")
	}

	min := objArr.Column("C").Min().ToInt()
	if min != 1 {
		t.Fatal("get min error")
	}
}
