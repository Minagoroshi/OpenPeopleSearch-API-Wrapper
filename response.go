package openpeoplesearch

type SearchResult struct {
	SearchId       string        `json:"searchId"`
	SearchType     string        `json:"searchType"`
	SearchCriteria []SearchParam `json:"searchCriteria"`
	StatusCode     string        `json:"statusCode"`
	StatusMessage  string        `json:"statusMessage"`
	ErrorMessages  []string      `json:"errorMessages"`
	Records        int           `json:"records"`
	SearchCost     float64       `json:"searchCost"`
	SearchDate     string        `json:"searchDate"`
	Results        []Result      `json:"results"`
}

type SearchParam struct {
	SearchName  string `json:"searchName"`
	SearchValue string `json:"searchValue"`
}

type Result struct {
	ReportedDate        string `json:"reportedDate"`
	FirstName           string `json:"firstName"`
	MiddleName          string `json:"middleName"`
	LastName            string `json:"lastName"`
	Address             string `json:"address"`
	City                string `json:"city"`
	State               string `json:"state"`
	Zip                 string `json:"zip"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	LineType            string `json:"lineType"`
	DOB                 string `json:"dob"`
	Occupation          string `json:"occupation"`
	Employer            string `json:"employer"`
	AssociatedDomain    string `json:"associatedDomain"`
	DataTypeName        string `json:"dataTypeName"`
	DataCategoryName    string `json:"dataCategoryName"`
	WebVerificationLink string `json:"webVerificationLink"`
	AgencyName          string `json:"agencyName"`
	Vol                 string `json:"vol"`
	RecordId            string `json:"recordId"`
}
