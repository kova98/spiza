package data

import (
	"database/sql"
	"log"
)

func InitDb(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	init := `
        CREATE TABLE IF NOT EXISTS restaurants (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );
        
        CREATE TABLE IF NOT EXISTS menu_categories (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            restaurant_id INTEGER REFERENCES restaurants(id)
        );
        
        CREATE TABLE IF NOT EXISTS items (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            category_id INTEGER REFERENCES menu_categories(id),
            order_num INTEGER,
            price NUMERIC,
            description TEXT,
            image TEXT
        );
        `
	_, err = db.Exec(init)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
