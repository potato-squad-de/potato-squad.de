<script lang="ts">
	import '../app.css'

	let { children } = $props()
</script>

<svelte:head>
	<title>Potato Squad</title>
	<meta name="description" content="Haters gonna hate, Potatoes gonna potate" />
</svelte:head>

<div
	class="grid min-h-screen select-none grid-rows-[auto_1fr] gap-4 bg-background text-text"
>
	<header
		class="flex w-full flex-col items-center justify-between gap-4 bg-none px-4 py-2 sm:flex-row 2xl:px-64"
	>
		<a href="/" class="flex items-center">
			<img src="/images/logo.png" alt="logo" class="h-24 sm:h-16" />
		</a>

		<div class="flex items-center gap-4">
			<a
				href="https://discord.gg/pHdZKq6mf9"
				target="_blank"
				rel="noopener noreferrer"
				class="hover:bg-discord/50 bg-discord flex min-w-min items-center gap-2 rounded-lg p-2 font-riffic text-5xl uppercase drop-shadow transition duration-100 ease-in-out sm:text-3xl"
			>
				<span class="icon-[mdi--discord]"></span>
				<span>Join</span>
			</a>
		</div>

		<div
			class="hidden items-center gap-4 rounded-lg bg-secondary px-4 py-2 text-3xl font-bold drop-shadow dark:bg-transparent dark:drop-shadow-none sm:flex"
		>
			<button
				aria-label="dark mode"
				class="flex justify-items-center"
				onclick={() => document.documentElement.classList.toggle('dark')}
			>
				<span
					class="icon-[material-symbols--dark-mode] dark:icon-[material-symbols--light-mode]"
				></span>
			</button>
		</div>
	</header>
	<main class="h-fit px-4 py-2 2xl:px-64">
		{@render children()}
	</main>
</div>
