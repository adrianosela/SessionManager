package datastore

import (
	"testing"
	"time"

	"github.com/adrianosela/SessionManager/session"
	"github.com/adrianosela/SessionManager/users"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMockStore(t *testing.T) {

	var db *MockDB

	Convey("Init Mock Store Test", t, func() {
		db = NewMockDB()
		So(db, ShouldNotEqual, nil)
	})

	Convey("Mock Store User Operations Tests", t, func() {

		Convey("CreateUserTest_Success", func() {
			usersToMake := []users.User{
				{
					Uname:       "User_A",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_B",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_C",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
			}

			for _, usr := range usersToMake {
				err := db.createUser(usr)
				So(err, ShouldBeNil)
				So(db.Users, ShouldContain, usr)
			}

			So(len(db.Users), ShouldEqual, len(usersToMake))

			//cleanup
			db.Users = []users.User{}
		})

		Convey("DeleteUserTest_Success", func() {

			db.Users = []users.User{
				{
					Uname:       "User_A",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_B",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_C",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
			}

			usersToDelete := []users.User{
				{
					Uname:       "User_A",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_B",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_C",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
			}

			So(len(db.Users), ShouldEqual, 3)

			//ensuring the users are being deleted
			for _, usr := range usersToDelete {
				err := db.deleteUser(usr)
				So(err, ShouldBeNil)
				So(db.Users, ShouldNotContain, usr)
			}

			So(len(db.Users), ShouldEqual, 0)

			//cleanup
			db.Users = []users.User{}
		})

		Convey("DeleteUserTest_Fail_DoesNotExist", func() {

			db.Users = []users.User{
				{
					Uname:       "User_A",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
				{
					Uname:       "User_B",
					LastUpdated: time.Now().UnixNano(),
					Status:      users.StatusOnline,
				},
			}

			userToDelete := users.User{
				Uname:       "User_C",
				LastUpdated: time.Now().UnixNano(),
				Status:      users.StatusOnline,
			}

			err := db.deleteUser(userToDelete)
			So(err, ShouldNotBeNil)

			//cleanup
			db.Users = []users.User{}
		})

	})

	Convey("Mock Store Session Operations Tests", t, func() {

		Convey("AddSessionTest_Success", func() {
			sessToMake := []session.Session{
				{
					User:          "User_A",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
				{
					User:          "User_B",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
			}

			for _, sess := range sessToMake {
				err := db.addSession(sess)
				So(err, ShouldBeNil)
				So(db.Sessions, ShouldContain, sess)
			}

			So(len(db.Sessions), ShouldEqual, len(sessToMake))

			//cleanup
			db.Sessions = []session.Session{}
		})

		Convey("DeleteSessionTest_Success", func() {
			db.Sessions = []session.Session{
				{
					User:          "User_A",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
				{
					User:          "User_B",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
			}

			sessToDelete := []session.Session{
				{
					User:          "User_A",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
				{
					User:          "User_B",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
			}

			So(len(db.Sessions), ShouldEqual, 2)

			//ensuring the users are being deleted
			for _, sess := range sessToDelete {
				err := db.deleteSession(sess)
				So(err, ShouldBeNil)
				So(db.Sessions, ShouldNotContain, sess)
			}

			So(len(db.Sessions), ShouldEqual, 0)

			//cleanup
			db.Sessions = []session.Session{}
		})

		Convey("DeleteSessionTest_Fail_DoesNotExist", func() {

			db.Sessions = []session.Session{
				{
					User:          "User_A",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
				{
					User:          "User_B",
					Token:         "test-token",
					LastRefreshed: time.Now().UnixNano(),
				},
			}

			sessToDelete := session.Session{
				User:          "User_C",
				Token:         "test-token",
				LastRefreshed: time.Now().UnixNano(),
			}

			err := db.deleteSession(sessToDelete)
			So(err, ShouldNotBeNil)

			//cleanup
			db.Sessions = []session.Session{}
		})

	})

}
