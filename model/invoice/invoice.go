// Package invoice provides access to the invoice table in the MySQL database.
package invoice

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	table = "invoice"
)

// Item defines the model.
type Item struct {
	ID        uint32         `db:"id"`
	Name      string         `db:"content"`
	UserID    uint32         `db:"user_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByID gets an invoice by ID.
func ByID(db Connection, ID string, userID string) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
		SELECT id, content, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ID, userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserID gets all invoices for a user.
func ByUserID(db Connection, userID string) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
		SELECT id, content, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDPaginate gets invoices for a user based on page and max variables.
func ByUserIDPaginate(db Connection, userID string, max int, page int) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
		SELECT id, content, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		LIMIT %v OFFSET %v
		`, table, max, page),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDCount counts the number of invoices for a user.
func ByUserIDCount(db Connection, userID string) (int, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
		SELECT count(*)
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err
}

// Create adds an invoice.
func Create(db Connection, content string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(content, user_id)
		VALUES
		(?,?)
		`, table),
		content, userID)
	return result, err
}

// Update makes changes to an existing invoice.
func Update(db Connection, content string, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET content = ?
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		content, ID, userID)
	return result, err
}

// DeleteHard deletes an invoice.
func DeleteHard(db Connection, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		DELETE FROM %v
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		`, table),
		ID, userID)
	return result, err
}

// DeleteSoft marks an invoice as deleted.
func DeleteSoft(db Connection, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET deleted_at = NOW()
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ID, userID)
	return result, err
}
