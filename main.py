from src import github
import aiohttp
import asyncio

async def main():
    async with aiohttp.ClientSession() as session:
        following = await github.get_github_following(session, "jckli")
        print(following)

if __name__ == "__main__":
    asyncio.run(main())

