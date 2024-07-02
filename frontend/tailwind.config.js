/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./{src,app}/**/*.{ts,tsx}"],
	theme: {
		extend: {
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
	variants: {
		extend: {},
	},
};
