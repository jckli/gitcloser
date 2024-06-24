package algorithm

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

var (
	githubUrl = "https://api.github.com/graphql"
	token     = os.Getenv("GITHUB_TOKEN")
)

func getUser(username, queryType string, c *fasthttp.Client) ([]UserNode, error) {
	var query string

	if queryType == "following" {
		query = fmt.Sprintf(`
		query {
			user(login: "%s") {
				following(first: 100) {
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
				}
			}
		}`, username)
	} else if queryType == "followers" {
		query = fmt.Sprintf(`
		query {
			user(login: "%s") {
				followers(first: 100) {
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
				}
			}
		}`, username)
	} else if queryType == "user" {
		query = fmt.Sprintf(`
		query {
			user(login: "%s") {
				login
				avatarUrl
				followers(first: 100) {
					nodes {
						login
						avatarUrl
						url
					}
					totalCount
				}
				following(first: 100) {
					nodes {
						login
						avatarUrl
						url
						
					}
					totalCount
				}
				url
			}
		}`, username)
	} else {
		return nil, fmt.Errorf("invalid query type: %s", queryType)
	}

	body := map[string]interface{}{
		"query": query,
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

	if queryType == "following" {
		return respBody.Data.User.Following.Nodes, nil
	} else if queryType == "followers" {
		return respBody.Data.User.Followers.Nodes, nil
	} else {
		return []UserNode{{
			Login:     respBody.Data.User.Login,
			AvatarUrl: respBody.Data.User.AvatarUrl,
			Url:       respBody.Data.User.Url,
			Followers: respBody.Data.User.Followers,
			Following: respBody.Data.User.Following,
		}}, nil

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
