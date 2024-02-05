<script lang="ts">
	import CfHeading from '$lib/Components/CFHeading.svelte';
	import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
	import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
	import IntersectionObserver from '$lib/Components/IntersectionObserver.svelte';
	import ParticipationCard from '$lib/Components/ParticipationCard.svelte';
	import TopSpacer from '$lib/Components/TopSpacer.svelte';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { DateTime } from 'luxon';

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
		if (fetchedItems.length == 0 || fetchedItems.length <= PER_PAGE) {
			endReached = true;
		}
		participations = [...participations, ...fetchedItems];
		currentPage = currentPage + 1;
	};

	const startDate = DateTime.fromSQL($currentUser?.created).toFormat('dd-MM-yyyy');



	// onMount(async () => {
	// 	// let _res = await pb.collection('participations').getList(currentPage, PER_PAGE,{ sort: "-date" });
	// 	// participations = _res.items as Participation[]
	// 	// console.log($currentUser);
	// 	// console.log(participations);
	// });
</script>

<TopSpacer />

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
	{#if endReached}
		<div class="bg-primary-800 text-white px-2 py-2 rounded-full font-bold text-center mb-8">
			Am {startDate} hast du mit
			<ClimateFriday invert />
			angefangen!
		</div>
	{/if}
	<IntersectionObserver onIntersect={fetchNextPage}>
		{#if !endReached}

			<div class="w-full flex flex-row justify-center">
				<span class="loading loading-dots text-primary-500 text-center text-4xl" />
			</div>
		{/if}
	</IntersectionObserver>
</div>
