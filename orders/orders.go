package orders

type OrderStore interface {
	Exec(query string, args ...any) error
}

type OrderService struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{store: store}
}

func (s *OrderService) PlaceOrder(orderID string, amount float64) error {
	return s.store.Exec(
		"INSERT INTO orders (id, amount) VALUES (?, ?)",
		orderID,
		amount,
	)
}
