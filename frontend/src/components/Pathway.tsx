"use client";

import Image from "next/image";
import { motion } from "framer-motion";

export const PathwayVisualization = (props: any) => {
	if (!props.json) {
		return (
			<>
				<div>
					<h1>
						Select a start and end user
						first
					</h1>
				</div>
			</>
		);
	}

	const pathway = props.json.data.pathway;
	console.log(pathway);

	return (
		<>
			<div className="flex gap-20">
				{pathway.map((user: any, index: number) => (
					<div key={index}>
						<div className="relative z-10 w-60 h-72 border-white/10 border-[1px] rounded-lg bg-ender-black text-zinc-300">
							<div
								className={`py-6 px-8 flex flex-col items-center ${user.bio == "" && "justify-center"} h-full`}
							>
								<Image
									src={
										user.avatarUrl
									}
									alt={
										user.login
									}
									width={
										100
									}
									height={
										100
									}
									className="rounded-full"
									unoptimized={
										true
									}
								/>
								<h1 className="text-2xl mt-1 text-center">
									{
										user.login
									}
								</h1>
								{user.bio !=
									"" && (
									<p className="text-base mt-4 text-zinc-500">
										{
											user.bio
										}
									</p>
								)}
							</div>
						</div>
						{index < pathway.length - 1 && (
							<motion.div
								className={`relative left-[280px] bottom-[200px] w-[2px] bg-ender-medium-gray z-0`}
								style={{
									height: 120,
									rotate: "-90deg",
									animation: "ease 1s infinite",
									backgroundImage:
										"linear-gradient(0deg,transparent 33%,hsla(0,0%,100%,.5) 50%,transparent 66%)",
									backgroundSize:
										"100% 300%",
								}}
								animate={{
									backgroundPosition:
										[
											"0% 100%",
											"0% 0%",
										],
								}}
								transition={{
									backgroundPosition:
										{
											duration: 1,
											repeat: Infinity,
											ease: "easeInOut",
										},
								}}
							/>
						)}
					</div>
				))}
			</div>
		</>
	);
};
