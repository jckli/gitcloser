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
		setCalculating(true);
		const start = document.getElementById(
			"startUser",
		) as HTMLInputElement;
		const end = document.getElementById(
			"endUser",
		) as HTMLInputElement;
		setStartUser(start.value);
		setEndUser(end.value);
		setUrl(
			`https://gitcloserapi.hayasaka.moe/v1/github/pathway/${start.value}/${end.value}/ws`,
		);
	};

	return (
		<>
			<div className="w-full flex flex-col gap-20 items-center justify-center text-3xl font-lexend text-zinc-300">
				<Dialog open={open} onOpenChange={setOpen}>
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
								Change Users
							</DialogTitle>
							<DialogDescription>
								Change the users
								you want to
								pathfind here.
								Click find path
								when you're
								done.
							</DialogDescription>
						</DialogHeader>
						{!calculating ? (
							<>
								<div className="grid gap-4 py-4">
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
									Messages:
								</h2>
								<ul className="list-disc pl-5">
									<li className="text-zinc-300">
										{
											lastMessage
										}
									</li>
								</ul>
							</div>
						)}
					</DialogContent>
				</Dialog>
				<div>
					<PathwayVisualization
						json={jsonRef.current}
					/>
				</div>
			</div>
		</>
	);
}
