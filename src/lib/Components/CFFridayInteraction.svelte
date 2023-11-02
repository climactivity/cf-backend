<script lang="ts">
	import { isFriday } from "$lib/Services/TimerStore";
	import CfSectionHeading from "./CFSectionHeading.svelte";
	import FridayCountdown from "./FridayCountdown.svelte";
	import NotificationSettings from "./NotificationSettings.svelte";
	import ParticipateForm from "./ParticipateForm.svelte";


</script>

{#if isFriday()}
    <div>
        <CfSectionHeading>Mitmachen</CfSectionHeading>
        <ParticipateForm/>
    </div>
{:else}
    <NotificationSettings/>
{/if}