package mockhttpserver

type ServerBuilder interface {
	Address(addr string) ServerBuilder
	Build() *server
}

// server implements Server
type serverBuilder struct {
	address string
}

func NewServerBuilder() ServerBuilder {
	return &serverBuilder{}
}

func (s *serverBuilder) Address(addr string) ServerBuilder {
	s.address = addr
	return s
}

func (s *serverBuilder) Build() *server {
	return &server{
		address: s.address,

	}
}
