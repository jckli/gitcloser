import type { Viewport } from "next";
import { config } from "@fortawesome/fontawesome-svg-core";

import "../globals.css";
import "@fortawesome/fontawesome-svg-core/styles.css";

import Head from "./head";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";

config.autoAddCss = false;

export const viewport: Viewport = {
  userScalable: false,
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <Head />
      <body className="min-h-svh flex flex-col">
        <Header />
        <div className="flex flex-col flex-grow justify-center">{children}</div>
        <Footer />
      </body>
    </html>
  );
}
