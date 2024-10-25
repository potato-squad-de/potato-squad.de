import { addDynamicIconSelectors } from '@iconify/tailwind'
import type { Config } from 'tailwindcss'

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	darkMode: 'class',

	theme: {
		extend: {
			fontFamily: {
				riffic: ['Riffic', 'cursive', 'sans-serif']
			},
			colors: {
				text: 'var(--text)',
				background: 'var(--background)',
				primary: 'var(--primary)',
				secondary: 'var(--secondary)',
				accent: 'var(--accent)'
			}
		}
	},

	plugins: [addDynamicIconSelectors()]
} as Config
