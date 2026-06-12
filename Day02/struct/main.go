package main

import "fmt"

// basic struct
type Person struct {

	// so basically when we initialized keys by simple letters it normally private. It also method name is simple letter it is private
	name        string
	age         int
	designation string
}

// struct inside struct withut slide
type Hobby struct {
	name       string
	hobby_type string
}

type User struct {
	name  string
	age   int
	job   string
	hobby Hobby
}

// getter with struct
func (u User) getName() string {
	return u.name
}

// setter with struct
func (u User) setName(name string) {
	u.name = name
	fmt.Println(u.name)
}

func main() {
	// basic struct call
	person1 := Person{name: "Kavishka", age: 26, designation: "DevOps Engineer"}
	fmt.Println(person1)

	// struct in struct without slide
	User := User{name: "Kavishka", age: 26, job: "Platform Engineer", hobby: Hobby{name: "Collecting Stamps", hobby_type: "indoor"}}
	fmt.Println(User)

	// struct with getter and setter
	// 1. getter
	username := User.getName()
	fmt.Println(username)

	// 2. setter
	User.setName("Hashini")
}
