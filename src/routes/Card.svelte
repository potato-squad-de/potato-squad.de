<script lang="ts">
	type Props = {
		name: string
		alias?: string
	}

	let props: Props = $props()
</script>

<div class="min-h-64 rounded-xl bg-zinc-700 p-3">
	<h2 class="font-riffic text-2xl uppercase">{props.name}</h2>
	<h3 class="pl-4 font-riffic text-lg italic">{props.alias}</h3>
</div>
