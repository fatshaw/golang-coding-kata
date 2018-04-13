package coding_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumSubArrayAllCorrect(t *testing.T) {
	t.Run("only one positive element in the array", func(t *testing.T) {
		arr := []int{1}
		assert.Equal(t, 1, MaxSumOfSubArray(arr))
	})

	t.Run("only one nagetive element in the array", func(t *testing.T) {
		arr := []int{-1}
		assert.Equal(t, -1, MaxSumOfSubArray(arr))
	})

	t.Run("two elements in the array", func(t *testing.T) {
		arr := []int{1, 2}
		assert.Equal(t, 3, MaxSumOfSubArray(arr))
	})

	t.Run("negative elements in the array", func(t *testing.T) {
		arr := []int{1, 2, -1}
		assert.Equal(t, 3, MaxSumOfSubArray(arr))
	})

	t.Run("negative elements in the array", func(t *testing.T) {
		arr := []int{1, 2, -1, 5}
		assert.Equal(t, 7, MaxSumOfSubArray(arr))
	})

	t.Run("sub array in the array", func(t *testing.T) {
		arr := []int{-1, 2, -1, 5, -10}
		assert.Equal(t, 6, MaxSumOfSubArray(arr))
	})
}
