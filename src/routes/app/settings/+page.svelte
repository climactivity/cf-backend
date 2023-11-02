<script>
	import { goto } from "$app/navigation";
	import CfButton from "$lib/Components/CFButton.svelte";
	import CfHeading from "$lib/Components/CFHeading.svelte";
	import { pb } from "$lib/Services/PocketbaseWrapper";

	const signOut = () => {
		pb.authStore.clear();
		goto("/");
	}
</script>

<CfHeading>Settings</CfHeading>

<CfButton transparent onclick={signOut}>Abmelden</CfButton>