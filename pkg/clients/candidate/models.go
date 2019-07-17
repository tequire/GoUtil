package candidate

import (
	"time"

	"github.com/google/uuid"
)

// Candidate defines a candidate
type Candidate struct {
	ID *uuid.UUID `json:"id"`

	UserID *uuid.UUID `json:"user_id"`

	CV        *string              `json:"cv"`                // CV for candidate.
	Degrees   []*Degree            `json:"degrees,omitempty"` // Degrees the candidate has earned or is earning.
	Languages []*LanguageCandidate `json:"languages,omitempty"`

	LinkedIn *string `json:"linkedin"` // Link to the LinkedIn profile of the candidate.
	Github   *string `json:"github"`   // Link to the Github profile of the candidate.

	Skills    []string `json:"skills,omitempty"`    // Skills the candidate has.
	Interests []string `json:"interests,omitempty"` // Interests the candidate has.

	NotificationFrequency *string `json:"notification_frequency"`
}

// Degree defines a degree
type Degree struct {
	ID *uuid.UUID `json:"id"`

	SchoolID int `json:"school_id" binding:"required"`

	Name string `json:"name"`

	School *School `json:"school"` // Degree has one School

	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`

	DegreeLevelID *uuid.UUID   `json:"level_id"`
	DegreeLevel   *DegreeLevel `json:"level"`

	DegreeFieldID *uuid.UUID   `json:"field_id"`
	DegreeField   *DegreeField `json:"field"`

	Description string   `json:"description"`
	Documents   []string `json:"documents"`
}

// DegreeLevel describes the level of an degree (bachelor, master, doctoral)
type DegreeLevel struct {
	ID   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

// DegreeField describes a major, for example (Computer & it...etc)
type DegreeField struct {
	ID   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

// School defines a school
type School struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`

	Name string `json:"name" sql:",unique"`
}

// LanguageCandidate is a type that defines a language skill of a given candidate.
type LanguageCandidate struct {
	ID *uuid.UUID `json:"id"`

	// Language fields
	LanguageID    *uuid.UUID           `json:"languageId"`
	Language      *Language            `json:"language"`
	ProficiencyID *uuid.UUID           `json:"proficiencyId"`
	Proficiency   *LanguageProficiency `json:"proficiency"`
}

// Language is a type storing a name of a language, like "Norwegian, Swedish, Korean...etc"
type Language struct {
	ID   *uuid.UUID `json:"id"`
	Code string     `json:"code"`
	Name string     `json:"name"`
}

// LanguageProficiency is a type that defines a proficiency level of a language
type LanguageProficiency struct {
	ID    *uuid.UUID `json:"id"`
	Level string     `json:"level"`
}
