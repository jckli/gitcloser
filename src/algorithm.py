import aiohttp
import asyncio
import os
from dotenv import load_dotenv
from . import github

async def bidirectional_bfs(user_1, user_2):
    queue_1 = [user_1]
    queue_2 = [user_2]
    visited_1 = {user_1: None}
    visited_2 = {user_2: None}
    
    load_dotenv()
    client_id = os.environ.get("CLIENT_ID")
    client_secret = os.environ.get("CLIENT_SECRET")

    if client_id is None or client_secret is None:
        print("Missing CLIENT_ID or CLIENT_SECRET")
        return
    
    async with aiohttp.ClientSession(auth=aiohttp.BasicAuth(client_id, client_secret)) as session:
        while queue_1 and queue_2:
            print(visited_1, visited_2)
            print("Queue 1:", queue_1)
            print("Queue 2:", queue_2)
            print("\n")

            user_1 = queue_1.pop(0)
            user_2 = queue_2.pop(0)

            print("Current User 1:", user_1)
            print("Current User 2:", user_2)

            following_1 = await github.get_github_following(session, user_1)
            following_2 = await github.get_github_following(session, user_2)

            if following_1[1] is None or following_2[1] is None:
                return None

            following_1_names = [user["login"] for user in following_1[1]]
            following_2_names = [user["login"] for user in following_2[1]]

            print("Following 1 Names:", following_1_names)
            print("Following 2 Names:", following_2_names)

            common_users = set(visited_1).intersection(visited_2)
            print("Common Users:", common_users)

            if common_users:
                common_user = common_users.pop()  # Pick one of the common users

                # Construct the path from user_1 to the common user
                path_from_user_1 = [user_1]
                current_user = user_1
                while current_user != common_user:
                    current_user = visited_1[current_user]
                    path_from_user_1.append(current_user)

                # Construct the path from user_2 to the common user
                path_from_user_2 = [user_2]
                current_user = user_2
                while current_user != common_user:
                    current_user = visited_2[current_user]
                    path_from_user_2.append(current_user)

                # Combine the paths to get the full path from user_1 to user_2
                path = path_from_user_1 + path_from_user_2[-2::-1]  # Exclude common_user and reverse path_from_user_2
                return path

            for user in following_1_names:
                if user not in visited_1:
                    visited_1[user] = user_1
                    queue_1.append(user)

            for user in following_2_names:
                if user not in visited_2:
                    visited_2[user] = user_2
                    queue_2.append(user)
    return None
