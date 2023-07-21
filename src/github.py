async def get_github_following(session, username):
    url = f"https://api.github.com/users/{username}/following"
    async with session.get(url) as response:
        if response.status == 200:
            return username, await response.json()
        elif response.status == 403:
            print("Rate limit exceeded")
            return username, None
    return username, []

