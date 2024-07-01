package github

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
)

var (
	githubUrl = "https://api.github.com/graphql"
	token     = os.Getenv("GITHUB_TOKEN")
)

func getSearchQuery(
	query string,
	c *fasthttp.Client,
) (*SearchUsers, *RateLimitInfo, error) {
	qlQuery := fmt.Sprintf(`
		{
			search(query: "%s", type: USER, first: 10) {
				nodes {
					... on User {
						login
						avatarUrl
						url
						bio
					}
				}
			}
		}
	`, query)

	body := map[string]interface{}{
		"query": qlQuery,
	}
	bodyJson, _ := json.Marshal(body)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(githubUrl)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.SetBody(bodyJson)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.Do(req, resp); err != nil {
		return nil, nil, err
	}

	respBody := &SearchUsersGithubResponse{}
	if err := json.Unmarshal(resp.Body(), respBody); err != nil {
		return nil, nil, err
	}

	if len(respBody.Errors) > 0 {
		return nil, nil, fmt.Errorf("GraphQL error: %s", respBody.Errors[0].Message)
	}
	rateLimitInfo := parseRateLimitInfo(resp)

	return &respBody.Data, rateLimitInfo, nil

}

func parseRateLimitInfo(resp *fasthttp.Response) *RateLimitInfo {
	limit, _ := strconv.Atoi(string(resp.Header.Peek("X-RateLimit-Limit")))
	remaining, _ := strconv.Atoi(string(resp.Header.Peek("X-RateLimit-Remaining")))
	reset, _ := strconv.Atoi(string(resp.Header.Peek("X-RateLimit-Reset")))

	return &RateLimitInfo{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}
}
