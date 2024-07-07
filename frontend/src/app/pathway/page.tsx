"use client";

import {
	SecondaryButton,
	SmallerSecondaryButton,
} from "@/components/SecondaryButton";
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/Dialog";
import { Input } from "@/components/ui/Input";
import useWebSocket from "@/hooks/useWebsocket";
import { useState, useRef, useEffect } from "react";
import { PathwayVisualization } from "@/components/Pathway";
import useDebounce from "@/hooks/useDebounce";
import Image from "next/image";

type User = {
	login: string;
	avatarUrl: string;
	url: string;
	bio: string;
	followers: {
		totalCount: number;
	};
	following: {
		totalCount: number;
	};
};

export default function Pathway() {
	const [open, setOpen] = useState<boolean>(false);
	const [calculating, setCalculating] = useState<boolean>(false);
	const [startUser, setStartUser] = useState<string>("");
	const [endUser, setEndUser] = useState<string>("");
	const [oldStartUser, setOldStartUser] = useState<string>("");
	const [oldEndUser, setOldEndUser] = useState<string>("");
	const [typingStartUser, setTypingStartUser] = useState<string>("");
	const [typingEndUser, setTypingEndUser] = useState<string>("");
	const [url, setUrl] = useState<string>("");
	const [showStartPopover, setShowStartPopover] =
		useState<boolean>(false);
	const [showEndPopover, setShowEndPopover] = useState<boolean>(false);
	const [searchResults, setSearchResults] = useState<User[]>([]);
	const { messages } = useWebSocket(url);
	const jsonRef = useRef(null);
	const [lastMessage, setLastMessage] = useState<string>("");
	const [error, setError] = useState<string>("");

	useEffect(() => {
		const lastMessage = messages[messages.length - 1];
		let json = null;

		try {
			json = JSON.parse(lastMessage);
		} catch (error) {
			if (lastMessage) {
				const formatMessage = lastMessage.replace(
					/^processing_user:\s*/,
					"",
				);
				setLastMessage(formatMessage);
			}
		}

		if (json) {
			if (json.error) {
				setError(json.error);
				setCalculating(false);
				return;
			}
			setError("");
			jsonRef.current = json;
			setLastMessage("");
			messages.splice(0, messages.length); // Clear messages
			setOpen(false);
			setTimeout(() => {
				setCalculating(false);
			}, 500);
		}
	}, [messages]);

	const fetchSearchResults = async (query: string) => {
		if (!query) {
			setSearchResults([]);
			return;
		}

		try {
			const response = await fetch(
				`https://gitcloserapi.hayasaka.moe/v1/github/search/${query}`,
			);
			if (!response.ok) {
				throw new Error("Network response was not ok");
			}
			const data = await response.json();
			data.data.results = data.data.results.filter(
				(user: User) => user.login && user.avatarUrl,
			);

			if (data.data.results.length > 5) {
				data.data.results = data.data.results.slice(
					0,
					5,
				);
			}

			setSearchResults(data.data.results);
		} catch (error) {
			console.error("Error fetching search results:", error);
			setSearchResults([]);
		}
	};

	const debouncedFetchSearchResults = useDebounce(
		fetchSearchResults,
		300,
	);

	const handleInputChange = (
		event: any,
		setTypingUser: any,
		setPopover: any,
	) => {
		const { value } = event.target;
		if (setPopover == setShowStartPopover) {
			setShowEndPopover(false);
		} else {
			setShowStartPopover(false);
		}
		setTypingUser(value);
		setSearchResults([]);
		debouncedFetchSearchResults(value);
		setPopover(true);
	};

	const handleFindPath = () => {
		if (!startUser || !endUser) {
			setError("Please enter a start and end user.");
			return;
		}

		if (startUser == endUser) {
			setError("Start and end users cannot be the same.");
			return;
		}

		if (startUser == oldStartUser && endUser == oldEndUser) {
			setError("Please enter a new start and end user.");
			return;
		}

		const start = document.getElementById(
			"startUser",
		) as HTMLInputElement;
		const end = document.getElementById(
			"endUser",
		) as HTMLInputElement;
		console.log(startUser, endUser, start.value, end.value);

		if (start.value != startUser || end.value != endUser) {
			setError("Please select only users from the dropdown.");
			return;
		}

		setOldStartUser(startUser);
		setOldEndUser(endUser);

		setCalculating(true);
		setUrl(
			`https://gitcloserapi.hayasaka.moe/v1/github/pathway/${startUser}/${endUser}/ws`,
		);
	};

	const handleSelectUser = (user: any, setUser: any, setPopover: any) => {
		setUser(user);
		if (setUser == setStartUser) {
			setTypingStartUser(user);
		} else {
			setTypingEndUser(user);
		}
		setPopover(false);
		setSearchResults([]);
	};

	return (
		<>
			<div className="md:h-[80vh] flex flex-col justify-center">
				<div className="w-full flex flex-col items-center justify-center text-3xl font-lexend text-zinc-300">
					<Dialog
						open={open}
						onOpenChange={setOpen}
					>
						<DialogTrigger>
							<div className="mt-4">
								<SecondaryButton
									text="Change Users"
									useClassicButton={
										true
									}
								/>
							</div>
						</DialogTrigger>
						<DialogContent className="sm:max-w-[425px]">
							<DialogHeader>
								<DialogTitle>
									Change
									Users
								</DialogTitle>
								<DialogDescription>
									Change
									the
									users
									you want
									to
									pathfind
									here.
									Click
									find
									path
									when
									you're
									done.
								</DialogDescription>
							</DialogHeader>
							{!calculating ? (
								<>
									{error && (
										<p className="text-red-400 text-sm text-lexend">
											Error:{" "}
											{
												error
											}
										</p>
									)}
									<div
										className={`grid gap-4 ${!error ? "py-4" : "pb-4"}`}
									>
										<div className="flex flex-col gap-1">
											<h1 className="font-lexend text-zinc-300 text-base text-left">
												Start
												User
											</h1>
											<Input
												id="startUser"
												placeholder="Enter a GitHub username"
												className="col-span-3"
												onChange={(
													event,
												) =>
													handleInputChange(
														event,
														setTypingStartUser,
														setShowStartPopover,
													)
												}
												value={
													typingStartUser
														? typingStartUser
														: ""
												}
											/>
											{showStartPopover && (
												<div
													className={`absolute ${error ? "top-[215px]" : "top-[195px]"} bg-ender-black border-[1px] border-white/10 rounded-lg text-zinc-300 sm:w-[375px] font-lexend`}
												>
													{searchResults.length >
													0 ? (
														searchResults.map(
															(
																user,
															) => (
																<div
																	key={
																		user.login
																	}
																	className="p-2 hover:bg-zinc-800/80 rounded-lg cursor-pointer flex gap-2 items-center transition-all ease-in-out duration-200"
																	onClick={() =>
																		handleSelectUser(
																			user.login,
																			setStartUser,
																			setShowStartPopover,
																		)
																	}
																>
																	<Image
																		src={
																			user.avatarUrl
																		}
																		alt={
																			user.login
																		}
																		width={
																			40
																		}
																		height={
																			40
																		}
																		className="rounded-full"
																		unoptimized={
																			true
																		}
																	/>
																	<p>
																		{
																			user.login
																		}
																	</p>
																</div>
															),
														)
													) : (
														<div className="p-2 text-center text-zinc-500">
															<p>
																No
																results
															</p>
														</div>
													)}
												</div>
											)}
										</div>
										<div className="flex flex-col gap-1">
											<h1 className="font-lexend text-zinc-300 text-base text-left">
												End
												User
											</h1>
											<Input
												id="endUser"
												placeholder="Enter a GitHub username"
												className="col-span-3"
												onChange={(
													event,
												) =>
													handleInputChange(
														event,
														setTypingEndUser,
														setShowEndPopover,
													)
												}
												value={
													typingEndUser
														? typingEndUser
														: ""
												}
											/>
											{showEndPopover && (
												<div
													className={`absolute ${error ? "top-[300px]" : "top-[280px]"} bg-ender-black border-[1px] border-white/10 rounded-lg text-zinc-300 sm:w-[375px] font-lexend`}
												>
													{searchResults.length >
													0 ? (
														searchResults.map(
															(
																user,
															) => (
																<div
																	key={
																		user.login
																	}
																	className="p-2 hover:bg-zinc-800/80 cursor-pointer flex gap-2 items-center transition-all ease-in-out duration-200 rounded-lg"
																	onClick={() =>
																		handleSelectUser(
																			user.login,
																			setEndUser,
																			setShowEndPopover,
																		)
																	}
																>
																	<Image
																		src={
																			user.avatarUrl
																		}
																		alt={
																			user.login
																		}
																		width={
																			40
																		}
																		height={
																			40
																		}
																		className="rounded-full"
																		unoptimized={
																			true
																		}
																	/>
																	<p>
																		{
																			user.login
																		}
																	</p>
																</div>
															),
														)
													) : (
														<div className="p-2 text-center text-zinc-500">
															<p>
																No
																results
															</p>
														</div>
													)}
												</div>
											)}
										</div>
									</div>
									<DialogFooter>
										<SmallerSecondaryButton
											text="Find Path"
											onClick={
												handleFindPath
											}
										/>
									</DialogFooter>
								</>
							) : (
								<div className="mt-4">
									<h2 className="font-lexend text-zinc-300 text-base text-left">
										Finding
										Pathway:
									</h2>
									<p className="text-zinc-500 font-lexend">
										{
											lastMessage
										}
									</p>
								</div>
							)}
						</DialogContent>
					</Dialog>
					<div className="mt-4">
						<PathwayVisualization
							json={jsonRef.current}
						/>
					</div>
				</div>
			</div>
		</>
	);
}
