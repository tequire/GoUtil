package gin

import (
	"github.com/google/uuid"
)

// TestUser creates a user for testing
func TestUser() User {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return User{ID: &id}
}

func KnownTestUser() User {
	id := uuid.MustParse("8fa840e6-2ec9-414e-a102-64bb4219e248")
	return User{ID: &id}
}

func KnownTestUser2() User {
	id := uuid.MustParse("b368f083-6674-48f6-8576-1fcb3600d82e")
	return User{ID: &id}
}
