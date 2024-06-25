package algorithm

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// implement a bidirectional bfs algorithm to find the shortest path between two nodes, which are startUser and endUser. strictly only go from startUser's following and recursively on to targetUser from targetUser's followers it should form a line like this: startUser -> people startUser follows/people following targetUser -> targetUser
func FindShortestPath(startUser, endUser string, c *fasthttp.Client) ([]UserNode, error) {

	startUserInfo, err := getBaseUser(startUser, c)
	if err != nil {
		return nil, err
	}

	if startUserInfo.Following.TotalCount == 0 {
		return nil, fmt.Errorf("no path found")
	}

	endUserInfo, err := getBaseUser(endUser, c)
	if err != nil {
		return nil, err
	}

	if endUserInfo.Followers.TotalCount == 0 {
		return nil, fmt.Errorf("no path found")
	}

	startNode := UserNode{
		Login:     startUser,
		Prev:      nil,
		AvatarUrl: startUserInfo.AvatarUrl,
		Url:       startUserInfo.Url,
	}
	startNode.Following.TotalCount = startUserInfo.Following.TotalCount
	startNode.Followers.TotalCount = startUserInfo.Followers.TotalCount
	endNode := UserNode{
		Login:     endUser,
		Prev:      nil,
		AvatarUrl: endUserInfo.AvatarUrl,
		Url:       endUserInfo.Url,
	}
	endNode.Following.TotalCount = endUserInfo.Following.TotalCount
	endNode.Followers.TotalCount = endUserInfo.Followers.TotalCount

	for _, v := range startUserInfo.Following.Nodes {
		if v.Login == endUser {
			return []UserNode{startNode, endNode}, nil
		}
	}

	startQueue := []UserNode{startNode}
	endQueue := []UserNode{endNode}

	startVisited := make(map[string]UserNode)
	endVisited := make(map[string]UserNode)

	startVisited[startUser] = startNode
	endVisited[endUser] = endNode

	for len(startQueue) > 0 && len(endQueue) > 0 {
		fmt.Println(startQueue)
		fmt.Println("-----")
		fmt.Println(endQueue)
		fmt.Println("-----")
		newSQ, err := bfs(&startQueue, &startVisited, "start", c)
		if err != nil {
			continue
		}

		newEQ, err := bfs(&endQueue, &endVisited, "end", c)
		if err != nil {
			continue
		}

		intersect, startNode, endNode := isIntersection(&startVisited, &endVisited)
		fmt.Println(intersect, startNode, endNode)

		if intersect {
			startPath := getPath(&startNode)
			reversePath(&startPath)
			endPath := getPath(&endNode)
			return append(startPath, endPath[1:]...), nil
		}
		startQueue = *newSQ
		endQueue = *newEQ
	}

	return nil, fmt.Errorf("no path found")
}

func bfs(
	queue *[]UserNode,
	visited *map[string]UserNode,
	direction string, c *fasthttp.Client,
) (*[]UserNode, error) {
	node := (*queue)[0]
	*queue = (*queue)[1:]

	if direction == "start" {
		following, rateLimitFollowing, err := getUser(node.Login, "following", c)
		if err != nil {
			return nil, err
		}
		fmt.Println(rateLimitFollowing)
		if rateLimitFollowing.Remaining == 0 {
			return nil, fmt.Errorf(
				"rate limit reached, try again in %d seconds",
				rateLimitFollowing.Reset,
			)
		}

		for _, v := range *following {
			if _, exists := (*visited)[v.Login]; !exists {
				v.Prev = &node
				(*visited)[v.Login] = v
				*queue = append((*queue), v)
			}
		}
	} else {
		followers, rateLimitFollowers, err := getUser(node.Login, "followers", c)
		if err != nil {
			return nil, err
		}
		fmt.Println(rateLimitFollowers)
		if rateLimitFollowers.Remaining == 0 {
			return nil, fmt.Errorf("rate limit reached, try again in %d seconds", rateLimitFollowers.Reset)
		}

		for _, v := range *followers {
			if _, exists := (*visited)[v.Login]; !exists {
				v.Prev = &node
				(*visited)[v.Login] = v
				*queue = append((*queue), v)
			}
		}
	}

	return queue, nil
}

func isIntersection(startVisited, endVisited *map[string]UserNode) (bool, UserNode, UserNode) {
	for k, v := range *startVisited {
		if u, exists := (*endVisited)[k]; exists {
			return true, v, u
		}
	}
	return false, UserNode{}, UserNode{}
}

func getPath(node *UserNode) []UserNode {
	path := []UserNode{*node}
	for node.Prev != nil {
		path = append(path, *node.Prev)
		node = node.Prev
	}
	return path
}

func reversePath(path *[]UserNode) {
	for i := 0; i < len(*path)/2; i++ {
		(*path)[i], (*path)[len(*path)-i-1] = (*path)[len(*path)-i-1], (*path)[i]
	}
}
