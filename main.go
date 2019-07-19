package main

import (
	"fmt"

	"github.com/tequire/GoUtil/pkg/clients/candidate"
)

var token = ""

func main() {
	client := candidate.New(token)
	client.SetProd(false)

	user, err := client.GetCandidateByUserID("da282e62-f2ed-4b47-8fc2-aa70456d8684")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user)

	fmt.Println(user.Languages)
	for _, lang := range user.Languages {
		fmt.Println(lang)
	}
	for _, degree := range user.Degrees {
		fmt.Println(degree)
	}

	/* for _, school := range orgs {
		fmt.Println(school)
	} */
}
