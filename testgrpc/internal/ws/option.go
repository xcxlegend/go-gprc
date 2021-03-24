package ws

type Option struct {
	Addr string
}

var DefaultOption = &Option{Addr: ":9999"}
