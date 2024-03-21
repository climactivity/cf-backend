<script lang="ts">
	import { canSubmit } from '$lib/Services/TimerStore';
	import CfSectionHeading from './CFSectionHeading.svelte';
	import NotificationSettings from './NotificationSettings.svelte';
	import ParticipateForm from './ParticipateForm.svelte';
	import { currentParticipation } from '$lib/Services/UserState';
	import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
	import CFButton from '$lib/Components/CFButton.svelte';
	import Modal from '$lib/Components/Modal.svelte';
	import CongratulationModal from '$lib/Components/CongratulationModal.svelte';

	let changeParticipation = false;

	currentParticipation.subscribe(_ => changeParticipation = false)
	let congratulateModal: Modal;
</script>
<CfSectionHeading>Mitmachen</CfSectionHeading>
{#if canSubmit()}
	<div>

		{#if $currentParticipation}
			<div class="font-bold py-4"> Danke das Du diese Woche beim <ClimateFriday/> mitgemacht hast! </div>
			{#if changeParticipation}
				<ParticipateForm {congratulateModal}/>
				{:else}
				<CFButton onclick={() => changeParticipation = true}>Ändern</CFButton>
			{/if}
		{:else}
			<ParticipateForm {congratulateModal}/>
		{/if}

	</div>


	<Modal bind:this={congratulateModal}>
		<CongratulationModal />
	</Modal>


{:else}
	<NotificationSettings />
{/if}