package github

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type RateLimitInfo struct {
	Limit     int
	Remaining int
	Reset     int
}

type PathwayResponse struct {
	Status int `json:"status"`
	Data   struct {
		Pathway []GithubUser `json:"pathway"`
	} `json:"data"`
}

type GithubUser struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatarUrl"`
	Url       string `json:"url"`
	Bio       string `json:"bio"`
	Followers struct {
		TotalCount int `json:"totalCount"`
	} `json:"followers"`
	Following struct {
		TotalCount int `json:"totalCount"`
	} `json:"following"`
}
type SearchUsersGithubResponse struct {
	Data   SearchUsers `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type SearchUsers struct {
	Search struct {
		Nodes []GithubUser `json:"nodes"`
	} `json:"search"`
}
type SearchUsersResponse struct {
	Status int `json:"status"`
	Data   struct {
		Results []GithubUser `json:"results"`
	} `json:"data"`
}
