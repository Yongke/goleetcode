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
	assert.Equal(t, lengthOfLongestSubstring(""), 0)
	assert.Equal(t, lengthOfLongestSubstring("a"), 1)
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("abba"), 2)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bpfbhmipx"), 7)
	assert.Equal(t, lengthOfLongestSubstring("tmmzuxt"), 5)
}

func TestP007(t *testing.T) {
	assert.Equal(t, reverse(123), 321)
	assert.Equal(t, reverse(-123), -321)
	assert.Equal(t, reverse(-1534236469), 0)
	assert.Equal(t, reverse(1534236469), 0)
}

func TestP008(t *testing.T) {
	assert.Equal(t, myAtoi("123"), 123)
	assert.Equal(t, myAtoi("+0123"), 123)
	assert.Equal(t, myAtoi("-123"), -123)
	assert.Equal(t, myAtoi("9223372036854775809"), 2147483647)
}

func TestP009(t *testing.T) {
	assert.Equal(t, isPalindrome(0), true)
	assert.Equal(t, isPalindrome(-1), false)
	assert.Equal(t, isPalindrome(12344321), true)
	assert.Equal(t, isPalindrome(123454321), true)
	assert.Equal(t, isPalindrome(1233210), false)
}

func TestP012(t *testing.T) {
	assert.Equal(t, intToRoman(11), "XI")
	assert.Equal(t, intToRoman(501), "DI")
	assert.Equal(t, intToRoman(954), "CMLIV")
	assert.Equal(t, intToRoman(401), "CDI")
}

func TestP013(t *testing.T) {
	assert.Equal(t, romanToInt("CMLIV"),954)
}

func TestP344(t *testing.T) {
	assert.Equal(t, reverseString("hello"), "olleh")
}

func TestP461(t *testing.T) {
	assert.Equal(t, hammingDistance(1, 4), 2)
	assert.Equal(t, hammingDistance(0, 15), 4)
}

func TestP541(t *testing.T) {
	assert.Equal(t, reverseStr("abcdefg", 2), "bacdfeg")
	assert.Equal(t, reverseStr("abcabcab", 3), "cbaabcba")
}

func TestP557(t *testing.T) {
	assert.Equal(t, reverseWords("Let's take LeetCode contest"), "s'teL ekat edoCteeL tsetnoc")
}

func TestP560(t *testing.T) {
	assert.Equal(t, subarraySum([]int{1},1), 1)
	assert.Equal(t, subarraySum([]int{0,0,0,0,0,0,0,0,0,0},0), 55)
	assert.Equal(t, subarraySum([]int{-1,-1,1},0), 1)
	assert.Equal(t, subarraySum([]int{-1,-1,1},1), 1)
	assert.Equal(t, subarraySum([]int{1,1,1},2), 2)
	assert.Equal(t, subarraySum([]int{1,2,3,2,1},3), 3)
	assert.Equal(t, subarraySum([]int{-1,-1,-1},-2), 2)
}

func TestP561(t *testing.T) {
	assert.Equal(t, arrayPairSum([]int{1, 4, 3, 2}), 4)
}


func TestP566(t *testing.T) {
	matrix1 := [][]int{{1, 2}, {3, 4}}
	assert.Equal(t, matrixReshape(matrix1, 1, 4), [][]int{{1, 2, 3, 4}})
	assert.Equal(t, matrixReshape(matrix1, 4, 1), [][]int{{1}, {2}, {3}, {4}})
	assert.Equal(t, matrixReshape(matrix1, 4, 2), matrix1)
}

