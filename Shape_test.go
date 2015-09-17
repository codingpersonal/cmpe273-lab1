package main

import ("testing";"fmt")

func TestCircleArea(t *testing.T) {
	fmt.Println("\n\nIn Shape->Circle Area Testing Function");
	
	//Test for positive Value 5
	circle_obj:= circle{5}
	carea := circle_obj.area()
	if carea != 78.53981633974483{
		t.Error("Error, area is:",carea,"but expected was 78.53981633974483")
	} else {
		fmt.Println("(i) Test Passed when radius = 5");
	}
	
	//Test for Value 0
	circle_obj = circle{0}
	carea = circle_obj.area()
	if carea != 0{
		t.Error("Error, area is:",carea,"but expected was 0")
	}else {
		fmt.Println("(ii) Test Passed when radius = 0");
	}
	
	//Test for negative value
	circle_obj = circle{-1}
	carea = circle_obj.area()
	if carea != -1{
		t.Error("Error,radius is negative so it should return -1")
	}else {
		fmt.Println("(iii) Test Passed when radius = -1");
	}
}

func TestRectangleArea(t *testing.T) {
	fmt.Println("\nIn Shape->Rectangle Area Testing Function");
	
	//Test for positive coordinates
	rectangle_obj:= rectangle{2,3,5,6}
	rarea := rectangle_obj.area()
	if rarea != 9 {
		t.Error("Error, area is:",rarea," but expected was 9")
	}else {
		fmt.Println("(i) Test Passed when coordinates = {2,3,5,6}");
	}
	
	//Test for 0 values
	rectangle_obj = rectangle{0,0,0,0}
	rarea = rectangle_obj.area()
	if rarea != 0 {
		t.Error("Error, area is:",rarea," but expected was 0")
	}else {
		fmt.Println("(ii) Test Passed when coordinates = {0,0,0,0}");
	}
	
	//Test for Negative Values
	rectangle_obj = rectangle{-2,-3,-5,-6}
	rarea = rectangle_obj.area()
	if rarea != 9 {
		t.Error("Error, area is:",rarea," but expected was 0")
	}else {
		fmt.Println("(iii) Test Passed when coordinates = {-2,-3,-5,-6}");
	}
}
