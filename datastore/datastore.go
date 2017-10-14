package datastore

import (
	"github.com/adrianosela/SessionManager/session"
	"github.com/adrianosela/SessionManager/users"
)

//Datastore will allow us to implement a database in multiple ways
type Datastore interface {
	//addSession will add a session to the database
	addSession(session.Session) error

	//removeSession will remove a session from the database
	removeSession(session.Session) error

	//createUser will add a new user to the database
	createUser(users.User) error

	//deleteUser will remove a user from the database
	deleteUser(users.User) error
}
