package algorithm

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// implement a bidirectional bfs algorithm to find the shortest path between two nodes, which are startUser and endUser. strictly only go from startUser's following and recursively on to targetUser from targetUser's followers it should form a line like this: startUser -> people startUser follows/people following targetUser -> targetUser
func FindShortestPath(startUser, endUser string, c *fasthttp.Client) ([]UserNode, error) {
	startNode := UserNode{Login: startUser, Visited: false}
	endNode := UserNode{Login: endUser, Visited: false}

	startQueue := []UserNode{startNode}
	endQueue := []UserNode{endNode}

	startVisited := make(map[string]UserNode)
	endVisited := make(map[string]UserNode)

	startVisited[startUser] = startNode
	endVisited[endUser] = endNode

	for len(startQueue) > 0 && len(endQueue) > 0 {
		newSQ, err := bfs(startQueue, startVisited, "start", c)
		if err != nil {
			return nil, err
		}

		newEQ, err := bfs(endQueue, endVisited, "end", c)
		if err != nil {
			return nil, err
		}

		intersect, startNode, endNode := isIntersection(startVisited, endVisited)
		if intersect {
			startPath := getPath(startNode)
			endPath := getPath(endNode)
			return append(startPath, endPath...), nil
		}
		startQueue = newSQ
		endQueue = newEQ

	}

	return nil, fmt.Errorf("no path found")
}

func bfs(
	queue []UserNode,
	visited map[string]UserNode,
	direction string, c *fasthttp.Client,
) ([]UserNode, error) {
	node := queue[0]
	queue = queue[1:]

	if direction == "start" {
		following, err := getUser(node.Login, "following", c)
		if err != nil {
			return nil, err
		}

		for _, v := range following {
			if !visited[v.Login].Visited {
				queue = append(queue, v)
				v.Prev = &node
				v.Visited = true
				visited[v.Login] = v
			}
		}
	} else {
		followers, err := getUser(node.Login, "followers", c)
		if err != nil {
			return nil, err
		}

		for _, v := range followers {
			if !visited[v.Login].Visited {
				queue = append(queue, v)
				v.Prev = &node
				v.Visited = true
				visited[v.Login] = v
			}
		}
	}

	return queue, nil
}

func isIntersection(startVisited, endVisited map[string]UserNode) (bool, UserNode, UserNode) {
	fmt.Println(startVisited)
	fmt.Println(endVisited)
	for k := range startVisited {
		if endVisited[k].Visited {
			return true, startVisited[k], endVisited[k]
		}
	}
	return false, UserNode{}, UserNode{}

}

func getPath(node UserNode) []UserNode {
	path := []UserNode{}
	for node.Prev != nil {
		path = append(path, node)
		node = *node.Prev
	}
	path = append(path, node)
	fmt.Println("path")
	fmt.Println(path)
	return path
}

func reversePath(path []UserNode) []UserNode {
	reversed := make([]UserNode, len(path))
	for i, v := range path {
		reversed[len(path)-1-i] = v
	}
	return reversed
}
