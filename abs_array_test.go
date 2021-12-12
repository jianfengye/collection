package collection

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestAbsCollection_DD(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl.DD()
}

func TestAbsCollection_Copy(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl2 := intColl.Copy()
	if intColl2.Count() != 2 {
		t.Fatal("Copy失败")
	}
	if reflect.TypeOf(intColl2) != reflect.TypeOf(intColl) {
		t.Fatal("Copy类型失败")
	}
}
func TestAbsCollection_NewEmpty(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl2 := intColl.NewEmpty()
	if intColl2.Count() != 0 {
		t.Fatal("NewEmpty失败")
	}
	if reflect.TypeOf(intColl2) != reflect.TypeOf(intColl) {
		t.Fatal("NewEmpty类型失败")
	}
}

func TestAbsCollection_Append(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl.Append(3)
	if intColl.Err() != nil {
		t.Fatal("error")
	}

	if intColl.Count() != 3 {
		t.Fatal("Append 失败")
	}
}

func TestAbsCollection_Index(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	foo := intColl.Index(1)
	i, err := foo.ToInt()
	if err != nil {
		t.Fatal("Index 类型错误")
	}
	if i != 2 {
		t.Fatal("Index 值错误")
	}
}

func TestAbsCollection_IsEmpty(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	// println(intColl.IsEmpty())
	if intColl.IsEmpty() != false {
		t.Fatal("IsEmpty 错误")
	}
}

func TestAbsCollection_IsNotEmpty(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	// println(intColl.IsNotEmpty())
	if intColl.IsNotEmpty() != true {
		t.Fatal("IsNotEmpty 错误")
	}
}

func TestAbsCollection_Search(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	if intColl.Search(2) != 1 {
		t.Fatal("Search 错误")
	}

	intColl = NewIntCollection([]int{1, 2, 3, 3, 2})
	if intColl.Search(3) != 2 {
		t.Fatal("Search 重复错误")
	}
}

func TestAbsCollection_Unique(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 3, 2})
	uniqColl := intColl.Unique()
	if uniqColl.Count() != 3 {
		t.Fatal("Unique 重复错误")
	}

}

func TestAbsCollection_Reject(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
	retColl := intColl.Reject(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 3
	})
	if retColl.Count() != 3 {
		t.Fatal("Reject 重复错误")
	}

}

func TestAbsCollection_Last(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 3, 2})
	last, err := intColl.Last().ToInt()
	if err != nil {
		t.Fatal("last get error")
	}
	if last != 2 {
		t.Fatal("last 获取错误")
	}

	last, err = intColl.Last(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 2
	}).ToInt()

	if err != nil {
		t.Fatal("last get error")
	}
	if last != 3 {
		t.Fatal("last 获取错误")
	}
}

func TestAbsCollection_Slice(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
	retColl := intColl.Slice(2)
	if retColl.Count() != 3 {
		t.Fatal("Slice 错误")
	}

	retColl = intColl.Slice(2, 2)
	if retColl.Count() != 2 {
		t.Fatal("Slice 两个参数错误")
	}

	// retColl.DD()

	retColl = intColl.Slice(2, -1)
	if retColl.Count() != 3 {
		t.Fatal("Slice第二个参数为-1错误")
	}

	// retColl.DD()
}

func TestAbsCollection_Merge(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})

	intColl2 := NewIntCollection([]int{3, 4})

	intColl.Merge(intColl2)

	if intColl.Err() != nil {
		t.Fatal(intColl.Err())
	}

	if intColl.Count() != 4 {
		t.Fatal("Merge 错误")
	}

}

func TestAbsCollection_Each(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4})
	sum := 0
	intColl.Each(func(item interface{}, key int) {
		v := item.(int)
		sum = sum + v
	})

	if intColl.Err() != nil {
		t.Fatal(intColl.Err())
	}

	if sum != 10 {
		t.Fatal("Each 错误")
	}

	sum = 0
	intColl.Each(func(item interface{}, key int) {
		v := item.(int)
		sum = sum + v
		if sum > 4 {
			intColl.SetErr(errors.New("stop the cycle"))
			return
		}
	})

	if sum != 6 {
		t.Fatal("Each 错误")
	}
}

