package order

type IorderService interface {
	GetOrderById(string) ([]byte, error)
}
