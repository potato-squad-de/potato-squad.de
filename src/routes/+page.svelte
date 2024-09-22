<script lang="ts">
	import { getRandomWebsite } from '$lib'
</script>

<img src="/images/logo.png" alt="logo" />

<div
	class="flex flex-col items-center gap-4 font-riffic text-2xl uppercase text-white md:flex-row"
>
	<a
		href="https://discord.gg/pHdZKq6mf9"
		target="_blank"
		rel="noopener noreferrer"
		class="flex min-w-fit items-center gap-2 rounded-lg bg-[#7289da] p-2 transition duration-100 ease-in-out hover:bg-[#97a4d4]"
	>
		<span class="icon-[mdi--discord]"></span>
		<span>Join</span>
	</a>
	<button
		on:click={() => window.open(getRandomWebsite(), '_blank')}
		class="min-w-fit rounded-lg bg-[#4b0082] p-2 uppercase transition duration-100 ease-in-out hover:bg-[#7700ce]"
	>
		Traust dich eh nicht
	</button>
</div>
