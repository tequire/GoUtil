package auth

import (
	"github.com/coreos/go-oidc"
)

type school struct {
	SchoolIds []string `json:"org.student"`
}

type claim struct {
	Value []string
}

// GetSchoolClaims gets the claim of "org.student" from a given token
func GetSchoolClaims(token *oidc.IDToken) []string {
	// Read scopes
	readClaims := school{}
	token.Claims(&readClaims)

	// The claim-value can either be an array or a single string
	if readClaims.SchoolIds != nil && len(readClaims.SchoolIds) > 0 {
		return readClaims.SchoolIds
	}

	// Read claim as string
	var claims struct {
		SchoolID string `json:"org.student"`
	}
	token.Claims(&claims)
	if claims.SchoolID == "" {
		return []string{}
	}
	return []string{claims.SchoolID}
}

// GetClaim returns an array of a wanted claim by name
/* func GetClaim(token *oidc.IDToken, name string) ([]string, error) {
	// Get claim as string array
	arrayClaim := claim{}

	//  Add struct-tag with given name
	f := reflect.TypeOf(arrayClaim)
	ff, found := f.FieldByName("Value")
	if !found {
		return nil, errors.New("unable to get field name")
	}
	ff.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, name))

	// Get claims as string array
	err := token.Claims(&s2)
	if err != nil {
		return nil, err
	}
	if s2.Value != nil && len(s2.Value) > 0 {
		return s2.Value, nil
	}

	// Get claim as string
	var stringClaim struct {
		Value string
	}

	// Add struct-tag with given name
	f = reflect.TypeOf(stringClaim)
	ff, found = f.FieldByName("Value")
	if !found {
		return nil, errors.New("unable to get field name")
	}
	ff.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, name))

	f2 = reflect.TypeOf(stringClaim)
	ff2, _ = f2.FieldByName("Value")
	fmt.Printf("Tag: %+v\n", ff2.Tag)

	// Get claims as string
	err = token.Claims(&stringClaim)
	if err != nil {
		return nil, err
	}
	if stringClaim.Value == "" {
		return nil, errors.New("no claim by that name")
	}
	return []string{stringClaim.Value}, nil
}
*/
