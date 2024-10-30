<script lang="ts">
	import { users } from '$lib/users'
	import Card from './Card.svelte'
</script>

<div class="flex flex-col gap-8">
	<div class="flex flex-col gap-4">
		<h1 class="font-riffic text-5xl drop-shadow">Lore</h1>

		<div class="flex flex-col gap-4">
			{#each Array(3) as user}
				<Card {...user} />
			{/each}
		</div>
	</div>
</div>
