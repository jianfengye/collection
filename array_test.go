package collection

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestNewCollection(t *testing.T) {
	// Test that NewCollection returns a non-nil pointer to a Collection struct
	arr := []int32{1, 2, 3}
	c := NewCollection[int32](arr)
	if c == nil {
		t.Error("NewCollection returned nil")
	}
}

func TestNewEmptyCollection(t *testing.T) {
	c := NewEmptyCollection[int32]()
	if c == nil {
		t.Error("NewEmptyCollection returned nil")
	}
	if len(c.value) != 0 {
		t.Error("NewEmptyCollection did not return an empty collection")
	}
}

func TestCopy(t *testing.T) {
	// Test that Copy returns a new Collection with the same values and error as the original
	arr := []int32{1, 2, 3}
	c := NewCollection[int32](arr).SetErr(errors.New("test error"))
	copied := c.Copy()
	if !reflect.DeepEqual(copied.value, c.value) {
		t.Error("Copy did not copy the values correctly")
	}
	if copied.err.Error() != c.err.Error() {
		t.Error("Copy did not copy the error correctly")
	}
}

func TestIsEmpty(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// check if IsEmpty() returns false
	if coll.IsEmpty() {
		t.Errorf("IsEmpty() returned true, expected false")
	}

	// remove all elements from the Collection
	coll.Remove(0).Remove(0).Remove(0)

	// check if IsEmpty() returns true
	if !coll.IsEmpty() {
		t.Errorf("IsEmpty() returned false, expected true")
	}
}

func TestIsNotEmpty(t *testing.T) {
	coll := NewCollection[int]([]int{1, 2, 3})
	if !coll.IsNotEmpty() {
		t.Errorf("Expected collection to be not empty, but it is empty")
	}

	coll = NewEmptyCollection[int]()
	if coll.IsNotEmpty() {
		t.Errorf("Expected collection to be empty, but it is not empty")
	}
}

func TestCollectionAppend(t *testing.T) {
	coll := NewCollection[int]([]int{1, 2, 3})
	newItem := 4
	coll.Append(newItem)
	if len(coll.value) != 4 {
		t.Errorf("Expected length of collection to be 4, but got %d", len(coll.value))
	}
	if coll.value[len(coll.value)-1] != newItem {
		t.Errorf("Expected last item in collection to be %d, but got %d", newItem, coll.value[len(coll.value)-1])
	}
}

func TestCollectionRemove(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// remove the first element
	coll.Remove(0)

	// check if the first element was removed
	if coll.value[0] != 2 {
		t.Errorf("Remove did not remove the element correctly")
	}

	// remove the last element
	coll.Remove(1)

	// check if the last element was removed
	if coll.value[0] != 2 || len(coll.value) != 1 {
		t.Errorf("Remove did not remove the element correctly")
	}
}

func TestCollectionInsert(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// insert a new element at the beginning
	coll.Insert(0, 0)

	// check if the element was inserted correctly
	if coll.value[0] != 0 {
		t.Errorf("Insert did not insert the element correctly")
	}

	// insert a new element at the end
	coll.Insert(4, 4)

	// check if the element was inserted correctly
	if coll.value[4] != 4 {
		t.Errorf("Insert did not insert the element correctly")
	}
}

func TestCollectionSearch(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// search for an existing element
	index := coll.Search(2)

	// check if the index of the element was returned correctly
	if index != 1 {
		t.Errorf("Search did not return the correct index")
	}

	// search for a non-existing element
	index = coll.Search(4)

	// check if -1 was returned
	if index != -1 {
		t.Errorf("Search did not return -1 for a non-existing element")
	}
}
func TestUnique(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 2, 3, 3, 3})

	// get the unique elements
	uniqueColl := coll.Unique()

	// check if the length of the unique collection is correct
	if len(uniqueColl.value) != 3 {
		t.Errorf("Unique did not return the correct number of unique elements")
	}

	// check if the unique collection contains the correct elements
	if uniqueColl.value[0] != 1 || uniqueColl.value[1] != 2 || uniqueColl.value[2] != 3 {
		t.Errorf("Unique did not return the correct unique elements")
	}
}

