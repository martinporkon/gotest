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

// Immediate failures
// t.FailNow()
// t.Fatal(args ...interface{})
// t.Fatalf(format string, args ...interface{})

// Non-immediate failures
// t.Fail()
// t.Error(args ...interface{})
// t.Errorf(format string, args ...interface{})
