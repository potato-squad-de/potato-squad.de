<script lang="ts">
	import {
		getDomainFromUrl,
		getIconClassFromDomain
	} from '$lib/domainIconClass'

	interface Props {
		displayName?: string
		aliases?: string[]
		bio?: string
		avatarRef?: string
		potatoAvatarRef?: string
		links?: {
			url: string
			name: string
		}[]
	}

	let {
		displayName = 'Mysterious Potato',
		aliases = ['Le Potato', 'Don Potato'],
		bio = `
        Est occaecat sit mollit eiusmod dolor nisi elit consectetur sunt commodo excepteur quis. Qui voluptate aliqua aute esse pariatur pariatur. Est consequat nulla ullamco dolore sit elit elit elit minim do nulla. Elit eiusmod sit fugiat Lorem adipisicing ut ad dolor ut voluptate deserunt ipsum tempor. Consectetur proident labore nisi ex cupidatat irure proident incididunt anim ad ad. Deserunt et duis ad nostrud nostrud labore cillum nulla laborum nisi officia eu. Sint aute adipisicing minim officia eiusmod quis in consequat sunt officia quis. Tempor non adipisicing labore ea consequat duis ut laboris esse consectetur labore labore. Mollit laborum dolor qui magna quis duis laboris ullamco. Aute Lorem fugiat elit non laboris anim magna consequat laborum minim eu dolore. Nostrud consequat elit nostrud nostrud nisi enim excepteur duis non duis consequat.
        `,
		avatarRef = '/images/avatar_placeholder.png',
		potatoAvatarRef = '/images/potato_avatar_placeholder.png',
		links = [
			{
				name: 'Homepage',
				url: 'https://potato-squad.de'
			},
			{
				name: 'Potato Squad',
				url: 'https://discord.gg/potatosquad'
			},
			{
				name: 'diesmal zieht er durch',
				url: 'https://twitch.tv/beluxo_'
			},
			{
				name: 'Herzschmerz :((',
				url: 'https://twitter.com/@DjHearthunterz'
			},
			{
				name: 'gute vogel diese',
				url: 'https://youtube.com/@Berd'
			},
			{
				name: 'uga uga rock',
				url: 'https://www.tiktok.com/@rockcleaning10'
			},
			{
				name: 'wo contributions?',
				url: 'https://github.com/timmehde'
			},
			{
				name: 'Halb mensch, Halb Nase',
				url: 'https://instagram.com/ssio'
			},
			{
				name: '2016 hat angerufen',
				url: 'https://open.spotify.com/artist/2l0AkkQ8XXCFFbbJJeDlcI'
			},
			{
				name: 'wo album?',
				url: 'https://soundcloud.com/eroshardstyle'
			}
		]
	}: Props = $props()
</script>

<div
	class="flex min-h-64 flex-col justify-between gap-8 rounded-xl bg-primary p-3 drop-shadow sm:flex-row"
>
	<div class="min-w-64">
		<h2 class="font-riffic text-2xl uppercase">{displayName}</h2>
		{#if aliases}
			<h3 class="text-md pl-4 font-riffic italic">
				aka {aliases?.join(', ')}
			</h3>
		{/if}
		<div class="relative mx-auto inline-block md:mx-0">
			<img
				src={avatarRef}
				alt="avatar"
				class="size-48 object-contain"
			/>
			<img
				src={potatoAvatarRef}
				alt="potato avatar"
				class="absolute bottom-2 right-2 h-16 w-16 object-contain"
			/>
		</div>
	</div>

	<div class="prosa-xs">
		<p class="text-md">{bio}</p>
	</div>

	<div class="flex min-w-fit flex-col gap-2">
		{#each links as link}
			<a
				href={link.url}
				target="_blank"
				rel="noopener noreferrer"
				class="flex flex-row items-center gap-4 rounded-xl bg-slate-100/10 p-2 sm:flex-row-reverse"
			>
				<span
					class={'size-4 ' + getIconClassFromDomain(getDomainFromUrl(link.url))}
				></span>
				<span class="underline">{link.name}</span>
			</a>
		{/each}
	</div>
</div>
