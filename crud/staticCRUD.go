package crud

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rebelnato/gosqlite3"
)

type UserData struct {
	id             int
	username       string
	passwordFromDb string
}

/*
 Below function will be used to perform user data insertion into associated db
 This can be beneficial when accompanied with a create user rest API

 Accepts user name and hashed password in ideal scenario which should be in string format
 Returns error based on execution status
*/

func InsertData(username, password string) error {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return dbError
	}
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

/*
 Below function will be used to fetch details related to a specific user
 This can be beneficial when accompanied with a user credential validation rest API

 Accepts user name in string format
 Returns data associated with user ( id , username and password )
*/

func QueryData(username string) (id int, user string, password string, err error) {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return 0, "", "", dbError
	}
	defer db.Close()

	var user1 UserData // Declaring var to store password temporarly as part of db query task

	userPassFromDB := db.QueryRow("SELECT id,username,password FROM users WHERE username = ?", username).Scan(&user1.id, &user1.username, &user1.passwordFromDb)
	if userPassFromDB != nil {
		if userPassFromDB == sql.ErrNoRows {
			log.Print("No user found")
			return 0, "404", "404", userPassFromDB
		}
	}

	return user1.id, user1.username, user1.passwordFromDb, userPassFromDB
}

/*
Below function will be used to fetch list of all users available in users table
This can be beneficial for admin user where in they want to get list of all users

Doesn't accepts any argument
Returns list of users available in "users" table and error based on execution status
*/
func QueryUserList() ([]string, error) {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return nil, dbError
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		users = append(users, username)
	}
	return users, nil
}

/*
 Below function will be used for updating the user name of an existing user
 This can be beneficial when user wants to update their user name

 Ideally this should be paired with a user validation function for user identification
 Accepts old and new user name
 Returns error based on function execution status
*/

func UpdateUsername(oldUsername, newUsername string) error {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return dbError
	}
	defer db.Close()
	id, _, _, _ := QueryData(oldUsername)
	_, err := db.Exec("UPDATE users SET username = ? WHERE id = ?", newUsername, id)
	return err
}

/*
 Below function will be used for updating the password for existing user
 This can be beneficial when user wants to update their password

 Ideally this should be paired with a user validation function for user identification
 Accepts user name and new password
 Returns error based on function execution status
*/

func UpdatePassword(username, newPassword string) error {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return dbError
	}
	defer db.Close()
	_, err := db.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, username)
	return err
}

/*
Below function will be used deleting a specific user from users table
This can be beneficial when user wants to delete his/her account

Internally performs user creds validation
Accepts username and password
Returns error based on function execution status
*/

func DeleteUser(username, password string) error {
	db, dbError := gosqlite3.ReadDbConfig()
	if dbError != nil {
		fmt.Println("Failed to read config details , stoping execution now")
		return dbError
	}
	defer db.Close()
	_, usernameFromDb, passwordFromDb, err := QueryData(username)
	if err != nil {
		fmt.Printf("Failed to fetch user data from db due to error message %q", err)
		return err
	} else {
		if username == usernameFromDb && password == passwordFromDb {
			_, err := db.Exec("DELETE FROM users where username = ?", username)
			if err != nil {
				fmt.Printf("Failed to delete row associated with user %q due to error message %q", username, err)
				return err
			}
			fmt.Printf("Successfully deleted entry for user %q", username)
		} else {
			fmt.Printf("\nProvided credentials doesn't match available credential in db")
		}
	}
	return err
}
