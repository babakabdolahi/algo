package queue

import "errors"

type Queue struct {
	Max   int
	Front int
	Rear  int
	Data  []int
}

func Init(max int) *Queue {
	q := &Queue{
		Max:   max,
		Front: -1,
		Rear:  -1,
	}

	q.Data = make([]int, max)

	return q
}

func Insert(q *Queue, v int) error {
	if q.Rear == q.Max-1 {
		return errors.New("queue overflow")
	}

	q.Rear++

	q.Data[q.Rear] = v

	if q.Front == -1 {
		q.Front++
	}

	return nil
}

func DeleteElement(q *Queue) (int, error) {
	if q.Front == -1 || q.Front > q.Rear {
		return 0, errors.New("queue underflow")
	}

	v := q.Data[q.Front]

	q.Front++

	return v, nil
}

func Peek(q *Queue) int {
	return q.Data[q.Front]
}
