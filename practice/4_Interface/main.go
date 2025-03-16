package main

import "fmt"

// Define an interface for playing an instrument
type Musician interface {
	PlayInstrument()
}

// Struct for a Drummer
type Drummer struct {
	Name string
}

// Struct for a Pianist
type Pianist struct {
	Name string
}

// Method for Drummer implementing Musician interface
func (d Drummer) PlayInstrument() {
	fmt.Printf("%s is rocking the drums!\n", d.Name)
}

// Method for Pianist implementing Musician interface
func (p Pianist) PlayInstrument() {
	fmt.Printf("%s is tickling the ivories on the piano!\n", p.Name)
}

// Function to describe any value
func Describe(a interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", a, a)
}

func main() {
	// Basic variable and function call
	instrumentCount := 2
	fmt.Println("Starting the band with", instrumentCount, "players!")
	Describe(instrumentCount)

	// Create a Drummer
	var drummer Drummer
	drummer.Name = "Alex"
	drummer.PlayInstrument()

	// Create a Pianist
	var pianist Pianist
	pianist.Name = "Sophie"
	pianist.PlayInstrument()

	// Use the interface to treat them polymorphically
	var band []Musician = []Musician{drummer, pianist}
	fmt.Println("\nFull band performance:")
	for _, member := range band {
		member.PlayInstrument()
	}
}