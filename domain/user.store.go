package domain

import (
	"context"

	"github.com/gocql/gocql"
)

type UserStore struct {
	session *gocql.Session
}

func NewUserStore(session *gocql.Session) *UserStore {
	return &UserStore{session}
}

type CreateUserParams struct {
	FirstName string
	LastName  string
	Email     string
	City      string
	Age       int
}

func (s *UserStore) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	var u User
	userID := gocql.TimeUUID()

	if err := s.session.Query(`
	INSERT INTO blog.users(id, firstname, lastname, email, city, age)
	VALUES (?, ?, ?, ?, ?, ?)
	`,
		userID,
		params.FirstName,
		params.LastName,
		params.Email,
		params.City,
		params.Age,
	).Exec(); err != nil {
		return u, err
	}

	u.ID = userID
	u.FirstName = params.FirstName
	u.LastName = params.LastName
	u.Email = params.Email
	u.City = params.City
	u.Age = params.Age

	return u, nil
}

func (s *UserStore) FindUsers(ctx context.Context) ([]User, error) {
	var users []User
	iter := s.session.Query(`SELECT * FROM blog.users`).Iter()
	for {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		users = append(users, User{
			ID:        row["id"].(gocql.UUID),
			FirstName: row["firstname"].(string),
			LastName:  row["lastname"].(string),
			Email:     row["email"].(string),
			City:      row["city"].(string),
			Age:       row["age"].(int),
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStore) FindUser(ctx context.Context, id gocql.UUID) (User, error) {

	var u User
	row := make(map[string]interface{})
	if err := s.session.Query(`
	SELECT *
	FROM blog.users
	WHERE id = ?
	`, id).Consistency(gocql.One).MapScan(row); err != nil {
		return u, err
	}
	u = User{
		ID:        row["id"].(gocql.UUID),
		FirstName: row["firstname"].(string),
		LastName:  row["lastname"].(string),
		Email:     row["email"].(string),
		City:      row["city"].(string),
		Age:       row["age"].(int),
	}
	return u, nil
}
