<script lang="ts">
	import CfButton from './CFButton.svelte';
	import Modal from './Modal.svelte';
	import NotificationSettingsModal from './NotificationSettingsModal.svelte';
	import { notificationsSetup } from '$lib/Services/UserState';

	let modal: Modal;

	notificationsSetup.subscribe((v) => console.log(v));
</script>

<div>
	{#if $notificationsSetup.notification_setup}
		<span class="font-bold text-lg">Danke, dass Du am nächsten Freitag dabei bist!</span>
		{#if !$notificationsSetup.notification_email && !$notificationsSetup.notification_push }
			<div class="text-sm text-neutral-800 py-4">
				<span class="font-bold ">Tipp:</span>
				<p class="">
					Falls Du doch eine Benachrichtigung bekommen möchtest klicke unten. Wir spammen dein Postfach auch nicht mit
					Müll voll, versprochen.
				</p>
			</div>
		{/if}
		<button class="underline font-bold" on:click={() => modal.show()}>Benachrichtigungseinstellungen ändern</button>

	{:else}
		<CfButton onclick={() => modal.show()}>Ich bin am nächsten Freitag dabei!</CfButton>

	{/if}
</div>

<Modal bind:this={modal}>
	<NotificationSettingsModal modal={modal} />
</Modal>
