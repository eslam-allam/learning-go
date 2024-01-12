package main

import (
	"fmt"
	"math"
)

type Employed interface {
	getSalary() float64
}

type Gig struct {
	name             string
	description      string
	paymentInDollars float64
}

type FreeLancer interface {
	getAcceptedGigs() []Gig
	requestServices(gig Gig) bool
	acceptGig(gigName string)
}

type Gender int

const (
	MALE Gender = iota
	FEMALE
)

type Person struct {
	name   string
	age    uint
	gender Gender
}

func (p Person) isAdult() bool {
	return p.age >= 21
}

type Doctor struct {
	hospital string
	Person
	salary float64
}

func (d Doctor) getSalary() float64 {
	return d.salary
}

type ItFreelancer struct {
	acceptedGigs  []Gig
	requestedGigs []Gig
	Person
	maxGigCapacity         int
	minimumAcceptedPayment float64
}

func (i ItFreelancer) getAcceptedGigs() []Gig {
	return i.acceptedGigs
}

func (i ItFreelancer) requestServices(gig Gig) (ItFreelancer, error) {
	if len(i.requestedGigs) == i.maxGigCapacity {
		return i, fmt.Errorf("%s already has too many jobs at the moment (%d)", i.name, len(i.requestedGigs))
	}
	i.requestedGigs = append(i.requestedGigs, gig)
	return i, nil
}

func (i ItFreelancer) acceptGig(gigName string) (ItFreelancer, error) {
	for x, gig := range i.requestedGigs {
		if gig.name == gigName {
			i.acceptedGigs = append(i.acceptedGigs, gig)
			i.requestedGigs = append(i.requestedGigs[:x], i.requestedGigs[x+1:]...)
			return i, nil
		}
	}
	return i, fmt.Errorf("could not accept gig '%s' as it does not exist", gigName)
}

func main() {
	people := []Person{}
	people = append(people,
		Person{name: "Eslam Allam", age: 24, gender: MALE},
		Person{name: "Immortal", age: math.MaxInt, gender: MALE})

	for _, p := range people {
		if p.isAdult() {
			fmt.Printf("%s is an adult. His age is %d\n", p.name, p.age)
		}
		switch p.gender {
		case MALE:
			fmt.Printf("%s is a gigachad\n", p.name)
		case FEMALE:
			fmt.Printf("%s is a woman lul\n", p.name)
		}
	}

	abood := struct {
		hospital string
		Person
	}{
		Person: Person{
			name:   "Eslam Allam",
			age:    24,
			gender: FEMALE,
		},
		hospital: "Al-Kindi",
	}

	fmt.Printf("Abood is %v\n", abood)
	fmt.Printf("Abood is %d years old\n", abood.age)
	fmt.Printf("Abood works at %s hospital\n", abood.hospital)

	eslam := ItFreelancer{
		maxGigCapacity: 2,
		Person: Person{
			name:   "Eslam Allam",
			gender: MALE,
			age:    24,
		},
		minimumAcceptedPayment: 64,
	}

	fmt.Printf("%s has %d acceptedGigs and %d requestedGigs\n", eslam.name, len(eslam.acceptedGigs), len(eslam.requestedGigs))

	gig := Gig{
		name:             "Gig 1",
		description:      "Do some stuff",
		paymentInDollars: 54.35,
	}

	eslam, err := eslam.requestServices(gig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gig2 := Gig{
		name:             "Gig 2",
		description:      "Do some more stuff",
		paymentInDollars: 84.5,
	}

	eslam, err = eslam.requestServices(gig2)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, gig := range eslam.requestedGigs {
		if gig.paymentInDollars < eslam.minimumAcceptedPayment {
			continue
		}
		eslam, err = eslam.acceptGig(gig.name)
		if err != nil {
			fmt.Printf("%s couldnt accept '%s' because: %s\n", eslam.name, gig.name, err.Error())
		}

		fmt.Printf("%s has accepted '%s'\n", eslam.name, gig.name)

	}

	fmt.Printf("%s now has %d requestServices and %d acceptedServices\n", eslam.name, len(eslam.requestedGigs), len(eslam.acceptedGigs))
}
