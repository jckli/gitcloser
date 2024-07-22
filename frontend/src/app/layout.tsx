import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";
import Head from "./head";
import "../globals.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="en">
			<Head />
			<body className="min-h-screen flex flex-col">
				<Header />
				<div className="flex flex-col flex-grow">
					{children}
				</div>
				<Footer />
			</body>
		</html>
	);
}
