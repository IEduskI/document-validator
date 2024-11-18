package validate

type Request struct {
	Document []Document
}

type Document struct {
	Type         string
	Value        string
	IssueCountry string
	IssueDate    string
	ExpiryDate   string
}
