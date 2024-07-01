package algorithm

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
	"sync"
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
				bio
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
				bio
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
		bio
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

func getPage(
	username, query string,
	endCursor *string,
	c *fasthttp.Client,
) (*UserNode, *RateLimitInfo, error) {

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

	rateLimitInfo := parseRateLimitInfo(resp)
	return &respBody.Data.User, rateLimitInfo, nil
}

func getUser(username, queryType string, c *fasthttp.Client) (*[]UserNode, *RateLimitInfo, error) {
	var query string

	if queryType == "following" {
		query = followingQueryTemplate
	} else if queryType == "followers" {
		query = followersQueryTemplate
	} else {
		return nil, nil, fmt.Errorf("invalid query type: %s", queryType)
	}

	firstPage, rateLimitInfo, err := getPage(username, query, nil, c)
	if err != nil {
		return nil, nil, err
	}

	if !firstPage.PageInfo.HasNextPage {
		if queryType == "following" {
			return &firstPage.Following.Nodes, rateLimitInfo, nil
		} else {
			return &firstPage.Followers.Nodes, rateLimitInfo, nil
		}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var allNodes []UserNode
	var endCursor = firstPage.PageInfo.EndCursor
	var pages []UserNode

	if queryType == "following" {
		allNodes = append(allNodes, firstPage.Following.Nodes...)
	} else {
		allNodes = append(allNodes, firstPage.Followers.Nodes...)
	}
	pageCount := 1

	wg.Add(1)

	go func() {
		defer wg.Done()
		for endCursor != "" {
			pageCount++
			page, _, err := getPage(username, query, &endCursor, c)
			if err != nil {
				fmt.Println("error fetching page: ", err)
				break
			}
			mu.Lock()
			if queryType == "following" {
				allNodes = append(pages, page.Following.Nodes...)
			} else {
				allNodes = append(pages, page.Followers.Nodes...)
			}
			mu.Unlock()
			endCursor = page.PageInfo.EndCursor
			if !page.PageInfo.HasNextPage {
				break
			}
		}

	}()
	wg.Wait()
	fmt.Println("Total pages fetched: ", pageCount)
	return &allNodes, rateLimitInfo, nil
}

func getBaseUser(
	username string,
	c *fasthttp.Client,
	following bool,
) (*UserNode, *RateLimitInfo, error) {
	query := userQueryTemplate

	firstPage, rateLimitInfo, err := getPage(username, query, nil, c)
	if err != nil {
		return nil, nil, err
	}

	if !firstPage.Following.PageInfo.HasNextPage || !following {
		return firstPage, rateLimitInfo, nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var allNodes []UserNode
	var endCursor = firstPage.Following.PageInfo.EndCursor
	var pages []UserNode

	allNodes = append(allNodes, firstPage.Following.Nodes...)
	pageCount := 1

	wg.Add(1)

	go func() {
		defer wg.Done()
		for endCursor != "" {
			pageCount++
			page, _, err := getPage(username, query, &endCursor, c)
			if err != nil {
				fmt.Println("error fetching page: ", err)
				break
			}
			mu.Lock()
			allNodes = append(pages, page.Following.Nodes...)
			mu.Unlock()
			endCursor = page.Following.PageInfo.EndCursor
			if !page.Following.PageInfo.HasNextPage {
				break
			}
		}
	}()
	wg.Wait()
	fmt.Println("Total pages fetched: ", pageCount)
	firstPage.Following.Nodes = allNodes

	return firstPage, rateLimitInfo, nil
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
