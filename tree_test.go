package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func Exampletree_Format() {
	file, err := ioutil.ReadFile("files.txt")
	if err != nil {
		log.Fatal(err)
	}

	g := New("/")
	lines := strings.Split(string(file), "\n")

	g.EatLines(lines)

	fmt.Print(g.Format())
	// Output:
	//
	// [34m1[0m
	// ├── [34m2[0m
	// │   └── [34m3[0m
	// │       ├── [34m4[0m
	// │       │   ├── [34m5[0m
	// │       │   └── [34mfisk2.txt[0m
	// │       └── [34mfisk2.txt[0m
	// ├── [34m3[0m
	// │   ├── [34m4[0m
	// │   │   ├── [34m5[0m
	// │   │   └── [34mfisk2.txt[0m
	// │   └── [34mfisk2.txt[0m
	// ├── [34m5[0m
	// │   ├── [34m4[0m
	// │   │   ├── [34m3[0m
	// │   │   │   ├── [34m2[0m
	// │   │   │   └── [34mfisk.txt[0m
	// │   │   └── [34mfisk.txt[0m
	// │   ├── [34mfisk.txt[0m
	// │   └── [34mfisk2.txt[0m
	// └── [34mfisk.txt[0m
	// [34mfisk.txt[0m
	//
}

func TestShallowTree(t *testing.T) {
	lines := []string{
		"one",
		"other",
		"this",
	}

	expected := `[34mone[0m
[34mother[0m
[34mthis[0m
`

	setup := func() *tree {
		return New("/")
	}
	RunFormatTestCase(lines, expected, setup, t)
}

func RunFormatTestCase(inputLines []string, expected string, setup func() *tree, t *testing.T) {
	tr := setup()

	tr.EatLines(inputLines)

	output := tr.Format()

	errorFormat := `Expected
===
%s===

Got
===
%s===`

	if output != expected {
		t.Errorf(
			errorFormat,
			expected,
			output,
		)
	}
}

func TestNodeFormat(t *testing.T) {
	lines := []string{
		"one",
		"other$retho",
		"this",
	}

	expected := `✓ one ⚡
✓ other ⚡
└── ✓ retho ⚡
✓ this ⚡
`
	setup := func() *tree {
		tr := New("$")
		tr.NodeFormat = "✓ %s ⚡"
		return tr
	}

	RunFormatTestCase(lines, expected, setup, t)
}

func TestSingleNode(t *testing.T) {
	lines := []string{"1"}
	expected := "1\n"
	setup := func() *tree {
		tr := New(".")
		tr.NodeFormat = "%s"
		return tr
	}
	RunFormatTestCase(lines, expected, setup, t)
}

func BenchmarkTreeFormat(b *testing.B) {
	file, err := ioutil.ReadFile("files.txt")
	lines := strings.Split(string(file), "\n")
	if nil != err {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		t := New("/")
		t.EatLines(lines)
		t.Format()
	}
}

func BenchmarkHeavyTreeFormat(b *testing.B) {
	file, err := ioutil.ReadFile("data.csv")
	lines := strings.Split(string(file), "\n")
	if nil != err {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		t := New(",")
		t.EatLines(lines)
		t.Format()
	}
}
