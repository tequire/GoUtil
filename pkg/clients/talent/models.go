package talent

import "time"

// Job defines a job from CTS
type Job struct {
	Addresses          []string                   `json:"addresses,omitempty"`
	ApplicationInfo    *ApplicationInfo           `json:"applicationInfo,omitempty"`
	CompanyDisplayName string                     `json:"companyDisplayName,omitempty"`
	CompanyName        string                     `json:"companyName,omitempty"`
	CustomAttributes   map[string]CustomAttribute `json:"customAttributes,omitempty"`
	DegreeTypes        []string                   `json:"degreeTypes,omitempty"`
	Department         string                     `json:"department,omitempty"`
	DerivedInfo        *JobDerivedInfo            `json:"derivedInfo,omitempty"`
	Description        string                     `json:"description,omitempty"`
	EmploymentTypes    []string                   `json:"employmentTypes,omitempty"`
	JobBenefits        []string                   `json:"jobBenefits,omitempty"`
	JobEndTime         string                     `json:"jobEndTime,omitempty"`
	JobLevel           string                     `json:"jobLevel,omitempty"`
	JobStartTime       string                     `json:"jobStartTime,omitempty"`
	LanguageCode       string                     `json:"languageCode,omitempty"`
	Name               string                     `json:"name,omitempty"`
	PostingCreateTime  string                     `json:"postingCreateTime,omitempty"`
	PostingExpireTime  string                     `json:"postingExpireTime,omitempty"`
	PostingPublishTime string                     `json:"postingPublishTime,omitempty"`
	PostingRegion      string                     `json:"postingRegion,omitempty"`
	PostingUpdateTime  string                     `json:"postingUpdateTime,omitempty"`
	PromotionValue     int64                      `json:"promotionValue,omitempty"`
	Qualifications     string                     `json:"qualifications,omitempty"`
	RequisitionID      string                     `json:"requisitionId,omitempty"`
	Responsibilities   string                     `json:"responsibilities,omitempty"`
	Title              string                     `json:"title,omitempty"`
	Visibility         string                     `json:"visibility,omitempty"`
	ForceSendFields    []string                   `json:"-"`
	NullFields         []string                   `json:"-"`
}

// ApplicationInfo defines the info of an application
type ApplicationInfo struct {
	Emails          []string `json:"emails,omitempty"`
	Instruction     string   `json:"instruction,omitempty"`
	Uris            []string `json:"uris,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

// JobDerivedInfo defines the dervied info of an job
type JobDerivedInfo struct {
	JobCategories   []string    `json:"jobCategories,omitempty"`
	Locations       []*Location `json:"locations,omitempty"`
	ForceSendFields []string    `json:"-"`
	NullFields      []string    `json:"-"`
}

// Location defines a location, with latlng position and addresses
type Location struct {
	LatLng          *LatLng        `json:"latLng,omitempty"`
	LocationType    string         `json:"locationType,omitempty"`
	PostalAddress   *PostalAddress `json:"postalAddress,omitempty"`
	RadiusInMiles   float64        `json:"radiusInMiles,omitempty"`
	ForceSendFields []string       `json:"-"`
	NullFields      []string       `json:"-"`
}

// LatLng defines a latlng position
type LatLng struct {
	Latitude        float64  `json:"latitude,omitempty"`
	Longitude       float64  `json:"longitude,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

// PostalAddress defines an address
type PostalAddress struct {
	AddressLines       []string `json:"addressLines,omitempty"`
	AdministrativeArea string   `json:"administrativeArea,omitempty"`
	LanguageCode       string   `json:"languageCode,omitempty"`
	Locality           string   `json:"locality,omitempty"`
	Organization       string   `json:"organization,omitempty"`
	PostalCode         string   `json:"postalCode,omitempty"`
	Recipients         []string `json:"recipients,omitempty"`
	RegionCode         string   `json:"regionCode,omitempty"`
	Revision           int64    `json:"revision,omitempty"`
	SortingCode        string   `json:"sortingCode,omitempty"`
	Sublocality        string   `json:"sublocality,omitempty"`
	ForceSendFields    []string `json:"-"`
	NullFields         []string `json:"-"`
}

// CustomAttribute defines a custom attribute
type CustomAttribute struct {
	Filterable      bool     `json:"filterable,omitempty"`
	LongValues      []string `json:"longValues,omitempty"`
	StringValues    []string `json:"stringValues,omitempty"`
	ForceSendFields []string `json:"-"`
	NullFields      []string `json:"-"`
}

// JobView defines a jobView. It only contains the essential parts of a job
type JobView struct {
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	JobID   int    `json:"jobId"`   // From JobAds
	JobName string `json:"jobName"` // From CTS

	Title             string `json:"title"`
	CompanyID         int    `json:"companyId"`
	CompanyLogo       string `json:"companyLogo"`
	CompanyName       string `json:"companyName"`
	PostingExpireTime string `json:"postingExpireTime"`
	Address           string `json:"address"`
}
