package collection

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type StrCollection struct {
	AbsCollection
	objs []string
}

func compareString(a interface{}, b interface{}) int {
	as := a.(string)
	bs := b.(string)
	return strings.Compare(as, bs)
}

func NewStrCollection(objs []string) *StrCollection {
	arr := &StrCollection{
		objs: objs,
	}
	arr.AbsCollection.compare = compareString
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = TYPE_STRING
	return arr
}

// Copy copy collection
func (arr *StrCollection) Copy() ICollection {
	return NewStrCollection(arr.objs)
}

func (arr *StrCollection) NewEmpty(err ...error) ICollection {
	return NewStrCollection([]string{})
}

func (arr *StrCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(string); ok {
		length := len(arr.objs)

		// 如果是append操作，直接调用系统的append，不新创建collection
		if index >= length {
			arr.objs = append(arr.objs, i)
			return arr
		}

		arr.objs = append(arr.objs, "0")
		copy(arr.objs[index+1:], arr.objs[index:])
		arr.objs[index] = i
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
	if i < 0 || i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1:len]...)
	return arr
}

func (arr *StrCollection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}

func (arr *StrCollection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(string)
	return arr
}

func (arr *StrCollection) Count() int {
	return len(arr.objs)
}

func (arr *StrCollection) DD() {
	ret := fmt.Sprintf("StrCollection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%s\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *StrCollection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *StrCollection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
