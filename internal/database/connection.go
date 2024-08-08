package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/heldercavalcante/api-bank/internal/configs"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *Connection
	connErr error
	once    sync.Once
)

type Connection struct {
	DB *sql.DB
}


func NewConnection() error {
	dbConfigs := configs.GetDBConfigs()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfigs.User, dbConfigs.Pass, dbConfigs.Host, dbConfigs.Port,dbConfigs.Database)
	db, err := sql.Open("mysql", connectionString)
    if err != nil {
		return err
    }

	db.SetMaxOpenConns(25) // Maximum number of open connections to the database
    db.SetMaxIdleConns(25) // Maximum number of connections in the idle connection pool
    db.SetConnMaxLifetime(0) // Maximum amount of time a connection may be reused (0 means unlimited)

	if err := db.Ping(); err != nil {
		return err
	}

	Conn = &Connection{DB: db}

	return nil
}

func InitDBConnection() {
    once.Do(func() {
        connErr = NewConnection()
    })
}

func GetDB() (*sql.DB, error) {
	if Conn == nil {
		InitDBConnection()
		if connErr != nil {
			log.Printf("failed to initialize database connection: %v", connErr)
			return nil, connErr
		}
	}
	return Conn.DB, nil
}