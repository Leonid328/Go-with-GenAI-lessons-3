package orders

import "testing"

type mockStore struct {
	query string
	args  []any
}

func (m *mockStore) Exec(query string, args ...any) error {
	m.query = query
	m.args = args
	return nil
}

func TestPlaceOrderAIMock(t *testing.T) {
	m := &mockStore{}
	svc := NewOrderService(m)
	err := svc.PlaceOrder("o-9", 15)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.args) != 2 {
		t.Fatalf("args=%d", len(m.args))
	}
}
