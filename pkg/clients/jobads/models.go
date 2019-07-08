package jobads

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
