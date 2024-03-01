package constant

type UserRole struct {
	Admin string
	User  string
}

var Role = UserRole{
	Admin: "admin",
	User:  "user",
}
