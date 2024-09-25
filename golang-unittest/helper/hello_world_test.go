package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeloWolrd1(t *testing.T) {
	result := HelloWorld("Razi")

	if result != "Hello Razi" {
		// t.Fail()
		t.Error("Error, result should be 'Hello Razi'")
	}
	fmt.Println("TestHeloWolrd1 done")
}

func TestHeloWolrd2(t *testing.T) {
	result := HelloWorld("Hiyah")

	if result != "HelloHiyah" {
		// t.FailNow()
		t.Fatal("Error, result should be 'HelloHiyah'")
	}
	fmt.Println("TestHeloWolrd2 done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Razi")
	assert.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestHelloWorldAssert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Razi")
	require.Equal(t, "HelloRazi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestHelloWorldRequire is done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Razi")
	require.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestSkip is done")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println(">>> Before Unit Test <<<")
	fmt.Println()

	m.Run()

	//after
	fmt.Println()
	fmt.Println(">>> After Unit Test <<<")
}

func TestSubTest(t *testing.T) {
	t.Run("Razi", func(t *testing.T) {
		result := HelloWorld("Razi")
		require.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	})

	t.Run("Hiyah", func(t *testing.T) {
		result := HelloWorld("Hiyah")
		require.Equal(t, "Hello Hiyah", result, "Error, result should be 'Hello Hiyah'")
	})

}

func TestTable(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Razi",
			request:  "Razi",
			expected: "Hello Razi",
		},
		{
			name:     "Aziz",
			request:  "Aziz",
			expected: "Hello Aziz",
		},
		{
			name:     "Syahputro",
			request:  "Syahputro",
			expected: "Hello Syahputro",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Error, result should be"+test.expected)
		})
	}
}

//if in Helper folder
//go test
//go test -v
//go test -v -run=<TestFunction>

//if in root folder
//go test -v ./...

/**
t.Fail() == t.Error() == assert : akan tetap dieksekusi kode selanjutnya jika gagal
t.FailNow() == t.Fatal() == require : tidak akan dieksekusi kode selanjutnya jika gagal
*/

// BENCHMARK
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Razi")
	}
}

func BenchmarkHelloWorldSyahputro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syahputro")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Razi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Razi")
		}
	})

	b.Run("Hiyah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hiyah")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name, request string
	}{
		{
			name:    "Razi",
			request: "Razi",
		},
		{
			name:    "Aziz",
			request: "Aziz",
		},
		{
			name:    "Syahputro",
			request: "Syahputro",
		},
		{
			name:    "Razi Aziz Syahputro",
			request: "Razi Aziz Syahputro",
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

/**
if u want to run Benchmark only not UnitTest the command is
go test -v -run=NotMatchTestFunc -bench=.

if u want to run Benchmark Specific only not UnitTest the command is
go test -v -run=TestNotMatchTestFunc -bench=Benchmarkxxx

if wants to run in root folder the command is
go test -v -run=TestNotMatchTestFunc -bench=. ./...
*/

//BENCHMARK
