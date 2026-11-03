package reflection

import (
	"reflect"
	"testing"
)

type point struct{ X, Y int }

var sink int

func BenchmarkDirectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = p.X
	}
}

func BenchmarkReflectField(b *testing.B) {
	p := point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = int(reflect.ValueOf(p).FieldByName("X").Int()) // name lookup every time
	}
}
