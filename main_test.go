package main

import (
	"testing"
)

/*
Writing a Simple Unit Test
Unit tests are a way to test small pieces of code, such as functions and methods.
This is useful because it allows you to find bugs early.
Unit tests make your testing strategies more efficient, since they are small and independent,
 and thus easy to maintain.
*/

func Test_isPrime_SimpleUnitTest(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 7)
	}

	if msg != "7 is a prime number!" {
		t.Error("wrong message returned:", msg)
	}
}

/*
Writing Table-Driven Tests

When writing tests, you may find yourself repeating a lot of code in order to
cover all the cases required. Think about how you would go about covering the
many cases involved in the Fooer example. You could write one test function per case,
but this would lead to a lot of duplication. You could also call the tested function
several times in the same test function and validate the output each time, but if the
test fails, it can be difficult to identify the point of failure. Instead, you can use a
table-driven approach to help reduce repetition.
As the name suggests, this involves organizing a test case as a table that contains the
inputs and the desired outputs.

This comes with two benefits:
--Table tests reuse the same assertion logic, keeping your test DRY.
--Table tests make it easy to know what is covered by a test, as you can easily

	see what inputs have been selected. Additionally, each row can be given a unique
	 name to help identify what’s being tested and express the intent of the test.
*/
func Test_isPrime_TableDrivenTests(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
