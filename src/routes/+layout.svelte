<script lang="ts">
	import '../app.css'

	let { children } = $props()
</script>

<svelte:head>
	<title>Potato Squad</title>
	<meta name="description" content="Haters gonna hate, Potatoes gonna potate" />
</svelte:head>

<div class="h-screen select-none bg-[#99ccff] flex flex-col items-center justify-center gap-4">
        {@render children()}
</div>
