package main

import "fmt"

func main() {
    person1 := map[string]interface{}{
        "name": "John Doe",
        "age":  30,
    }

    fmt.Println("Map:", person1)
}
