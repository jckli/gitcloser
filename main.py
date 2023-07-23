import aiohttp
import asyncio
from src import algorithm

async def main():
    user_1 = "sindresorhus"
    user_2 = "kentcdodds"
    path = await algorithm.bidirectional_bfs(user_1, user_2)
    print(path)

if __name__ == "__main__":
    asyncio.run(main())

