package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}

// no assertions in go
// go test
// testing package
// testing/quick package
// testing/iotest package
// net/http/httptest package
// test servers with end-to-end tests and real http messages
// to endpoints
// Community projects:
// Testify <- community project
// Ginkgo <- BDD testing community project
// GoConvey <- generates results for browser format
// httpexpect
// gomock
// go-sqlmock <- for databases. Im memory mockable database

//_test must be added to filenames as the testing framework wont find them otherwise
// Prefix function names with "Test"
// one parameter *testing.T <- pointer?
// white box test -> as the tests are part of the same package, we've got access to all of the
// package level variables within our test. (package main)
// add _test suffic to package for blackbox tests (package main_test)
// best practise is to run black box testing
// tests the API of the package and does not consider the internal implementatino details of the
// package, more robust and more representative to the consumers of the package.
