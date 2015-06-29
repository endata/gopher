package contracts

type Renderable interface {
	NewRenderer() Renderable
}
