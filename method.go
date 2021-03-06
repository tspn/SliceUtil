package sliceutils

func (c SliceUtils) Sort(f func (x, y interface{}) bool) SliceUtils{
	for i := 0; i < c.Len(); i++{
		for j := 0; j < c.Len(); j++{
			if f(c[i], c[j]){
				c.Swap(i, j)
			}
		}
	}
	return c
}

func (c SliceUtils) Swap(i, j int)SliceUtils{
	t := c[i]
	c[i] = c[j]
	c[j] = t

	return c
}

func (c SliceUtils) Reverse()SliceUtils{
	new := SliceUtils{}
	for _, d := range c{
		new = new.UnShift(d)
	}
	return new
}

func (c SliceUtils) Reduce(f func(value, sum interface{}) interface{}) interface{}{
	var v interface{}
	if c.Len() > 1 {
		v = c[0]
		for _, d := range c[1:]{
			v = f(d, v)
		}
	}else if c.Len() == 1{
		return c[0]
	}else{
		return nil
	}

	return v
}

func (c SliceUtils) Map (f func(i interface{}) interface{}) SliceUtils{
	new := SliceUtils{}
	for _, d := range c{
		new = append(new, f(d))
	}

	return new
}

func (c SliceUtils) Filter (f func(i interface{}) bool) SliceUtils{
	new := SliceUtils{}
	for _, d := range c{
		if f(d){
			new = append(new, d)
		}
	}

	return new
}

func (c SliceUtils) ForEach(f func (i interface{})) SliceUtils{
	for _, d := range c{
		f(d)
	}
	return c
}

func (c SliceUtils) UnShift(i interface{})SliceUtils{
	new := append(SliceUtils{i}, c...)
	return new
}

func (c SliceUtils) Shift() (SliceUtils, interface{}){
	if c.Len() > 0{
		t := c.Head()
		new := c[1:]
		return new, t
	}else{
		return c, nil
	}
}

func (c SliceUtils) DeQueue() (SliceUtils, interface{}){
	return c.Shift()
}

func (c SliceUtils) Queue(i interface{}) SliceUtils{
	new := c.Push(i)
	return new
}

func (c SliceUtils) Pop() (SliceUtils, interface{}){
	if c.Len() > 0{
		t := c.Tail()
		new := c[:c.Len() - 1]
		return new, t
	}else{
		return c, nil
	}
}

func (c SliceUtils) Push(i interface{}) SliceUtils{
	new := append(c, i)
	return new
}

func (c SliceUtils) Tail() interface{}{
	if len(c) > 0{
		return c[len(c) - 1]
	}else{
		return nil
	}
}

func (c SliceUtils) Head() interface{}{
	if len(c) > 0 {
		return c[0]
	}else{
		return nil
	}

}

func (c SliceUtils) Count() int{
	return c.Len()
}

func (c SliceUtils) Len() int{
	return len(c)
}