package leetcode

// Queue is the queue
type Queue struct {
	queue []int
	size  int
}

func (q *Queue) add(val int) {
	if len(q.queue) == q.size {
		return
	}
	q.queue = append(q.queue, val)
	q.size++
}

func (q *Queue) pop() int {
	if len(q.queue) == 0 {
		return -1
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return result
}
