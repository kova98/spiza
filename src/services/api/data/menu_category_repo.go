package data

import "database/sql"

type MenuCategoryRepo struct {
	db *sql.DB
}

func NewMenuCategoryRepo(db *sql.DB) *MenuCategoryRepo {
	return &MenuCategoryRepo{db}
}

func (r *MenuCategoryRepo) CreateMenuCategory(menuCategory MenuCategory) error {
	_, err := r.db.Exec(`INSERT INTO menu_categories (name, restaurant_id) VALUES ($1, $2)`, menuCategory.Name, menuCategory.RestaurantId)
	if err != nil {
		return err
	}
	return nil
}

func (r *MenuCategoryRepo) DeleteMenuCategory(id int64) error {
	_, err := r.db.Exec(`DELETE FROM menu_categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
