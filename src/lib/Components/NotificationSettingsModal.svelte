<script lang="ts">
    import CfButton from './CFButton.svelte';
    import CfSectionHeading from './CFSectionHeading.svelte';
    import type Modal from "$lib/Components/Modal.svelte";
    import {pb} from "$lib/Services/PocketbaseWrapper";
    import {onMount} from "svelte";

    let notification_email = false;
    let notification_push = false;

    let updateSuccess = false;
    const submit = async () => {
        console.log('sent');
        console.log(pb.authStore.model?.id)
        await pb.collection("users").update(pb.authStore.model?.id,
            {
                "notification_email": notification_email,
                "notification_push": notification_push,
                "notification_setup": true

            }
        )
        updateSuccess = true;
        //modal.hide()
    };
    export let modal: Modal;
    const markParticipateWithoutNotifications = () => {
        modal.hide()
    };

    onMount(async () => {
        notification_email = pb.authStore.model?.notification_email
        notification_push = pb.authStore.model?.notification_push
    });
</script>

<CfSectionHeading>Benachrichtigungen einrichten</CfSectionHeading>
{#if updateSuccess}
    <div>
        Super, danke dass Du dabei bist!
        <CfButton share={true}/>

        <div class="flex flex-row justify-end pt-8">
            <CfButton rounded onclick={() => {
                setTimeout(() => updateSuccess = false, 500);
                modal.hide();
            }}>Schließen</CfButton>
        </div>

    </div>
{:else}
    <div class="flex flex-col gap-4 py-8">
        <div>
            Wir wollen dich auf dem Laufenden halten, damit du keine #ClimateFriday Aktionen verpasst. Gib
            uns bitte dein Einverständnis. Danke!
        </div>

        <form>
            <fieldset class="flex flex-col gap-4">
                <legend class="sr-only"> Benachrichtigungseinstellungen</legend>
                <div>
                    <label for="push_notifications"> Push Benachrichtigungen aktivieren </label>
                    <input id="push_notifications" type="checkbox" bind:checked={notification_push}/>
                    <label for="push_notifications" class="col-span-2 text-sm text-neutral-900 text-opacity-90">

                    </label>
                </div>
                <div>
                    <label for="email_notification"> E-Mail Benachrichtigungen aktivieren </label>
                    <input id="email_notification" type="checkbox" bind:checked={notification_email}/>
                    <label for="email_notification" class="col-span-2 text-sm text-neutral-900 text-opacity-90">

                    </label>
                </div>
            </fieldset>
        </form>

        <button
                class="text-right text-sm text-neutral-700 underline pt-4"
                on:click={markParticipateWithoutNotifications}>Neee, ach lass mal!
        </button
        >

        <!-- Wie machen wir die Will-ich-nicht-Option? Einfach über Speichern ohne das etwas angekreuzt wurde?
        Ich fände es besser, wenn eine aktive Ablehnung erforderlich ist. Also zwei Tickboxen. Einen
        Togglebutton der auf an voreingestellt ist, ist ja nicht erlaubt, oder? -->
    </div>
    <div class="flex flex-row justify-end pt-8">
        <CfButton rounded onclick={submit}>Speichern</CfButton>
    </div>
{/if}