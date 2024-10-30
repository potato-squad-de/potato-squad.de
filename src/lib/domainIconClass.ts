const subDomains = ['www.', 'm.', 'mobile.', 'm.', 'open.']

export function getDomainFromUrl(url: string) {
	const urlObject = new URL(url)
	const domain = urlObject.hostname
	for (const subDomain of subDomains) {
		if (domain.startsWith(subDomain)) {
			return domain.replace(subDomain, '')
		}
	}
	return domain
}

export function getIconClassFromDomain(domain: string) {
	switch (domain) {
		default:
			return 'icon-[mdi--web]'
		case 'discord.gg':
			return 'icon-[logos--discord-icon]'
		case 'twitch.tv':
			return 'icon-[logos--twitch]'
		case 'twitter.com':
		case 'x.com':
			return 'icon-[logos--x]'
		case 'instagram.com':
			return 'icon-[skill-icons--instagram]'
		case 'youtube.com':
			return 'icon-[logos--youtube-icon]'
		case 'tiktok.com':
			return 'icon-[logos--tiktok-icon]'
		case 'github.com':
			return 'icon-[logos--github-icon]'
		case 'spotify.com':
			return 'icon-[logos--spotify-icon]'
		case 'soundcloud.com':
			return 'icon-[logos--soundcloud]'
		case 'linkedin.com':
			return 'icon-[logos--linkedin-icon]'
		case 'facebook.com':
			return 'icon-[logos--facebook]'
	}
}
