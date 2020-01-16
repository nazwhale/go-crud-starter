package dao

import (
	"database/sql"
	"fmt"
	"os"
)

// Create

func CreateWork(name, imageURL string, artistID int) (int, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	return createWork(db, name, imageURL, artistID), nil
}

func createWork(db *sql.DB, name, imageURL string, artistID int) int {
	sqlStatement := `
INSERT INTO works (name, image_url, artist_id)
VALUES ($1, $2, $3)
RETURNING id`
	id := 0

	err := db.QueryRow(sqlStatement, name, imageURL, artistID).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

// List

func ListWorks(limit int) ([]Work, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return listWorks(db, limit), nil
}

func listWorks(db *sql.DB, limit int) []Work {
	sqlStatement := `
SELECT *
FROM works
LIMIT $1;`

	rows, err := db.Query(sqlStatement, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	works := []Work{}
	for rows.Next() {
		var work Work
		err = rows.Scan(&work.ID, &work.Name, &work.ImageURL, &work.ArtistID)
		if err != nil {
			panic(err)
		}
		works = append(works, work)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return works
}

// Read

// consider returning a dao object with these CRUD methods on it
func ReadWork(workID int) (Work, error) {

	// consider popping this in main()
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return Work{}, err
	}
	defer db.Close()

	return readWork(db, workID), nil
}

func readWork(db *sql.DB, workID int) Work {
	sqlStatement := `
SELECT *
FROM works
WHERE id = $1;`

	var work Work
	row := db.QueryRow(sqlStatement, workID)
	switch err := row.Scan(&work.ID, &work.Name, &work.ImageURL, &work.ArtistID); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned!")
	case nil:
		fmt.Printf("%+v", work)
	default:
		panic(err)
	}

	return work
}

// Update

func UpdateWork(id, artistID int, name, imageURL string) (Work, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return Work{}, err
	}
	defer db.Close()

	return updateWork(db, id, artistID, name, imageURL), nil
}

func updateWork(db *sql.DB, id, artistID int, name, imageURL string) Work {
	sqlStatement := `
UPDATE works
SET name = $2, image_url = $3, artist_id = $4
WHERE id = $1
RETURNING *;`

	var work Work
	err := db.QueryRow(sqlStatement, id, name, imageURL, artistID).Scan(&work.ID, &work.ImageURL, &work.Name, &work.ArtistID)
	if err != nil {
		panic(err)
	}

	return work
}

// Delete

func deleteWork(db *sql.DB) {
	sqlStatement := `
DELETE from works
WHERE id = $1;`

	rsp, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := rsp.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("rows affected: ", count)
}
