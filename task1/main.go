package main

import (
	"fmt"
	"strings"
)

func main() {
	// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
	// 其余每个元素均出现两次。找出那个只出现了一次的元素。
	// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	singleNumber()
	// 9. 回文数：判断一个整数是否是回文数
	// 可以使用将整数转换为字符串，然后使用双指针法判断是否是回文数。
	palindromeNumber()
	// 20. 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	// 可以使用栈来解决，遍历字符串，如果遇到左括号，则将其压入栈中，如果遇到右括号，则从栈中弹出元素并判断是否匹配。
	validParentheses()
	// 14. 最长公共前缀：查找字符串数组中的最长公共前缀
	// 可以使用双指针法，遍历字符串数组，找到最长公共前缀。
	longestCommonPrefix()
	// 66. 加一：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	// 可以使用从后往前遍历数组，如果当前位是9，则需要进位，否则直接加一。
	plusOne()
	// 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
	removeDuplicatesFromSortedArray()
	// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中
	mergeIntervals()
	// 1. 两数之和：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。你可以按任意顺序返回答案。
	// 可以使用两层 for 循环，遍历数组，如果两个数的和为目标值，则返回它们的下标。
	twoSum()
}

/*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber() {
	nums := []int{4, 1, 2, 1, 2}
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			fmt.Println(num)
			return
		}
	}
}

/*
回文数
考察：数字操作、条件判断
题目：判断一个整数是否是回文数*
*/
func palindromeNumber() {
	num := 12121
	if num < 0 {
		fmt.Println("不是回文数")
		return
	}
	originalNum := num
	reversedNum := 0
	for num > 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}
	if originalNum == reversedNum {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文数")
	}
}

/*
有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
*/
func validParentheses() bool {
	s := "([)]"
	l := len(s)
	if l%2 != 0 {
		fmt.Println("无效")
		return false
	}
	stack := []string{}
	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			if s[i] == '(' || s[i] == '{' || s[i] == '[' {
				stack = append(stack, string(s[i]))
				continue
			} else {
				break
			}
		} else {
			p := stack[len(stack)-1]
			if p == "(" && s[i] == ')' || p == "{" && s[i] == '}' || p == "[" && s[i] == ']' {
				stack = stack[:len(stack)-1]
				continue
			}
			if s[i] == '(' || s[i] == '{' || s[i] == '[' {
				stack = append(stack, string(s[i]))
				continue
			} else {
				break
			}
		}
	}
	if len(stack) == 0 {
		fmt.Println("有效")
		return true
	} else {
		fmt.Println("无效")
		return false
	}
}

/*
最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
*/
func longestCommonPrefix() {
	strs := []string{"flower", "flow", "flowight"}
	minLen := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	var prefixs []string
Label:
	for j := 0; j < minLen; j++ {
		for i := 0; i < len(strs); i++ {
			if len(prefixs) < j+1 {
				prefixs = append(prefixs, string(strs[i][j]))
			} else {
				if prefixs[len(prefixs)-1] != string(strs[i][j]) {
					prefixs = prefixs[:len(prefixs)-1]
					break Label
				} else {
					continue
				}
			}
		}
	}
	fmt.Println(strings.Join(prefixs, ""))
}

/*
*
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne() []int {
	digits := []int{9, 9, 9}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		p := 1 * digits[len(digits)-i-1]
		for j := 1; j < i+1; j++ {
			p *= 10
		}
		carry += p
	}
	fmt.Println(digits)
	l := len(digits)
	if carry%10 == 0 {
		l++
	}
	result := make([]int, l)
	i := l - 1
	for carry > 0 {
		result[i] = carry % 10
		carry /= 10
		i--
	}
	fmt.Println(result)
	return result
}

/*
*
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
*/
func removeDuplicatesFromSortedArray() {
	nums := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(nums)
	slen := len(nums)
	for i := 0; i < slen; i++ {
		j := i + 1
		k := j
		for {
			if j < slen && nums[i] == nums[j] {
				for ; k < slen; k++ {
					nums[k-1] = nums[k]
				}
				k = j
				slen--
			} else {
				break
			}
		}
	}
	fmt.Println(nums[:slen])
}

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中
*/
func mergeIntervals() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {10, 14}, {15, 18}, {1, 4}, {4, 5}, {16, 20}}
	var _merge = func(intervals [][]int) ([][]int, bool) {
		for i := 0; i < len(intervals); i++ {
			for j := i + 1; j < len(intervals); j++ {
				c, ok := doMerge(intervals[i], intervals[j])
				if ok {
					var t [][]int
					t = append(t, c)
					for _i := 0; _i < len(intervals); _i++ {
						if _i == i || _i == j {
							continue
						}
						t = append(t, intervals[_i])
					}
					return t, true
				}
			}
		}
		return intervals, false
	}
	ok := true
	for ok {
		intervals, ok = _merge(intervals)
	}
	fmt.Println(intervals)
}

func doMerge(a []int, b []int) ([]int, bool) {
	if a[0] <= b[1] && a[0] >= b[0] || b[0] <= a[1] && b[1] >= a[0] { // 说明有交集
		min := a[0]
		max := a[1]
		for i := 0; i < len(b); i++ {
			if b[i] < min {
				min = b[i]
			}
			if b[i] > max {
				max = b[i]
			}
			if a[i] < min {
				min = a[i]
			}
			if a[i] > max {
				max = a[i]
			}
		}

		// fmt.Println(min, max)
		return []int{min, max}, true
	}
	return nil, false
}

/*
两数之和：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
func twoSum() []int {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
				return []int{i, j}
			}
		}
	}
	return nil
}
