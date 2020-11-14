package mockhttpserver

import (
	"io/ioutil"
	"net/http"
)

// server implements Server
type server struct {
	address  string
	host     string
	port     string
	server   *http.Server
	matchers []*matcher
}

func Server() *server {
	return &server{}
}

func (s *server) Address(addr string) *server {
	s.address = addr
	return s
}

func (s *server) When(req *request) *matcher {
	return &matcher{
		req: req,
		s:   s,
	}
}

func (s *server) addMatcher(m *matcher) {
	s.matchers = append(s.matchers, m)
}

func (s *server) Start() error {
	s.server = &http.Server{Addr: s.address}

	s.server.Handler = s

	return s.server.ListenAndServe()
}

func (s *server) StartInBackground() {
	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()
}

func (s *server) Stop() error {
	return s.server.Shutdown(nil)
}

// ServeHTTP is implementation of http.Handler
func (s *server) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	for _, m := range s.matchers {
		bs, _ := ioutil.ReadAll(r.Body)
		if m.MatchRequest(r.URL.Path, r.Method, bs) {
			writer.WriteHeader(m.res.status)
			writer.Write(m.res.body)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}
