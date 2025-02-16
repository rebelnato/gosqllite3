package crud

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserData struct {
	id             int
	username       string
	passwordFromDb string
}

// SQLlite 3 db related functions
func InsertData(db *sql.DB, username, password string) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

func QueryData(db *sql.DB, username string) (id int, user string, password string, err error) {

	// Declaring var to store password temporarly as part of db query task
	var user1 UserData

	// Fetch user password from db
	userPassFromDB := db.QueryRow("SELECT id,username,password FROM users WHERE username = ?", username).Scan(&user1.id, &user1.username, &user1.passwordFromDb)
	if userPassFromDB != nil {
		if userPassFromDB == sql.ErrNoRows {
			log.Print("No user found")
			return 0, "404", "404", userPassFromDB
		}
	}

	return user1.id, user1.username, user1.passwordFromDb, userPassFromDB
}

func QueryUserList(db *sql.DB) ([]string, error) {
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

func UpdateUsername(db *sql.DB, oldUsername, newUsername string) error {
	id, _, _, _ := QueryData(db, oldUsername)
	_, err := db.Exec("UPDATE users SET username = ? WHERE id = ?", newUsername, id)
	return err
}

func UpdatePassword(db *sql.DB, username, newPassword string) error {
	_, err := db.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func DeleteUser(db *sql.DB, username, password string) error {
	_, usernameFromDb, passwordFromDb, err := QueryData(db, username)
	if err != nil {
		fmt.Printf("Failed to fetch user data from db due to error message %q", err)
		// fmt.Printf("Reinitiating the example test flow due to failure while fetching the user data .")
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
