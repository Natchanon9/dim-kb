---
title: "Testing Best Practices"
description: "Habits that keep Go tests fast and trustworthy"
sidebar:
  order: 10
status: done
lastReviewed: 2026-11-03
---

Good tests are fast, independent and test behaviour through small interfaces.

## Fake dependencies through interfaces

```go title="welcome_test.go"
type Mailer interface{ Send(to, body string) error }

type fakeMailer struct{ sent []string }

func (f *fakeMailer) Send(to, _ string) error {
	f.sent = append(f.sent, to)
	return nil
}

func TestWelcome(t *testing.T) {
	m := &fakeMailer{}
	if err := Welcome(m, "a@b.com"); err != nil {
		t.Fatal(err)
	}
	if len(m.sent) != 1 {
		t.Errorf("sent = %v", m.sent)
	}
}
```

## Checklist

- **Table-driven + subtests** for anything with several cases.
- **Hand-written fakes** for small interfaces; mock generators (`mockgen`) for big ones.
- **No sleeps;** wait on channels or poll with a deadline.
- **Real dependencies in integration tests** (testcontainers, a test database), behind a build tag or `-short` check.
- **Test the public API** of a package (`package foo_test`) when you can.
- **Run `-race`** in CI.

:::caution[Pitfalls]
- Tests that depend on map order, time or test order are flaky; fix them, don't retry them.
:::
