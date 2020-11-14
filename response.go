package mockhttpserver

type response struct {
	status  int
	body    []byte
	headers map[string]string
}

func Response() *response {
	return &response{}
}

func (r *response) Status(status int) *response {
	r.status = status
	return r
}

func (r *response) Body(body []byte) *response {
	r.body = body
	return r
}

func (r *response) Header(key, value string) *response {
	r.headers[key] = value
	return r
}
