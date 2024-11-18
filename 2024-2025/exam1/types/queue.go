package types

type Queue []*Node

func (q *Queue) Len() int { return len(*q) }

func (q *Queue) Less(i, j int) bool {
	return (*q)[i].Distance < (*q)[j].Distance
}

func (q *Queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *Queue) Push(x *Node) {
	if q.Len() == 0 {
		*q = append(*q, x)
		return
	}
	var insertFlag bool
	for k, v := range *q {
		if x.Distance < v.Distance {
			if k > 0 {
				*q = append((*q)[:k+1], (*q)[k:]...)
				(*q)[k] = x
				insertFlag = true
			} else {
				*q = append([]*Node{x}, *q...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		*q = append(*q, x)
	}
}

func (q *Queue) Pop() *Node {
	firstEl := (*q)[0]
	*q = (*q)[1:q.Len()]
	return firstEl
}
