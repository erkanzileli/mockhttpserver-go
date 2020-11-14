package mockhttpserver

type request struct {
	path   string
	method string
	body   []byte
}

func Request() *request {
	return &request{}
}

func (r *request) Path(path string) *request {
	r.path = path
	return r
}

func (r *request) Method(method string) *request {
	r.method = method
	return r
}

func (r *request) Body(body []byte) *request {
	r.body = body
	return r
}
