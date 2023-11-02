<script lang="ts">
	import { menuOpen } from '$lib/Services/MenuState';
	import { currentUser } from '$lib/Services/PocketbaseWrapper';
	import { clamp, mapToRange } from '$lib/Util';
	import MenuButton from './MenuButton.svelte';
	import SideNavigationMenu from './SideNavigationMenu.svelte';
	import StreakCounter from './StreakCounter.svelte';
	import WarmingstripesBackground from './WarmingstripesBackground.svelte';

	export let bgShadowOffset = 15.0;
	export let bgOpacityStartOffset = 15.0;
	export let bgOpacityFullOffest = 40.0;

	let bgOpacity = 0.0;
	let bgCastShadow = false;
	let y = 0;

	$: bgOpacity = clamp(mapToRange(y, bgOpacityStartOffset, bgOpacityFullOffest));
	$: bgCastShadow = y > bgShadowOffset;
</script>

<header class="h-32">
	<WarmingstripesBackground adaptAngle opacity={ 1 - bgOpacity } />
	<div class="fixed  z-20       
	left-0 right-0 top-0 px-4 pt-5 pb-4 overflow-clip
		{bgCastShadow || $menuOpen ? 'shadow-md' : ''}">
		<div
		style="grid-template-columns: 5rem 1fr 10rem;"
		class="grid grid-flow-col items-stretch  max-w-xl mx-auto"
	>
		{#if $currentUser}
			<MenuButton/>
			<div/>
			<StreakCounter /> 
		{/if}

	</div>
	<WarmingstripesBackground opacity={bgCastShadow ? 1 : 0}   />


	</div>
	<!-- <div class="h-24" /> -->
	<SideNavigationMenu/>
</header>

<svelte:window bind:scrollY={y} />

<style>
	@media (min-width: 480px) {
		header {
			padding: 12px 0;
		}
	}
</style>
