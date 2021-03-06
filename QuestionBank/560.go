/*
 * @Time     : 2020/5/16 20:58
 * @Author   : cancan
 * @File     : 560.go
 * @Function : 和为K的子数组
 */

/*
 * Question:
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 *
 * Example 1 :
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 *
 * Note :
 * 1.数组的长度为 [1, 20,000]。
 * 2.数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 */

package QuestionBank

func subarraySum(nums []int, k int) int {
	count := 0
	tmp := map[int]int{}
	tmp[0] = 1
	pre := 0
	for _, v := range nums {
		pre += v
		if _, ok := tmp[pre-k]; ok {
			count += tmp[pre-k]
		}
		tmp[pre] += 1
	}
	return count
}
