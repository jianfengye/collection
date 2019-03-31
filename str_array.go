package collection

import (
	"github.com/derekparker/trie"
	"strings"
)

type StrArray struct{
	AbsArray
	objs []string
	tri *trie.Trie // 使用trie树增加查找效率，但是仅仅用于Has函数

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

	arr.AbsArray.Parent = arr
	return arr
}
func (arr *StrArray) mustBeString(obj interface{}) string {
	if i, ok := obj.(string); ok {
		return i
	} else {
		panic("obj must be int")
	}
}


func (arr *StrArray) Append(obj interface{}) {
	if str, ok := obj.(string); ok {
		arr.objs = append(arr.objs, str)
		arr.tri.Add(str, 1)
	} else {
		panic("can not append none string to StrArray")
	}
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *StrArray) Search(obj interface{}) int {
	ob := arr.mustBeString(obj)
	for i, o := range arr.objs {
		if strings.Compare(o, ob) == 0 {
			return i
		}
	}
	return -1
}

func (arr *StrArray) ToString() []string {
	return arr.objs
}

func (arr *StrArray) Unique() IArray {
	objs := arr.ToString()
	ret := make([]string, 0 ,len(objs))
	strArr := NewStrArray(ret)

	for _, s := range objs {
		if strArr.Search(s) < 0 {
			strArr.Append(s)
		}
	}

	return strArr
}


func (arr *StrArray) Len() int {
	return len(arr.objs)
}

func (arr *StrArray) Has(obj interface{}) bool {
	ob := arr.mustBeString(obj)
	_, isExist := arr.tri.Find(ob)
	return isExist
}

func (arr *StrArray) NewEmptyIArray() IArray {
	return NewStrArray([]string{})
}