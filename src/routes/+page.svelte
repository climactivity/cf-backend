<script lang="ts">
	import { goto } from '$app/navigation';
	import { currentUser } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';

	onMount(() => {
		console.log($currentUser);
		if ($currentUser) {
			goto('/app');
		} else {
			goto('/welcome');
		}
	});
</script>

<h1>This will be the app</h1>
