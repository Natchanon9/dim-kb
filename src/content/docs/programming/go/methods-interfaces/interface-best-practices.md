---
title: "Interface Best Practices"
description: "Small interfaces, defined by the consumer"
sidebar:
  order: 11
status: done
lastReviewed: 2026-11-03
---

Good Go interfaces are small and declared where they are used, not next to the implementation.

## Define the interface at the consumer

```go title="notify.go"
// Declared by the code that needs it, with only the method it uses.
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

type fakeStore struct{}

func (fakeStore) Email(int) (string, error) { return "a@b.com", nil }

Notify(fakeStore{}, 1) // sent to a@b.com
```

## Rules of thumb

- **Accept interfaces, return structs.**
- **Smaller is better:** one method is ideal.
- **Don't create an interface before a second implementation (or a test fake) needs it.**
- **No `IUserService` naming;** name by behaviour.

:::caution[Pitfalls]
- A package that exports an interface plus its only implementation usually just needs the struct.
:::
