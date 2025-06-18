package main

import "fmt"

// Person struct represents a contact
type Person struct {
    Name  string
    Phone string
}

// Greet method makes Person implement the Greeter interface
func (p Person) Greet() {
    fmt.Println("Hi, I'm", p.Name, "and my phone is", p.Phone)
}

// Greeter interface requires a Greet() method
type Greeter interface {
    Greet()
}

// Function to update phone number using pointer
func updatePhone(p *Person, newPhone string) {
    p.Phone = newPhone
}

func main() {
    // 1. Use a map to store contacts by name
    contactBook := make(map[string]string)
    contactBook["Alice"] = "123-4567"
    contactBook["Bob"] = "987-6543"

    fmt.Println("Original contact book:")
    fmt.Println(contactBook)

    // 2. Create a Person and greet
    person := Person{Name: "Charlie", Phone: "555-1234"}
    person.Greet()

    // 3. Update phone using pointer
    updatePhone(&person, "999-9999")
    fmt.Println("After phone update:")
    person.Greet()

    // 4. Add the updated person to the contact book
    contactBook[person.Name] = person.Phone

    // 5. Loop through and print contact book
    fmt.Println("\nUpdated contact book:")
    for name, phone := range contactBook {
        fmt.Printf("Name: %s, Phone: %s\n", name, phone)
    }

    // 6. Try using interface
    var g Greeter = person
    fmt.Print("\nCalling Greet through interface: ")
    g.Greet()
}
