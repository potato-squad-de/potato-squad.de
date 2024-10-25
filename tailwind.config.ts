import { addDynamicIconSelectors } from '@iconify/tailwind'
import type { Config } from 'tailwindcss'

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
} as Config
