<script lang="ts">
	import Card from './Card.svelte'

	const cards = [
		{ name: 'Hurensohn 1', alias: 'Alpha Hurensohn' },
		{ name: 'Hurensohn 2', alias: 'Beta Hurensohn' },
		{ name: 'Hurensohn 3', alias: 'Gamma Hurensohn' },
		{ name: 'Hurensohn 4', alias: 'Delta Hurensohn' }
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
