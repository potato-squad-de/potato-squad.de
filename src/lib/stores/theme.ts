import { browser } from '$app/environment'
import { writable } from 'svelte/store'

type Theme = 'light' | 'dark'

export const getTheme = (): Theme | undefined => {
	if (browser) {
		if (window.localStorage.getItem('theme')) {
			return window.localStorage.getItem('theme') as Theme
		}

		if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
			return 'dark'
		}
		return 'light'
	}
}
const theme = writable<string>(getTheme())

theme.subscribe((value) => {
	if (browser) {
		if (value === 'dark') {
			document.documentElement.classList.add('dark')
		}

		if (value === 'light') {
			document.documentElement.classList.remove('dark')
		}

		window.localStorage.setItem('theme', value)
	}
})

export default theme
