package repository

import (
	"context"

	"github.com/heldercavalcante/api-bank/internal/database"
	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
)


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

func DeleteUserAddress(userId int32) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.DeleteUsersAddressByUserId(ctx, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func UpdateUserAddress(userAddress sqlc.UpdateUserAddressParams) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.UpdateUserAddress(ctx, userAddress)
	if err != nil {
		return err
	}
	
	return nil
}