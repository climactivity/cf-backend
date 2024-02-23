<script>
	import Header from '$lib/Components/Header.svelte';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(async () => {
		console.log(pb.authStore.isValid);
		if (!$currentUser || !pb.authStore.isValid) {
			pb.authStore.clear();
			goto('/app/auth');
		}
	});
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
