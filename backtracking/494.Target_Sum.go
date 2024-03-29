// package golang
// name 494.(dfs)Target_Sum

package backtracking

/*
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols +
and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
Note:
The length of the given array is positive and will not exceed 20.
The sum of elements in the given array will not exceed 1000.
Your output answer is guaranteed to be fitted in a 32-bit integer.
*/

var targetSumCount = 0

func dfsTargetSum(target, index, sum int, numbers []int) {
	if index == len(numbers) {
		if sum == target {
			targetSumCount++
		}
		return
	}

	dfsTargetSum(target, index+1, sum+numbers[index], numbers)
	dfsTargetSum(target, index+1, sum-numbers[index], numbers)
}

func TargetSum(numbers []int, target int) int {
	dfsTargetSum(target, 0, 0, numbers)
	defer func() {
		targetSumCount = 0
	}()
	return targetSumCount
}
