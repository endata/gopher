package contracts

//type Router interface{}

type Routerable interface {
	NewRouter(ctx interface{}) Routerable
	Get(path string, fn interface{}) Routerable
}
