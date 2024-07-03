"use client";

import { motion } from "framer-motion";
import { useEffect, useState } from "react";
import Image from "next/image";

export const Tree = () => {
	const [highlightPath, setHighlightPath] = useState<number | null>(null);

	useEffect(() => {
		// Highlight a random path after 7 seconds and never again
		setTimeout(() => {
			setHighlightPath(Math.floor(Math.random() * 5));
		}, 6000);
	}, []);
	console.log(highlightPath);

	const boxWidth = 560;
	const boxHeight = 288;

	const users = [
		"https://avatars.githubusercontent.com/u/83573936",
		"https://avatars.githubusercontent.com/u/57842793",
		"https://avatars.githubusercontent.com/u/35376349",
		"https://avatars.githubusercontent.com/u/89956908",
		"https://avatars.githubusercontent.com/u/92439990",
	];

	return (
		<div className="relative justify-center items-center flex-col flex">
			<motion.div
				className="relative z-10 rounded-full w-40 h-40 flex items-center justify-center text-white border-ender-medium-gray border-[1px]"
				initial={{ scale: 1, translateY: 100 }}
				animate={{ scale: 0.5, translateY: 0 }}
				transition={{
					duration: 0.5,
					delay: 2,
					ease: "easeInOut",
					type: "spring",
				}}
			>
				<Image
					src="https://avatars.githubusercontent.com/u/39673993"
					alt="jckli"
					width={160}
					height={160}
					unoptimized={true}
					className="rounded-full"
				/>
			</motion.div>
			<motion.div
				className={`absolute left[${boxWidth / 2}px] w-[2px] bg-ender-medium-gray`}
				style={{
					height: boxHeight,
					rotate: "48deg",
					x: -110,
					animation: "ease 1s infinite",
					backgroundImage:
						"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
					backgroundSize: "100% 300%",
				}}
				initial={{
					opacity: highlightPath == null ? 0 : 1,
				}}
				animate={{
					opacity:
						highlightPath === 0 ||
						highlightPath == null
							? 1
							: 0,
					backgroundPosition: [
						"0% 100%",
						"0% 0%",
					],
				}}
				transition={{
					opacity: {
						delay:
							highlightPath === 0 ||
							highlightPath != null
								? 0
								: 2.5,
						duration: 0.5,
					},
					backgroundPosition: {
						duration: 1,
						repeat: Infinity,
						ease: "easeInOut",
					},
				}}
			/>
			<motion.div
				className={`absolute left[${boxWidth / 2}px] w-[2px] bg-ender-medium-gray`}
				style={{
					height: boxHeight - 60,
					rotate: "35deg",
					x: -55,
					animation: "ease 1s infinite",
					backgroundImage:
						"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
					backgroundSize: "100% 300%",
				}}
				initial={{
					opacity: highlightPath == null ? 0 : 1,
				}}
				animate={{
					opacity:
						highlightPath === 1 ||
						highlightPath == null
							? 1
							: 0,
					backgroundPosition: [
						"0% 100%",
						"0% 0%",
					],
				}}
				transition={{
					opacity: {
						delay:
							highlightPath === 1 ||
							highlightPath != null
								? 0
								: 3,
						duration: 0.5,
					},
					backgroundPosition: {
						duration: 1,
						repeat: Infinity,
						ease: "easeInOut",
					},
				}}
			/>
			<motion.div
				className={`absolute left[${boxWidth / 2}px] w-[2px] bg-ender-medium-gray`}
				style={{
					height: boxHeight - 100,
					rotate: "0deg",
					animation: "ease 1s infinite",
					backgroundImage:
						"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
					backgroundSize: "100% 300%",
				}}
				initial={{
					opacity: highlightPath == null ? 0 : 1,
				}}
				animate={{
					opacity:
						highlightPath === 2 ||
						highlightPath == null
							? 1
							: 0,
					backgroundPosition: [
						"0% 100%",
						"0% 0%",
					],
				}}
				transition={{
					opacity: {
						delay:
							highlightPath === 2 ||
							highlightPath != null
								? 0
								: 3.5,
						duration: 0.5,
					},
					backgroundPosition: {
						duration: 1,
						repeat: Infinity,
						ease: "easeInOut",
					},
				}}
			/>
			<motion.div
				className={`absolute left[${boxWidth / 2}px] w-[2px] bg-ender-medium-gray`}
				style={{
					height: boxHeight - 60,
					rotate: "-35deg",
					x: 55,
					animation: "ease 1s infinite",
					backgroundImage:
						"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
					backgroundSize: "100% 300%",
				}}
				initial={{
					opacity: highlightPath == null ? 0 : 1,
				}}
				animate={{
					opacity:
						highlightPath === 3 ||
						highlightPath == null
							? 1
							: 0,
					backgroundPosition: [
						"0% 100%",
						"0% 0%",
					],
				}}
				transition={{
					opacity: {
						delay:
							highlightPath === 3 ||
							highlightPath != null
								? 0
								: 4,
						duration: 0.5,
					},
					backgroundPosition: {
						duration: 1,
						repeat: Infinity,
						ease: "easeInOut",
					},
				}}
			/>
			<motion.div
				className={`absolute left[${boxWidth / 2}px] w-[2px] bg-ender-medium-gray`}
				style={{
					height: boxHeight,
					rotate: "-48deg",
					x: 110,
					animation: "ease 1s infinite",
					backgroundImage:
						"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
					backgroundSize: "100% 300%",
				}}
				initial={{
					opacity: highlightPath == null ? 0 : 1,
				}}
				animate={{
					opacity:
						highlightPath === 4 ||
						highlightPath == null
							? 1
							: 0,
					backgroundPosition: [
						"0% 100%",
						"0% 0%",
					],
				}}
				transition={{
					opacity: {
						delay:
							highlightPath === 4 ||
							highlightPath != null
								? 0
								: 4.5,
						duration: 0.5,
					},
					backgroundPosition: {
						duration: 1,
						repeat: Infinity,
						ease: "easeInOut",
					},
				}}
			/>
			<div className="flex gap-[40px] mt-20">
				{Array.from({ length: 5 }).map((_, index) => (
					<motion.div
						key={index}
						initial={{
							opacity:
								highlightPath ==
								null
									? 0
									: 1,
						}}
						animate={{
							opacity:
								highlightPath ===
								index
									? 1
									: highlightPath ==
										  null
										? 1
										: 0.2,
						}}
						transition={{
							duration: 0.5,
							delay:
								highlightPath ===
								index
									? 0
									: highlightPath ==
										  null
										? 2.5 +
											index *
												0.5
										: 0,
						}}
					>
						<div className="relative z-10 rounded-full w-20 h-20 border-ender-medium-gray border-[1px] flex items-center justify-center text-white">
							<Image
								src={
									users[
										index
									]
								}
								className="rounded-full"
								alt="user"
								width={80}
								height={80}
								unoptimized={
									true
								}
							/>
						</div>
					</motion.div>
				))}
			</div>
		</div>
	);
};
