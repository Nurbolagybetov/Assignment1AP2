package frontend

type FrontendServer interface {
	Run(port string) error
}
