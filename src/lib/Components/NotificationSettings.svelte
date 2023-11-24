<script lang="ts">
	import CfButton from './CFButton.svelte';
	import Modal from './Modal.svelte';
	import NotificationSettingsModal from './NotificationSettingsModal.svelte';
	let modal: Modal;
</script>

<div>
	<CfButton onclick={() => modal.show()}>Ich bin dabei am nächsten Freitag</CfButton>
</div>

<Modal bind:this={modal}>
	<NotificationSettingsModal />
</Modal>
