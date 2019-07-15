package main

import (
	"fmt"

	"github.com/tequire/GoUtil/pkg/clients/jobads"
)

func main() {
	client := jobads.New("")
	client.SetProd(true)

	orgs, err := client.GetCompanies()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(orgs)

	for _, school := range orgs {
		fmt.Println(school)
	}
}
