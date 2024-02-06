<script lang="ts">
	import { goto } from '$app/navigation';
	import CfButton from '$lib/Components/CFButton.svelte';
	import TopSpacer from '$lib/Components/TopSpacer.svelte';
	import { pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';
	import Modal from '$lib/Components/Modal.svelte';
	import Imprint from '$lib/Components/Imprint.svelte';
	import CFSectionHeading from '$lib/Components/CFSectionHeading.svelte';

	onMount(async () => {
		if (await pb.authStore.isValid) {
			goto('/app');
		}
	});

	let imprintModal: Modal;
</script>

<TopSpacer logo={true} />
<div class="flex flex-col h-full">
	<div>
		<h1>Bitte trage dich als Mitmacherin bzw. Mitmacher ein,</h1>

		<ul>
			<li>... damit wir zeigen können, wie viele wir sind und sehen können, wo wir wachsen.</li>
			<li>
				... damit wir dir am Donnerstag eine kleine Erinnerung schicken können, dass am nächsten Tag
				wieder Climate Friday ist. So reißt dein Streak nicht so leicht ab ;-)
			</li>
			<li>Danke, dass du mitmachst!</li>
		</ul>
	</div>

	<div class="flex flex-col mx-auto w-[16rem] py-4 gap-4">
		<CfButton onclick={() => goto('/app/auth/withEmail')}>Mit Email anmelden</CfButton>
		<CfButton disabled onclick={() => goto('/app/auth/withApple')}>Mit Apple anmelden</CfButton>
		<CfButton disabled onclick={() => goto('/app/auth/withGoogle')}>Mit Google anmelden</CfButton>
	</div>

	<span class="text-center text-sm text-neutral-800 underline mt-20"><button class="underline"
																																						 on:click={imprintModal.show()}>Impressum</button></span>
</div>

<Modal bind:this={imprintModal}>
	<div>
		<CFSectionHeading>Impressum</CFSectionHeading>
		<Imprint />
	</div>
</Modal>