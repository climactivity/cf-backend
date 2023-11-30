<script lang="ts">
	import CfButton from './CFButton.svelte';
	import Modal from './Modal.svelte';
	import NotificationSettingsModal from './NotificationSettingsModal.svelte';
	let modal: Modal;
</script>

<div>
	<CfButton onclick={() => modal.show()}>Ich bin am nächsten Freitag dabei!</CfButton>
</div>

<Modal bind:this={modal}>
	<NotificationSettingsModal modal={modal} />
</Modal>
