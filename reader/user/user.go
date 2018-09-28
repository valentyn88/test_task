package user

// User struct for describing user
type User struct {
	ID string
	Name string
	Email string
	Mobile_number string
}

// NewUser create new user
func NewUser(row []string) User {
	u := User{}
	if row[0] != "" {
		u.ID = row[0]
	}
	if row[1] != "" {
		u.Name = row[1]
	}
	if row[2] != "" {
		u.Email = row[2]
	}
	if row[3] != "" {
		u.Mobile_number = row[3]
	}

	return u
}