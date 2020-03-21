# collection

Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求极致性能，追求业务开发效能的场景。

| 版本 | 说明 |
| ------| ------ |
| 1.1.2 |  增加一些空数组的判断，解决一些issue |
| 1.1.1 |  对collection包进行了json解析和反解析的支持，对mix类型支持了SetField和RemoveFields的类型设置 |
| 1.1.0 |  增加了对int32的支持，增加了延迟加载，增加了Copy函数，增加了compare从ICollection传递到IMix，使用快排加速了Sort方法 |
| 1.0.1 |  第一次发布 |

`go get github.com/jianfengye/collection`

创建collection库的说明文章见：[一个让业务开发效率提高10倍的golang库](https://www.cnblogs.com/yjf512/p/10818089.html)

Collection包目前支持的元素类型：int32, int, int64, float32, float64, string, struct。



使用下列几个方法进行初始化Collection:

```go
NewIntCollection(objs []int) *IntCollection

NewInt64Collection(objs []int64) *Int64Collection

NewInt32Collection(objs []int32) *Int32Collection

NewFloat64Collection(objs []float64) *Float64Collection

NewFloat32Collection(objs []float32) *Float32Collection

NewStrCollection(objs []string) *StrCollection

NewObjCollection(objs interface{}) *ObjCollection
```

Collection的Error是随着Collection对象走，或者下沉到IMix中，所以可以放心在ICollection和IMix进行链式调用，只需要最后进行一次错误检查即可。

```
ret, err := objColl.Map(func(item interface{}, key int) IMix {
    foo := item.(Foo)
    return NewMix(foo.A)
}).Reduce(func(carry IMix, item IMix) IMix {
    ret, _ := carry.ToString()
    join, _ := item.ToString()
    return NewMix(ret + join)
}).ToString()
if err != nil {
    ...
}
```

支持的方法有:

[Append](#Append)

[Avg](#Avg)

[Contain](#Contain)

[Copy](#Copy)

[DD](#DD)

[Diff](#Diff)

[Each](#Each)

[Every](#Every)

[ForPage](#ForPage)

[Filter](#Filter)

[First](#First)

[Index](#Index) 

[IsEmpty](#IsEmpty)

[IsNotEmpty](#IsNotEmpty)

[Join](#Join)

[Last](#Last)

[Last](#Last)

[Merge](#Merge)

[Map](#Map)

[Mode](#Mode)

[Max](#Max)

[Min](#Min)

[Median](#Median)

[NewEmpty](#NewEmpty)

[Nth](#Nth)

[Pad](#Pad)

[Pop](#Pop)

[Push](#Push)

[Prepend](#Prepend)

[Pluck](#Pluck)

[Reject](#Reject)

[Reduce](#Reduce)

[Random](#Random)

[Reverse](#Reverse)

[Slice](#Slice)

[Search](#Search)

[Sort](#Sort)

[SortDesc](#SortDesc)

[Sum](#Sum)

[Shuffle](#Shuffle)

[SortBy](#SortBy)

[SortByDesc](#SortByDesc)

[ToInts](#ToInts)

[ToInt64s](#ToInt64s)

[ToFloat64s](#ToFloat64s)

[ToFloat32s](#ToFloat32s)

[ToMixs](#ToMixs)

[Unique](#Unique)

### Append

`Append(item interface{}) ICollection`

Append挂载一个元素到当前Collection，如果挂载的元素类型不一致，则会在Collection中产生Error

```
intColl := NewIntCollection([]int{1,2})
intColl.Append(3)
if intColl.Err() == nil {
    intColl.DD()
}

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

### Avg

`Avg() IMix`

返回Collection的数值平均数，这里会进行类型降级，int,int64,float64的数值平均数都是返回float64类型。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
mode, err := intColl.Avg().ToFloat64()
if err != nil {
    t.Error(err.Error())
}
if mode != 2.0 {
    t.Error("Avg error")
}
```

### Copy

Copy方法根据当前的数组，创造出一个同类型的数组，有相同的元素

```
func TestAbsCollection_Copy(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl2 := intColl.Copy()
	intColl2.DD()
	if intColl2.Count() != 2 {
		t.Error("Copy失败")
	}
	if reflect.TypeOf(intColl2) != reflect.TypeOf(intColl) {
		t.Error("Copy类型失败")
	}
}
```



### Contain

`Contains(obj interface{}) bool`

判断一个元素是否在Collection中，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
if intColl.Contains(1) != true {
    t.Error("contain 错误1")
}
if intColl.Contains(5) != false {
    t.Error("contain 错误2")
}
```

### Diff

`Diff(arr ICollection) ICollection`

获取前一个Collection不在后一个Collection中的元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl2 := NewIntCollection([]int{2, 3, 4})

diff := intColl.Diff(intColl2)
diff.DD()
if diff.Count() != 1 {
    t.Error("diff 错误")
}

/*
IntCollection(1):{
	0:	1
}
*/
```

### DD 

`DD()`

DD方法按照友好的格式展示Collection

```
a1 := Foo{A: "a1"}
a2 := Foo{A: "a2"}

objColl := NewObjCollection([]Foo{a1, a2})
objColl.DD()

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a1}
	1:	{A:a2}
}
*/

intColl := NewIntCollection([]int{1,2})
intColl.DD()

/*
IntCollection(2):{
	0:	1
	1:	2
}
*/
```


### Each

`Each(func(item interface{}, key int))`

对Collection中的每个函数都进行一次函数调用。传入的参数是回调函数。

如果希望在某次调用的时候中止，在此次调用的时候设置Collection的Error，就可以中止调用。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
sum := 0
intColl.Each(func(item interface{}, key int) {
    v := item.(int)
    sum = sum + v
})

if intColl.Err() != nil {
    t.Error(intColl.Err())
}

if sum != 10 {
    t.Error("Each 错误")
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
    t.Error("Each 错误")
}

/*
PASS
*/
```

### Every

`Every(func(item interface{}, key int) bool) bool`

判断Collection中的每个元素是否都符合某个条件，只有当每个元素都符合条件，才整体返回true，否则返回false。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
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
```

### ForPage

`ForPage(page int, perPage int) ICollection`

将Collection函数进行分页，按照每页第二个参数的个数，获取第一个参数的页数数据。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
ret := intColl.ForPage(1, 2)
ret.DD()

if ret.Count() != 2 {
    t.Error("For page错误")
}

/*
IntCollection(2):{
	0:	3
	1:	4
}
*/
```


### Filter

`Filter(func(item interface{}, key int) bool) ICollection`

根据过滤函数获取Collection过滤后的元素。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.Filter(func(obj interface{}, index int) bool {
    val := obj.(int)
    if val == 2 {
        return true
    }
    return false
}).DD()

/*
IntCollection(2):{
	0:	2
	1:	2
}
*/
```

### First

`First(...func(item interface{}, key int) bool) IMix`

获取符合过滤条件的第一个元素，如果没有填写过滤函数，返回第一个元素。

注：只能传递0个或者1个过滤函数，如果传递超过1个过滤函数，只有第一个过滤函数起作用

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.First(func(obj interface{}, index int) bool {
    val := obj.(int)
    if val > 2 {
        return true
    }
    return false
}).DD()

/*
IMix(int): 3 
*/

```

```
func TestIntCollection_Filter(t *testing.T) {
	intColl := NewIntCollection([]int{1,2,3})
	a, err := intColl.First().ToInt()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(a, 1) {
		t.Error("filter error")
	}
}
```

### Index

`Index(i int) IMix`

Index获取元素中的第几个元素，下标从0开始，如果i超出了长度，则Collection记录错误。

```
intColl := NewIntCollection([]int{1,2})
foo := intColl.Index(1)
foo.DD()

/*
IMix(int): 2 
*/
```

### IsEmpty

`IsEmpty() bool`

判断一个Collection是否为空，为空返回true, 否则返回false

```go
intColl := NewIntCollection([]int{1,2})
println(intColl.IsEmpty())  // false
```

### IsNotEmpty

`IsNotEmpty() bool`

判断一个Collection是否为空，为空返回false，否则返回true
```go
intColl := NewIntCollection([]int{1,2})
println(intColl.IsNotEmpty()) // true
```


### Join

`Join(split string, format ...func(item interface{}) string) string`

将Collection中的元素按照某种方式聚合成字符串。该函数接受一个或者两个参数，第一个参数是聚合字符串的分隔符号，第二个参数是聚合时候每个元素的格式化函数，如果没有设置第二个参数，则使用`fmt.Sprintf("%v")`来该格式化

```go
intColl := NewIntCollection([]int{2, 4, 3})
out := intColl.Join(",")
if out != "2,4,3" {
    t.Error("join错误")
}
out = intColl.Join(",", func(item interface{}) string {
    return fmt.Sprintf("'%d'", item.(int))
})
if out != "'2','4','3'" {
    t.Error("join 错误")
}
```

### Last

`Last(...func(item interface{}, key int) bool) IMix`

获取该Collection中满足过滤的最后一个元素，如果没有填写过滤条件，默认返回最后一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 3, 2})
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
```


### Merge

`Merge(arr ICollection) ICollection`

将两个Collection的元素进行合并，这个函数会修改原Collection。

```go
intColl := NewIntCollection([]int{1, 2 })

intColl2 := NewIntCollection([]int{3, 4})

intColl.Merge(intColl2)

if intColl.Err() != nil {
    t.Error(intColl.Err())
}

if intColl.Count() != 4 {
    t.Error("Merge 错误")
}

intColl.DD()

/*
IntCollection(4):{
	0:	1
	1:	2
	2:	3
	3:	4
}
*/
```

### Map

`Map(func(item interface{}, key int) interface{}) ICollection`

对Collection中的每个函数都进行一次函数调用，并将返回值组装成ICollection

这个回调函数形如： `func(item interface{}, key int) interface{}`

如果希望在某此调用的时候中止，就在此次调用的时候设置Collection的Error，就可以中止，且此次回调函数生成的结构不合并到最终生成的ICollection。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
newIntColl := intColl.Map(func(item interface{}, key int) interface{} {
    v := item.(int)
    return v * 2
})
newIntColl.DD()

if newIntColl.Count() != 4 {
    t.Error("Map错误")
}

newIntColl2 := intColl.Map(func(item interface{}, key int) interface{} {
    v := item.(int)

    if key > 2 {
        intColl.SetErr(errors.New("break"))
        return nil
    }

    return v * 2
})
newIntColl2.DD()

/*
IntCollection(4):{
	0:	2
	1:	4
	2:	6
	3:	8
}
IntCollection(3):{
	0:	2
	1:	4
	2:	6
}
*/
```

### Mode

`Mode() IMix`

获取Collection中的众数，如果有大于两个的众数，返回第一次出现的那个。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3, 4, 5, 6})
mode, err := intColl.Mode().ToInt()
 if err != nil {
     t.Error(err.Error())
 }
 if mode != 2 {
     t.Error("Mode error")
 }
 
 intColl = NewIntCollection([]int{1, 2, 2, 3, 4, 4, 5, 6})
 
 mode, err = intColl.Mode().ToInt()
 if err != nil {
     t.Error(err.Error())
 }
 if mode != 2 {
     t.Error("Mode error")
 }
```



### Max

`Max() IMix`

获取Collection中的最大元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
max, err := intColl.Max().ToInt()
if err != nil {
    t.Error(err)
}

if max != 3 {
    t.Error("max错误")
}

```

### Min

`Min() IMix`

获取Collection中的最小元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
min, err := intColl.Min().ToInt()
if err != nil {
    t.Error(err)
}

if min != 1 {
    t.Error("min错误")
}

```


### Median

`Median() IMix`

获取Collection的中位数，如果Collection个数是单数，返回排序后中间的元素，如果Collection的个数是双数，返回排序后中间两个元素的算数平均数。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
median, err := intColl.Median().ToFloat64()
if err != nil {
    t.Error(err)
}

if median != 2.0 {
    t.Error("Median 错误" + fmt.Sprintf("%v", median))
}
```

### NewEmpty

`NewEmpty(err ...error) ICollection`

NewEmpty方法根据当前的数组，创造出一个同类型的数组，但长度为0

```
intColl := NewIntCollection([]int{1,2})
intColl2 := intColl.NewEmpty()
intColl2.DD()

/*
IntCollection(0):{
}
*/
```


### Nth

`Nth(n int, offset int) ICollection`

Nth(n int, offset int) 获取从offset偏移量开始的每第n个，偏移量offset的设置为第一个。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
ret := intColl.Nth(4, 1)
ret.DD()

if ret.Count() != 2 {
    t.Error("Nth 错误")
}

/*
IntCollection(2):{
	0:	2
	1:	6
}
*/
```

### Pad

`Pad(start int, def interface{}) ICollection` 

填充Collection数组，如果第一个参数大于0，则代表往Collection的右边进行填充，如果第一个参数小于零，则代表往Collection的左边进行填充。

```go
intColl := NewIntCollection([]int{1, 2, 3})
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

/*
IntCollection(5):{
	0:	1
	1:	2
	2:	3
	3:	0
	4:	0
}
IntCollection(5):{
	0:	0
	1:	0
	2:	1
	3:	2
	4:	3
}
*/
```

### Pop

`Pop() IMix`

从Collection右侧弹出一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
pop := intColl.Pop()
in, err := pop.ToInt()
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

/*
IntCollection(5):{
	0:	1
	1:	2
	2:	3
	3:	4
	4:	5
}
*/
```

### Push

`Push(item interface{}) ICollection`

往Collection的右侧推入一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
intColl.Push(7)
intColl.DD()
if intColl.Count() != 7 {
    t.Error("Push 后本体错误")
}

/*
IntCollection(7):{
	0:	1
	1:	2
	2:	3
	3:	4
	4:	5
	5:	6
	6:	7
}
*/
```

### Prepend

`Prepend(item interface{}) ICollection`

往Collection左侧加入元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
intColl.Prepend(0)
if intColl.Err() != nil {
    t.Error(intColl.Err().Error())
}

intColl.DD()
if intColl.Count() != 7 {
    t.Error("Prepend错误")
}

/*
IntCollection(7):{
	0:	0
	1:	1
	2:	2
	3:	3
	4:	4
	5:	5
	6:	6
}
*/
```

### Pluck

`Pluck(key string) ICollection`

将对象数组中的某个元素提取出来组成一个新的Collection。这个元素必须是Public元素

注：这个函数只对ObjCollection生效。

```go
type Foo struct {
	A string
}

func TestObjCollection_Pluck(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})

	objColl.Pluck("A").DD()
}

/*
StrCollection(2):{
	0:	a1
	1:	a2
}
*/
```

### Reject

`Reject(func(item interface{}, key int) bool) ICollection`

将满足过滤条件的元素删除

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
retColl := intColl.Reject(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 3
})
if retColl.Count() != 3 {
    t.Error("Reject 重复错误")
}

retColl.DD()

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

### Reduce

`Reduce(func(carry IMix, item IMix) IMix) IMix`

对Collection中的所有元素进行聚合计算。

如果希望在某次调用的时候中止，在此次调用的时候设置Collection的Error，就可以中止调用。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
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

/*
IMix(int): 10 
*/
```

### Random

`Random() IMix`

随机获取Collection中的元素，随机数种子使用时间戳

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
out := intColl.Random()
out.DD()

_, err := out.ToInt()
if err != nil {
    t.Error(err.Error())
}

/*
IMix(int): 5 
*/
```

### Reverse

`Reverse() ICollection`

将Collection数组进行转置

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
vs := intColl.Reverse()
vs.DD()

/*
IntCollection(6):{
	0:	6
	1:	5
	2:	4
	3:	3
	4:	2
	5:	1
}
*/
```

### Search

`Search(item interface{}) int`

查找Collection中第一个匹配查询元素的下标，如果存在，返回下标；如果不存在，返回-1

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2})
if intColl.Search(2) != 1 {
    t.Error("Search 错误")
}

intColl = NewIntCollection([]int{1,2, 3, 3, 2})
if intColl.Search(3) != 2 {
    t.Error("Search 重复错误")
}
```

### Slice

`Slice(...int) ICollection`

获取Collection中的片段，可以有两个参数或者一个参数。

如果是两个参数，第一个参数代表开始下标，第二个参数代表结束下标，当第二个参数为-1时候，就代表到Collection结束。

如果是一个参数，则代表从这个开始下标一直获取到Collection结束的片段。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
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

/*
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
IntCollection(2):{
	0:	3
	1:	4
}
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
*/

```


### Shuffle

`Shuffle() ICollection`

将Collection中的元素进行乱序排列，随机数种子使用时间戳

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
newColl := intColl.Shuffle()
newColl.DD()
if newColl.Err() != nil {
    t.Error(newColl.Err())
}

/*
IntCollection(4):{
	0:	1
	1:	3
	2:	2
	3:	2
}
*/
```

### Sort

`Sort() ICollection`

将Collection中的元素进行升序排列输出，必须设置compare函数

```go
intColl := NewIntCollection([]int{2, 4, 3})
intColl2 := intColl.Sort()
if intColl2.Err() != nil {
    t.Error(intColl2.Err())
}
intColl2.DD()

/*
IntCollection(3):{
	0:	2
	1:	3
	2:	4
}
*/
```

### SortDesc

`SortDesc() ICollection`

将Collection中的元素按照降序排列输出，必须设置compare函数

```go
intColl := NewIntCollection([]int{2, 4, 3})
intColl2 := intColl.SortDesc()
if intColl2.Err() != nil {
    t.Error(intColl2.Err())
}
intColl2.DD()

/*
IntCollection(3):{
	0:	4
	1:	3
	2:	2
}
*/
```


### Sum

`Sum() IMix`

返回Collection中的元素的和

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.Sum().DD()
sum, err := intColl.Sum().ToInt()
if err != nil {
    t.Error(err)
}

if sum != 8 {
    t.Error("sum 错误")
}

/*
IMix(int): 8 
*/
```

### SortBy

`SortBy(key string) ICollection`

根据对象数组中的某个元素进行Collection升序排列。这个元素必须是Public元素

注：这个函数只对ObjCollection生效。这个对象数组的某个元素必须是基础类型。

```go
type Foo struct {
	A string
	B int
}

func TestObjCollection_SortBy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})

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

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a2 B:2}
	1:	{A:a1 B:3}
}
*/
```

### SortByDesc

`SortByDesc(key string) ICollection`

根据对象数组中的某个元素进行Collection降序排列。这个元素必须是Public元素

注：这个函数只对ObjCollection生效。这个对象数组的某个元素必须是基础类型。

```go
type Foo struct {
	A string
	B int
}

func TestObjCollection_SortByDesc(t *testing.T) {
	a1 := Foo{A: "a1", B: 2}
	a2 := Foo{A: "a2", B: 3}

	objColl := NewObjCollection([]Foo{a1, a2})

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

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a2 B:3}
	1:	{A:a1 B:2}
}
*/
```

------------

### ToInts

`ToInts() ([]int, error)`

将Collection变化为int数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
arr, err := intColl.ToInts()
if err != nil {
    t.Error(err)
}
if len(arr) != 4 {
    t.Error(errors.New("ToInts error"))
}
```

### ToInt64s

`ToInt64s() ([]int64, error)`

将Collection变化为int64数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误。

```go
intColl := NewInt64Collection([]int{1, 2, 2, 3})
arr, err := intColl.ToInts()
if err != nil {
    t.Error(err)
}
if len(arr) != 4 {
    t.Error(errors.New("ToInts error"))
}
```

### ToFloat64s

`ToFloat64s() ([]float64, error)`

将Collection变化为float64数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误。

```go
arr := NewFloat64Collection([]float64{1.0 ,2.0,3.0,4.0,5.0})

arr.DD()

max, err := arr.Max().ToFloat64()
if err != nil {
    t.Error(err)
}

if max != 5 {
    t.Error(errors.New("max error"))
}


arr2 := arr.Filter(func(obj interface{}, index int) bool {
    val := obj.(float64)
    if val > 2.0 {
        return true
    }
    return false
})
if arr2.Count() != 3 {
    t.Error(errors.New("filter error"))
}

out, err := arr2.ToFloat64s()
if err != nil || len(out) != 3 {
    t.Error(errors.New("to float64s error"))
}

```

### ToFloat32s

`ToFloat32s() ([]float32, error)`

将Collection变化为float32数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误。

```go
arr := NewFloat32Collection([]float32{1.0 ,2.0,3.0,4.0,5.0})

arr.DD()

max, err := arr.Max().ToFloat32()
if err != nil {
    t.Error(err)
}

if max != 5 {
    t.Error(errors.New("max error"))
}


arr2 := arr.Filter(func(obj interface{}, index int) bool {
    val := obj.(float32)
    if val > 2.0 {
        return true
    }
    return false
})
if arr2.Count() != 3 {
    t.Error(errors.New("filter error"))
}

out, err := arr2.ToFloat32s()
if err != nil || len(out) != 3 {
    t.Error(errors.New("to float32s error"))
}
```

### ToMixs

`ToMixs() ([]IMix, error)`

将Collection变化为Mix数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
arr, err := intColl.ToMixs()
if err != nil {
    t.Error(err)
}
if len(arr) != 4 {
    t.Error(errors.New("ToInts error"))
}
```

### Unique

`Unique() ICollection`

将Collection中重复的元素进行合并，返回唯一的一个数组。

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2, 3, 3, 2})
uniqColl := intColl.Unique()
if uniqColl.Count() != 3 {
    t.Error("Unique 重复错误")
}

uniqColl.DD()
/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

























License
------------
`collection` is licensed under [Apache License](LICENSE).

## Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/jianfengye/collection/graphs/contributors"><img src="https://opencollective.com/collection/contributors.svg?width=890&button=false" /></a>


## Backers

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/collection#backer)]

<a href="https://opencollective.com/collection#backers" target="_blank"><img src="https://opencollective.com/collection/backers.svg?width=890"></a>


## Sponsors

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/collection#sponsor)]

<a href="https://opencollective.com/collection/sponsor/0/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/1/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/2/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/3/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/4/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/5/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/6/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/7/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/8/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/9/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/9/avatar.svg"></a>


