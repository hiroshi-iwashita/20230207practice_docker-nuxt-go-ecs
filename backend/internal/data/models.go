package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package.
// It returns the type Model, which embeds all of the types we want to
// be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User:   User{},
	}
}

// Models is the type for this package. Note that any model that is
// included as a member in this type is available to us throughout the
// application, anywhere that the app variable is used, provided that the
// model is also added in the New function.
type Models struct {
	User   User
}

// User is the stucture which holds one user from the database. Note
// that it embeds a token type.
type User struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

// GetAll returns a slice of all users, sorted by last name
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM users`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
