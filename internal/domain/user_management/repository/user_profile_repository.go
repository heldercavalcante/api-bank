package repository

import (
	"context"

	"github.com/heldercavalcante/api-bank/internal/database"
	sqlc "github.com/heldercavalcante/api-bank/internal/database/sqlc"
)

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

func DeleteUserProfile(userId int32) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.DeleteUsersProfileByUserId(ctx, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func UpdateUserProfile(userProfile sqlc.UpdateUserProfileParams) error {
	dbConn, err := database.GetDB()
	if err != nil {
		return err
	}

	dt := sqlc.New(dbConn)
	ctx := context.Background()
	err =  dt.UpdateUserProfile(ctx, userProfile)
	if err != nil {
		return err
	}
	
	return nil
}