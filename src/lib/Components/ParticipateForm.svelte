<script lang="ts">
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { slugify } from '$lib/Util';
	import { DateTime } from 'luxon';
	import Modal from './Modal.svelte';
	import { fade } from 'svelte/transition';
	import ClimateFriday from './ClimateFriday.svelte';
	import CfButton from './CFButton.svelte';
	import CongratulationModal from './CongratulationModal.svelte';

    let gratulateModal: Modal

    let submitting = false;
    let submitted = false; 
	let selectedOption: string | undefined = undefined;
    let note: string | undefined = undefined;
	const options: { value: string | undefined; label: string }[] = [
		{
			value: 'banner',
			label: 'Mit Banner'
		},
		{
			value: 'window_foil',
			label: 'Mit Fensterfolie'
		},
		{
			value: 'button',
			label: 'Mit Button'
		},
		{
			value: 'sticker',
			label: 'Mit Sticker'
		},
		{
			value: 'flag',
			label: 'Mit Fähnchen'
		},
		{
			value: '',
			label: 'Leider gar nicht'
		}
	];
    
    const postParticipation = async () => {
        let participation: Partial<Participation> = {
			date: DateTime.now().toHTTP() || undefined, // <- typescript, am I right?
			note,
			option: selectedOption,
			weeknumber: DateTime.now().weekNumber,
            owner: $currentUser?.id ||''
		};
        return pb.collection("participations").create(participation);
    }

    const submit = async () => {
        submitting = true;

        try {
            let _res = await postParticipation();
            submitted = true;
            console.log(_res);
            gratulateModal.show();
        } catch (e) {
            console.error(e);
        } finally {
            submitting = false;
        }

        // setTimeout(() => {
        //     submitting = false;
        //     submitted = true
        //     gratulateModal.show();
        // }, 1000)
    }

    let disabled = true;
    $: disabled = submitting || selectedOption === undefined || selectedOption === '';
</script>

<form class="flex flex-col gap-4 max-w-4xl relative">
    {#if submitting}
        <div transition:fade class="absolute inset-0 bg-opacity-50 bg-white flex justify-center items-center">
            <span  class="loading loading-dots w-20 text-primary-500"></span>
        </div>
    {/if}
	<fieldset class="flex flex-col gap-4 ">
		<legend class="pb-4">Wie machst du heute beim Climate Friday mit?</legend>
		<div style="grid-template-columns: 1fr 1rem;" class="grid grid-flow-row gap-2">
			{#each options as { value, label }}
				<label for={slugify(label)}>{label}</label>
				<input id={slugify(label)} type="radio" name={label} {value} bind:group={selectedOption} />
			{/each}
		</div>
		<!-- <div class="form-control w-full max-w-xs">
			<label class="label px-0 pb-" for="note">
				<span class="label-text">Notiz</span>
			</label>
			<input
				placeholder="Notiz zu diesem Freitag"
				class="input input-bordered w-full max-w-xs"
				type="text"
				id="note"
				bind:value={note}
			/>
		</div> -->

	</fieldset>
    <div class="flex flex-row justify-end w-full">
        <button
		class="
            px-8
            py-4
            transition-all
            hover:duration-250
            active:duration-100
            font-semibold
            font-serif
            rounded-lg
            {disabled
                ? 'bg-neutral-100 text-neutral-400 hover:bg-neutral-100'
                : 'bg-primary-200 hover:bg-primary-300 text-primary-900'}
            {disabled
                ? ''
                : 'active:text-opacity-80 active:scale-[0.98] active:translate-y-0.5'}         
            "
		type="submit" {disabled}
		on:click|preventDefault={submit}>Abschicken</button
	>
    </div>

</form>


<Modal  bind:this={gratulateModal}>
    <CongratulationModal/>
</Modal>
