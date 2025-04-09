package backend

type Options struct {
	Port string
}

func NewOptions(port string) *Options {
	return &Options{Port: port}
}
