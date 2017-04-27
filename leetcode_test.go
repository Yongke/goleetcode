package goleetcode

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestP001(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	assert.Equal(t, twoSum(nums, target), []int{0, 1})
}

func TestP002(t *testing.T) {
	CheckP002(t, []int{3}, []int{7}, []int{0, 1})
	CheckP002(t, []int{1}, []int{9, 9}, []int{0, 0, 1})
	CheckP002(t, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8})
}

func CheckP002(t *testing.T, a []int, b []int, expect []int) {
	nl1 := makeNodeList(a)
	nl2 := makeNodeList(b)
	nl_new := addTwoNumbers(nl1, nl2)
	assert.Equal(t, NodeListToSlice(nl_new), expect)
}

func makeNodeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	return &ListNode{nums[0], makeNodeList(nums[1:])}
}

func NodeListToSlice(node *ListNode) []int {
	if node == nil {
		return []int{}
	}

	t := []int{node.Val}
	return append(t, NodeListToSlice(node.Next)...)
}

func TestP003(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bpfbhmipx"), 7)

}
