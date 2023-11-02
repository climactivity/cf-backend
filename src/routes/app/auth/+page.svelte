<script>
	import { goto } from '$app/navigation';
	import { pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';

	onMount(async () => {
		if (await pb.authStore.isValid) {
			goto('/app');
		}
	});
</script>

<h1>Bitte trage dich als Mitmacherin bzw. Mitmacher ein,</h1>

<ul>
	<li>... damit wir zeigen können, wie viele wir sind und sehen können, wo wir wachsen.</li>
	<li>
		... damit wir dir am Donnerstag eine kleine Erinnerung schicken können, dass am nächsten Tag
		wieder Climate Friday ist. So reißt dein Streak nicht so leicht ab ;-)
	</li>
	<li>Danke, dass du mitmachst!</li>
</ul>

<div id="login-options" class="flex flex-col absolute bottom-0 left-0 right-0 py-10">
	<button id="login-email" on:click={() => goto('/app/auth/withEmail')}> Mit Email anmelden </button>
	<button id="login-apple"> Mit Apple anmelden </button>
	<button id="login-google"> Mit Google anmelden </button>
</div>
