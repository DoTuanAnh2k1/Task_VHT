package model

type Pool struct {
	objects     *LinkedList
	initialSize int
}

func NewObjectPool(initialSize int) *Pool {
	pool := &Pool{
		objects:     NewLinkedList(),
		initialSize: initialSize,
	}
	for i := 0; i < initialSize; i++ {
		pool.objects.PushBack(&Item{})
	}
	return pool
}

func (p *Pool) Acquire(fId int, pri int64) *Item {
	if p.objects.Len() == 0 {
		return &Item{}
	}

	obj := p.objects.PopBack()
	obj.FileId = fId
	obj.Priority = pri
	return obj
}

func (p *Pool) Release(obj *Item) {
	p.objects.PushBack(obj)
}