func TestAbsCollection_Map(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4})
	newIntColl := intColl.Map(func(item interface{}, key int) interface{} {
		v := item.(int)
		return v * 2
	})

	if newIntColl.Count() != 4 {
		t.Fatal("Map错误")
	}

	newIntColl2 := intColl.Map(func(item interface{}, key int) interface{} {
		v := item.(int)

		if key > 2 {
			intColl.SetErr(errors.New("break"))
			return nil
		}

		return v * 2
	})
	_, err := newIntColl2.ToInts()
	if err == nil {
		t.Fatal("error should not be empty")
	}

	intColl.SetErr(nil)
	newIntColl3 := intColl.Map(func(item interface{}, key int) interface{} {
		v := item.(int)

		if key == 2 {
			return nil
		}

		return v * 2
	})
	out3, err := newIntColl3.ToInts()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(out3, []int{2, 4, 8}) {
		t.Fatal("continue error")
	}
}

func TestAbsCollection_Reduce(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4})
	sumMix := intColl.Reduce(func(carry IMix, item IMix) IMix {
		carryInt, _ := carry.ToInt()
		itemInt, _ := item.ToInt()
		return NewMix(carryInt + itemInt)
	})

	// sumMix.DD()

	sum, err := sumMix.ToInt()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sum != 10 {
		t.Fatal("Reduce计算错误")
	}
}

func TestAbsCollection_Every(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4})
	if intColl.Every(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 1
	}) != false {
		t.Fatal("Every错误")
	}

	if intColl.Every(func(item interface{}, key int) bool {
		i := item.(int)
		return i > 0
	}) != true {
		t.Fatal("Every错误")
	}
}

func TestAbsCollection_ForPage(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	ret := intColl.ForPage(1, 2)

	if ret.Count() != 2 {
		t.Fatal("For page错误")
	}
}

func TestAbsCollection_Nth(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	ret := intColl.Nth(4, 1)

	if ret.Count() != 2 {
		t.Fatal("Nth 错误")
	}
}

func TestAbsCollection_Pad(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3})
	ret := intColl.Pad(5, 0)
	if ret.Err() != nil {
		t.Fatal(ret.Err().Error())
	}

	if ret.Count() != 5 {
		t.Fatal("Pad 错误")
	}

	ret = intColl.Pad(-5, 0)
	if ret.Err() != nil {
		t.Fatal(ret.Err().Error())
	}
	if ret.Count() != 5 {
		t.Fatal("Pad 错误")
	}
}

func TestAbsCollection_Pop(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	pop := intColl.Pop()
	in, err := pop.ToInt()
	if err != nil {
		t.Fatal(err.Error())
	}
	if in != 6 {
		t.Fatal("Pop 错误")
	}
	if intColl.Count() != 5 {
		t.Fatal("Pop 后本体错误")
	}
}

func TestAbsCollection_Push(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	intColl.Push(7)
	if intColl.Count() != 7 {
		t.Fatal("Push 后本体错误")
	}
}

func TestAbsCollection_Prepend(t *testing.T) {
	old := []int{1, 2, 3, 4, 5, 6}
	intColl := NewIntCollection(old)
	intColl.Prepend(4)
	if intColl.Err() != nil {
		t.Fatal(intColl.Err().Error())
	}

	if intColl.Count() != 7 {
		t.Fatal("Prepend错误")
	}

	intColl.Prepend(12)
	if intColl.Count() != 8 {
		t.Fatal("Prepend 第二次错误")
	}
	if len(old) != 6 {
		t.Fatal("Prepend 修改了原数组")
	}
}

func TestAbsCollection_Random(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	out := intColl.Random()

	_, err := out.ToInt()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAbsCollection_Reverse(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
	vs := intColl.Reverse()
	if vs.Count() != 6 {
		t.Fatal("Reverse error")
	}
}

func TestAbsCollection_Mode(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3, 4, 5, 6})
	mode, err := intColl.Mode().ToInt()
	if err != nil {
		t.Fatal(err.Error())
	}
	if mode != 2 {
		t.Fatal("Mode error")
	}

	intColl = NewIntCollection([]int{1, 2, 2, 3, 4, 4, 5, 6})

	mode, err = intColl.Mode().ToInt()
	if err != nil {
		t.Fatal(err.Error())
	}
	if mode != 2 {
		t.Fatal("Mode error")
	}
}

