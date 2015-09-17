package main 

import ("fmt";"time"

)

//Prints 2 incoming strings after a gap of "delay" seconds time
// delay is in seconds
func sleep(delay uint) {
	fmt.Println("Time Now is:",time.Now());
	//Time.After returns current time
	delayTime := <-time.After(time.Second * time.Duration(delay))
	fmt.Println("Timestamp after delay:",delay," by time.After is:", delayTime);
}

func sleep_driver() {
	var delay uint
	fmt.Println("Enter the delay:");
	fmt.Scanf("%d",&delay);
	sleep(delay)
}

