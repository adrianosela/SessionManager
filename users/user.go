package users

//User represents a user of our service
type User struct {
	Uname       string `json:"username"`
	LastUpdated int64  `json:"lastupdated"`
	Status      string `json:"status"`
}

const (
	//StatusOffline represents the offline status
	StatusOffline = "offline"
	//StatusOnline represents the online status
	StatusOnline = "online"
)
