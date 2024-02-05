<script lang="ts">
	import { goto } from '$app/navigation';
	import BlogEmbed from '$lib/Components/BlogEmbed.svelte';
	import CfFridayInteraction from '$lib/Components/CFFridayInteraction.svelte';
	import FridayCountdown from '$lib/Components/FridayCountdown.svelte';
	import { currentUser } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';
	import StatCTABlock from '$lib/Components/StatCTABlock.svelte';

	onMount(async () => {
		if (!$currentUser) {
			goto('/app/auth');
		}
	});
</script>

<div class="h-36" />

<main class="flex flex-col gap-4">
	<FridayCountdown />
	<CfFridayInteraction />

	<StatCTABlock />
	<BlogEmbed />
</main>
