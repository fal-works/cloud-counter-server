package database

import (
	"database/sql"
	"math"
	"os"
	"strconv"
)

// Connect returns a new connection to the database.
// Also creates `counter` table if not exists.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	err = prepareTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// CountValue is a value object for the count.
type CountValue struct {
	Count int `json:"count"`
}

func prepareTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS counter (
			id     SERIAL PRIMARY KEY,
			count  INTEGER DEFAULT 0
		);
	`)
	return err
}

func insertRow(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO counter DEFAULT VALUES;`)
	return err
}

func selectRow(db *sql.DB) *sql.Row {
	return db.QueryRow(`SELECT count FROM counter WHERE id = 1;`)
}

func updateRow(db *sql.DB, countInstance CountValue) error {
	_, err := db.Exec(`UPDATE counter SET count = ` + strconv.Itoa(countInstance.Count) + ` WHERE id = 1`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return err
}

// GetCount returns the current count value.
// Also inserts the row if not exists.
func GetCount(db *sql.DB) CountValue {
	countInstance := CountValue{}

	err := selectRow(db).Scan(&countInstance.Count)

	if err != nil {
		err = insertRow(db)

		if err != nil {
			return CountValue{
				Count: -1,
			}
		}

		return CountValue{
			Count: 0,
		}
	}

	return countInstance
}

// GetIncrementedCount increments the current value and returns it.
// Also inserts the row if not exists.
func GetIncrementedCount(db *sql.DB) CountValue {
	countInstance := GetCount(db)

	if countInstance.Count >= math.MaxInt32 {
		countInstance.Count = 0
	} else {
		countInstance.Count++
	}

	updateRow(db, countInstance)

	return countInstance
}
