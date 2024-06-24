package algorithm

import (
	"time"
)

type GraphQlResponse struct {
	Data   Data `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type Data struct {
	User UserNode `json:"user"`
}

type UserNode struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatarUrl"`
	Url       string `json:"url"`
	Followers struct {
		Nodes      []UserNode `json:"nodes"`
		TotalCount int        `json:"totalCount"`
	} `json:"followers"`
	Following struct {
		Nodes      []UserNode `json:"nodes"`
		TotalCount int        `json:"totalCount"`
	} `json:"following"`
	Prev *UserNode
}

type RateLimitResponse struct {
	Data struct {
		RateLimit struct {
			Remaining int       `json:"remaining"`
			ResetAt   time.Time `json:"resetAt"`
		} `json:"rateLimit"`
	} `json:"data"`
}
