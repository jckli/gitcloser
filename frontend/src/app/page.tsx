import { PrimaryButton } from "@/components/PrimaryButton";
import { SecondaryButton } from "@/components/SecondaryButton";
import { Tree } from "@/components/Tree";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

export default function Home() {
	return (
		<>
			<div className="w-full h-[600px] flex lg:flex-row flex-col gap-20 items-center justify-center text-3xl font-lexend text-zinc-300 px-6">
				<div className="text-left max-w-[400px] pt-48 lg:pt-0">
					<h3 className="mb-2 text-base text-zinc-400">
						Currently in Beta
					</h3>
					<h1 className="mb-4 font-semibold">
						Find how close you are to
						another GitHub user in 20
						seconds or less
					</h1>
					<PrimaryButton
						text="Start Now"
						link="/pathway"
						useNextLink={true}
					/>
					<SecondaryButton
						text="GitHub"
						icon={faGithub}
						link="https://github.com/jckli/gitcloser"
					/>
				</div>
				<div className="sm:block hidden">
					<Tree />
				</div>
			</div>
		</>
	);
}
