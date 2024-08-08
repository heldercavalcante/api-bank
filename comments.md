#Go Commands
    -install mysql drivers
    go get -u github.com/go-sql-driver/mysql

    #install importeds
    go mod tidy

#dbMate Commands
    -Install
    curl -fsSL https://github.com/amacneil/dbmate/releases/download/v1.12.0/dbmate-linux-amd64 -o dbmate
    chmod +x dbmate
    sudo mv dbmate /usr/local/bin/dbmate

    -Need to create a .env file with the db base_ul
    DATABASE_URL="mysql://user:password@tcp(localhost:3306)/dbname"

    -Create a migration
    dbmate new create_users_table

    -Update Table
    dbmate new add_age_to_users

    -Run the migration
    dbmate up

    -Rollback
    dbmate down


#sqlc commands
    -Install
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

    -Gerar
    sqlc generate
