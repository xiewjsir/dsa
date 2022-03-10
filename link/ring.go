package link

type Ring struct {
	Value      interface{}
	next, prev *Ring
}

type iRing interface {
	init() *Ring
	Next() *Ring
	Prev() *Ring
	Link(ring *Ring) *Ring
	UnLink(n int64) *Ring
	Len() int64
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}

	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}

	return r.prev
}

func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}

	if n > 0 {
		for ; n > 0; n-- {
			r = r.Next()
		}
	} else {
		for ; n < 0; n++ {
			r = r.Prev()
		}
	}

	return r
}

func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()

		r.next = s
		s.prev = r

		p.next = n
		n.prev = p
	}

	return n
}

func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}

	return r.Link(r.Move(n + 1))
}

func (r *Ring) Len() (n int64) {
	if r == nil {
		return
	}

	p := r
	for {
		n++
		if p.next == r {
			break
		}
		p = p.next
	}
	return
}

func NewRing(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := new(Ring)
	r.Value = 0
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p, Value: i}
		p = p.next
	}
	p.next = r
	r.prev = p

	return r
}
