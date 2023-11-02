<script lang="ts">
	import { pb } from '$lib/Services/PocketbaseWrapper';
	import { onDestroy, onMount } from 'svelte';

	export let region: string;

	let unsubscribe: Function[] = [];

	let userCount = 0;
	let regionCount = 0;

	const globalUpdated = (e: any) => {
		console.log('global update', e);
	};

	const regionUpdated = (e: any) => {
		console.log('retion update', e);
	};

	onMount(async () => {
		try {
			let _resGlobal = await pb.collection('user_count').getOne('usercount');
			userCount = _resGlobal.count;
		} catch (e) {}
		try {
			let _resRegion = await pb.collection('user_count').getOne(region);
			regionCount = _resRegion.count;
		} catch (e) {}

		unsubscribe.push(await pb.collection('user_count').subscribe('usercount', globalUpdated));
		unsubscribe.push(await pb.collection('user_count').subscribe(region, globalUpdated));
	});

	onDestroy(() => {
		unsubscribe.map((v) => v());
	});
</script>

Teilnehmer:innen am letzten Freitag: {userCount} Aus deinem PLZ-Gebiet waren es {regionCount}
