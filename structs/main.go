package main

import "fmt"

// Embed this struct to Person
type contactInfo struct {
	email       string
	phoneNumber int
}
type Person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	contactInfoSuraj := contactInfo{
		email:       "suraj@gmail.com",
		phoneNumber: 1092345656,
	}
	// first way to declare a struct
	suraj := Person{firstName: "Suraj",
		lastName:    "Arya",
		contactInfo: contactInfoSuraj,
	}
	fmt.Println(suraj)

	//second way to declare struct
	var diksha Person
	diksha.firstName = "Diksha"
	diksha.lastName = "Tiwari"
	diksha.contactInfo.email = "diksha@gmail.com"
	diksha.contactInfo.phoneNumber = 12345678

	diksha.updateName("Priya")
	diksha.print()

}

func (p Person) print() {
	fmt.Println(p)

}

func (p *Person) updateName(firstNameUpdate string) {
	(*p).firstName = firstNameUpdate
}

//Go is Pass by Value function language
/*There are two types:-
1. Value type:- it requires pointers to modify the value.
	eg:-int,float,bool,string, struct
2. Refernce type: it does not require pointers to modify value
	eg:- slices,map,channels,pointers,function
*/
