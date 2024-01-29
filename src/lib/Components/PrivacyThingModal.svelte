<script lang="ts">
    import CfButton from './CFButton.svelte';
    import CfSectionHeading from './CFSectionHeading.svelte';
    import type Modal from "$lib/Components/Modal.svelte";
    import {pb} from "$lib/Services/PocketbaseWrapper";
    import {onMount} from "svelte";

    let notification_email = false;
    let notification_push = false;

    let updateSuccess = false;

    export let modal: Modal;


</script>

<CfSectionHeading>Datenschutzerklärung</CfSectionHeading>
<div class="flex flex-col gap-4 py-8">
    <div>
        Datenschutzerklärung hier einfügen
    </div>
</div>
