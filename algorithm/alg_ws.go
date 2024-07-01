package algorithm

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/valyala/fasthttp"
	"sync"
)

var wsMutex sync.Mutex

func FindShortestPathWS(
	startUser, endUser string,
	conn *websocket.Conn,
	c *fasthttp.Client,
) ([]UserNode, error) {
	var startUserInfo, endUserInfo *UserNode
	var startUserErr, endUserErr error
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		startUserInfo, _, startUserErr = getBaseUser(startUser, c, true)
	}()
	go func() {
		defer wg.Done()
		endUserInfo, _, endUserErr = getBaseUser(endUser, c, false)
	}()
	wg.Wait()

	if startUserErr != nil {
		sendWebSocketMessage(conn, "error: "+startUserErr.Error())
		return nil, startUserErr
	}
	if startUserInfo.Following.TotalCount == 0 {
		sendWebSocketMessage(conn, "error: no path found")
		return nil, fmt.Errorf("no path found")
	}
	if endUserErr != nil {
		sendWebSocketMessage(conn, "error: "+endUserErr.Error())
		return nil, endUserErr
	}
	if endUserInfo.Followers.TotalCount == 0 {
		sendWebSocketMessage(conn, "error: no path found")
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
		var wg sync.WaitGroup
		errChan := make(chan error, 2)

		var newSQ, newEQ *[]UserNode

		wg.Add(2)
		go func() {
			defer wg.Done()
			var err error
			newSQ, err = bfsWS(&startQueue, &startVisited, "start", conn, c)
			if err != nil {
				errChan <- err
			}
		}()
		go func() {
			defer wg.Done()
			var err error
			newEQ, err = bfsWS(&endQueue, &endVisited, "end", conn, c)
			if err != nil {
				errChan <- err
			}
		}()

		wg.Wait()
		close(errChan)

		if len(errChan) > 0 {
			for err := range errChan {
				if err != nil {
					fmt.Println(err)
				}
			}
			continue
		}

		intersect, startNode, endNode := isIntersection(&startVisited, &endVisited)

		if intersect {
			startPath := getPath(&startNode)
			reversePath(&startPath)
			endPath := getPath(&endNode)
			return append(startPath, endPath[1:]...), nil
		}
		startQueue = *newSQ
		endQueue = *newEQ
	}

	sendWebSocketMessage(conn, "error: no path found")
	return nil, fmt.Errorf("no path found")
}

func bfsWS(
	queue *[]UserNode,
	visited *map[string]UserNode,
	direction string, conn *websocket.Conn, c *fasthttp.Client,
) (*[]UserNode, error) {
	node := (*queue)[0]
	*queue = (*queue)[1:]

	sendWebSocketMessage(conn, fmt.Sprintf("processing_user: %s", node.Login))

	if direction == "start" {
		following, rateLimitFollowing, err := getUser(node.Login, "following", c)
		if err != nil {
			return nil, err
		}
		fmt.Println(rateLimitFollowing)
		if rateLimitFollowing.Remaining == 0 {
			sendWebSocketMessage(
				conn,
				fmt.Sprintf(
					"error: rate limit reached, try again in %d seconds",
					rateLimitFollowing.Reset,
				),
			)
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
			sendWebSocketMessage(
				conn,
				fmt.Sprintf(
					"error: rate limit reached, try again in %d seconds",
					rateLimitFollowers.Reset,
				),
			)
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

func sendWebSocketMessage(conn *websocket.Conn, message string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	conn.WriteMessage(websocket.TextMessage, []byte(message))
}
