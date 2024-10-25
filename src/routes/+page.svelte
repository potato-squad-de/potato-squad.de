<script lang="ts">
	import type { CardProps } from '$lib'
	import Card from './Card.svelte'

	const cards: CardProps[] = [
		{
			name: 'Hurensohn 1',
			aliases: ['Alpha Hurensohn', 'Richtiger Kevin'],
			bio: 'Id cillum culpa esse exercitation cillum commodo est sunt nostrud cillum aute magna quis. Eiusmod occaecat sit nulla ullamco minim cillum minim. Esse consequat pariatur sit cillum pariatur laborum dolore proident Lorem sint cillum ut. Exercitation consectetur elit eiusmod fugiat commodo.',
			socials: [
				{
					platform: 'discord',
					username: 'Kevin',
					url: 'https://discord.gg/pHdZKq6mf9'
				},
				{
					platform: 'github',
					username: 'Kevin',
					url: 'https://github.com/KevinHurensohn'
				},
				{
					platform: 'twitter',
					username: 'KevinHurensohn',
					url: 'https://twitter.com/KevinHurensohn'
				}
			],
			otherRefs: [{ name: 'Potato Squad', url: 'https://potato-squad.org' }]
		},
		{ name: 'Hurensohn 2', aliases: ['Beta Hurensohn'] },
		{ name: 'Hurensohn 3' },
		{ name: 'Hurensohn 4' }
	]
</script>

<div class="flex flex-col gap-4">
	<h1 class="font-riffic text-5xl drop-shadow">Lore</h1>

	<div class="flex flex-col gap-4">
		{#each cards as card}
			<Card {...card} />
		{/each}
	</div>
</div>
