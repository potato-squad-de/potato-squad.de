<script lang="ts">
	import '../app.css'

	let { children } = $props()
</script>

<svelte:head>
	<title>Potato Squad</title>
	<meta name="description" content="Haters gonna hate, Potatoes gonna potate" />
</svelte:head>

<div
	class="bg-background text-text grid min-h-screen select-none grid-rows-[auto_1fr] gap-4"
>
	<header
		class="flex w-full flex-col items-center justify-between gap-4 bg-none px-4 py-2 md:flex-row 2xl:px-64"
	>
		<a href="/" class="flex items-center">
			<img src="/images/logo.png" alt="logo" class="h-24 md:h-16" />
		</a>

		<a
			href="https://discord.gg/pHdZKq6mf9"
			target="_blank"
			rel="noopener noreferrer"
			class="flex min-w-fit items-center gap-2 rounded-lg bg-[#7289da] p-2 font-riffic text-2xl uppercase transition duration-100 ease-in-out hover:bg-[#97a4d4]"
		>
			<span class="icon-[mdi--discord]"></span>
			<span>Join</span>
		</a>

		<div
			class="bg-secondary flex items-center gap-4 rounded-lg px-4 py-2 text-3xl font-bold"
		>
			<button aria-label="dark mode" class="flex justify-items-center">
				<span
					class="icon-[material-symbols--dark-mode] dark:icon-[material-symbols--light-mode]"
				></span>
			</button>

			<a href="/login" class="flex justify-items-center" aria-label="Login">
				<span class="icon-[mdi--login]"></span>
			</a>
		</div>
	</header>
	<main class="h-fit px-4 py-2 2xl:px-64">
		{@render children()}
	</main>
</div>
