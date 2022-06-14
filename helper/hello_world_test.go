package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "imam", request: "imam"},
		{name: "ardi", request: "ardi"},
		{name: "ihab", request: "ihab"},
		{name: "kami", request: "kami"},
		{name: "iwal", request: "iwal"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)

			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("imam", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("imam")
		}
	})
	b.Run("david", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("david")
		}
	})
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("imam")
	}
}

func BenchmarkSayHello2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("fahrezi")
	}
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "imam", request: "imam", expected: "hello imam"},
		{name: "ardi", request: "ardi", expected: "hello ardi"},
		{name: "shaqil", request: "shaqil", expected: "hello shaqil"},
		{name: "ihab", request: "ihab", expected: "hello ihab"},
		{name: "iwal", request: "iwal", expected: "hello iwal"},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			result := SayHello(v.request)
			require.Equal(t, v.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("imam", func(t *testing.T) {
		result := SayHello("imam")
		require.Equal(t, "hello imam", result, "Result must be 'Hello imam'")

	})
	t.Run("ihab", func(t *testing.T) {
		result := SayHello("ihab")
		require.Equal(t, "hello ihab", result, "Result must be 'Hello ihab'")

	})
}

func TestMain(t *testing.M) {
	fmt.Println("BEFORE RUN")
	t.Run()
	fmt.Println("AFTER RUN")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("only run mac os")
	}
	result := SayHello("ihab")
	require.Equal(t, "hello ihab", result, "Result must be 'Hello ihab'")

}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("ihab")
	require.Equal(t, "hello ihab", result, "Result must be 'Hello ihab'")
	fmt.Println("say hello with require done")

}
func TestSayHelloAssert(t *testing.T) {
	result := SayHello("ihab")
	assert.Equal(t, "hello ihab", result, "Result must be 'Hello ihab'")
	fmt.Println("say hello assert done")

}

func TestSayHello(t *testing.T) {
	result := SayHello("shaqil")
	if result != "hello shaqil" {
		//t.Fail()
		t.Error("return must be shaqil")

	}
	fmt.Println("say hello shaqil done")
}
func TestSayHelloImam(t *testing.T) {
	result := SayHello("imam")
	if result != "hello imam" {
		//t.FailNow()
		t.Fatal("return must be imam")
	}
	fmt.Println("test sayhello imam done")
}
