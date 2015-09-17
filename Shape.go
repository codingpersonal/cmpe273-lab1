package main

import (
	"fmt"
	"math"
)

//interface defining area and perimeter of shapes
type shape interface {
	area() float64
	perimeter() float64
}

//struct with radius of circle
type circle struct {
	rad float64
}

//struct with (x,y) corrdinates of length and width of the rectangle
type rectangle struct {
	len_x, len_y, wid_x, wid_y float64
}

//takes 2 (x,y) coordinates and returns the distance between them
func distance(len_x, len_y, wid_x, wid_y float64) float64 {
	dis_x := wid_x - len_x
	dis_y := wid_y - len_y
	return math.Sqrt((dis_x * dis_x) + (dis_y * dis_y))
}

//returns the area of the circle
func (c circle) area() float64 {
	if c.rad < 0 {
		return -1
	} else {
		return math.Pi * c.rad * c.rad
	}
}

//returns the area of the rectangle
func (r rectangle) area() float64 {
	len := distance(r.len_x, r.len_y, r.len_x, r.wid_y)
	wid := distance(r.len_x, r.len_y, r.wid_x, r.len_y)
	return len * wid
}

//returns the perimeter of the circle
func (c circle) perimeter() float64 {
	if c.rad < 0 {
		return -1
	} else {
		return math.Pi * 2 * c.rad
	}
}

//returns the perimeter of the rectangle
func (r rectangle) perimeter() float64 {
	len := distance(r.len_x, r.len_y, r.len_x, r.wid_y)
	wid := distance(r.len_x, r.len_y, r.wid_x, r.len_y)
	return (2 * (len + wid))
}

//computes the area and perimeter of the passed shape object
func computeShape(s_obj shape) {
	fmt.Println("Area:", s_obj.area())
	fmt.Println("Perimeter:", s_obj.perimeter())
}

func shape_driver() {
	fmt.Println("Implementing the Shape Interface\n")
	
	var rad float64
	fmt.Println("Enter the radius of the circle\n")
	fmt.Scanf("%f\n",&rad)

	var x1,y1,x2,y2 float64
	fmt.Println("Enter the (x,y) coordinates for the length and width\n")
	fmt.Scanf("%f %f %f %f",&x1,&y1,&x2,&y2)
	
	c_obj := circle{rad}
	r_obj := rectangle{x1, y1, x2, y2}

	fmt.Println("Computing Area and Perimeter for shape circle\n")
	computeShape(c_obj)

	fmt.Println("\nComputing Area and Perimeter for shape rectange")
	computeShape(r_obj)
}