func TestAbsCollection_Avg(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	mode, err := intColl.Avg().ToFloat64()
	if err != nil {
		t.Fatal(err.Error())
	}
	if mode != 2.0 {
		t.Fatal("Avg error")
	}
}

func TestAbsCollection_Shuffle(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	newColl := intColl.Shuffle()
	// newColl.DD()
	if newColl.Err() != nil {
		t.Fatal(newColl.Err())
	}
}

func TestAbsCollection_Max(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	max, err := intColl.Max().ToInt()
	if err != nil {
		t.Fatal(err)
	}

	if max != 3 {
		t.Fatal("max错误")
	}
}

func TestAbsCollection_Min(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	min, err := intColl.Min().ToInt()
	if err != nil {
		t.Fatal(err)
	}

	if min != 1 {
		t.Fatal("min错误")
	}
}

func TestAbsCollection_Contains(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	if intColl.Contains(1) != true {
		t.Fatal("contain 错误1")
	}
	if intColl.Contains(5) != false {
		t.Fatal("contain 错误2")
	}

}

func TestAbsCollection_Diff(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	intColl2 := NewIntCollection([]int{2, 3, 4})

	diff := intColl.Diff(intColl2)
	// diff.DD()
	if diff.Count() != 1 {
		t.Fatal("diff 错误")
	}
}

func TestAbsCollection_Sort(t *testing.T) {
	intColl := NewIntCollection([]int{2, 4, 3})
	intColl2 := intColl.Sort()
	if intColl2.Err() != nil {
		t.Fatal(intColl2.Err())
	}
	ins, err := intColl2.ToInts()
	if err != nil {
		t.Fatal(err)
	}
	if ins[1] != 3 || ins[0] != 2 {
		t.Fatal("sort error")
	}
}

func TestAbsCollection_SortDesc(t *testing.T) {
	intColl := NewIntCollection([]int{2, 4, 3})
	intColl2 := intColl.SortDesc()
	if intColl2.Err() != nil {
		t.Fatal(intColl2.Err())
	}
	ins, err := intColl2.ToInts()
	if err != nil {
		t.Fatal(err)
	}
	if ins[1] != 3 || ins[0] != 4 {
		t.Fatal("sort error")
	}
}

func TestAbsCollection_Join(t *testing.T) {
	intColl := NewIntCollection([]int{2, 4, 3})
	out := intColl.Join(",")
	if out != "2,4,3" {
		t.Fatal("join错误")
	}
	out = intColl.Join(",", func(item interface{}) string {
		return fmt.Sprintf("'%d'", item.(int))
	})
	if out != "'2','4','3'" {
		t.Fatal("join 错误")
	}
}

func TestAbsCollection_Median(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	median, err := intColl.Median().ToFloat64()
	if err != nil {
		t.Fatal(err)
	}

	if median != 2.0 {
		t.Fatal("Median 错误" + fmt.Sprintf("%v", median))
	}
}

func TestAbsCollection_Sum(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	sum, err := intColl.Sum().ToInt()
	if err != nil {
		t.Fatal(err)
	}

	if sum != 8 {
		t.Fatal("sum 错误")
	}
}

func TestAbsCollection_Filter(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	ret := intColl.Filter(func(obj interface{}, index int) bool {
		val := obj.(int)
		return val == 2
	})
	if ret.Count() != 2 {
		t.Fatal("filter error")
	}
}

func TestAbsCollection_First(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	ret := intColl.First(func(obj interface{}, index int) bool {
		val := obj.(int)
		return val > 2
	})
	if ret.MustToInt() != 3 {
		t.Fatal("first error")
	}
}

func TestAbsCollection_ToInts(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	arr, err := intColl.ToInts()
	if err != nil {
		t.Fatal(err)
	}
	if len(arr) != 4 {
		t.Fatal(errors.New("ToInts error"))
	}
}

func TestAbsCollection_ToMixs(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	arr, err := intColl.ToMixs()
	if err != nil {
		t.Fatal(err)
	}
	if len(arr) != 4 {
		t.Fatal(errors.New("ToInts error"))
	}
}

func TestAbsCollection_ContainsCount(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	count := intColl.ContainsCount(2)
	if count != 2 {
		t.Fatal(errors.New("contains count error"))
	}
}
