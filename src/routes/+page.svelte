<script lang="ts">
	import { goto } from '$app/navigation';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';


	onMount(() => {
		console.log($currentUser);

		if (!$currentUser || !pb.authStore.isValid) {
			pb.authStore.clear();
			goto('/welcome');
		}

		if ($currentUser) {
			goto('/app');
		} else {
			goto('/welcome');
		}
	});
</script>

<div class="py-20 min-w-full flex align-middle justify-center">
	Redirecting &nbsp;
	<span class="loading loading-dots text-primary-500 text-center text-4xl" />
</div>