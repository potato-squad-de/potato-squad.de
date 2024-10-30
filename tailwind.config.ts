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
				accent: 'var(--accent)',
				discord: '#5765f2'
			}
		}
	},

	safelist: [
		'icon-[mdi--web]',
		'icon-[logos--discord-icon]',
		'icon-[logos--twitch]',
		'icon-[logos--x]',
		'icon-[logos--youtube-icon]',
		'icon-[logos--tiktok-icon]',
		'icon-[skill-icons--instagram]',
		'icon-[logos--github-icon]',
		'icon-[logos--spotify-icon]',
		'icon-[logos--soundcloud]',
		'icon-[logos--linkedin-icon]',
		'icon-[logos--facebook]'
	],

	plugins: [addDynamicIconSelectors()]
} as Config
