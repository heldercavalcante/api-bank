package service

import (
	"database/sql"
	"time"

	"github.com/heldercavalcante/api-bank/internal/common"
	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/entity"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/repository"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(requestData entity.UserRequest) (int64, error) {
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

func GetUserById(userId int32) (sqlc.User, error){
	user, err := repository.GetUserById(userId)
	if err != nil {
		return sqlc.User{}, err
	}
	
	return user, nil	
}

func GetUsers() ([]sqlc.User, error){
	users, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}
	
	return users, nil
}

func Validate(requestData *entity.UserRequest) (common.ValidationReturn ,error) {
	if requestData.UserName == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "username is required",
		}, nil
	}
	if requestData.Password == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "password is required",
		}, nil
	}
	if requestData.Email == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "email is required",
		}, nil
	}
	if requestData.FirstName == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "firstName is required",
		}, nil
	}
	if requestData.LastName == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "lastName is required",
		}, nil
	}
	if requestData.DateOfBirth == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "dateOfBirth is required",
		}, nil
	}
	if 
		requestData.Address.StreetAddress == "" || 
		requestData.Address.City == "" || 
		requestData.Address.PostalCode == "" || 
		requestData.Address.Zone == "" ||
		requestData.Address.Country == "" {
		return common.ValidationReturn{
			Valid: false,
			Message: "complete address with streetAddress, city, zone, country, and postalCode is required",
		}, nil
	}

	exists, err := repository.UserNameOrEmailExists(sqlc.UserNameOrEmailExistsParams{
		Username: requestData.UserName,
		Email: requestData.Email,
	})
	if err != nil {
		return common.ValidationReturn{}, err
	}
	if exists {
		return common.ValidationReturn{
			Valid: false,
			Message: "UserName or Email is already being used!",
		}, nil
	}

	return common.ValidationReturn{
		Valid: true,
		Message: "Valid!",
	}, nil
}


func GetUserCompleteDataById(userId int32) (entity.UserCompleteData, error){
	user, err := repository.GetUserCompleteDataById(userId)
	if err != nil {
		return entity.UserCompleteData{}, err
	}
	
	return user, nil	
}

func GetUsersCompleteData() ([]entity.UserCompleteData, error){
	users, err := repository.GetUsersCompleteData()
	if err != nil {
		return []entity.UserCompleteData{}, err
	}
	
	return users, nil	
}

func DeleteUser(userId int32) error{
	err := repository.DeleteUserProfile(userId)
	if err != nil {
		return err
	}

	err = repository.DeleteUserAddress(userId)
	if err != nil {
		return err
	}

	err = repository.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(userRequest entity.UserRequest) error {

	parsedDateOfBirth, err := time.Parse("2006-01-02", userRequest.DateOfBirth)
	if err != nil {
		return err
	}
	userProfile := sqlc.UpdateUserProfileParams{
		FirstName: sql.NullString{String: userRequest.FirstName, Valid: true},
		LastName: sql.NullString{String: userRequest.LastName, Valid: true},
		PhoneNumber: sql.NullString{String: userRequest.PhoneNumber, Valid: true},
		DateOfBirth: sql.NullTime{Time: parsedDateOfBirth, Valid: true},
		UserID: int32(userRequest.UserID),
	}
	err = repository.UpdateUserProfile(userProfile)
	if err != nil {
		return err
	}

	userAddress := sqlc.UpdateUserAddressParams{
		StreetAddress: userRequest.Address.StreetAddress,
		HouseNumber: sql.NullInt32{Int32: int32(userRequest.Address.Number), Valid: true},
		Complement: sql.NullString{String: userRequest.Address.Complement, Valid: true},
		City: userRequest.Address.City,
		Zone: userRequest.Address.Zone,
		District: sql.NullString{String: userRequest.Address.District, Valid: true},
		PostalCode: userRequest.Address.PostalCode,
		Country: userRequest.Address.Country,
		UserID: int32(userRequest.UserID),
	}
	err = repository.UpdateUserAddress(userAddress)
	if err != nil {
		return err
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)	
	if err != nil {
		return err
	}
	
	curDate := time.Now()

	user := sqlc.UpdateUserParams{
		Username: userRequest.UserName,
		PasswordHash: string(encpw),
		Email: userRequest.Email,
		UpdatedAt: sql.NullTime{Time: curDate, Valid: true},
		UserID: int32(userRequest.UserID),
	}

	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}