package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

// https://blog.gobuffalo.io/pop-v4-5-9-released-f8e3461dbe81
type Friendship struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	PersonID  int64     `json:"person_id" db:"person_id"`
	FriendID  int64     `json:"friend_id" db:"friend_id" fk_id:"person_id"`
	Person    Person    `belongs_to:"person"`
	Friend    Person    `belongs_to:"person" fk_id:"FriendID"`
}

// String is not required by pop and may be deleted
func (f Friendship) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Friendships is not required by pop and may be deleted
type Friendships []Friendship

// String is not required by pop and may be deleted
func (f Friendships) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Friendship) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Friendship) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Friendship) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// func (friendships []*Friendship) DestroyAll() error {
// 	DB, error := pop.Connect("development")
// 	if error != nil {
// 		log.Panic(error)
// 	}

// 	friendships := models.Friendships{}
// 	error = DB.All(&friendships)
// 	if error != nil {
// 		fmt.Print("FriendshipDestroyAll ERROR!\n")
// 		fmt.Printf("%v\n", error)
// 		return error
// 	}

// 	for i := 0; i < len(friendships); i++ {
// 		friendship := friendships[i]
// 		if error = DB.Destroy(&friendship); error != nil {
// 			log.Panic(error)
// 		}
// 	}

// 	return nil
// }
