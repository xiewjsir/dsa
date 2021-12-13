package linknode

type Ring struct {
	Value     interface{}
	next, pre *Ring
}

func (r *Ring) init() *Ring {
	r.next = r
	r.pre = r
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil{
		return r.init()
	}
	
	return r.next
}

func (r *Ring) Move(n int) *Ring{
	if r.next == nil{
		return r.init()
	}

	if n>0{
		for ;n>0;n--{
			r = r.next
		}
	}else{
		for ;n<0;n++{
			r = r.pre
		}
	}

	return r
}

func (r *Ring) Link(s *Ring) *Ring{
	n := r.Next()
	if s != nil{
		p := s.Prev()

		s.next = n
		s.pre = r

		p.next = n
		n.pre = p
	}

	return n
}

func (r *Ring) Prev() *Ring  {
	if r.next == nil{
		return r.init()
	}
	
	return r.pre
}

func New(n int) *Ring {
	if n<=0{
		return nil
	}
	
	r := new(Ring)
	p := r
	for i:=1;i<n;i++{
		p.next = &Ring{pre: p,Value: i}
		p = p.next
	}
	p.next = r
	r.pre = p
	
	return r
}