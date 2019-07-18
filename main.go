package main

import (
	"fmt"

	"github.com/tequire/GoUtil/pkg/clients/identity"
)

func main() {
	client := identity.New("")
	client.SetProd(false)

	user, err := client.CustomUserToken("ad04a462-3c65-4f99-a877-22d64d8fe0a7", 24)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user)

	/* for _, school := range orgs {
		fmt.Println(school)
	} */
}
