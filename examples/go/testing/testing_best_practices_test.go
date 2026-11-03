package gotesting

import (
	"errors"
	"testing"
)

// A small interface makes the dependency easy to fake.
type Mailer interface{ Send(to, body string) error }

type fakeMailer struct{ sent []string }

func (f *fakeMailer) Send(to, _ string) error {
	if to == "" {
		return errors.New("empty recipient")
	}
	f.sent = append(f.sent, to)
	return nil
}

func Welcome(m Mailer, email string) error { return m.Send(email, "welcome!") }

func TestWelcome(t *testing.T) {
	m := &fakeMailer{}
	if err := Welcome(m, "a@b.com"); err != nil {
		t.Fatal(err)
	}
	if len(m.sent) != 1 || m.sent[0] != "a@b.com" {
		t.Errorf("sent = %v", m.sent)
	}
	if err := Welcome(m, ""); err == nil {
		t.Error("expected an error for an empty address")
	}
}
