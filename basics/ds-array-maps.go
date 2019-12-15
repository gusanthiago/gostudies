package main

import "fmt"

func main() {

  // array
    values := []int{2, 2, 44, 12}
    values = append(values, (10))

    for _, number := range values {
        fmt.Println("Number ", number)
    }

    slicedValues := values[1:3]  

    for _, number := range slicedValues {
        fmt.Println("Number Sliced ", number)
    }

    // map
    mapTest := map[string]string {
        "test": "TestGolang",
        "map": "I'm the map",
    }

    tVar, exists := mapTest["testa"]
    if exists != true {
        fmt.Println(tVar)
        fmt.Println("Is key is not exists")
    }

    fmt.Println(mapTest["test"])

    for key, value := range mapTest {
        fmt.Println("Key", key, "Value", value)
    }

}