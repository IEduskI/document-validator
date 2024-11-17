package rest

type Request struct {
	Type         string `json:"type"`
	Value        string `json:"value"`
	IssueCountry string `json:"issueCountry"`
	IssueDate    string `json:"issueDate"`
	ExpiryDate   string `json:"expiryDate"`
}

type Response struct {
	Message string `json:"message"`
}

type Error struct {
	Description string `json:"description"`
}
