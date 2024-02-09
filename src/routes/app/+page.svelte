<script lang="ts">
	import BlogEmbed from '$lib/Components/BlogEmbed.svelte';
	import CfFridayInteraction from '$lib/Components/CFFridayInteraction.svelte';
	import FridayCountdown from '$lib/Components/FridayCountdown.svelte';
	import StatCTABlock from '$lib/Components/StatCTABlock.svelte';
	import RegionSettingPrompt from '$lib/Components/RegionSettingPrompt.svelte';
</script>

<div class="h-28" />

<main class="flex flex-col gap-4">
	<FridayCountdown />

	<RegionSettingPrompt />

	<CfFridayInteraction />

	<StatCTABlock />
	<BlogEmbed />
</main>
