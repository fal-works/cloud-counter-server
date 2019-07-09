package database

import (
	"database/sql"
	"os"
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

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS counter (
      id     SERIAL PRIMARY KEY,
      count  BIGINT DEFAULT 0
    );
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// CountValue is a value object for the count.
type CountValue struct {
	Count int64 `json:"count"`
}

// GetCount returns the current count value.
// Also inserts the row if not exists.
func GetCount(db *sql.DB) (CountValue, error) {
	countInstance := CountValue{}

	row := db.QueryRow(
		`SELECT count FROM counter WHERE id = 0`,
	)

	err := row.Scan(&countInstance.Count)

	if err != nil {
		_, err = db.Exec(`
			INSERT INTO counter DEFAULT VALUES;
		`)

		return CountValue{
			Count: 0,
		}, nil
	}

	return countInstance, nil
}
