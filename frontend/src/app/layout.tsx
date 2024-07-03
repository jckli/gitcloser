import { Header } from "@/components/Header";
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
			<body>
				<Header />
				{children}
			</body>
		</html>
	);
}
