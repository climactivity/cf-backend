<script>
    import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
    import TopSpacer from '$lib/Components/TopSpacer.svelte';
    import AwardDisplay from "$lib/Components/AwardDisplay.svelte";
</script>

<TopSpacer/>
<CfSectionHeading>Highlights</CfSectionHeading>
<div class="grid grid-flow-col">
    <AwardDisplay icon="fire" tint="300" title="5er Serie" achieved={true} reason="fünf aufeinander folgende Teilnahmen"/>
    <AwardDisplay icon="medal" tint="300" title="Frischling" achieved={true} reason="insgesamt zwei Teilnahmen" />

</div>
<CfSectionHeading>Awards</CfSectionHeading>
<div>
    <div class="grid grid-flow-col overflow-x-scroll breakout snap-x">
        <AwardDisplay icon="fire" tint="200" title="2er Serie" achieved={true} reason="zwei aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="300" title="5er Serie" achieved={true} reason="fünf aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="500" title="10er Serie" achieved={false} reason="10 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="700" title="20er Serie" achieved={false} reason="20 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="800" title="50er Serie" achieved={false} reason="50 aufeinander folgende Teilnahmen"/>
        <AwardDisplay icon="fire" tint="900" title="100er Serie" achieved={false} reason="100 aufeinander folgende Teilnahmen"/>
    </div>

    <div class="grid grid-flow-col overflow-x-scroll breakout snap-x">
        <AwardDisplay icon="medal" tint="300" title="Frischling" achieved={true} reason="insgesamt zwei Teilnahmen" />
        <AwardDisplay icon="medal" tint="500" title="Talent" achieved={false} reason="insgesamt fünf Teilnahmen" />
        <AwardDisplay icon="medal" tint="700" title="Lehrling" achieved={false} reason="insgesamt 10 Teilnahmen" />
        <AwardDisplay icon="medal" tint="800" title="Geselle" achieved={false} reason="insgesamt 20 Teilnahmen" />
        <AwardDisplay icon="medal" tint="900" title="Meister" achieved={false} reason="insgesamt 50 Teilnahmen" />
    </div>
</div>

<div class="italic text-neutral-300">
    *komplett fake, lol
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