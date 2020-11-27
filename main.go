package main

import (
	"context"
	"log"
	"os"

	"github.com/gocql/gocql"

	"github.com/alextanhongpin/go-cassandra/domain"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	store := domain.NewUserStore(session)

	var createUser bool
	if createUser {
		u, err := store.CreateUser(context.Background(), domain.CreateUserParams{
			FirstName: "john",
			LastName:  "doe",
			Email:     "john.doe@mail.com",
			City:      "Singapore",
			Age:       21,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%#v\n", u)

	}

	log.Println("Finding users")
	users, err := store.FindUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		log.Printf("%#v\n", u)
	}

	userID, err := gocql.ParseUUID("ed1f0210-3096-11eb-8608-acde48001122")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finding user with id", userID)
	user, err := store.FindUser(context.Background(), userID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", user)
}
