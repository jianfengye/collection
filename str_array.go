package collection

import (
	"fmt"
	"github.com/derekparker/trie"
	"github.com/pkg/errors"
	"strings"
)

type StrArray struct{
	AbsArray
	objs []string

	tri *trie.Trie // 使用trie树增加查找效率，但是仅仅用于Has函数
}

func compareString(a interface{}, b interface{}) int {
	as := a.(string)
	bs := b.(string)
	return strings.Compare(as, bs)
}

func NewStrArray(objs []string) *StrArray {
	tri := trie.New()
	for i, obj := range objs {
		tri.Add(obj, i)
	}
	arr := &StrArray{
		objs:objs,
		tri: tri,
	}
	arr.AbsArray.compare = compareString
	arr.AbsArray.Parent = arr
	return arr
}

func (arr *StrArray) NewEmpty(err ...error) IArray {
	arr2 := NewStrArray(arr.objs)
	if len(err) != 0 {
		arr2.SetErr(err[0])
	}
	return arr2
}

func (arr *StrArray) Insert(index int, item interface{}) IArray {
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


func (arr *StrArray) Remove(i int) IArray {
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


func (arr *StrArray) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *StrArray) Count() int {
	return len(arr.objs)
}

func (arr *StrArray) DD() {
	ret := fmt.Sprintf("StrArray(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%s\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}