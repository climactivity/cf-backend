<script lang="ts">
	import { goto } from '$app/navigation';
	import BlogEmbed from '$lib/Components/BlogEmbed.svelte';
	import CfFridayInteraction from '$lib/Components/CFFridayInteraction.svelte';
	import CfSection from '$lib/Components/CFSection.svelte';
	import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
	import CountUp from '$lib/Components/CountUp.svelte';
	import FridayCountdown from '$lib/Components/FridayCountdown.svelte';
	import ParticipateForm from '$lib/Components/ParticipateForm.svelte';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';

	onMount(async () => {
		if (!$currentUser) {
			goto('/app/auth');
		}
	});
</script>

<div class="h-36" />

<main class="flex flex-col gap-4">
	<CfSection>
		<!-- <span slot="heading">Hello</span> -->
		<div>Teilnehmer:innen am letzten Freitag:<br /> <CountUp number={1200} /></div>
		<div>Aus deinem PLZ-Gebiet waren es<br /> <CountUp number={234} /></div>
	</CfSection>
	<FridayCountdown />

	<CfFridayInteraction />
	<BlogEmbed />
</main>
