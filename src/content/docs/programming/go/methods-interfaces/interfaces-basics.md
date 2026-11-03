---
title: "Interfaces Basics"
description: "A set of method signatures; any type with those methods satisfies it"
sidebar:
  order: 4
status: done
lastReviewed: 2026-11-03
---

An interface lists method signatures. Any type that has those methods can be used as that interface, which lets one function work with many types.

## Example

```go title="speaker.go"
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "woof" }

type Robot struct{ ID int }

func (r Robot) Speak() string { return fmt.Sprintf("beep %d", r.ID) }

for _, s := range []Speaker{Dog{}, Robot{7}} {
	fmt.Println(s.Speak()) // woof, beep 7
}
```

## What an interface value holds

An interface value is a pair: (dynamic type, value). It is `nil` only when both are unset.

:::caution[Pitfalls]
- Calling a method on a nil interface panics.
- Name one-method interfaces with `-er`: `Reader`, `Stringer`, `Speaker`.
:::
