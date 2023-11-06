<script lang="ts">
	import _logo from '$lib/Images/an_logo.png';
	export let logo = false;
</script>

{#if logo}
	<div class="h-36 flex flex-col items-center">
		<img class="w-24 h-24" alt="ClimateFriday logo" src={_logo} />
	</div>
{:else}
	<div class="h-36" />
{/if}
