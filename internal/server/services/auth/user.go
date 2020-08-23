package auth

type Role string

const (
	AdminRole     Role = "Admin"
	PowerUserRole      = "PowerUser"
	UserRole           = "User"
)

type User struct {
	ID         int
	Username   string
	Password   string
	Role       Role
	Department string
}

// Fake user database... no storage, management, or password encryption
var users = []User{
	{1, "admin", "1234", AdminRole, "IT"},
	{2, "poweruser", "2345", PowerUserRole, "Accounting"},
	{3, "user", "3456", UserRole, "Sales"},
}

func AuthenticateUser(username, password string) *User {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user
		}
	}
	return nil
}

func GetUser(id int) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
