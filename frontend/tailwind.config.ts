import type { Config } from "tailwindcss";

const config = {
	darkMode: ["class"],
	content: [
		"./pages/**/*.{ts,tsx}",
		"./components/**/*.{ts,tsx}",
		"./app/**/*.{ts,tsx}",
		"./src/**/*.{ts,tsx}",
	],
	prefix: "",
	theme: {
		container: {
			center: true,
			padding: "2rem",
			screens: {
				xxsm: "380px",
				xsm: "460px",
				sxsm: "510px",
				sm: "640px",
				md: "768px",
				mmlg: "810px",
				mlg: "894px",
				lg: "1024px",
				xl: "1280px",
				"2xl": "1400px",
			},
		},
		extend: {
			keyframes: {
				"accordion-down": {
					from: { height: "0" },
					to: {
						height: "var(--radix-accordion-content-height)",
					},
				},
				"accordion-up": {
					from: {
						height: "var(--radix-accordion-content-height)",
					},
					to: { height: "0" },
				},
			},
			animation: {
				"accordion-down":
					"accordion-down 0.2s ease-out",
				"accordion-up": "accordion-up 0.2s ease-out",
			},
			colors: {
				ender: {
					"main-blue": "#4E8386",
					"dark-green": "#005500",
					"dark-blue": "#1B4D3E",
					"dark-purple": "#331355",
					black: "#0F0F0F",
					"neon-green": "#00FF00",
					"neon-blue": "#00CED1",
					"dark-gray": "#2B2B2B",
					"medium-gray": "#4D4D4D",
				},
			},
			fontFamily: {
				lexend: ["Lexend"],
			},
		},
	},
	plugins: [require("tailwindcss-animate")],
} satisfies Config;

export default config;
