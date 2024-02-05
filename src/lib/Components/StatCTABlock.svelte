<script lang="ts">

    import CountUp from '$lib/Components/CountUp.svelte';
    import CfSection from '$lib/Components/CFSection.svelte';
    import { pb } from '$lib/Services/PocketbaseWrapper';

    const FAKE_STATS = false;

    let participantsLastFriday = async () => {
        const _res = await pb.collection("user_count").getOne("usercount");
        return FAKE_STATS ? 1200 : _res.count
    }

    let participantsSameRegion = async () => {
        const _res = await pb.collection("user_count").getOne(pb.authStore.model?.region);
        console.log(_res)
        return FAKE_STATS ? 234 : _res.count
    }

</script>

<CfSection>
    {#await participantsLastFriday() then p}
			<div>Schon registrierte Teilnehmer:innen:<br />
            <CountUp number={p}/>
        </div>
    {/await}
    {#await participantsSameRegion() then p}
        <div>Aus deinem PLZ-Gebiet waren es<br/>
            <CountUp number={p}/>
        </div>
    {/await}
</CfSection>