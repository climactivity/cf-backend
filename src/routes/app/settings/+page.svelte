<script>
	import { goto } from '$app/navigation';
	import CfButton from '$lib/Components/CFButton.svelte';
	import CfHeading from '$lib/Components/CFHeading.svelte';
	import TopSpacer from '$lib/Components/TopSpacer.svelte';
	import { pb } from '$lib/Services/PocketbaseWrapper';
    import {onMount} from "svelte";

	const signOut = () => {
		pb.authStore.clear();
		goto('/');
	};

    const save = async () => {
        let rec = await pb.collection("users").update(pb.authStore.model?.id, {region: plz});

    }

    let plz = "";

    onMount(async () => {
        let rec = await pb.collection("users").getOne(pb.authStore.model?.id);
        plz = rec.region;
    })
</script>

<TopSpacer />

<CfHeading>Settings</CfHeading>
<div>

    <label class="label" for="email">
        <span class="label-text">Postleitzahl</span>
    </label>
    <input
            type="text"
            placeholder="plz"
            class="input input-bordered w-[16rem]"
            id="plz"
            bind:value={plz}
    />

<CfButton onclick={save}>Speichern</CfButton>
</div>

<CfButton transparent onclick={signOut}>Abmelden</CfButton>
