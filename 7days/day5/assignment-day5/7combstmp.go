package main
import "fmt"
type Student struct {
	name string
	age int
	grade string
}

func main() {
	//defining a student
	stud :=Student{name: "Yaash", age: 21, grade: "S"}
	//creating map for a student
	students := make(map[string]Student)
	//add students to map
	students["S059"] = stud

	

	fmt.Println(stud)

}
