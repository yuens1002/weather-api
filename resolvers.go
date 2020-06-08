package main

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

/*
 * 	User GQL type

		type User {
		userID: ID!
		username: String!
		email: String!
		password: String!
	}
*/

// User type should match the exact shape of the schema commented above
type User struct {
	UserID   graphql.ID
	Username string
	Email    string
	Password string
}

// RootResolver ingests Db to run queries (getters) against it
type RootResolver struct {
	*Db
}

// String prints pretty structs ie. fmt.Printf doesn't do it on its own
func (r *UserResolver) String() string {
	return fmt.Sprintf("User{userID: %s, username: %s, email: %s, password: %s}",
		r.u.UserID, r.u.Username, r.u.Email, r.u.Password,
	)
}

// Users returns all users from Db
func (r *RootResolver) Users() ([]*UserResolver, error) {
	var userRxs []*UserResolver
	users := r.Db.Users()
	for _, u := range users {
		u := u
		userRxs = append(userRxs, &UserResolver{&u})
	}
	return userRxs, nil
}

// User returns a single user from Db
func (r *RootResolver) User(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	// Find user:
	u, err := r.Db.User(args.UserID)
	if err != nil {
		// Didn’t find user:
		return nil, nil
	}
	return &UserResolver{&u}, nil
}

// UserResolver ingests properties from User
type UserResolver struct{ u *User }

// UserID returns the userId of the user
func (r *UserResolver) UserID() graphql.ID {
	return r.u.UserID
}

// Username returns the username of the user
func (r *UserResolver) Username() string {
	return r.u.Username
}

// Email returns the email of the user
func (r *UserResolver) Email() string {
	return r.u.Email
}

// Password returns the password of the user
func (r *UserResolver) Password() string {
	return r.u.Password
}

/*************************************************************************
qgl syntax

mutation {
  signUp(signUpInput: {
		email: "newuser10@gmail.com",
			username: "newUser10",
			password: "asdfasdfawerawer",
		}) {
    addedUser {
      email
      userID
      username
    }
    ok
    error
  }
}

*************************************************************************/

// SignUpArgs provides the shape of the input types
type SignUpArgs struct {
	Username string
	Email    string
	Password string
}

// SignUp returns a new User from Db and its responses
func (r *RootResolver) SignUp(args struct{ SignUpInput *SignUpArgs }) (*SignUpResolver, error) {
	// Find user:
	u, err := r.Db.CreateUser(args.SignUpInput)
	// need to deal with this different, so sort of error if we can't create the user
	// a. user already exists
	// b. email already exists
	if err != nil {
		// error creating the user
		msg := "already signed up"
		return &SignUpResolver{
			Status: false,
			Msg:    &msg,
			User:   nil,
		}, nil
	}

	return &SignUpResolver{
		Status: true,
		Msg:    nil,
		User:   &UserResolver{u},
	}, nil
}

// SignUpResolver is the response type
type SignUpResolver struct {
	Status bool
	Msg    *string
	User   *UserResolver
}

// Ok for SignUpResponse
func (r *SignUpResolver) Ok() bool {
	return r.Status
}

// Error for SignUpResponse
func (r *SignUpResolver) Error() *string {
	return r.Msg
}

// AddedUser for SignUpResponse
func (r *SignUpResolver) AddedUser() *UserResolver {
	return r.User
}
