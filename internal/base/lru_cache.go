package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Length() int {
	size := 0

	if l.Head == nil {
		return size
	}

	size = 1
	current := l.Head
	for current.Prev != nil {
		current = current.Prev
		size++
	}

	return size
}

func (l *LruCache) Put(key string, value string) {
	existing := l.GetNode(key)
	if existing != nil {
		existing.Value = value
		return
	}

	node := &Node{
		Key:   key,
		Value: value,
	}

	if l.Tail == nil {
		l.Tail = node
	}

	if l.Head == nil {
		l.Head = node
	}

	if l.Tail.Value != node.Value {
		node.Prev = l.Head
	}

	l.Head.Next = node
	l.Head = node

	if l.Length() > l.size {
		l.Tail = l.Tail.Next
		l.Tail.Prev = nil
	}
}

func (l *LruCache) Get(key string) *string {
	if l.Head == nil {
		return nil
	}
	current := l.Head

	for current != nil {
		if current.Key == key {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
				current.Prev = l.Head
				l.Head = current
			}

			return &current.Value
		}
		current = current.Prev
	}

	return nil
}

func (l *LruCache) GetNode(key string) *Node {
	if l.Head == nil {
		return nil
	}
	current := l.Head

	for current != nil {
		if current.Key == key {
			return current
		}
		current = current.Prev
	}

	return nil
}
