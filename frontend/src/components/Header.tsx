import Image from "next/image";

export const Header = () => {
	return (
		<>
			<div className="bg-ender-black/80 sticky top-0 navbar backdrop-blur z-50 border-b-white/10 border-b-[1px]">
				<div className="flex items-center px-6 py-3 mx-auto max-w-[1300px] text-2xl justify-between md:justify-start font-lexend text-white">
					<div className="flex items-center">
						<Image
							src="/eye.png"
							alt="GitCloser Logo"
							width={50}
							height={50}
							draggable={false}
						/>
						<h1 className="text-ender-main-blue font-bold ml-1">GitCloser</h1>
					</div>
					<ul className="hidden md:flex gap-[20px] list-none text-white/40 ml-[40px] text-xl flex-1 font-bold">
						<li>
							<a
								href="#"
								className="hover:text-white/90 transition-all duration-200 ease-in-out"
							>
								home
							</a>
						</li>
						<li>
							<a
								href="#documentation"
								className="hover:text-white/90 transition-all duration-200 ease-in-out"
							>
								documentation
							</a>
						</li>
					</ul>
				</div>
			</div>
		</>
	);
};