// TestFilter tests the Filter method of the Collection struct
func TestFilter(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// define a filter function that only keeps even numbers
	filterFunc := func(item int, key int) bool {
		return item%2 == 0
	}

	// apply the filter function to the collection
	filteredColl := coll.Filter(filterFunc)

	// check if the length of the filtered collection is correct
	if len(filteredColl.value) != 2 {
		t.Errorf("Filter did not return the correct number of elements")
	}

	// check if the filtered collection contains the correct elements
	if filteredColl.value[0] != 2 || filteredColl.value[1] != 4 {
		t.Errorf("Filter did not return the correct elements")
	}
}

// TestReject tests the Reject method of the Collection struct
func TestReject(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// define a reject function that rejects even numbers
	rejectFunc := func(item int, key int) bool {
		return item%2 == 0
	}

	// apply the reject function to the collection
	rejectedColl := coll.Reject(rejectFunc)

	// check if the length of the rejected collection is correct
	if len(rejectedColl.value) != 3 {
		t.Errorf("Reject did not return the correct number of elements")
	}

	// check if the rejected collection contains the correct elements
	if rejectedColl.value[0] != 1 || rejectedColl.value[1] != 3 || rejectedColl.value[2] != 5 {
		t.Errorf("Reject did not return the correct elements")
	}
}

// TestFirst tests the First method of the Collection struct
func TestFirst(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// get the first element
	first := coll.First()

	// check if the first element is correct
	if first != 1 {
		t.Errorf("First did not return the correct first element")
	}

	// create a new empty Collection
	emptyColl := NewEmptyCollection[int]()

	// get the first element of an empty Collection
	first = emptyColl.First()

	// check if the first element is the zero value of the type
	if first != 0 {
		t.Errorf("First did not return the zero value of the type for an empty Collection")
	}
}

func TestLast(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// get the last element
	last := coll.Last()

	// check if the last element is correct
	if last != 3 {
		t.Errorf("Last did not return the correct last element")
	}

	// create a new empty Collection
	emptyColl := NewEmptyCollection[int]()

	// get the last element of an empty Collection
	last = emptyColl.Last()

	// check if the last element is the zero value of the type
	if last != 0 {
		t.Errorf("Last did not return the zero value of the type for an empty Collection")
	}
}

func TestSlice(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get a slice of the collection
	slice := coll.Slice(1, 3)

	// check if the length of the slice is correct
	if len(slice.value) != 2 {
		t.Errorf("Slice did not return the correct number of elements")
	}

	// check if the slice contains the correct elements
	if slice.value[0] != 2 || slice.value[1] != 3 {
		t.Errorf("Slice did not return the correct elements")
	}
}

func TestIndex(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// get the element at index 1
	element := coll.Index(1)

	// check if the element is correct
	if element != 2 {
		t.Errorf("Index did not return the correct element")
	}
}

func TestSetIndex(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// set the element at index 1 to a new value
	coll.SetIndex(1, 4)

	// check if the element was set correctly
	if coll.value[1] != 4 {
		t.Errorf("SetIndex did not set the element correctly")
	}
}

func TestCount(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// get the count of elements
	count := coll.Count()

	// check if the count is correct
	if count != 3 {
		t.Errorf("Count did not return the correct count")
	}
}

func TestMerge(t *testing.T) {
	// create two new Collections with some elements
	coll1 := NewCollection([]int{1, 2, 3})
	coll2 := NewCollection([]int{4, 5, 6})

	// merge the two collections
	mergedColl := coll1.Merge(coll2)

	// check if the length of the merged collection is correct
	if len(mergedColl.value) != 6 {
		t.Errorf("Merge did not return the correct number of elements")
	}

	// check if the merged collection contains the correct elements
	if mergedColl.value[0] != 1 || mergedColl.value[1] != 2 || mergedColl.value[2] != 3 || mergedColl.value[3] != 4 || mergedColl.value[4] != 5 || mergedColl.value[5] != 6 {
		t.Errorf("Merge did not return the correct elements")
	}
}

