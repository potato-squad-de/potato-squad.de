<script lang="ts">
	import { type CardProps } from '$lib'

	let props: CardProps = $props()
</script>

<div
	class="bg-primary flex min-h-64 justify-between gap-8 rounded-xl p-3 drop-shadow"
>
	<div class="min-w-64">
		<h2 class="font-riffic text-2xl uppercase">{props.name}</h2>
		{#if props.aliases}
			<h3 class="text-md pl-4 font-riffic italic">
				aka {props.aliases?.join(', ')}
			</h3>
		{/if}
		<img
			src={props.imgUrl ? props.imgUrl : '/images/avatar_placeholder.png'}
			alt="avatar"
			class="size-64"
		/>
	</div>
	<div class="font-['Comic_Sans_MS'] italic">
		{#if props.bio}
			<p class="text-md">{props.bio}</p>
		{/if}
	</div>
	<div class="flex flex-col justify-between gap-4 text-xl">
		<div class="flex flex-col gap-2">
			{#if props.socials}
				{#each props.socials as social}
					<a href={social.url} target="_blank" rel="noopener noreferrer">
						<span class="icon-[mdi--{social.platform}]"></span>
						{social.username}
					</a>
				{/each}
			{/if}
		</div>
		<div class="flex flex-col gap-2">
			{#if props.otherRefs}
				{#each props.otherRefs as ref}
					<a href={ref.url} target="_blank" rel="noopener noreferrer">
						<span class="icon-[mdi--link-variant]"></span>
						{ref.name}
					</a>
				{/each}
			{/if}
		</div>
	</div>
</div>
