package main

import ("testing";"fmt")

	var fibTestSet = []struct {
		input          int64
		expectedOutput int64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

func TestFibo_Positive(t *testing.T) {
	fmt.Println("\n\nIn Fibonacci Test , Testing for positive values");
	for _, id := range fibTestSet {
		valFromFunc := fibo_rec(id.input)
		if valFromFunc != id.expectedOutput {
			t.Errorf("Fib(%d): expected %d, actual %d", id.input, id.expectedOutput, valFromFunc)
		}else {
		fmt.Println("Test Passed");
		}
	}
}

func TestFibo_Zero(t *testing.T) {
	fmt.Println("\nIn Fibonacci Test , Testing for Val = 0");
	output := fibo_rec(0)
	if output != 0 {
		t.Error("Expected 0, but got ", output)
	} else {
		fmt.Println("Test Passed");
	}
}

func TestFibo_Negative(t *testing.T) {
	fmt.Println("\nIn Fibonacci Test , Testing for Value = -1");
	output := fibo_rec(-1)
	if output != -1 {
		t.Error("Expected -1[error] since fibonacci of negative numbers is not possible, but got ", output)
	}else {
		fmt.Println("Test Passed");
	}
}
