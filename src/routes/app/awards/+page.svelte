<script>
    import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
    import TopSpacer from '$lib/Components/TopSpacer.svelte';
    import AwardDisplay from '$lib/Components/AwardDisplay.svelte';

    import { participationCount, streakCount } from '$lib/Services/UserState';
</script>

<TopSpacer/>
<!--
<CfSectionHeading>Highlights</CfSectionHeading>
<div class="grid grid-flow-col">
    <AwardDisplay icon="fire" tint="300" title="5er Serie" achieved={true} reason="fünf aufeinander folgende Teilnahmen"/>
    <AwardDisplay icon="medal" tint="300" title="Frischling" achieved={true} reason="insgesamt zwei Teilnahmen" />

</div>
-->
<CfSectionHeading>Awards</CfSectionHeading>
<div>
    <CfSectionHeading>Serien</CfSectionHeading>

    <div class="grid grid-flow-col overflow-x-scroll breakout snap-x">
        <AwardDisplay icon="fire" tint="200" title="2er Serie" achieved={$streakCount >= 2} reason="zwei aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="300" title="5er Serie" achieved={$streakCount >= 5} reason="fünf aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="500" title="10er Serie" achieved={$streakCount >= 10} reason="10 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="700" title="20er Serie" achieved={$streakCount >= 20} reason="20 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="800" title="50er Serie" achieved={$streakCount >= 50} reason="50 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="900" title="100er Serie" achieved={$streakCount >= 100} reason="100 aufeinander folgende Teilnahmen"/>
    </div>

    <CfSectionHeading>Teilnahmen</CfSectionHeading>
    <div class="grid grid-flow-col overflow-x-scroll breakout snap-x">
        <AwardDisplay icon="medal" tint="300" title="Frischling" achieved={$participationCount >= 2} reason="insgesamt zwei Teilnahmen" />
        <AwardDisplay icon="medal" tint="500" title="Talent" achieved={$participationCount >= 5} reason="insgesamt fünf Teilnahmen" />
        <AwardDisplay icon="medal" tint="700" title="Lehrling" achieved={$participationCount >= 10} reason="insgesamt 10 Teilnahmen" />
        <AwardDisplay icon="medal" tint="800" title="Geselle" achieved={$participationCount >= 20} reason="insgesamt 20 Teilnahmen" />
        <AwardDisplay icon="medal" tint="900" title="Meister" achieved={$participationCount >= 50} reason="insgesamt 50 Teilnahmen" />
    </div>
</div>

<style>
    .breakout {
        margin:1em -100%; /* old browsers fallback */
        margin:1em calc(50% - 50vw);
    }
    .breakout::before, .breakout::after {
        content: "";
        width: 1rem;
    }
</style>