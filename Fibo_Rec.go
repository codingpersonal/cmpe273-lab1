package main 

import ("fmt")

//Prints the fibonacci series till the given range
func fibo_rec_driver() {
	var num int64;
	
	fmt.Println("Welcome to Fibonacci series program");
	fmt.Println("Enter a range for the series to print");
	fmt.Scanf("%d",&num);
	
	if num < 0 {
		fmt.Println("Error! Fibonacci for Negative  Number is not possible");
		return
	}
	
	var index int64;
	fmt.Println("Fibonacci Series:");
	for index = 0; index <= num; index = index + 1 {
		fmt.Println("[",index,"]","=>", fibo_rec(index))
	}
}

//Take a input range and returns the fibonacci output
func fibo_rec(num int64) int64 {
	if num == 1 {
		return 1
	} else if num == 0 {
	return 0
	}	else if num < 0 {
	return -1
	}	else {
		return fibo_rec(num -1) + fibo_rec(num - 2) }
}