func TestEach(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// define a function that adds 1 to each element
	addOne := func(item int, key int) {
		coll.value[key] = item + 1
	}

	// apply the function to each element of the collection
	coll.Each(addOne)

	// check if the elements were modified correctly
	if coll.value[0] != 2 || coll.value[1] != 3 || coll.value[2] != 4 {
		t.Errorf("Each did not modify the elements correctly")
	}
}

func TestMap(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// define a function that multiplies each element by 2
	multiplyByTwo := func(item int, key int) int {
		return item * 2
	}

	// apply the function to each element of the collection
	mappedColl := coll.Map(multiplyByTwo)

	// check if the length of the mapped collection is correct
	if len(mappedColl.value) != 3 {
		t.Errorf("Map did not return the correct number of elements")
	}

	// check if the mapped collection contains the correct elements
	if mappedColl.value[0] != 2 || mappedColl.value[1] != 4 || mappedColl.value[2] != 6 {
		t.Errorf("Map did not return the correct elements")
	}
}

// TestReduce tests the Reduce method of the Collection struct
func TestReduce(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// define a function that sums two integers
	sum := func(carry int, item int) int {
		return carry + item
	}

	// apply the function to the collection
	result := coll.Reduce(sum)

	// check if the result is correct
	if result != 6 {
		t.Errorf("Reduce did not return the correct result")
	}
}

// TestEvery tests the Every method of the Collection struct
func TestEvery(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{2, 4, 6})

	// define a function that checks if an integer is even
	isEven := func(item int, key int) bool {
		return item%2 == 0
	}

	// check if every element of the collection is even
	if !coll.Every(isEven) {
		t.Errorf("Every did not return true for a collection where every element satisfies the condition")
	}

	// create a new Collection with some elements
	coll = NewCollection([]int{2, 3, 6})

	// check if every element of the collection is even
	if coll.Every(isEven) {
		t.Errorf("Every did not return false for a collection where not every element satisfies the condition")
	}
}

// TestForPage tests the ForPage method of the Collection struct
func TestForPage(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the first page with 2 elements per page
	pageColl := coll.ForPage(1, 2)

	// check if the length of the page collection is correct
	if len(pageColl.value) != 2 {
		t.Errorf("ForPage did not return the correct number of elements for the first page")
	}

	// check if the page collection contains the correct elements
	if pageColl.value[0] != 1 || pageColl.value[1] != 2 {
		t.Errorf("ForPage did not return the correct elements for the first page")
	}

	// get the second page with 2 elements per page
	pageColl = coll.ForPage(2, 2)

	// check if the length of the page collection is correct
	if len(pageColl.value) != 2 {
		t.Errorf("ForPage did not return the correct number of elements for the second page")
	}

	// check if the page collection contains the correct elements
	if pageColl.value[0] != 3 || pageColl.value[1] != 4 {
		t.Errorf("ForPage did not return the correct elements for the second page")
	}

	// get the third page with 2 elements per page
	pageColl = coll.ForPage(3, 2)

	// check if the length of the page collection is correct
	if len(pageColl.value) != 1 {
		t.Errorf("ForPage did not return the correct number of elements for the third page")
	}

	// check if the page collection contains the correct elements
	if pageColl.value[0] != 5 {
		t.Errorf("ForPage did not return the correct elements for the third page")
	}
}

