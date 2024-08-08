package repository

import (
	"context"

	"github.com/heldercavalcante/api-bank/internal/database"
	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
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

func CreateUserAddress(userAddress sqlc.CreateUserAddressParams) (int64, error) {
	dbConn, err := database.GetDB()
	if err != nil {
		return 0, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	result, err := dt.CreateUserAddress(ctx, userAddress)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId() 
}

func CreateUserProfile(userProfile sqlc.CreateUserProfileParams) (int64, error) {
	dbConn, err := database.GetDB()
	if err != nil {
		return 0, err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	result, err := dt.CreateUserProfile(ctx, userProfile)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId() 
}