package methodsinterfaces

import "fmt"

// UserStore is declared by the code that needs it, with only the method it uses.
type UserStore interface {
	Email(id int) (string, error)
}

func Notify(s UserStore, id int) string {
	e, err := s.Email(id)
	if err != nil {
		return "no email"
	}
	return "sent to " + e
}

type fakeStore struct{} // a test double is trivial with a one-method interface

func (fakeStore) Email(int) (string, error) { return "a@b.com", nil }

func Example_bestPractices() {
	fmt.Println(Notify(fakeStore{}, 1))
	// Output: sent to a@b.com
}
