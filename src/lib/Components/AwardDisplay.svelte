<script lang="ts">
    import { DateTime } from 'luxon';

    import Fire from '$lib/Components/fire.svelte';
    import Medal from '$lib/Components/medal.svelte';

    import fire_svg from '$lib/Images/fire.svg';
    import medal_svg from '$lib/Images/medal.svg';
    import Modal from '$lib/Components/Modal.svelte';

    export let icon = "medal";
    export let title = "Awa awa";
    export let tint = "500";
    export let date = "2023-11-24 12:32:01.000Z";
    let dateStr = ""
    export let reason;
    export let achieved = false;
    let modal: Modal;

    $:dateStr = DateTime.fromSQL(date).toFormat('dd LLL.', {locale: "de"});

</script>

<button class="h-40 w-40 border-success-{tint} {achieved ? '' : 'border-neutral-500' }  border m-2 rounded-2xl grid grid-flow-row items-center justify-center"
        on:click={() => modal.show()}>
    {#if achieved}
        <div class="text-success-{tint}">
            {#if icon === "fire" }
                <Fire/>
            {:else}
                <Medal/>
            {/if}
        </div>
    {:else}
        {#if icon === "fire" }
            <img src={fire_svg} alt="award">
        {:else}
            <img src={medal_svg} alt="award">
        {/if}
    {/if}
    <span class="text-center">{title}</span>
</button>

<Modal bind:this={modal}>
    <div class="grid grid-flow-row items-center justify-center gap-4">
        {#if achieved}
            <div class="text-success-{tint} mx-auto">
                {#if icon === "fire" }
                    <Fire/>
                {:else}
                    <Medal/>
                {/if}
            </div>
            <div class="text-center text-lg font-bold">
                {title}
            </div>
            <div class="text-center text-lg">
                Verliehen für {reason}
            </div>
        {:else}
            <div class="mx-auto">
                {#if icon === "fire" }
                    <img src={fire_svg} alt="award">
                {:else}
                    <img src={medal_svg} alt="award">
                {/if}
            </div>
            <div class="text-center text-lg font-bold">
                {title}
            </div>
            <div class="text-center text-lg">
                Wird verliehen für {reason}
            </div>
        {/if}

    </div>
</Modal>

<style>
    button {
        -webkit-tap-highlight-color: transparent;
    }
</style>