package src

import "sync"

type skipListInterface interface {
	Insert(Node)
	Remove(int2 int)
	Top(int2 int)[]Node
	Range(condition Condition)[]Node
}

type SkipList struct {
	Head *ListNode
	Tail *ListNode
	Length int
	Size int
	sync.RWMutex
}

func (s *SkipList)Remove(score int64){
	s.Lock()
	defer s.Unlock()



}

func (s *SkipList)Insert(node Node){
	s.Lock()
	defer s.Unlock()


}

func (s *SkipList)Top(k int) (values []Node) {
	s.RLock()
	defer s.RUnlock()

	if k <=0 {
		return
	}

	if k > s.Length{
		k = s.Length
	}

	var node *ListNode
	for i:=0;i<k;i++{
		if i == 0{
			node = s.Head
		}

		if node == nil{
			return
		}

		values = append(values, node.Value)
		node = node.Next
	}

	return
}

type Condition struct {
	Start int
	Length int
	Reverse bool
}



