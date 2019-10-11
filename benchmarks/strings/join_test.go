package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const (
	benchN = 1000
)

const (
	format    = "firstname: %s lastname: %s age: %d"
	firstname = "john"
	lastname  = "doe"
	age       = 12
)

var ageString = strconv.Itoa(age)

//var d = data{firstname: firstname, lastname: lastname, age: age}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf(format, firstname, lastname, age)
	}
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "firstname: " + firstname + " lastname: " + lastname + " age: " + ageString
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	w := strings.Builder{}
	w.Grow(len("firstname: ") + len(firstname) + len(" lastname: ") + len(lastname) + len(" age: ") + len(ageString))
	for i := 0; i < b.N; i++ {
		w.WriteString("firstname: ")
		w.WriteString(firstname)
		w.WriteString(" lastname: ")
		w.WriteString(lastname)
		w.WriteString(" age: ")
		w.WriteString(ageString)
		_ = w.String()
	}
}
