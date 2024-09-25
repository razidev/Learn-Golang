package main

import "fmt"

type job struct {
	Position string
	Salary   int
}

type MySelf struct {
	Name, Address string
	Age           int
	job           //embeded struct
}

//struct function
func (me MySelf) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", me.Name)
}

type People struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	var Me MySelf
	Me.Name = "Razi"
	Me.Address = "Kp. Mangga"
	Me.Age = 25
	Me.job.Salary = 1000
	Me.Position = "Fullstack Developer"
	fmt.Println("Me:", Me)
	Me.sayHi("Joko")

	aziz := MySelf{
		Name:    "Aziz",
		Address: "Gg. Teladan 5",
		job: job{
			Salary:   3000,
			Position: "Frontend Developer",
		},
	}
	fmt.Println("aziz:", aziz)
	//not recommended not using .job because it can be redundant if there are another Position field
	fmt.Println("aziz position:", aziz.Position)
	//recommended
	fmt.Println("aziz salary:", aziz.job.Salary)

	//not recommended because it depends on position
	syahputro := MySelf{"Syahputro", "Jakarta", 23, job{"Cloud Developer", 20000}}
	fmt.Println("syahputro:", syahputro)
	fmt.Println("syahputro salary:", syahputro.Salary)
	fmt.Println("syahputro position:", syahputro.job.Position)

	p1 := People{
		first: "Razi",
		last:  "Syahputro",
		favFlavors: []string{
			"Hot",
			"Coke",
			"Chocolate",
		},
	}

	p2 := People{
		first: "Hiyah",
		last:  "Iza",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
