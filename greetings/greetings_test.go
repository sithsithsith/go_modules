package greetings

import (
	"regexp"
	"testing"
)

func TestGreetingName(t *testing.T) {
	name := "itachi"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("itachi")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("itachi") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetingEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
