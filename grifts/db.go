package grifts

import (
	"fmt"
	"log"

	"github.com/conradwt/zero-to-graphql-using-buffalo/models"
	"github.com/gobuffalo/pop"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// reset DB table(s)
		DB, error := pop.Connect("development")
		if error != nil {
			log.Panic(error)
		}

		if error = FriendshipDestroyAll(); error != nil {
			log.Panic(error)
		}

		if error = PersonDestroyAll(); error != nil {
			log.Panic(error)
		}

		// create people
		me := models.Person{
			FirstName: "Conrad",
			LastName:  "Taylor",
			Email:     "conradwt@gmail.com",
			Username:  "conradwt"}
		_, error = DB.ValidateAndSave(&me)
		if error != nil {
			log.Panic(error)
		}

		dhh := models.Person{
			FirstName: "David",
			LastName:  "Heinemeier Hansson",
			Email:     "dhh@37signals.com",
			Username:  "dhh"}
		_, error = DB.ValidateAndSave(&dhh)
		if error != nil {
			log.Panic(error)
		}

		ezra := models.Person{
			FirstName: "Ezra",
			LastName:  "Zygmuntowicz",
			Email:     "ezra@merbivore.com",
			Username:  "ezra"}
		_, error = DB.ValidateAndSave(&ezra)
		if error != nil {
			log.Panic(error)
		}

		matz := models.Person{
			FirstName: "Yukihiro",
			LastName:  "Matsumoto",
			Email:     "matz@heroku.com",
			Username:  "matz"}
		_, error = DB.ValidateAndSave(&matz)
		if error != nil {
			log.Panic(error)
		}

		// create friendships
		meFriends := models.Friendships{
			{PersonID: me.ID, FriendID: matz.ID},
		}

		error = DB.Create(&meFriends)
		if error != nil {
			log.Panic(error)
		}

		dhhFriends := models.Friendships{
			{PersonID: dhh.ID, FriendID: ezra.ID},
			{PersonID: dhh.ID, FriendID: matz.ID},
		}

		error = DB.Create(&dhhFriends)
		if error != nil {
			log.Panic(error)
		}

		ezraFriends := models.Friendships{
			{PersonID: ezra.ID, FriendID: dhh.ID},
			{PersonID: ezra.ID, FriendID: matz.ID},
		}

		error = DB.Create(&ezraFriends)
		if error != nil {
			log.Panic(error)
		}

		matzFriends := models.Friendships{
			{PersonID: matz.ID, FriendID: me.ID},
			{PersonID: matz.ID, FriendID: ezra.ID},
			{PersonID: matz.ID, FriendID: dhh.ID},
		}

		error = DB.Create(&matzFriends)
		if error != nil {
			log.Panic(error)
		}

		// people count
		people := models.People{}
		error = DB.All(&people)
		fmt.Println("People count:", len(people))

		// friendships count
		friendships := models.Friendships{}
		error = DB.All(&friendships)
		fmt.Println("Friendships count:", len(friendships))

		return nil
	})

})

// PersonDestroyAll removes all the person records from the DB.
func PersonDestroyAll() error {
	DB, error := pop.Connect("development")
	if error != nil {
		log.Panic(error)
	}

	people := models.People{}
	error = DB.All(&people)
	if error != nil {
		fmt.Print("PersonDestroyAll ERROR!\n")
		fmt.Printf("%v\n", error)
		return error
	}

	for i := 0; i < len(people); i++ {
		person := people[i]
		if error = DB.Destroy(&person); error != nil {
			log.Panic(error)
		}
	}

	return nil
}

// FriendshipDestroyAll removes all the friendship records from the DB.
func FriendshipDestroyAll() error {
	DB, error := pop.Connect("development")
	if error != nil {
		log.Panic(error)
	}

	friendships := models.Friendships{}
	error = DB.All(&friendships)
	if error != nil {
		fmt.Print("FriendshipDestroyAll ERROR!\n")
		fmt.Printf("%v\n", error)
		return error
	}

	for i := 0; i < len(friendships); i++ {
		friendship := friendships[i]
		if error = DB.Destroy(&friendship); error != nil {
			log.Panic(error)
		}
	}

	return nil
}
