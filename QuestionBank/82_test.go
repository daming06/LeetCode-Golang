// @Time     : 2019/4/28 11:23
// @Author   : cancan
// @File     : 82_test.go
// @Function : 删除排序链表中的重复元素 II tet

package QuestionBank

import "testing"

func Test_82(t *testing.T) {
	tests := []struct {
		head *ListNode
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 2, 3, 3, 4, 4, 5}),
			InitSinglyLinkList([]int{1, 2, 5}),
		},
		{
			InitSinglyLinkList([]int{1, 1, 1, 2, 3}),
			InitSinglyLinkList([]int{2, 3}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(deleteDuplicatesII(test.head), test.ans) {
			t.Errorf("failure head %v ans %v",
				SinglyLinkList2Slice(test.head),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
