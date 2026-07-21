package main

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
)

func main() {

}

func testDatabase() {
	psql := postgres.New()
	u, err := psql.Register(entity.User{ID: 0, Name: "ali", PhoneNumber: "09123223"})
	fmt.Println(u, err)

	isUnique, err := psql.IsPhoneNumberUnique(u.PhoneNumber + "23")
	fmt.Println(isUnique, err)
}
