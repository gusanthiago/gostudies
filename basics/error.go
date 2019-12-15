package main

import (
  "errors"
  "fmt"
)

type User struct {
    Name, Role, Email string
    Age               int
}

func (u User) Salary() (int, error) {
    switch u.Role {
        case "SW":
            return 100, nil
        case "CTO":
            return 500, nil
    }
    return 0, errors.New("I'm not able to handle this role")
}

func (u *User) UpdateAge(age int) {
    u.Age = age
}


func main() {
    user := User {Name: "gusanthiago", Role: "SW", Age: 30}
    salary, err := user.Salary()
    fmt.Println(salary, err)
}