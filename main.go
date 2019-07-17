package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/tequire/GoUtil/pkg/clients/candidate"
)

func main() {
	client := candidate.New("")
	client.SetProd(false)

	candidateID := "46783d67-80ef-4d1f-bf23-44a76374560b"
	langID := uuid.MustParse("d6067b41-5433-4964-97c9-885ba4b52849")
	profID := uuid.MustParse("82edcac6-693b-42ea-bf24-0a52774f395c")

	candidate, err := client.CreateLanguage(&candidateID, candidate.LanguageCandidate{
		LanguageID:    &langID,
		ProficiencyID: &profID,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(candidate)

	/* for _, school := range orgs {
		fmt.Println(school)
	} */
}