// TestNth tests the Nth method of the Collection struct
func TestNth(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get every second element starting from the first
	nthColl := coll.Nth(2, 0)

	// check if the length of the nth collection is correct
	if len(nthColl.value) != 3 {
		t.Errorf("Nth did not return the correct number of elements")
	}

	// check if the nth collection contains the correct elements
	if nthColl.value[0] != 1 || nthColl.value[1] != 3 || nthColl.value[2] != 5 {
		t.Errorf("Nth did not return the correct elements")
	}

	// get every second element starting from the second
	nthColl = coll.Nth(2, 1)

	// check if the length of the nth collection is correct
	if len(nthColl.value) != 2 {
		t.Errorf("Nth did not return the correct number of elements")
	}

	// check if the nth collection contains the correct elements
	if nthColl.value[0] != 2 || nthColl.value[1] != 4 {
		t.Errorf("Nth did not return the correct elements")
	}
}

// TestPad tests the Pad method of the Collection struct
func TestPad(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// pad the collection with 2 elements
	paddedColl := coll.Pad(5, 0)

	// check if the length of the padded collection is correct
	if len(paddedColl.value) != 5 {
		t.Errorf("Pad did not return the correct number of elements")
	}

	// check if the padded collection contains the correct elements
	if paddedColl.value[0] != 1 || paddedColl.value[1] != 2 || paddedColl.value[2] != 3 || paddedColl.value[3] != 0 || paddedColl.value[4] != 0 {
		t.Errorf("Pad did not return the correct elements")
	}
}

// TestPop tests the Pop method of the Collection struct
func TestPop(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// pop the last element
	popped := coll.Pop()

	// check if the popped element is correct
	if popped != 3 {
		t.Errorf("Pop did not return the correct popped element")
	}

	// check if the length of the collection is correct
	if len(coll.value) != 2 {
		t.Errorf("Pop did not modify the length of the collection correctly")
	}

	// check if the collection contains the correct elements
	if coll.value[0] != 1 || coll.value[1] != 2 {
		t.Errorf("Pop did not modify the elements of the collection correctly")
	}

	// create a new empty Collection
	emptyColl := NewEmptyCollection[int]()

	// try to pop an element from an empty Collection
	popped = emptyColl.Pop()

	// check if the popped element is the zero value of the type
	if popped != 0 {
		t.Errorf("Pop did not return the zero value of the type for an empty Collection")
	}
}

// TestPush tests the Push method of the Collection struct
func TestPush(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// push a new element to the collection
	coll.Push(4)

	// check if the length of the collection is correct
	if len(coll.value) != 4 {
		t.Errorf("Push did not modify the length of the collection correctly")
	}

	// check if the collection contains the correct elements
	if coll.value[0] != 1 || coll.value[1] != 2 || coll.value[2] != 3 || coll.value[3] != 4 {
		t.Errorf("Push did not modify the elements of the collection correctly")
	}
}

// TestPrepend tests the Prepend method of the Collection struct
func TestPrepend(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// prepend a new element to the collection
	coll.Prepend(0)

	// check if the length of the collection is correct
	if len(coll.value) != 4 {
		t.Errorf("Prepend did not modify the length of the collection correctly")
	}

	// check if the collection contains the correct elements
	if coll.value[0] != 0 || coll.value[1] != 1 || coll.value[2] != 2 || coll.value[3] != 3 {
		t.Errorf("Prepend did not modify the elements of the collection correctly")
	}
}

// TestRandom tests the Random method of the Collection struct
func TestRandom(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// get a random element from the collection
	random := coll.Random()

	// check if the random element is in the collection
	if random != 1 && random != 2 && random != 3 {
		t.Errorf("Random did not return an element from the collection")
	}
}

// TestReverse tests the Reverse method of the Collection struct
func TestReverse(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// reverse the collection
	reversedColl := coll.Reverse()

	// check if the length of the reversed collection is correct
	if len(reversedColl.value) != 3 {
		t.Errorf("Reverse did not return the correct number of elements")
	}

	// check if the reversed collection contains the correct elements
	if reversedColl.value[0] != 3 || reversedColl.value[1] != 2 || reversedColl.value[2] != 1 {
		t.Errorf("Reverse did not return the correct elements")
	}
}

