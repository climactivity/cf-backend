<script>
	import { goto } from '$app/navigation';
	import CfButton from '$lib/Components/CFButton.svelte';
	import TopSpacer from '$lib/Components/TopSpacer.svelte';
	import { pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';

	onMount(async () => {
		if (await pb.authStore.isValid) {
			goto('/app');
		}
	});
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
		<CfButton onclick={() => goto('/app/auth/withApple')}>Mit Apple anmelden</CfButton>
		<CfButton onclick={() => goto('/app/auth/withGoogle')}>Mit Google anmelden</CfButton>
	</div>
</div>
