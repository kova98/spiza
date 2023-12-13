package data

import "github.com/jmoiron/sqlx"

type MenuCategoryRepo struct {
	db *sqlx.DB
}

func NewMenuCategoryRepo(db *sqlx.DB) *MenuCategoryRepo {
	return &MenuCategoryRepo{db}
}

func (r *MenuCategoryRepo) CreateMenuCategory(menuCategory *MenuCategory) (int64, error) {
	var id int64
	err := r.db.QueryRow(`
		INSERT INTO menu_categories (name, restaurant_id) VALUES ($1, $2) RETURNING id`,
		menuCategory.Name, menuCategory.RestaurantId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MenuCategoryRepo) DeleteMenuCategory(id int64) error {
	_, err := r.db.Exec(`DELETE FROM menu_categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
