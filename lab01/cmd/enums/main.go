package main

type Role int

const (
	User Role = 1 << iota
	Manager
	Admin
	Customer
	RoleMask = (1 << iota) - iota
)

type Status int

const (
	Registered = iota
	Active
	Disabled
)

func (r Role) String() string {
	switch r {
	case User:
		return "User"
	case Manager:
		return "Manager"
	case Admin:
		return "Admmin"
	case Customer:
		return "Customer"
	default:
		return "Invalid role"
	}
}

func main() {

}
