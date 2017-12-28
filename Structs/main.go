package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // equivalent to contactInfo contactInfo
}

// *person - byref instead of byval
func (pPointer *person) changeFirstName(newName string) {
	(*pPointer).firstName = newName //*p - the value in the memory address of pointer pPointer
}

func (pPointer *person) print() {
	fmt.Printf("%+v\n", (*pPointer))
}

func main() {
	// p := person{"tomer", "yakir"} // based on field order
	// p := person{firstName: "tomer", lastName: "yakir"}
	/* var p person
	fmt.Printf("%+v", p)
	p.firstName = "tomer"
	p.lastName = "yakir"
	fmt.Println(p)
	fmt.Printf("%+v", p)
	*/
	p := person{firstName: "tomer",
		lastName: "yakir",
		contactInfo: contactInfo{
			email:   "tomer.yakir@mongodb.com",
			zipCode: 4580800,
		},
	}
	pPointer := &p // get the pointer to p (memory address of p)
	pPointer.changeFirstName("tom")
	pPointer.print()

	// go allows to receive both on the var and the pointer
	p.changeFirstName("tomeragain")
	p.print()
}
