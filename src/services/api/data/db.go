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

	return db
}

/*


CREATE TABLE restaurants (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE menu_categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    restaurant_id INTEGER REFERENCES restaurants(id)
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category_id INTEGER REFERENCES menu_categories(id),
    order_num INTEGER,
    price NUMERIC,
    description TEXT,
    image TEXT
);

*/
