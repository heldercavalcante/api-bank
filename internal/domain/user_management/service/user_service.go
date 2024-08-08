package service

import (
	"database/sql"
	"time"

	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/entity"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/repository"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(requestData entity.CreateUserRequest) (int64, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)	
	if err != nil {
		return 0, err
	}

	parsedDateOfBirth, err := time.Parse("2006-01-02", requestData.DateOfBirth)
	if err != nil {
		return 0, err
	}
	curDate := time.Now()

	user := sqlc.CreateUserParams{
		Username: requestData.UserName,
		PasswordHash: string(encpw),
		Email: requestData.Email,
		CreatedAt: sql.NullTime{Time: curDate, Valid: true},
		UpdatedAt: sql.NullTime{Time: curDate, Valid: true},
	}

	userId, err := repository.CreateUser(user)
	if err != nil {
		return 0, err
	}

	userAddress := sqlc.CreateUserAddressParams{
		UserID: int32(userId),
		StreetAddress: requestData.Address.StreetAddress,
		HouseNumber: sql.NullInt32{Int32: int32(requestData.Address.Number), Valid: true},
		Complement: sql.NullString{String: requestData.Address.Complement, Valid: true},
		City: requestData.Address.City,
		Zone: requestData.Address.Zone,
		District: sql.NullString{String: requestData.Address.District, Valid: true},
		PostalCode: requestData.Address.PostalCode,
		Country: requestData.Address.Country,
	}

	userAddressId, err := repository.CreateUserAddress(userAddress)
	if err != nil {
		return 0, err
	}

	userProfile := sqlc.CreateUserProfileParams{
		UserID: int32(userId),
		FirstName: sql.NullString{String: requestData.FirstName, Valid: true},
		LastName: sql.NullString{String: requestData.LastName, Valid: true},
		PhoneNumber: sql.NullString{String: requestData.PhoneNumber, Valid: true},
		AddressID: int32(userAddressId),
		DateOfBirth: sql.NullTime{Time: parsedDateOfBirth, Valid: true},
		CreatedAt: sql.NullTime{Time: curDate, Valid: true},
		UpdatedAt: sql.NullTime{Time: curDate, Valid: true},
	}
	_, err = repository.CreateUserProfile(userProfile)
	if err != nil {
		return 0, err
	}
	
	return userId, nil	
}