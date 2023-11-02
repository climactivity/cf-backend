<script lang="ts">
	import CfHeading from '$lib/Components/CFHeading.svelte';
	import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
	import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
	import IntersectionObserver from '$lib/Components/IntersectionObserver.svelte';
	import ParticipationCard from '$lib/Components/ParticipationCard.svelte';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';

	let participations: Participation[] = [];

	let currentPage = 1;
	let fetching = false;
	let endReached = false;
	const PER_PAGE = 20;

	const fetchNextPage = async () => {
		if (endReached || fetching) return;
		fetching = true;
		const _res = await pb
			.collection('participations')
			.getList(currentPage, PER_PAGE, { sort: '-date' });
		fetching = false;

		const fetchedItems = _res.items as Participation[];
		if (fetchedItems.length == 0) {
			endReached = true;
		}
		participations = [...participations, ...fetchedItems];
		currentPage = currentPage + 1;
	};

	// onMount(async () => {
	// 	// let _res = await pb.collection('participations').getList(currentPage, PER_PAGE,{ sort: "-date" });
	// 	// participations = _res.items as Participation[]
	// 	// console.log($currentUser);
	// 	// console.log(participations);
	// });


</script>

<div>
	<CfSectionHeading>Highlights</CfSectionHeading>
</div>
<div>
	<CfHeading>Verlauf</CfHeading>

	<div>
		{#each participations as participation}
			<ParticipationCard {participation} />
		{/each}
	</div>
	<IntersectionObserver onIntersect={fetchNextPage}>
		{#if endReached}
			<div class="bg-primary-800 text-white px-2 py-2 rounded-full font-bold  text-center mb-8">Hier hast du mit <ClimateFriday invert/> angefangen!</div>
		{:else}
			<div class="w-full flex flex-row justify-center">
				<span class="loading loading-dots text-primary-500 text-center text-4xl" />
			</div>
		{/if}
	</IntersectionObserver>
</div>
