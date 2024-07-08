<img src="https://github.com/jckli/gitcloser/blob/master/gitcloser.png" alt="GitCloser Logo" width="300"/>

_**Git** as in the version control system "Git" & **Closer** as in "how close one is"_

# 🔎 Find how close you are to another GitHub user in less than 20 seconds

> GitCloser Algorithm is located at `/backend/algorithm`

GitCloser is a web application that makes it super easy to find out how close you are to another GitHub user. It utilizes a Bidirectional Breadth First Search (BFS) algorithm to find the (mostly shortest) path between two users on GitHub. Implemented in Go with go routines for concurrency, the algorithm takes around 20 seconds or less for most users on GitHub.

You can try it out at [gitcloser.hayasaka.moe](https://gitcloser.hayasaka.moe) without deploying anything youself - just type in two GitHub usernames and click "Find Path"! If you want to self host it though, it'll take a bit of configuration.

## Self-hosting

### Prerequisites

- Go 1.22
- Node.js 22.3.0
- GitHub account

### Steps

1. Clone the repository
2. Create a GitHub OAuth App at [https://github.com/settings/tokens?type=beta](https://github.com/settings/tokens?type=beta). You can just generate one with no scopes or permissions. Copy the API key.
3. Create a `.env` file in the `/backend` directory with the following content:

```bash
GITHUB_TOKEN=YOUR_GITHUB_API_KEY
PORT=3001
```

4. Run the backend server with `go run main.go` in the `/backend` directory
5. Install the frontend dependencies with `npm install` in the `/frontend` directory
6. Run the frontend server with `npm run start` in the `/frontend` directory
7. You can find GitCloser at `http://localhost:3000` in your browser
