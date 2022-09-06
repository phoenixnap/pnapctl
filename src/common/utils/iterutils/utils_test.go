package iterutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSuccess(test_framework *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	doubled := Map(nums, func(n int) int { return n * 2 })

	assert.Equal(test_framework, doubled, []int{2, 4, 6, 8, 10})
}

func TestCurrySuccess(test_framework *testing.T) {
	multiply := func(num1 int, num2 int) int {
		return num1 * num2
	}

	double := Curry(multiply, 2)

	assert.Equal(test_framework, double(5), 10)
}

func TestFilterSuccess(test_framework *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })

	assert.Equal(test_framework, evens, []int{2, 4, 6})
}

func TestContainsSuccess(test_framework *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	hasFive := Contains(nums, 5)
	hasSix := Contains(nums, 6)

	assert.True(test_framework, hasFive)
	assert.False(test_framework, hasSix)
}

func TestFindElementThatSuccess(test_framework *testing.T) {
	nums := []int{1, 2, 50, 100, 20}

	moreThanTen := FindElementThat(nums, func(n int) bool { return n > 10 })
	moreThanThousand := FindElementThat(nums, func(n int) bool { return n > 1000 })

	assert.Equal(test_framework, moreThanTen, &nums[2])
	assert.Nil(test_framework, moreThanThousand)
}

func TestDerefSuccess(test_framework *testing.T) {
	num := 5

	pointers := []*int{&num, &num, &num}
	values := Deref(pointers)

	inputType := fmt.Sprintf("%T", pointers)
	outputType := fmt.Sprintf("%T", values)

	assert.Equal(test_framework, "[]*int", inputType)
	assert.Equal(test_framework, "[]int", outputType)
}

func TestDerefInterfaceSuccess(test_framework *testing.T) {
	num := 5

	pointers := []interface{}{&num, &num, &num}
	values := DerefInterface(pointers)

	_, ok1 := pointers[0].(*int)
	_, ok2 := values[0].(int)

	assert.True(test_framework, ok1)
	assert.True(test_framework, ok2)
}

func TestMapInterfaceSuccess(test_framework *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	doubled := MapInterface(nums, func(n int) int { return n * 2 })

	for i := range doubled {
		assert.Equal(test_framework, expected[i], doubled[i].(int))
	}
}

func TestNotPredicateSuccess(test_framework *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isOdd := Not(isEven)

	assert.True(test_framework, isEven(2))
	assert.False(test_framework, isOdd(2))
}

func TestIsNilPredicateSuccess(test_framework *testing.T) {
	var nilpointer *int

	assert.True(test_framework, IsNil(nil))
	assert.True(test_framework, IsNil(nilpointer))
	assert.False(test_framework, IsNil(0))
}
