package db

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

type InfoForConenct struct {
	Host            string
	Port            int64
	User            string
	Password        string
	DBname          string
	ApplicationName string //тип реплики "database=%s&role=%s" master sync async
	Timeout         time.Duration
}

func ConnectDB(info InfoForConenct) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s application_name=%s", //role=%s
		info.Host, info.Port, info.User, info.Password, info.DBname, info.ApplicationName) //sslmode=disable

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("sql open : %w", err)
	}

	db.SetConnMaxLifetime(info.Timeout)
	//defer db.Close() //если раскомментить, то после выхода из функции коннект будет закрыт, поэтому данную строку нужно указывать вне функции для закрытия коннекта с бд
	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("ping : %w", errPing)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