// TestShuffle tests the Shuffle method of the Collection struct
func TestShuffle(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// shuffle the collection
	shuffledColl := coll.Shuffle()

	// check if the length of the shuffled collection is correct
	if len(shuffledColl.value) != 3 {
		t.Errorf("Shuffle did not return the correct number of elements")
	}

	// check if the shuffled collection contains the same elements as the original collection
	if !coll.Contains(shuffledColl.value[0]) || !coll.Contains(shuffledColl.value[1]) || !coll.Contains(shuffledColl.value[2]) {
		t.Errorf("Shuffle did not return the correct elements")
	}
}

// TestGroupBy tests the GroupBy method of the Collection struct
func TestGroupBy(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// define a function that groups even and odd numbers
	groupBy := func(item int, key int) interface{} {
		if item%2 == 0 {
			return "even"
		} else {
			return "odd"
		}
	}

	// group the collection by even and odd numbers
	groupedColl := coll.GroupBy(groupBy)

	// check if the length of the grouped collection is correct
	if len(groupedColl) != 2 {
		t.Errorf("GroupBy did not return the correct number of groups")
	}

	// check if the grouped collection contains the correct elements
	if len(groupedColl["even"].value) != 2 || len(groupedColl["odd"].value) != 3 {
		t.Errorf("GroupBy did not return the correct elements")
	}
}

// TestSplit tests the Split method of the Collection struct
func TestSplit(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// split the collection into two parts
	splitColl := coll.Split(2)

	// check if the length of the split collection is correct
	if len(splitColl) != 3 {
		t.Errorf("Split did not return the correct number of parts")
	}

	// check if the split collection contains the correct elements
	if len(splitColl[0].value) != 2 || len(splitColl[1].value) != 2 || len(splitColl[2].value) != 1 {
		t.Errorf("Split did not return the correct elements")
	}
}

// TestDD tests the DD method of the Collection struct
func TestDD(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3})

	// call the DD method
	coll.DD()

	// the DD method should not return anything, so there is nothing to check
}

// TestPluckString tests the PluckString method of the Collection struct
func TestPluckString(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]string{
		{"name": "Alice", "age": "20"},
		{"name": "Bob", "age": "30"},
		{"name": "Charlie", "age": "40"},
	})

	// pluck the "name" field from the collection
	pluckedColl := coll.PluckString("name")

	// check if the length of the plucked collection is correct
	if len(pluckedColl.value) != 3 {
		t.Errorf("PluckString did not return the correct number of elements")
	}

	// check if the plucked collection contains the correct elements
	if pluckedColl.value[0] != "Alice" || pluckedColl.value[1] != "Bob" || pluckedColl.value[2] != "Charlie" {
		t.Errorf("PluckString did not return the correct elements")
	}
}

// TestPluckInt64 tests the PluckInt64 method of the Collection struct
func TestPluckInt64(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]int64{
		{"id": 1, "age": 20},
		{"id": 2, "age": 30},
		{"id": 3, "age": 40},
	})

	// pluck the "id" field from the collection
	pluckedColl := coll.PluckInt64("id")

	// check if the length of the plucked collection is correct
	if len(pluckedColl.value) != 3 {
		t.Errorf("PluckInt64 did not return the correct number of elements")
	}

	// check if the plucked collection contains the correct elements
	if pluckedColl.value[0] != 1 || pluckedColl.value[1] != 2 || pluckedColl.value[2] != 3 {
		t.Errorf("PluckInt64 did not return the correct elements")
	}
}

// TestPluckFloat64 tests the PluckFloat64 method of the Collection struct
func TestPluckFloat64(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]float64{
		{"price": 1.99, "quantity": 10},
		{"price": 2.99, "quantity": 20},
		{"price": 3.99, "quantity": 30},
	})

	// pluck the "price" field from the collection
	pluckedColl := coll.PluckFloat64("price")

	// check if the length of the plucked collection is correct
	if len(pluckedColl.value) != 3 {
		t.Errorf("PluckFloat64 did not return the correct number of elements")
	}

	// check if the plucked collection contains the correct elements
	if pluckedColl.value[0] != 1.99 || pluckedColl.value[1] != 2.99 || pluckedColl.value[2] != 3.99 {
		t.Errorf("PluckFloat64 did not return the correct elements")
	}
}

