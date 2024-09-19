package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"John", 75000, "Backend Developer"}
	emp3 := employee{"Emily", 70000, "Frontend Developer"}

	emplys := []employee{emp1, emp2, emp3}

	comp := company{"Tetra", emplys}

	fmt.Println("Company Name:", comp.companyName)
	fmt.Println("Employees:")
	for _, emp := range comp.employees {
		fmt.Println("Name:", emp.name)
		fmt.Println("Salary:", emp.salary)
		fmt.Println("Position:", emp.position)
		fmt.Println()
	}
}