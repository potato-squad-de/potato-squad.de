import { addDynamicIconSelectors } from "@iconify/tailwind"

/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./web/ui/**/*.go"],
	theme: {
		extend: {
			fontFamily: {
				riffic: ["Riffic", "cursive"]
			}
		}
	},
	plugins: [addDynamicIconSelectors(), require("tailwind-htmx")]
}
