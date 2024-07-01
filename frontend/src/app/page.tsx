import { GitHubButton } from "@/components/GitHubButton";

export default function Home() {
	return (
		<>
			<div className="w-full h-[100vh] flex flex-col items-center justify-center text-white text-3xl font-semibold text-center">
				<h1 className="mb-4">jckli&apos;s next.js template</h1>
				<GitHubButton />
			</div>
		</>
	);
}
