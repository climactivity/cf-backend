<script lang="ts">
	import CfButton from './CFButton.svelte';
	import Modal from './Modal.svelte';
	import NotificationSettingsModal from './NotificationSettingsModal.svelte';
	import { notificationsSetup } from '$lib/Services/UserState';
	import { currentWeek, justSetupWeek, postCurrentWeek } from '$lib/Components/weekApi';
	import { onMount } from 'svelte';

	let modal: Modal;

	notificationsSetup.subscribe((v) => console.log(v));


	// let currentWeekPromise: Promise<Week | undefined>
	// let currentWeek: Week | undefined

	const setupWeek = async () => {
		// if notifications are off, show the setup modal
		console.log($notificationsSetup);
		if (!$notificationsSetup.notification_email && !$notificationsSetup.notification_push) {
			modal.show();
			return;
		}
		await postCurrentWeek();
		justSetupWeek.set(true);
	};

	onMount(() => justSetupWeek.set(false));

</script>


{#if $currentWeek}
	<span id="c1" class="font-bold text-lg">Danke, dass Du am nächsten Freitag dabei bist!</span>
{:else }

	{#if $justSetupWeek}
		<!--TODO something fancier-->
		<span id="c3" class="font-bold text-lg">Danke, dass Du am nächsten Freitag dabei bist!</span>
	{:else}
		<CfButton onclick={() => setupWeek()}>Ich bin am nächsten Freitag dabei!</CfButton>
		{/if}

{/if}


<div>

</div>

<!--<div>-->
<!--	{#if $notificationsSetup.notification_setup}-->
<!--		<span class="font-bold text-lg">Danke, dass Du am nächsten Freitag dabei bist!</span>-->
<!--		{#if !$notificationsSetup.notification_email && !$notificationsSetup.notification_push }-->
<!--			<div class="text-sm text-neutral-800 py-4">-->
<!--				<span class="font-bold ">Tipp:</span>-->
<!--				<p class="">-->
<!--					Falls Du doch eine Benachrichtigung bekommen möchtest klicke unten. Wir spammen dein Postfach auch nicht mit-->
<!--					Müll voll, versprochen.-->
<!--				</p>-->
<!--			</div>-->
<!--		{/if}-->
<!--		<button class="underline font-bold" on:click={() => modal.show()}>Benachrichtigungseinstellungen ändern</button>-->

<!--	{:else}-->
<!--		<CfButton onclick={() => modal.show()}>Ich bin am nächsten Freitag dabei!</CfButton>-->

<!--	{/if}-->
<!--</div>-->

<Modal bind:this={modal}>
	<NotificationSettingsModal modal={modal} />
</Modal>
