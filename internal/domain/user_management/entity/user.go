package entity

type Address struct {
	StreetAddress string `json:"streetAddress"`
	Number int64 `json:"number"`
	Complement string `json:"complement"`
	City string `json:"city"`
	Zone string `json:"zone"`
	District string `json:"district"`
	PostalCode string `json:"postalCode"`
	Country string `json:"country"`
}

type UserRequest struct {
	UserID int64 `json:"userId"`
	UserName string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address Address `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserCompleteData struct {
	UserID int64 `json:"userId"`
	UserName string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address Address `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}