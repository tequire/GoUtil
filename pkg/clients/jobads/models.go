package jobads

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

// Organization defines an organization
type Organization struct {
	ID                    int       `json:"id"`
	OneComID              int       `json:"oneComId"`
	ParentCompanyID       int       `json:"parentCompanyId"`
	Title                 string    `json:"title"`
	Description           string    `json:"description"`
	WebAddress            string    `json:"webAddress"`
	Logo                  string    `json:"logo"`
	OwnerOrganizationName string    `json:"ownerOrganizationName"`
	Verified              bool      `json:"verified"`
	Category              *Category `json:"category"`
}

// Category defines an organization category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OrganizationSetting defines the setting for an organization
type OrganizationSetting struct {
	OrganizationID int    `json:"organizationSettingsId"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
	TertiaryColor  string `json:"tertiaryColor"`
	BackgroundURL  string `json:"backgroundUrl"`
	LogoURL        string `json:"logoUrl"`
	Domains        string `json:"domains"`
	SubDomain      string `json:"subDomain"`
}

// Job is a jobAd
type Job struct {
	JobAdID             int       `json:"jobAdId"`
	ID                  int       `json:"id"`
	OwnerOrganizationID uuid.UUID `json:"ownerOrganizationId"`
	CompanyID           int       `json:"companyId"`
	CompanyName         string    `json:"companyName"`
	JobTitle            string    `json:"jobTitle"`
	NumberOfVacancies   int       `json:"numberOfVacancies"`
	ApplicationDeadline *Time     `json:"applicationDeadline"`
	JobDescription      string    `json:"jobDescription"`
	TaskDescription     string    `json:"taskDescription"`
	OfferDescription    string    `json:"offerDescription"`
	ApplicationLink     string    `json:"applicationLink"`
	EmploymentLevel     int       `json:"employmentLevel"`
	EmploymentStartDate *Time     `json:"employmentStartDate"`
	EmploymentEndDate   *Time     `json:"employmentEndDate"`
	PublishDate         *Time     `json:"publishDate"`
	PublishEndDate      *Time     `json:"publishEndDate"`
	Qualifications      string    `json:"qualifications"`
	PersonalQualities   string    `json:"personalQuailites"`
	StreetAddress       string    `json:"streetAdress"`
	ZipCode             string    `json:"zipCode"`
	PostalAddress       string    `json:"postalAddress"`
	Latitude            float32   `json:"latitude"`
	Longitude           float32   `json:"longitude"`
	CountryList         []Country `json:"country"`
	CountryID           int       `json:"countryId"`
	ReferenceNumber     string    `json:"referenceNumber"`
	Visibility          string    `json:"visibility"`
	PaymentRole         string    `json:"paymentRole"`
	JobAdState          int       `json:"jobAdState"`
	EducationFields     []string  `json:"educationFields"`
	Departments         []string  `json:"departments"`
	EducationLevels     []string  `json:"educationLevels"`
	EmploymentTypes     []string  `json:"employmentTypes"`
	Languages           []string  `json:"languages"`
	Locations           []string  `json:"locations"`
	Sectors             []string  `json:"sectors"`
	Schools             []School  `json:"schools"`
}

// Country defines a country
type Country struct {
	ID   string
	Text string
}

// School defines a school by id and name
type School struct {
	ID   string
	Text string
}

// Time defines a ISO-datetime and supports empty strings
type Time struct {
	time.Time
}

// ToString converts to time...
func (t *Time) String() string {
	ts, err := ptypes.TimestampProto(t.Time)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ptypes.TimestampString(ts)
}

// UnmarshalJSON returns time.Now() no matter what!
func (t *Time) UnmarshalJSON(b []byte) error {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	timestamp, err := time.Parse("\"2006-01-02T15:04:05\"", string(b))
	if err == nil {
		*t = Time{timestamp}
	}
	return nil
}

// Article defines an article
type Article struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Ingress         string    `json:"ingress"`
	FeaturedImage   string    `json:"featuredImage"`
	Slug            string    `json:"slug"`
	Tags            []Tag     `json:"tags"`
	OrganizationIds []int     `json:"organizationIds"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Video defines a video
type Video struct {
	ID              string    `json:"id"`
	VideoLink       string    `json:"videoLink"`
	Thumbnail       string    `json:"thumbnail,omitempty"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	OrganizationIds []int     `json:"organizationIds"`
	Tags            []Tag     `json:"tags"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Tag defines a tag
type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
