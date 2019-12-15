package main

import "fmt"

type User struct {
    Name, Role, Email string
    Age               int
}

func (u User) Salary() int {
    switch u.Role {
        case "SW":
            return 100
        case "CTO":
            return 500
    }
    return 0
}

func (u *User) UpdateAge(age int) {
    u.Age = age
}


func main() {
    user := User {Name: "gusanthiago", Role: "SW", Age: 30}
    test := User {Name: "gusanthiago", Role: "CTO", Age: 30}
    fmt.Println(user.Name, user.Age)
    fmt.Println(test.Name, test.Age)
    test.UpdateAge(10)
    fmt.Println(test.Name, test.Age)
}
