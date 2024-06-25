package algorithm

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
	"time"
)

var (
	githubUrl = "https://api.github.com/graphql"
	token     = os.Getenv("GITHUB_TOKEN")
)

const followingQueryTemplate = `
query ($username: String!, $after: String) {
	user(login: $username) {
		following(first: 100, after: $after) {
			nodes {
				login
				avatarUrl
				url
				followers {
					totalCount
				}
				following {
					totalCount
				}
			}
			pageInfo {
				endCursor
				hasNextPage
			}
		}
	}
}`

const followersQueryTemplate = `
query ($username: String!, $after: String) {
	user(login: $username) {
		followers(first: 100, after: $after) {
			nodes {
				login
				avatarUrl
				url
				followers {
					totalCount
				}
				following {
					totalCount
				}
			}
			pageInfo {
				endCursor
				hasNextPage
			}
		}
	}
}`

const userQueryTemplate = `
query ($username: String!, $after: String) {
	user(login: $username) {
		login
		avatarUrl
		followers(first: 100) {
			nodes {
				login
				avatarUrl
				url
			}
			totalCount
			pageInfo {
				endCursor
				hasNextPage
			}
			
		}
		following(first: 100, after: $after) {
			nodes {
				login
				avatarUrl
				url
			}
			totalCount
			pageInfo {
				endCursor
				hasNextPage
			}
		}
		url
	}
}`

func getUser(username, queryType string, c *fasthttp.Client) (*[]UserNode, *RateLimitInfo, error) {
	var query string

	if queryType == "following" {
		query = followingQueryTemplate
	} else if queryType == "followers" {
		query = followersQueryTemplate
	} else {
		return nil, nil, fmt.Errorf("invalid query type: %s", queryType)
	}

	var endCursor *string
	var nodes []UserNode
	var rateLimitInfo *RateLimitInfo

	for {
		body := map[string]interface{}{
			"query": query,
			"variables": map[string]interface{}{
				"username": username,
				"after":    endCursor,
			},
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

		respBody := &GraphQlResponse{}
		if err := json.Unmarshal(resp.Body(), respBody); err != nil {
			return nil, nil, err
		}

		if len(respBody.Errors) > 0 {
			return nil, nil, fmt.Errorf("GraphQL error: %s", respBody.Errors[0].Message)
		}

		rateLimitInfo = parseRateLimitInfo(resp)

		if queryType == "following" {
			nodes = append(nodes, respBody.Data.User.Following.Nodes...)
			if !respBody.Data.User.Following.PageInfo.HasNextPage {
				break
			}
			endCursor = &respBody.Data.User.Following.PageInfo.EndCursor
		} else {
			nodes = append(nodes, respBody.Data.User.Followers.Nodes...)
			if !respBody.Data.User.Followers.PageInfo.HasNextPage {
				break
			}
			endCursor = &respBody.Data.User.Followers.PageInfo.EndCursor
		}
	}
	return &nodes, rateLimitInfo, nil
}

func getBaseUser(username string, c *fasthttp.Client) (*UserNode, error) {
	var allFollowing []UserNode
	var endCursor *string

	for {
		query := userQueryTemplate

		body := map[string]interface{}{
			"query": query,
			"variables": map[string]interface{}{
				"username": username,
				"after":    endCursor,
			},
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
			return nil, err
		}

		respBody := &GraphQlResponse{}
		if err := json.Unmarshal(resp.Body(), respBody); err != nil {
			return nil, err
		}

		if len(respBody.Errors) > 0 {
			return nil, fmt.Errorf("GraphQL error: %s", respBody.Errors[0].Message)
		}

		allFollowing = append(allFollowing, respBody.Data.User.Following.Nodes...)
		if !respBody.Data.User.Following.PageInfo.HasNextPage {
			respBody.Data.User.Following.Nodes = allFollowing
			return &respBody.Data.User, nil
		}

		endCursor = &respBody.Data.User.Following.PageInfo.EndCursor
	}
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

func checkRateLimit(c *fasthttp.Client) (int, time.Time, error) {
	query := `
		query {
			rateLimit {
				remaining
				resetAt
			}
		}
	`

	body := map[string]interface{}{
		"query": query,
	}
	bodyJSON, _ := json.Marshal(body)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(githubUrl)
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.SetBody(bodyJSON)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.Do(req, resp); err != nil {
		return 0, time.Time{}, err
	}

	var rateLimitResponse RateLimitResponse
	if err := json.Unmarshal(resp.Body(), &rateLimitResponse); err != nil {
		return 0, time.Time{}, err
	}

	return rateLimitResponse.Data.RateLimit.Remaining, rateLimitResponse.Data.RateLimit.ResetAt, nil
}

func handleRateLimit(c *fasthttp.Client) error {
	remaining, resetAt, err := checkRateLimit(c)
	if err != nil {
		return err
	}

	if remaining < 10 {
		sleepDuration := time.Until(resetAt)
		fmt.Printf("Rate limit approaching, sleeping for %v...\n", sleepDuration)
		time.Sleep(sleepDuration)
	}

	return nil
}
