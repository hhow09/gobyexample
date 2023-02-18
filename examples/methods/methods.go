// Go supports _methods_ defined on struct types.

// see [pointer vs. value receiver](https://github.com/hhow09/gobyexample/blob/master/docs/pointer-value-receiver.md)

package main

import "fmt"

type user struct {
	name  string
	email string
}

// value receiver will only works on the copy of struct
func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
	// will not change because this method is a value receivere.
	// warning: ineffective assignment to field user.name
	u.name = "will not change"
}

func (u *user) changeEmail(email string) {
	// will de-reference to (*u).email under the hood
	u.email = email
}

func main() {
	ua := user{name: "Tom", email: "tom@gmail.com"}
	ub := user{name: "Ben", email: "ben@gmail.com"}
	ua.notify()
	fmt.Println("ua.name", ua.name)
	ua.changeEmail("tom@newdomain.com")
	fmt.Println("ua.email", ua.email)
	ub.changeEmail("ben@newdomain.com")
	fmt.Println("ub.email", ub.email)
}
