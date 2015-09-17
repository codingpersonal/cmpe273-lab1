package main

import ("testing"; "fmt"; "time")

func TestSleep(t *testing.T) {
	fmt.Println("\n\nIn sleep Test Case");
	currTime := time.Now().Second();
	sleep(4)
	delayTime := time.Now().Second();
	
	if delayTime - currTime != 4 {	
		t.Error("Error,Sleep should be of 4 seconds");
	} else {
		fmt.Println("Sleep works fine, Test Passed");
	}
}