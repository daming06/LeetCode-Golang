// @Time     : 2019/5/5 18:11
// @Author   : cancan
// @File     : 509.go
// @Function : 斐波那契数

/*
Question:
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

    F(0) = 0,   F(1) = 1
    F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

给定 N，计算 F(N)。

Example 1：
输入：2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2：
输入：3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3：
输入：4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3.

提示：
0 ≤ N ≤ 30
*/

package QuestionBank

func fib(N int) int {
	pre, cur := 0, 1

	for i := 0; i < N; i++ {
		pre, cur = cur, pre+cur
	}

	return pre
}
