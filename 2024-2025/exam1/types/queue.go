package types

type Queue []*Node

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	return q[i].Distance < q[j].Distance
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x *Node) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	oldQueue := *q
	lengthOld := len(oldQueue)
	lastEl := oldQueue[lengthOld-1]
	*q = oldQueue[0 : lengthOld-1]
	return lastEl
}
