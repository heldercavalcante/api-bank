package repository

import (
	"context"

	"github.com/heldercavalcante/api-bank/internal/database"
	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/entity"
)


func CreateUser(user sqlc.CreateUserParams) (int64, error) {
	dbConn, err := database.GetDB()
	if err != nil {
		return 0, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	result, err :=  dt.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId() 
}

func GetUserById(userId int32) (sqlc.User, error){
	dbConn, err := database.GetDB()
	if err != nil {
		return sqlc.User{}, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	user, err :=  dt.GetUser(ctx, userId)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil 
}

func GetUsers() ([]sqlc.User, error){
	dbConn, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	users, err :=  dt.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil 
}

func DeleteUser(userId int32) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func UpdateUser(user sqlc.UpdateUserParams) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	
	return nil
}

func UserNameOrEmailExists(params sqlc.UserNameOrEmailExistsParams) (bool, error) {
	dbConn, err := database.GetDB()
	if err != nil {
		return false, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	count, err := dt.UserNameOrEmailExists(ctx, params)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}


func GetUserCompleteDataById(userId int32) (entity.UserCompleteData, error){
	dbConn, err := database.GetDB()
	if err != nil {
		return entity.UserCompleteData{}, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	user, err :=  dt.GetUserCompleteDataById(ctx, userId)
	if err != nil {
		return entity.UserCompleteData{}, err
	}

	address := entity.Address{
		StreetAddress: user.StreetAddress,
		Complement: user.Complement.String,
		District: user.District.String,
		Number: int64(user.HouseNumber.Int32),
		City: user.City,
		Zone: user.Zone,
		PostalCode: user.PostalCode,
		Country: user.Country,
	}

	userCompleteData := entity.UserCompleteData{
		UserID: int64(user.UserID),
		UserName: user.Username,
		FirstName: user.FirstName.String,
		LastName: user.LastName.String,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber.String,
		Address: address,
		DateOfBirth: user.DateOfBirth.Time.Format("2006-01-02"),
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02"),
	}

	return userCompleteData, nil 
}

func GetUsersCompleteData() ([]entity.UserCompleteData, error){
	dbConn, err := database.GetDB()
	if err != nil {
		return []entity.UserCompleteData{}, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	users, err :=  dt.GetUsersCompleteData(ctx)
	if err != nil {
		return []entity.UserCompleteData{}, err
	}

	var usersCompleteData []entity.UserCompleteData
	for _, user := range users {
		address := entity.Address{
			StreetAddress: user.StreetAddress,
			Complement: user.Complement.String,
			District: user.District.String,
			Number: int64(user.HouseNumber.Int32),
			City: user.City,
			Zone: user.Zone,
			PostalCode: user.PostalCode,
			Country: user.Country,
		}
	
		userCompleteData := entity.UserCompleteData{
			UserID: int64(user.UserID),
			UserName: user.Username,
			FirstName: user.FirstName.String,
			LastName: user.LastName.String,
			Email: user.Email,
			PhoneNumber: user.PhoneNumber.String,
			Address: address,
			DateOfBirth: user.DateOfBirth.Time.Format("2006-01-02"),
			CreatedAt: user.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02"),
		}
		usersCompleteData = append(usersCompleteData, userCompleteData)
	}

	return usersCompleteData, nil 
}