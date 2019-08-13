package Response

type Res struct {
	Code int
	Message interface{}
}

func (r *Res)SetCode(c int) *Res {
	r.Code = c
	return r
}

func (r *Res)SetMessage(m interface{}) *Res {
	r.Message = m
	return r
}
