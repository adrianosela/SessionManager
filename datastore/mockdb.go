package datastore

import (
	"fmt"

	"github.com/adrianosela/SessionManager/session"
	"github.com/adrianosela/SessionManager/users"
)

const (
	//SessionDoesNotExistError represents an error trying to access a nonexistent session
	SessionDoesNotExistError = "[ERROR] : Session does not exist"
	//UserDoesNotExistError represents an error trying to access a nonexistent user
	UserDoesNotExistError = "[ERROR] : User does not exist"
)

//MockDB is a DB used for simple local runs, tests and as a proof of concept
type MockDB struct {
	Users    []users.User      `json:"users"`
	Sessions []session.Session `json:"sessions"`
}

//NewMockDB initializes a new mock datastore in memory and returns its address
func NewMockDB() *MockDB {
	return &MockDB{
		Users:    []users.User{},
		Sessions: []session.Session{},
	}
}

func (db *MockDB) addSession(sess session.Session) error {
	db.Sessions = append(db.Sessions, sess)
	return nil
}

func (db *MockDB) createUser(usr users.User) error {
	db.Users = append(db.Users, usr)
	return nil
}

func (db *MockDB) deleteSession(deadSession session.Session) error {
	// removing by checking the username as only active sessions will
	// ever exist and users may have only one active session
	for idx, sess := range db.Sessions {
		if deadSession.User == sess.User {
			if idx+1 < len(db.Sessions) {
				db.Sessions = append(db.Sessions[:idx], db.Sessions[idx+1:]...)
			} else {
				db.Sessions = db.Sessions[:idx]
			}
			return nil
		}
	}
	return fmt.Errorf("%s. Could not delete session with token: %s", SessionDoesNotExistError, deadSession.Token)
}

func (db *MockDB) deleteUser(usr users.User) error {
	for idx, user := range db.Users {
		if user.Uname == usr.Uname {
			if idx+1 < len(db.Users) {
				db.Users = append(db.Users[:idx], db.Users[idx+1:]...)
			} else {
				db.Users = db.Users[:idx]
			}
			return nil
		}
	}
	return fmt.Errorf("%s. Could not delete user: %s", UserDoesNotExistError, usr.Uname)
}
