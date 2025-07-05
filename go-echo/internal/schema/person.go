package schema

// TODO: do it properly with Gorm
type Person struct{}

type PersonSchema struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Version   uint   `json:"version,omitempty"`
}

func (p *Person) New() PersonSchema {
	return PersonSchema{
		UserID:    1,
		Email:     "maail@mail.com",
		FirstName: "John",
		LastName:  "Doe",
		Version:   1,
	}
}
