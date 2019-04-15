package collection

import (
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

func (arr *StrArray) NewEmpty() IArray {
	return NewStrArray(arr.objs)
}

func (arr *StrArray) mustBeString(obj interface{}) string {
	if i, ok := obj.(string); ok {
		return i
	} else {
		panic("obj must be int")
	}
}

func (arr *StrArray) Append(obj interface{}) (IArray,error) {
	if str, ok := obj.(string); ok {
		arr.objs = append(arr.objs, str)
		arr.tri.Add(str, 1)
		return arr, nil
	}
	return arr, errors.New("can not append none string to StrArray")
}

func (arr *StrArray) ToString() ([]string, error) {
	return arr.objs, nil
}

func (arr *StrArray) Len() int {
	return len(arr.objs)
}

func (arr *StrArray) Has(obj interface{}) bool {
	ob := arr.mustBeString(obj)
	_, isExist := arr.tri.Find(ob)
	return isExist
}
