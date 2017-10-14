package session

import (
	"time"

	"github.com/satori/go.uuid"
)

/* Improvements TODO:
 * * include the user id in the token
 * * have tokens be JWTs emitted by an auth endpoint (not uuids)
 */

// Session will represent an active user session
type Session struct {
	User          string `json:"username"`
	Token         string `json:"session_token"`
	LastRefreshed int64  `json:"expires"`
}

// New initializes a new session
func New(username string) *Session {

	token := uuid.NewV4().String()

	sess := &Session{
		User:          username,
		Token:         token,
		LastRefreshed: time.Now().UnixNano(),
	}

	return sess
}
