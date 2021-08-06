package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
//Struct examples
type Student struct {
	id    int
	name  string
	marks float64
}

func (s *Student) percentage() float64 {
	return (s.marks / 60) * 100
}

type Catherine struct {
	Student
	modelMarks string
}

//Interface example
type Givenname interface {
	name() string
}

type Student1 struct {
	studentName string
}

type Teacher struct {
	teacherName string
}

func (s1 Student1) name() string {
	return s1.studentName
}

func (t Teacher) name() string {
	return t.teacherName
}

func getName(givenName Givenname) string {
	return givenName.name()
}
*/

type Radius interface {
	area() float64
	perimeter() float64
}
type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	circRadius float64
}

func (c Circle) area() float64 {
	area := math.Pi * c.circRadius * c.circRadius
	return area
}

func (c Circle) perimeter() float64 {
	perimeter := 2 * math.Pi * c.circRadius
	return perimeter
}

func (r Rectangle) area() float64 {
	area := r.height * r.width
	return area
}

func (r Rectangle) perimeter() float64 {
	perimeter := 2 * (r.height + r.width)
	return perimeter
}

func display(radius Radius) {
	fmt.Println(radius.area())
	fmt.Println(radius.perimeter())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	w, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	input2, _ := reader.ReadString('\n')
	l, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	input3, _ := reader.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(input3), 64)
	r1 := Rectangle{width: w, height: l}
	display(r1)
	var c1 Circle
	c1.circRadius = value
	display(c1)

	/*
		//Struct Examples
		var s Student
		s.id = 2
		s.name = "Ann"
		s.marks = 55.86
		fmt.Printf("Student ID : %d\n", s.id)
		fmt.Printf("Student Name : %s\n", s.name)
		fmt.Printf("Student Marks : %f\n", s.marks)
		fmt.Printf("Student Percentage : %f\n", s.percentage())
		c := new(Catherine)
		c.marks = 44.9
		c.modelMarks = "Alexander"
		fmt.Printf("Catherine's Percentage : %f\n", c.percentage())
		fmt.Println("Catherine's Percentage :", c.modelMarks)

		//Interface examples
		var s1 Student1
		s1.studentName = "Ann"
		t := Teacher{teacherName: "Roopa Madam"}
		fmt.Printf("Name of Student: %s\n", getName(s1))
		fmt.Printf("Name of Teacher: %s\n", getName(t))
	*/
}
