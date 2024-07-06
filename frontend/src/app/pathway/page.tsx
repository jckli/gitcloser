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

export default function Pathway() {
	const [open, setOpen] = useState<boolean>(false);
	const [calculating, setCalculating] = useState<boolean>(false);
	const [startUser, setStartUser] = useState<string>("");
	const [endUser, setEndUser] = useState<string>("");
	const [url, setUrl] = useState<string>("");
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

	const handleFindPath = () => {
		const start = document.getElementById(
			"startUser",
		) as HTMLInputElement;
		const end = document.getElementById(
			"endUser",
		) as HTMLInputElement;

		if (!start.value || !end.value) {
			setError("Please enter a start and end user.");
			return;
		}

		if (start.value == end.value) {
			setError("Start and end users cannot be the same.");
			return;
		}

		if (start.value == startUser && end.value == endUser) {
			setError("Start and end users are the same as before.");
			return;
		}

		setCalculating(true);
		setStartUser(start.value);
		setEndUser(end.value);
		setUrl(
			`https://gitcloserapi.hayasaka.moe/v1/github/pathway/${start.value}/${end.value}/ws`,
		);
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
										<p className="text-red-400 text-sm">
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
												defaultValue={
													startUser
														? startUser
														: ""
												}
											/>
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
												defaultValue={
													endUser
														? endUser
														: ""
												}
											/>
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
									<p className="text-zinc-500">
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
