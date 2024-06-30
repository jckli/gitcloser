package pathway

type PathwayResponse struct {
	Status int `json:"status"`
	Data   struct {
		Pathway []PathwayUser `json:"pathway"`
	} `json:"data"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type PathwayUser struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatarUrl"`
	Url       string `json:"url"`
	Followers struct {
		TotalCount int `json:"totalCount"`
	} `json:"followers"`
	Following struct {
		TotalCount int `json:"totalCount"`
	} `json:"following"`
}
