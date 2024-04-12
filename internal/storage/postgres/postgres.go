package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Banner struct {
	ID        int
	Content   string
	FeatureID int
}

type Tag struct {
	ID int
}

type Feature struct {
	ID int
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		fmt.Println("Error opening database connection: ", err)
		return nil, err
	}
	defer db.Close()

	createBannerTable := `CREATE TABLE IF NOT EXISTS banners(id SERIAL PRIMARY KEY, 
    													content JSONB,
  														feature_id INT);`
	_, err = db.Exec(createBannerTable)
	if err != nil {
		fmt.Println("Error creating banners table: ", err)
		return nil, err
	}

	createTagTable := `CREATE TABLE IF NOT EXISTS tags(
  id SERIAL PRIMARY KEY
 );`
	_, err = db.Exec(createTagTable)
	if err != nil {
		fmt.Println("Error creating tags table: ", err)
		return nil, err
	}

	createFeatureTable := `CREATE TABLE IF NOT EXISTS features(
  id SERIAL PRIMARY KEY
 );`
	_, err = db.Exec(createFeatureTable)
	if err != nil {
		fmt.Println("Error creating features table: ", err)
		return nil, err
	}

	fmt.Println("Tables created successfully")
	return &Storage{db: db}, nil
}
