package candidate

import "encoding/json"

// Client to access the Candidate API
type Client struct {
	token  string
	isProd bool
}

// New creates a new organization client
func New(token string) *Client {
	return &Client{
		token:  token,
		isProd: false,
	}
}

// SetProd sets if the client should request against the prod or dev environment
func (c *Client) SetProd(isProd bool) {
	c.isProd = isProd
}

// CreateCandidate creates a candidate - requires an Admin-token
func (c *Client) CreateCandidate(userID string) (*Candidate, error) {
	var reqBody struct {
		UserID string `json:"user_id"`
	}
	reqBody.UserID = userID
	result, err := handleRequest(c, "POST", "api/v1/candidate", reqBody)
	if err != nil {
		return nil, err
	}

	// Read response body
	candidate := &Candidate{}
	err = json.Unmarshal(result.Body, candidate)
	if err != nil {
		return nil, err
	}

	return candidate, nil
}

// CreateDegree creates a degree for a given candidate
// If 'candidateID' is nil, the degree will be assign the candidate with beloning userID in token
func (c *Client) CreateDegree(candidateID *string, degree Degree) (*Degree, error) {
	var reqBody struct {
		Degree
		CandidateID *string `json:"candidate_id,omitempty"`
	}
	reqBody.Degree = degree
	reqBody.CandidateID = candidateID
	result, err := handleRequest(c, "POST", "api/v1/candidate/degree", reqBody)
	if err != nil {
		return nil, err
	}

	// Read response body
	resDegree := &Degree{}
	err = json.Unmarshal(result.Body, resDegree)
	if err != nil {
		return nil, err
	}
	return resDegree, nil
}

// CreateLanguage creates a language for a given candidate.
// If 'candidateID' is nil, the language will be assign the candidate with beloning userID in token
func (c *Client) CreateLanguage(candidateID *string, lang LanguageCandidate) (*LanguageCandidate, error) {
	var reqBody struct {
		LanguageCandidate
		CandidateID *string `json:"candidateId,omitempty"`
	}
	reqBody.CandidateID = candidateID
	reqBody.LanguageCandidate = lang
	result, err := handleRequest(c, "POST", "api/v1/candidate/language", reqBody)
	if err != nil {
		return nil, err
	}

	// Read response body
	resLang := &LanguageCandidate{}
	err = json.Unmarshal(result.Body, resLang)
	if err != nil {
		return nil, err
	}
	return resLang, nil
}
