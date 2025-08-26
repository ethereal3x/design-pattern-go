package objectpool

type Object struct {
	ID int
}

type ObjectPool struct {
	objects chan *Object
}

func NewObjectPool(size int) *ObjectPool {
	pool := &ObjectPool{
		objects: make(chan *Object, size),
	}
	for i := 0; i < size; i++ {
		obj := &Object{
			ID: i,
		}
		pool.objects <- obj
	}
	return pool
}

func (p *ObjectPool) AcquireObject() *Object {
	return <-p.objects
}

func (p *ObjectPool) ReleaseObject(object *Object) {
	p.objects <- object
}
