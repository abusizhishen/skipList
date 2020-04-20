package src

type Node struct {
	Key string `json:"key"`
	Score int64 `json:"score"`
	Value interface{} `json:"value"`
}

type ListNode struct {
	Value Node `json:"value"`
	Pre *ListNode
	Next *ListNode
}
