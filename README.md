# collection

这个包目标是用于替换golang原生的Slice和Map，使用场景是在大量不追求性能，追求业务开发效能的场景。

# 不要忌讳panic

golang写业务代码经常会被吐槽，写业务太慢了，其中最大的吐槽点就是，处理各种error太麻烦了。一个项目中，会有30%或者更多的是在处理error。

对于golang的error这个事情，golang的官方也说的很详细，建议函数返回error，并且让上层调用处理。

error和panic实际上就是以前写PHP业务的时候争论的使用errno还是exception的争论。实际上，后续在PHP世界里面，大家都倾向于会使用exception来做错误处理。不知道为何，在golang这个环境中，好像网络上更倾向于使用error的机制。

我思考，理由是处理的问题语境不同。在服务端工具世界里面，我们并不希望跑出来一个程序告诉我coredump了，而希望能准确告诉我为什么coredump，这样才是一个更健壮的工具。而在业务，特别是web业务中，我们更希望的是快速实现主流程功能，一些错误机制，我们希望能后续逐步加入。基本上，在web业务中，速度决定市场，从有一个idea，到具体的实现，整个过程的实现路径越短越好，但是到实现上线后，还可以有后续的迭代版本。有充足的时间来增加健壮性。但是这个第一步上线的速度应该是第一位的。

所以我觉得，如果大多数gopher都是觉得要立即处理error，我倒并不这么看。panic的意思更像是一种延后处理的TODO机制：“TODO：我不知道这个错误现在怎么处理，后续有时间再完善。”。这种逻辑，在业务代码中是非常必要的，而且会让写业务的思维更多的聚焦在正确的实现路径上。就像一棵树，先用最快的步骤实现了树干，再实现具体的枝枝叶叶。所以我强烈建议，在写业务代码的时候，对于不确定，或者基本不想实现的error，直接panic出去，在需要捕获的地方recover这个panic，这才是写业务代码，golang应该有的姿态。

即使golang2.0出来的error handler机制，也就在一个函数范围内做错误的handle，我认为还是不够，我们希望能在goroutine有一个统一的地方进行recover。

所以，不要忌讳panic，在写业务代码的时候。

# 让人发愁的slice & map

写过PHP的人自然就知道，php中的array是个多么强大的存在，如果你写过laravel，laravel中的collection更是神一般的存在。但是在golang中，slice和map是基础的结构，但是由于强类型的关系，对一个slice进行查找,都是有一定思维负担的工作。

这里说的思维负担，是说我的思维逻辑正处在“实现这个业务，我第一步需要从A数据库拿取a数据，从B数据库中拿取b数据，然后这两个数据进行合并去重”这种逻辑中抽离出来，思考如何进行“合并去重”。如果我们在写业务逻辑过程中，屡屡被这种“从array中获取一个最大值”、“从map中获取所有key”等等的逻辑，这个无意对于业务代码的速度是个不小的拖累。

