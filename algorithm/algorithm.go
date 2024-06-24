package algorithm

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// implement a bidirectional bfs algorithm to find the shortest path between two nodes, which are startUser and endUser. strictly only go from startUser's following and recursively on to targetUser from targetUser's followers it should form a line like this: startUser -> people startUser follows/people following targetUser -> targetUser
func FindShortestPath(startUser, endUser string, c *fasthttp.Client) ([]UserNode, error) {

	startUserInfo, err := getUser(startUser, "user", c)
	if err != nil {
		return nil, err
	}

	if startUserInfo[0].Following.TotalCount == 0 {
		return nil, fmt.Errorf("no path found")
	}

	endUserInfo, err := getUser(endUser, "user", c)
	if err != nil {
		return nil, err
	}

	if endUserInfo[0].Followers.TotalCount == 0 {
		return nil, fmt.Errorf("no path found")
	}

	startNode := UserNode{
		Login:     startUser,
		Prev:      nil,
		AvatarUrl: startUserInfo[0].AvatarUrl,
		Url:       startUserInfo[0].Url,
	}
	startNode.Following.TotalCount = startUserInfo[0].Following.TotalCount
	startNode.Followers.TotalCount = startUserInfo[0].Followers.TotalCount
	endNode := UserNode{
		Login:     endUser,
		Prev:      nil,
		AvatarUrl: endUserInfo[0].AvatarUrl,
		Url:       endUserInfo[0].Url,
	}
	endNode.Following.TotalCount = endUserInfo[0].Following.TotalCount
	endNode.Followers.TotalCount = endUserInfo[0].Followers.TotalCount

	for _, v := range startUserInfo[0].Following.Nodes {
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
			return nil, err
		}

		newEQ, err := bfs(&endQueue, &endVisited, "end", c)
		if err != nil {
			return nil, err
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
		following, err := getUser(node.Login, "following", c)
		if err != nil {
			return nil, err
		}

		for _, v := range following {
			if _, exists := (*visited)[v.Login]; !exists {
				v.Prev = &node
				(*visited)[v.Login] = v
				*queue = append((*queue), v)
			}
		}
	} else {
		followers, err := getUser(node.Login, "followers", c)
		if err != nil {
			return nil, err
		}

		for _, v := range followers {
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
