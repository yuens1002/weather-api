package main

import (
	"database/sql"
	"fmt"
	"log"

	// postgres driver
	graphql "github.com/graph-gophers/graphql-go"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	c "github.com/yuens1002/server-graphql-go-weather/config"
	util "github.com/yuens1002/server-graphql-go-weather/utils"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// ConnectToDB creates a new DB connection
func ConnectToDB(connString string) *Db {
	db, err := sql.Open("postgres", connString)
	util.Check(err, "sql.Open")

	// Check that our connection is good
	err = db.Ping()
	util.Check(err, "DB.Ping")
	log.Println("Database is connected ")
	return &Db{db}
}

// ConnString returns a connection string based on the parameters it's given
func ConnString(s c.DBcfg) string {
	fmt.Println(s.Host)
	if viper.GetString("HOST") == "localhost" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			s.Host, s.Port, s.User, s.Password, s.DBName)
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		s.Host, s.Port, s.User, s.Password, s.DBName)

}

// Users corresponds to Users query from graphql
func (d *Db) Users() []User {
	var (
		rows  *sql.Rows
		err   error
		u     User   // Create User struct for holding each row's data
		users []User // Create slice of Users for our response
	)
	rows, err = d.Query(`SELECT * FROM users`)
	util.Check(err, "rows")

	defer rows.Close()

	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() { // order matters here from db table
		err = rows.Scan(
			&u.UserID,
			&u.Username,
			&u.Email,
			&u.Password,
		)
		util.Check(err, "rows.Next")
		users = append(users, u)
	}

	return users
}

// User returns a single user
func (d *Db) User(uid graphql.ID) (User, error) {
	var (
		sqlStatement = `SELECT * FROM users WHERE user_id=$1;`
		row          *sql.Row
		err          error
		u            User
	)
	row = d.QueryRow(sqlStatement, uid)
	err = row.Scan(
		&u.UserID,
		&u.Username,
		&u.Email,
		&u.Password,
	)
	util.Check(err, "row.Scan")
	return u, nil
}

// CreateUser - inserts a new user
// *User --> in case of error to return nil
func (d *Db) CreateUser(i *SignUpArgs) (*User, error) {
	var (
		sqlStatement = `
			INSERT INTO users (email, username, password)
			VALUES ($1, $2, $3)
			RETURNING user_id`
		userID graphql.ID
		row    *sql.Row
		err    error
		u      User
	)
	/***************************************************************************
		* retrieve the UserID of the newly inserted record
		* db.Exec() requires the Result interface with the
				LastInsertId() method which relies on a returned value from postgresQL
		* lib/pq does not however return the last inserted record
	****************************************************************************/
	row = d.QueryRow(sqlStatement, i.Email, i.Username, i.Password)
	if err = row.Scan(&userID); err != nil {
		// err: username or email is not unqiue --> user already exsits
		return nil, err
	}
	u, _ = d.User(userID)
	return &u, nil
}
