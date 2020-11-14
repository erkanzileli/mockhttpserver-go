package mockhttpserver

import "bytes"

type matcher struct {
	req *request
	res *response
	s   *server
}

func (m *matcher) Respond(res *response) {
	m.res = res
	m.s.addMatcher(m)
}

func (m *matcher) MatchRequest(path, method string, body []byte) bool {
	return m.req.path == path && m.req.method == method && bytes.Equal(m.req.body, body)
}
