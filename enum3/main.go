package main

import "fmt"

type UserRole int

const (
	Admin UserRole = iota
	Moderator
	User
	Guest
)

// Kullanıcı rolünü string olarak döndüren bir String () metodu tanımlıyoruz...

func (r UserRole) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Moderator:
		return "Moderator"
	case User:
		return "User"
	case Guest:
		return "guest"

	default:
		return "Unknow"
	}

} // Kullanıcı rolleri için fonksiyon

func getPermissions(role UserRole) string {
	switch role {
	case Admin:
		return "Full access to all resources"
	case Moderator:
		return "Access to moderate content"
	case User:
		return "Access to own content"
	case Guest:
		return "Limited access to public content"

	default:
		return "No access"
	}

}
func main() {
	role := Admin
	fmt.Println("Role..: ", role)
	r2 := getPermissions(role)
	fmt.Println("Permission..:", r2)

}
