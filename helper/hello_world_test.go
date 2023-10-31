package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Vicry",
			request: "Vicry",
		},
		{
			name:    "Fahreza",
			request: "Fahreza",
		},
		{
			name:    "VicryFahreza",
			request: "Vicry Fahreza",
		},
		{
			name:    "Holland",
			request: "Holland Bakery Tray",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Vicry", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Vicry")
		}
	})
	b.Run("Fahreza", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fahreza")
		}
	})
}

func BenchmarkHelloWorldVicry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Vicry")
	}
}

func BenchmarkHelloWorldFahreza(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fahreza")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Vicry",
			Request:  "Vicry",
			Expected: "Hello Vicry",
		},
		{
			Name:     "Fahreza",
			Request:  "Fahreza",
			Expected: "Hello Fahreza",
		},
		{
			Name:     "Walker",
			Request:  "Walker",
			Expected: "Hello Walker",
		},
		{
			Name:     "Tom",
			Request:  "Tom",
			Expected: "Hello Tom",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Vicry", func(t *testing.T) {
		result := HelloWorld("Vicry")
		require.Equal(t, "Hello Vicry", result, "Result is not 'Hello Vicry'")
	})
	t.Run("Fahreza", func(t *testing.T) {
		result := HelloWorld("Fahreza")
		require.Equal(t, "Hello Fahreza", result, "Result is not 'Hello Fahreza'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	result := HelloWorld("Vicry")
	if runtime.GOOS == "windows" {
		t.Skip("Test Cannot Be Run on Windows")
	}
	require.Equal(t, "Hello Vicry", result, "Result is not 'Hello Vicry'")
	fmt.Println("Test HelloWorld with Assert Done")
}

func TestHelloWorldWithRequire(t *testing.T) {
	result := HelloWorld("Vicry")
	require.Equal(t, "Hello Vicry", result, "Result is not 'Hello Vicry'")
	fmt.Println("Test HelloWorld with Require Done")
}

func TestHelloWorldWithAssert(t *testing.T) {
	result := HelloWorld("Vicry")
	assert.Equal(t, "Hello Vicry", result, "Result is not 'Hello Vicry'")
	fmt.Println("Test HelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Vicry")
	if result != "Hello Vicry" {
		// panic("Result is not 'Hello Vicry'")
		t.Error("Seharusnya 'Hello Vicry'")
	}
	fmt.Println("Dieksekusi")
}

func TestHelloWorldFahreza(t *testing.T) {
	result := HelloWorld("Fahreza")
	if result != "Hello Fahreza" {
		t.Error("Seharusnya 'Hello Fahreza'")
	}
	fmt.Println("Dieksekusi")
}