// TestPluckUint64 tests the PluckUint64 method of the Collection struct
func TestPluckUint64(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]uint64{
		{"id": 1, "age": 20},
		{"id": 2, "age": 30},
		{"id": 3, "age": 40},
	})

	// pluck the "id" field from the collection
	pluckedColl := coll.PluckUint64("id")

	// check if the length of the plucked collection is correct
	if len(pluckedColl.value) != 3 {
		t.Errorf("PluckUint64 did not return the correct number of elements")
	}

	// check if the plucked collection contains the correct elements
	if pluckedColl.value[0] != 1 || pluckedColl.value[1] != 2 || pluckedColl.value[2] != 3 {
		t.Errorf("PluckUint64 did not return the correct elements")
	}
}

// TestPluckBool tests the PluckBool method of the Collection struct
func TestPluckBool(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]bool{
		{"is_active": true, "is_admin": false},
		{"is_active": false, "is_admin": true},
		{"is_active": true, "is_admin": true},
	})

	// pluck the "is_active" field from the collection
	pluckedColl := coll.PluckBool("is_active")

	// check if the length of the plucked collection is correct
	if len(pluckedColl.value) != 3 {
		t.Errorf("PluckBool did not return the correct number of elements")
	}

	// check if the plucked collection contains the correct elements
	if pluckedColl.value[0] != true || pluckedColl.value[1] != false || pluckedColl.value[2] != true {
		t.Errorf("PluckBool did not return the correct elements")
	}
}

// TestSortBy tests the SortBy method of the Collection struct
func TestSortBy(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]int{
		{"id": 3, "age": 20},
		{"id": 1, "age": 30},
		{"id": 2, "age": 40},
	})

	// sort the collection by the "id" field
	sortedColl := coll.SortBy("id")

	// check if the length of the sorted collection is correct
	if len(sortedColl.value) != 3 {
		t.Errorf("SortBy did not return the correct number of elements")
	}

	// check if the sorted collection contains the correct elements
	if sortedColl.value[0]["id"] != 1 || sortedColl.value[1]["id"] != 2 || sortedColl.value[2]["id"] != 3 {
		t.Errorf("SortBy did not return the correct elements")
	}
}

// TestSortByDesc tests the SortByDesc method of the Collection struct
func TestSortByDesc(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]int{
		{"id": 3, "age": 20},
		{"id": 1, "age": 30},
		{"id": 2, "age": 40},
	})

	// sort the collection by the "id" field in descending order
	sortedColl := coll.SortByDesc("id")

	// check if the length of the sorted collection is correct
	if len(sortedColl.value) != 3 {
		t.Errorf("SortByDesc did not return the correct number of elements")
	}

	// check if the sorted collection contains the correct elements
	if sortedColl.value[0]["id"] != 3 || sortedColl.value[1]["id"] != 2 || sortedColl.value[2]["id"] != 1 {
		t.Errorf("SortByDesc did not return the correct elements")
	}
}

// TestKeyByStrField tests the KeyByStrField method of the Collection struct
func TestKeyByStrField(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]map[string]string{
		{"name": "Alice", "age": "20"},
		{"name": "Bob", "age": "30"},
		{"name": "Charlie", "age": "40"},
	})

	// key the collection by the "name" field
	keyedColl, err := coll.KeyByStrField("name")

	// check if the error is nil
	if err != nil {
		t.Errorf("KeyByStrField returned an error")
	}

	// check if the length of the keyed collection is correct
	if len(keyedColl) != 3 {
		t.Errorf("KeyByStrField did not return the correct number of elements")
	}

	// check if the keyed collection contains the correct elements
	if _, ok := keyedColl["Alice"]; !ok {
		t.Errorf("KeyByStrField did not return the correct elements")
	}

	if _, ok := keyedColl["Bob"]; !ok {
		t.Errorf("KeyByStrField did not return the correct elements")
	}
}

