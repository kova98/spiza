package data

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InitDb(connStr string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", connStr)
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
            category_id INTEGER REFERENCES menu_categories(id),
            name TEXT NOT NULL DEFAULT '',
            order_num INTEGER DEFAULT 0,
            price NUMERIC NOT NULL DEFAULT 0,
            description TEXT NOT NULL DEFAULT '',
            image TEXT NOT NULL DEFAULT ''
        );        
        
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			restaurant_id INTEGER NOT NULL REFERENCES restaurants(id),
			status INTEGER NOT NULL DEFAULT 0,
			items INTEGER[] NOT NULL,
			date_created timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
		);
        `
	_, err = db.Exec(init)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
