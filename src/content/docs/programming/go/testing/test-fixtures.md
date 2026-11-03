---
title: "Test Fixtures"
description: "testdata folders and golden files"
sidebar:
  order: 8
status: done
lastReviewed: 2026-11-03
---

Put input files and expected outputs in a `testdata/` folder. The go tool ignores it as a package, and tests run with the package directory as the working directory.

## Golden files

```go title="golden_test.go"
func TestGreetGolden(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "greeting.golden"))
	if err != nil {
		t.Fatal(err)
	}
	if got := Greet("Ann") + "\n"; got != string(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

A common addition is an `-update` flag that rewrites the golden files when output changes on purpose:

```go
var update = flag.Bool("update", false, "rewrite golden files")
// if *update { os.WriteFile(goldenPath, []byte(got), 0o644) }
```

Good for rendered templates, generated code, big JSON responses.

:::caution[Pitfalls]
- Review golden file diffs carefully; `-update` accepts whatever the code produces now.
:::