// TestMax tests the Max method of the Collection struct
func TestMax(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the maximum element from the collection
	max := coll.Max()

	// check if the maximum element is correct
	if max != 5 {
		t.Errorf("Max did not return the correct element")
	}
}

// TestMin tests the Min method of the Collection struct
func TestMin(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the minimum element from the collection
	min := coll.Min()

	// check if the minimum element is correct
	if min != 1 {
		t.Errorf("Min did not return the correct element")
	}
}

// TestContains tests the Contains method of the Collection struct
func TestContains(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// check if the collection contains the element 3
	if !coll.Contains(3) {
		t.Errorf("Contains did not return the correct result")
	}

	// check if the collection contains the element 6
	if coll.Contains(6) {
		t.Errorf("Contains did not return the correct result")
	}
}

// TestContainsCount tests the ContainsCount method of the Collection struct
func TestContainsCount(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5, 3})

	// check the count of the element 3 in the collection
	count := coll.ContainsCount(3)

	// check if the count is correct
	if count != 2 {
		t.Errorf("ContainsCount did not return the correct count")
	}
}

// TestDiff tests the Diff method of the Collection struct
func TestDiff(t *testing.T) {
	// create a new Collection with some elements
	coll1 := NewCollection([]int{1, 2, 3, 4, 5})
	coll2 := NewCollection([]int{3, 4, 5, 6, 7})

	// get the difference between the two collections
	diffColl := coll1.Diff(coll2)

	// check if the length of the difference collection is correct
	if len(diffColl.value) != 2 {
		t.Errorf("Diff did not return the correct number of elements")
	}

	// check if the difference collection contains the correct elements
	if diffColl.value[0] != 1 || diffColl.value[1] != 2 {
		t.Errorf("Diff did not return the correct elements")
	}
}

// TestSort tests the Sort method of the Collection struct
func TestSort(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{3, 1, 4, 2, 5})

	// sort the collection
	sortedColl := coll.Sort()

	// check if the length of the sorted collection is correct
	if len(sortedColl.value) != 5 {
		t.Errorf("Sort did not return the correct number of elements")
	}

	// check if the sorted collection contains the correct elements
	if sortedColl.value[0] != 1 || sortedColl.value[1] != 2 || sortedColl.value[2] != 3 || sortedColl.value[3] != 4 || sortedColl.value[4] != 5 {
		t.Errorf("Sort did not return the correct elements")
	}
}

// TestSortDesc tests the SortDesc method of the Collection struct
func TestSortDesc(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{3, 1, 4, 2, 5})

	// sort the collection in descending order
	sortedColl := coll.SortDesc()

	// check if the length of the sorted collection is correct
	if len(sortedColl.value) != 5 {
		t.Errorf("SortDesc did not return the correct number of elements")
	}

	// check if the sorted collection contains the correct elements
	if sortedColl.value[0] != 5 || sortedColl.value[1] != 4 || sortedColl.value[2] != 3 || sortedColl.value[3] != 2 || sortedColl.value[4] != 1 {
		t.Errorf("SortDesc did not return the correct elements")
	}
}

// TestJoin tests the Join method of the Collection struct
func TestJoin(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]string{"apple", "banana", "cherry"})

	// join the elements of the collection with a comma separator
	joinedStr := coll.Join(",")

	// check if the joined string is correct
	if joinedStr != "apple,banana,cherry" {
		t.Errorf("Join did not return the correct string")
	}

	// join the elements of the collection with a comma separator and a custom format function
	joinedStr = coll.Join(",", func(item interface{}) string {
		return strings.ToUpper(item.(string))
	})

	// check if the joined string with custom format is correct
	if joinedStr != "APPLE,BANANA,CHERRY" {
		t.Errorf("Join did not return the correct string with custom format")
	}
}