所以基于这个逻辑，我使用的方法是自己实现了一个IArray和IMap的接口的实现。这两个接口的实现尽量能完成对于golang的Slice和Map的一些更为通用的操作，但是同时又能保持一定的扩展性。[collection包](https://github.com/jianfengye/collection)

这个包所有的错误“类型不对”，“结构不对”等，均使用panic的方式来返回错误，这样，把这个包强制处理成可以链式模式处理的逻辑。增加代码可读性。

比如下面这个需求：
```
1 我需要从一个inoutLinks的Slice对象中，获取所有的LogicLinkID的的值，排重，得到一个数组。
2 我需要从一个inoutLinks的Slice对象中，获取其中最大的ID值。
```
```
// 需求一：
// 从junctionMap表批量根据junction_id批量获取
objArray := collection.NewObjArray(reflect.ValueOf(inoutLinks))
logicLinkIds := objArray.Column("LogicLinkID").Unique().ToString()

// 需求二：
// 获取最大的一个元素作为start
startId = objArray.Column("ID").Max().ToInt64()
```

上面的两个代码看起来确实清爽不少，并且意思很清晰。

当然也有副作用：

## 错误

一旦其中有错误，或者有“未想到的传入错误”，就会在链条的任何一个地方panic出错误，需要及时recover。

这个在上一节“不要忌讳panic”就思考过，我现在的代码是为了正常的业务逻辑，比如这里的其他异常的业务逻辑，比如“如果这个inoutLinks没有一个字段叫做ID”这种错误，就直接panic出错误了。

所以强烈建议在统一的一个地方进行recover，比如：
```
func() {
			defer func() {
				if err := recover(); err != nil{
					log.Println("painc error:", err, string(debug.Stack()))
				}
			}()
      doSomeThings()
      ...
}
```
这里建议使用debug.Stack来将错误堆栈打印出来，以便于调试。

其实try...catch...的本质就是在于代码逻辑的中断和goto。在golang中，只有使用defer+recover才能进行这个处理。

这里还是要强调一下，我建议错误直接panic，是在web业务处理场景下。

## 性能

看到上面我的例子，一定有人会诟病，这里的Column方法，本质里面是不是使用了反射啊？那么性能是不是没有保证？

对于golang的反射，我的态度也是，并不要惧怕使用。golang是一个工具，它的出现本质是为了解决问题，而不是要求所有代码的性能。换而言之，如果我代码中大量使用反射，增加了我代码的灵活度，减少了开发周期，更早的占据了市场的份额，那么这个工具在这个事情上的使用，就是成功的。不要被性能所绑架。在准确评估市场，项目，访问量的情况下，大部分的业务项目应该来说，都可以牺牲一定的性能来满足业务的实现速度的。

所以，我这里的ObjectArray中的Column方法等，使用了反射等原理。

当然有人会质疑，我并不是所有的业务接口都不追求性能。当我需要追求性能的时候，难道让我重写一遍？

所以这里的Collection包使用的是接口和继承设计，换而言之，如果你对某个ObjectArray的性能确实需要非常追求的话，当你对某个Object获取ID字段的值是非常需要性能的话，你完全可以自定义一个继承ObjectArray的结构，并且覆盖实现其中的Column方法，直接写
```
func (arr *ObjArray) Column(key string) IArray {
  ...

  if key == "ID" {
    for _, obj := arr {
      result = append(result, obj.ID)
    }
    return result
  }
}
```

所以这个Collection包的基本思想是“提供加速golang业务代码的能力，同时提供足够扩展追求性能的能力”。

Collection包实现了两个通用数据类型，希望能在追求业务代码速度的场景中替换golang中的slice和map:
```

type IArray interface {
	// 放入一个元素到数组中，对所有Array生效
	Append(obj interface{})

	// 查找数据中是否包含，-1不包含，>=0 返回数组中元素下标，对所有Array生效
	Search(obj interface{}) int
	// 返回数组中对象的某个key组成的数组，仅对ObjectArray生效
	Column(key string) IArray
	// 过滤数组中重复的元素，仅对基础Array生效
	Unique() IArray

	// 将数组中对象某个key作为map的key，整个对象作为value，作为map返回，如果key有重复会进行覆盖，仅对ObjectArray生效
	KeyBy(key string) *Map

	// 数组中最大的元素，仅对基础Array生效
	Max() *Mix
	// 数组中最小的元素，仅对基础Array生效
	Min() *Mix

	// 获取数组片段，对所有Array生效
	Slice(start, end int) IArray
	// 获取某个下标，对所有Array生效
	Index(i int) *Mix
	// 获取数组长度，对所有Array生效
	Len() int
	// 判断是否包含某个元素，（并不进行定位），对基础Array生效
	Has(obj interface{}) bool
	// 将两个数组进行合并，参数的数据挂在当前数组中，返回当前数组，对所有Array生效
	Merge(arr IArray) IArray

	// 转化为golang原生的字符数组，仅对StrArray生效
	ToString() []string
	// 转化为golang原生的Int64数组，仅对Int64Array生效
	ToInt64() []int64
	// 转化为golang原生的Int数组，仅对IntArray生效
	ToInt() []int
}
```
```
type IMap interface {

	// 设置一个Map的key和value，如果key存在，则覆盖
	Set(key interface{}, value interface{})
	// 删除一个Map的key
	Remove(key interface{})
	// 根据key获取一个Map的value
	Get(key interface{}) *Mix
	// 获取一个Map的长度
	Len() int

	// 获取Map的所有key组成的集合
	Keys() IArray
	// 获取Map的所有value组成的集合
	Values() IArray
}
```

其中有些地方不确定的单个元素的类型，使用的是*Mix结构，这个结构提供一系列的ToXxx接口，使用方需要对这个Mix对象所代表的数据结构负责。

