<script>
	import Footer from '$lib/Components/Footer.svelte';
	import Header from '$lib/Components/Header.svelte';
	import { pb, currentUser } from '$lib/Services/PocketbaseWrapper';
</script>

	<Header/>
	<main>
		<slot />
	</main>



<style>

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 0 var(--x-pad, 1.5rem);
		width: 100%;
		max-width: 38rem;
		margin: 0 auto;
		box-sizing: border-box;
	}
</style>
