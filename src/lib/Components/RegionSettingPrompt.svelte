<script>
	import { onMount } from 'svelte';
	import { pb } from '$lib/Services/PocketbaseWrapper';
	import Modal from '$lib/Components/Modal.svelte';
	import CFSectionHeading from '$lib/Components/CFSectionHeading.svelte';
	import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
	import { fade } from 'svelte/transition';

	let plz;
	onMount(async () => {
		const userReccord = await pb.collection('users').getOne(pb.authStore.model?.id);
		plz = userReccord.region;
		console.log(plz);
	});

	const savePlz = async () => {
		let rec = await pb.collection('users').update(pb.authStore.model?.id, { region: plz });
		success = true;
		setTimeout(
			() => {
				success = false;
				modal.hide();
			},
			2000);
	};

	let modal;

	let success = false;
</script>


{#if !plz}
	<div>
		<span>ping</span>
		<button on:click={() => modal.show()}>click</button>
	</div>
{/if}

<Modal bind:this={modal}>
	<div class="flex-col flex gap-4 ">
		<CFSectionHeading>
			Postleitzahl
		</CFSectionHeading>
		<p>
			Bitte gib deine Postleitzahl ein damit wir sehen wo Menschen beim
			<ClimateFriday />
			mitmachen
		</p>
		<label class="label" for="plz">
			<span class="label-text">Postleitzahl</span>
			{#if success}
				<span transition:fade class="label-text text-success-700">Gespeichert!</span>
			{/if}
		</label>
		<input
			type="text"
			placeholder="bsp. 12345"
			class="input input-bordered"
			id="plz"
			bind:value={plz}
		/>
		<button class="text-right underline font-bold" on:click={savePlz}>Speichern</button>
	</div>
</Modal>