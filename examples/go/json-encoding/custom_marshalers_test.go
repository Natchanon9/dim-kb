package jsonencoding

import (
	"encoding/json"
	"fmt"
)

type Money int64 // cents

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d.%02d"`, m/100, m%100)), nil
}

func (m *Money) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	var whole, frac int64
	if _, err := fmt.Sscanf(s, "%d.%d", &whole, &frac); err != nil {
		return fmt.Errorf("money %q: %w", s, err)
	}
	*m = Money(whole*100 + frac)
	return nil
}

type Price struct {
	Amount Money `json:"amount"`
}

func Example_customMarshalers() {
	b, _ := json.Marshal(Price{1999})
	fmt.Println(string(b))

	var p Price
	json.Unmarshal([]byte(`{"amount":"5.50"}`), &p)
	fmt.Println(int64(p.Amount))
	// Output:
	// {"amount":"19.99"}
	// 550
}
