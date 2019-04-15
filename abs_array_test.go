package collection

import (
	"reflect"
	"testing"
)

func TestAbsArray_DD(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	intColl.DD()
}

func TestAbsArray_NewEmpty(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	intColl2 := intColl.NewEmpty()
	if intColl2.Count() != 0 {
		t.Error("NewEmpty失败")
	}
	if reflect.TypeOf(intColl2) != reflect.TypeOf(intColl) {
		t.Error("NewEmpty类型失败")
	}
}

func TestAbsArray_Append(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	intColl.Append(3)
	if intColl.Count() != 3 {
		t.Error("Append 失败")
	}
	intColl.DD()
}

func TestAbsArray_Index(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	foo := intColl.Index(1)
	i, err := foo.ToInt()
	if err != nil {
		t.Error("Index 类型错误")
	}
	if i != 2 {
		t.Error("Index 值错误")
	}
}

func TestAbsArray_IsEmpty(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	if intColl.IsEmpty() != false {
		t.Error("IsEmpty 错误")
	}
}

func TestAbsArray_IsNotEmpty(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	if intColl.IsNotEmpty() != true {
		t.Error("IsNotEmpty 错误")
	}
}

func TestAbsArray_Search(t *testing.T) {
	intColl := NewIntArray([]int{1,2})
	if intColl.Search(2) != 1 {
		t.Error("Search 错误")
	}

	intColl = NewIntArray([]int{1,2, 3, 3, 2})
	if intColl.Search(3) != 2 {
		t.Error("Search 重复错误")
	}
}

func TestAbsArray_Unique(t *testing.T) {
	intColl := NewIntArray([]int{1,2, 3, 3, 2})
	uniqColl := intColl.Unique()
	if uniqColl.Count() != 3 {
		t.Error("Unique 重复错误")
	}

	uniqColl.DD()
}

func TestAbsArray_Reject(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5})
	retColl := intColl.Reject(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 3
	})
	if retColl.Count() != 3 {
		t.Error("Reject 重复错误")
	}

	retColl.DD()
}

func TestAbsArray_Last(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 3, 2})
	last, err := intColl.Last().ToInt()
	if err != nil {
		t.Error("last get error")
	}
	if last != 2 {
		t.Error("last 获取错误")
	}

	last, err = intColl.Last(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 2
	}).ToInt()

	if err != nil {
		t.Error("last get error")
	}
	if last != 3 {
		t.Error("last 获取错误")
	}
}

func TestAbsArray_Slice(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5})
	retColl := intColl.Slice(2)
	if retColl.Count() != 3 {
		t.Error("Slice 错误")
	}

	retColl.DD()

	retColl = intColl.Slice(2,2)
	if retColl.Count() != 2 {
		t.Error("Slice 两个参数错误")
	}

	retColl.DD()

	retColl = intColl.Slice(2, -1)
	if retColl.Count() != 3 {
		t.Error("Slice第二个参数为-1错误")
	}

	retColl.DD()
}

func TestAbsArray_Merge(t *testing.T) {
	intColl := NewIntArray([]int{1, 2 })

	intColl2 := NewIntArray([]int{3, 4})

	intColl.Merge(intColl2)

	if intColl.Count() != 4 {
		t.Error("Merge 错误")
	}

	intColl.DD()
}

func TestAbsArray_Combine(t *testing.T) {
	intColl := NewIntArray([]int{1, 2 })

	intColl2 := NewIntArray([]int{3, 4})

	m, err := intColl.Combine(intColl2)
	if err != nil {
		t.Error("Combine错误: " + err.Error())
	}

	m.DD()

	intColl3 := NewIntArray([]int{3, 4, 5})

	m, err = intColl.Combine(intColl3)
	if err == nil {
		t.Error("Combine应该出现个数错误")
	}
}

func TestAbsArray_CrossJoin(t *testing.T) {
	intColl := NewIntArray([]int{1, 2 })

	intColl2 := NewIntArray([]int{3, 4})

	m, err := intColl.CrossJoin(intColl2)
	if err != nil {
		t.Error("CrossJoin错误: " + err.Error())
	}

	if m.Len() != 4 {
		t.Error("CrossJoin错误: " + err.Error())
	}

	m.DD()
}

func TestAbsArray_Each(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4})
	sum := 0
	intColl.Each(func(item interface{}, key int) {
		v := item.(int)
		sum = sum + v
	})
	if sum != 10 {
		t.Error("Each 错误")
	}
}

func TestAbsArray_Map(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4})
	newIntColl := intColl.Map(func(item interface{}, key int) IMix {
		v := item.(int)
		return NewMix(v * 2)
	})
	newIntColl.DD()

	if newIntColl.Count() != 4 {
		t.Error("Map错误")
	}
}

func TestAbsArray_Reduce(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4})
	sumMix := intColl.Reduce(func(carry IMix, item IMix) IMix {
		carryInt, _ := carry.ToInt()
		itemInt, _ := item.ToInt()
		return NewMix(carryInt + itemInt)
	})

	sumMix.DD()

	sum, err := sumMix.ToInt()
	if err != nil {
		t.Error(err.Error())
	}
	if sum != 10 {
		t.Error("Reduce计算错误")
	}
}

func TestAbsArray_Every(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4})
	if intColl.Every(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 1
	}) != false {
		t.Error("Every错误")
	}

	if intColl.Every(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 0
	}) != true {
		t.Error("Every错误")
	}
}

func TestAbsArray_ForPage(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	ret := intColl.ForPage(1, 2)
	ret.DD()

	if ret.Count() != 2 {
		t.Error("For page错误")
	}
}

func TestAbsArray_Nth(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	ret := intColl.Nth(4, 1)
	ret.DD()

	if ret.Count() != 2 {
		t.Error("Nth 错误")
	}
}

func TestAbsArray_Pad(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3})
	ret := intColl.Pad(5, 0)
	if ret.Err() != nil {
		t.Error(ret.Err().Error())
	}

	ret.DD()
	if ret.Count() != 5 {
		t.Error("Pad 错误")
	}

	ret = intColl.Pad(-5, 0)
	if ret.Err() != nil {
		t.Error(ret.Err().Error())
	}
	ret.DD()
	if ret.Count() != 5 {
		t.Error("Pad 错误")
	}
}

func TestAbsArray_Pop(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	pop := intColl.Pop()
	in, err :=  pop.ToInt()
	if err != nil {
		t.Error(err.Error())
	}
	if in != 6 {
		t.Error("Pop 错误")
	}
	intColl.DD()
	if intColl.Count() != 5 {
		t.Error("Pop 后本体错误")
	}
}

func TestAbsArray_Push(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	intColl.Push(7)
	intColl.DD()
	if intColl.Count() != 7 {
		t.Error("Push 后本体错误")
	}
}

func TestAbsArray_Prepend(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	intColl.Prepend(0)
	if intColl.Err() != nil {
		t.Error(intColl.Err().Error())
	}

	intColl.DD()
	if intColl.Count() != 7 {
		t.Error("Prepend错误")
	}
}

func TestAbsArray_Random(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	out := intColl.Random()
	out.DD()

	_, err := out.ToInt()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAbsArray_Reverse(t *testing.T) {
	intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
	vs := intColl.Reverse()
	vs.DD()
}

func TestAppend(t *testing.T) {
	a := []int{1, 2}
	a = append(a, 4)
}