// TestUnion tests the Union method of the Collection struct
func TestUnion(t *testing.T) {
	// create two new Collections with some elements
	coll1 := NewCollection([]int{1, 2, 3})
	coll2 := NewCollection([]int{3, 4, 5})

	// get the union of the two collections
	unionColl := coll1.Union(coll2)

	// check if the length of the union collection is correct
	if len(unionColl.value) != 5 {
		t.Errorf("Union did not return the correct number of elements")
	}

	// check if the union collection contains the correct elements
	if unionColl.value[0] != 1 || unionColl.value[1] != 2 || unionColl.value[2] != 3 || unionColl.value[3] != 4 || unionColl.value[4] != 5 {
		t.Errorf("Union did not return the correct elements")
	}
}

// TestIntersect tests the Intersect method of the Collection struct
func TestIntersect(t *testing.T) {
	// create two new Collections with some elements
	coll1 := NewCollection([]int{1, 2, 3})
	coll2 := NewCollection([]int{3, 4, 5})

	// get the intersection of the two collections
	intersectColl := coll1.Intersect(coll2)

	// check if the length of the intersection collection is correct
	if len(intersectColl.value) != 1 {
		t.Errorf("Intersect did not return the correct number of elements")
	}

	// check if the intersection collection contains the correct element
	if intersectColl.value[0] != 3 {
		t.Errorf("Intersect did not return the correct element")
	}
}

// TestAvg tests the Avg method of the Collection struct
func TestAvg(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the average of the elements in the collection
	avg := coll.Avg()

	// check if the average is correct
	if avg != 3 {
		t.Errorf("Avg did not return the correct value")
	}
}

// TestMedian tests the Median method of the Collection struct
func TestMedian(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the median of the elements in the collection
	median := coll.Median()

	// check if the median is correct
	if median != 3 {
		t.Errorf("Median did not return the correct value")
	}
}

// TestMode tests the Mode method of the Collection struct
func TestMode(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5})

	// get the mode of the elements in the collection
	mode := coll.Mode()

	// check if the mode is correct
	if mode != 4 {
		t.Errorf("Mode did not return the correct value")
	}
}

// TestSum tests the Sum method of the Collection struct
func TestSum(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the sum of the elements in the collection
	sum := coll.Sum()

	// check if the sum is correct
	if sum != 15 {
		t.Errorf("Sum did not return the correct value")
	}
}

// TestValues tests the Values method of the Collection struct
func TestValues(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]int{1, 2, 3, 4, 5})

	// get the values of the collection
	values := coll.Values()

	// check if the length of the values slice is correct
	if len(values) != 5 {
		t.Errorf("Values did not return the correct number of elements")
	}

	// check if the values slice contains the correct elements
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Errorf("Values did not return the correct elements")
	}
}

// TestToJson tests the ToJson method of the Collection struct
func TestToJson(t *testing.T) {
	// create a new Collection with some elements
	coll := NewCollection([]string{"apple", "banana", "cherry"})

	// convert the collection to JSON
	jsonData, err := coll.ToJson()

	// check if the conversion was successful
	if err != nil {
		t.Errorf("ToJson returned an error: %v", err)
	}

	// check if the JSON data is correct
	expectedJson := `["apple","banana","cherry"]`
	if string(jsonData) != expectedJson {
		t.Errorf("ToJson did not return the correct JSON data")
	}
}

// TestFromJson tests the FromJson method of the Collection struct
func TestFromJson(t *testing.T) {
	// create a JSON string
	jsonData := `["apple","banana","cherry"]`

	// create a new empty Collection
	coll := NewCollection([]string{})

	// populate the collection from the JSON data
	err := coll.FromJson([]byte(jsonData))

	// check if the population was successful
	if err != nil {
		t.Errorf("FromJson returned an error: %v", err)
	}

	// check if the collection contains the correct elements
	if len(coll.value) != 3 || coll.value[0] != "apple" || coll.value[1] != "banana" || coll.value[2] != "cherry" {
		t.Errorf("FromJson did not populate the collection correctly")
	}
}
