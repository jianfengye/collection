package collection

import (
	"fmt"
	"github.com/derekparker/trie"
	"github.com/pkg/errors"
	"strings"
)

type StrCollection struct{
	AbsCollection
	objs []string
}

func compareString(a interface{}, b interface{}) int {
	as := a.(string)
	bs := b.(string)
	return strings.Compare(as, bs)
}

func NewStrCollection(objs []string) *StrCollection {
	tri := trie.New()
	for i, obj := range objs {
		tri.Add(obj, i)
	}
	arr := &StrCollection{
		objs:objs,
	}
	arr.AbsCollection.compare = compareString
	arr.AbsCollection.Parent = arr
	return arr
}

func (arr *StrCollection) NewEmpty(err ...error) ICollection {
	arr2 := NewStrCollection(arr.objs)
	if len(err) != 0 {
		arr2.SetErr(err[0])
	}
	return arr2
}

func (arr *StrCollection) Insert(index int, item interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := item.(string); ok {
		length := len(arr.objs)
		tail := arr.objs[index:length]
		arr.objs = append(arr.objs[0:index], i)
		arr.objs = append(arr.objs, tail...)
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}


func (arr *StrCollection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1: len]...)
	return arr
}


func (arr *StrCollection) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *StrCollection) Count() int {
	return len(arr.objs)
}

func (arr *StrCollection) DD() {
	ret := fmt.Sprintf("StrCollection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%s\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}