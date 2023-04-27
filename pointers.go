package main

import "fmt"

type Students struct {
	Name string
	Course string
	Gender string
	Age int
}
func main () {
	//pointers()
	arrays()
	//slices()
}
func pointers() {
	student1 := Students{"Naveena", "Golang", "F", 25}
	student2 := Students{"Deepthi", "Bigdata", "F", 26}
	student3 := Students{"laxmi", "java", "F", 28}
	var ptr *Students
	ptr = &student1
	fmt.Println(student1, student2, student3)
	fmt.Println(&ptr, *ptr)
	ptr.Age = 26
	ptr.Name = "Manikanta"
	ptr.Gender = "M"
	fmt.Println(student1, student2, student3)
	fmt.Println(&ptr, *ptr)
	student2.Name = "Akhila"
	student2.Age = 24
	fmt.Println(student1, student2, student3)
}

func arrays() {
	var myarray = [9] string{"Naveena", "Manikanta", "Soumya", "Sohan", "Praharsha", "Mounika", "Shanmukh"}
	for i :=0; i < len(myarray) ;i++ {
		if myarray[i] == "Soumya" {
			myarray[i] = "Phani"
		}
		fmt.Println(myarray[i])
	}
	//myarray[7] = "Sai"
	//myarray = append(myarray, "Sai") we cannot add element to array ----- first argument to append must be a slice
	fmt.Println(myarray)
	var s []string = myarray[:6]
	fmt.Println(len(s), cap(s)) //Here the length is 6 and capacity is 7
	fmt.Println(s)
	s = append(s, "Deepthi", "Akhila")
	fmt.Println(s, len(s), cap(s)) // slice length increased automatically len 8 and capacity 14
}

func slices() {
	s := make([]int, 5) //[0 0 0 0 0] 5 5
    fmt.Println(s,len(s),cap(s)) //len 5 and cap 5
	s[0] = 1
    s[1] = 2
    s[2] = 3
	//s[6] = 4 ------ we get an error index out of range
    s = append(s, 4, 5, 6, 7) //like this we won't get any error 
	fmt.Println(s, len(s),cap(s)) //[1 2 3 0 0 4 5 6 7] 9 10
	a := s[3:8]
	fmt.Println(a, len(a), cap(a)) //[0 0 4 5 6] 5 7
	a = s[2:7]
	fmt.Println(a, len(a), cap(a)) //[3 0 0 4 5] 5 8
	a = s[4:10]
	fmt.Println(a, len(a), cap(a)) //[0 4 5 6 7 0] 6 6
	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))
	if s1 == nil {
		fmt.Println("nil!")
	}
	s1 = append(s1, 2,3)
	fmt.Println(s1, len(s1), cap(s1)) //[2 3] 2 2
}