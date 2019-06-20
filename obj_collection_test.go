package collection

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

type Foo struct {
	A string
	B int
}


func TestObjCollection_Pluck(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})

	strColl := objColl.Pluck("A")

	strColl.DD()

	str, err := strColl.Index(0).ToString()
	if err != nil {
		t.Error(err)
	}

	if str != "a1" {
		t.Error(errors.New("Pluck error"))
	}
}

func TestObjCollection_SortBy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.SortBy("B")

	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Error(err)
	}

	foo := obj.(Foo)
	if foo.B != 2 {
		t.Error("SortBy error")
	}
}
func TestObjCollection_Sort(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})
	compare := func(m interface{}, n interface{}) int {
		m1, _ := m.(Foo)
		n1, _ := n.(Foo)
		return m1.B - n1.B
	}
	objColl.SetCompare(compare)

	newObjColl := objColl.Sort()
	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Error(err)
	}

	foo := obj.(Foo)
	if foo.B != 2 {
		t.Error("Sort error")
	}
}
func TestObjCollection_Copy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.Copy()

	newObjColl.DD()

	if newObjColl.Count() != 3 {
		t.Error("Copy count error")
	}
	inewA1, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Error("Copy get first element error" + err.Error())
	}
	newA1 := inewA1.(Foo)
	if newA1.B != 3 {
		t.Error("Copy get first element error")
	}
}

func TestObjCollection_SortByDesc(t *testing.T) {
	a1 := Foo{A: "a1", B: 2}
	a2 := Foo{A: "a2", B: 3}
	a3 := Foo{A: "a3", B: 1}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.SortByDesc("B")

	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Error(err)
	}

	foo := obj.(Foo)
	if foo.B != 3 {
		t.Error("SortBy error")
	}
}

func TestObjCollection(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})
	objColl.DD()

	if objColl.IsNotEmpty() != true {
		t.Error("Is Not Empty error")
	}

	if objColl.Count() != 2 {
		t.Error("Count error")
	}

	a3 := Foo{A: "a3"}
	a4 := Foo{A: "a4"}
	objColl.Append(a3).Append(a4)
	if objColl.Count() != 4 {
		t.Error("Append Error")
	}

	objColl.SetCompare(func(a interface{}, b interface{}) int {
		aObj := a.(Foo)
		bObj := b.(Foo)
		if aObj.A > bObj.A {
			return 1
		}
		if aObj.A == bObj.A {
			return 0
		}
		if aObj.A < bObj.A {
			return -1
		}
		return 0
	})

	objColl.DD()
	if objColl.Search(Foo{A: "a3"}) != 2 {
		t.Error("Search error")
	}

	objColl2 := objColl.Filter(func(obj interface{}, index int) bool {
		foo := obj.(Foo)
		if foo.A == "a3" {
			return true
		}
		return false
	})
	if objColl2.Count() != 1 {
		t.Error("Filter Error")
	}

	obj, _ := objColl.Last().ToInterface()
	if foo, ok := obj.(Foo); !ok || foo.A != "a4" {
		t.Error("Last error")
	}

	ret, err := objColl.Map(func(item interface{}, key int) interface{} {
		foo := item.(Foo)
		return foo.A
	}).Reduce(func(carry IMix, item IMix) IMix {
		ret, _ := carry.ToString()
		join, _ := item.ToString()
		return NewMix(ret + join)
	}).ToString()
	if err != nil {
		t.Error("Map error")
	}
	if ret != "a1a2a3a4" {
		t.Error("Reduce error")
	}

	objColl.ForPage(1, 2).DD()

	aColl := objColl.Pluck("A")
	aColl.DD()

	a0 := Foo{A: "a0"}
	objColl.Append(a0)

	objColl.Sort().DD()
	o, err := objColl.Index(0).ToInterface()
	if err != nil {
		t.Error(err)
	}
	fooOut := o.(Foo)
	if fooOut.A != "a0" {
		t.Error("sort result error")
	}

	objColl.DD()
	objColl.SetCompare(func(a interface{}, b interface{}) int {
		aFoo := a.(Foo)
		bFoo := b.(Foo)
		return strings.Compare(aFoo.A, bFoo.A)
	})

	objColl3 := objColl.SortBy("A")
	objColl3.DD()

	objColl3.SortByDesc("A")
	objColl3.DD()

	objColl3.Remove(2)
	objColl3.DD()
}

func TestObjCollection_ToJson(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})
	byt, err := objColl.ToJson()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(byt))

	out := NewEmptyMixCollection()
	for i := 0; i < objColl.Count(); i++ {
		if i == 1 {
			out.Append(objColl.Index(i).SetField("d", 12).RemoveFields("B"))
			continue
		}
		out.Append(objColl.Index(i).SetField("c", "test").RemoveFields("A"))
	}
	byt, err = out.ToJson()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(byt))
}

func TestObjCollection_FromJson(t *testing.T) {
	data := `[{"A":"a1","B":1},{"A":"a2","B":2}]`
	dataByte := []byte(data)
	objColl := NewObjCollection([]Foo{})
	err := objColl.FromJson(dataByte)
	if err != nil {
		t.Error(err)
	}
	objColl.DD()
}

func TestObjCollection_Marshal(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})
	byt, err := json.Marshal(objColl)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(byt))

	out := NewEmptyMixCollection()
	for i := 0; i < objColl.Count(); i++ {
		if i == 1 {
			out.Append(objColl.Index(i).SetField("d", 12).RemoveFields("B"))
			continue
		}
		out.Append(objColl.Index(i).SetField("c", "test").RemoveFields("A"))
	}
	byt, err = json.Marshal(out)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(byt))
}