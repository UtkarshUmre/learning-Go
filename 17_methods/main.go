package main

import "fmt"

// func main(){
// 	fmt.Println("methods in golang")
// }

func main() {

	fmt.Println("struct in go")
	Nathan := User{"Nathan", "Nathan@gmail", 23, true}
	fmt.Printf("nathans details are %+v\n", Nathan)
	fmt.Printf("name is %v and emai is %v \n", Nathan.Name, Nathan.Email)
	Nathan.GetStatus()
	Nathan.NewMail()
	fmt.Printf("name is %v and emai is %v \n", Nathan.Name, Nathan.Email)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus(){
	fmt.Println("is user active ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("email of this user is : ", u.Email)
}