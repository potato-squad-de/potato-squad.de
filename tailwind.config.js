import { addDynamicIconSelectors } from '@iconify/tailwind'

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				riffic: ['Riffic', 'cursive', 'sans-serif']
			}
		}
	},
	plugins: [addDynamicIconSelectors()]
}
