package leetcode

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 6},
						},
					},
				},
			},
		},
	}
	exp := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 6},
					},
				},
			},
		},
	}

	if r := deleteMiddle(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4},
			},
		},
	}
	exp = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	if r := deleteMiddle(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	l = &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	exp = &ListNode{Val: 2}
	if r := deleteMiddle(l); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
