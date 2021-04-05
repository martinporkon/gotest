// integrtion test or end to end tests it maske ssense to have folder
package messages

import "testing"

// Stub of two test functions. Pointer to testing.T objects, convetionally as referred to withing
// the test of variable T
func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got) // this is a non fatal error, so we dont need to exit the test immediately but do need to report
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got) // this is a non fatal error, so we dont need to exit the test immediately but do need to report
	}
}

// go test -coverprofile cover.out <- generates a file
// go tool cover -func cover.out <- analyses coverage file
// go tool cover -html cover.out <- performs in a web browser

// go test -coverprofile count.out -covermode count <- generates another profile

// Error <- signasl a failed test , but does not stop the rest of the test from executing
// Fatal <- will fail the test and stop its execution

// Immediate failures
// t.FailNow()
// t.Fatal(args ...interface{})
// t.Fatalf(format string, args ...interface{})

// Non-immediate failures
// t.Fail()
// t.Error(args ...interface{})
// t.Errorf(format string, args ...interface{})

// Test specified packages
// go test {pkg1} {pkg2} ...., {pkgn}

// go test ./... <- run tests in current package and descendats
// go test -v <- generate verbose output
// gp test -run {regexp} <- run only tests matching {regexp}
// go help test
// go help testflag

// table-driven test

func TestGreetTableDriven(t *testing.T) {
	scenarious := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}

	for _, s := range scenarious {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got: %q\n",
				s.input, s.expect, got)
		}
	}
}

// Useful functions
// Log and Logf, Helper, Skip, Skipf and SkipNow
// Run
// parallel
// golang.org/pkg/testing
