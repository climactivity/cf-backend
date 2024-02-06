<script>
    import CfHeading from '$lib/Components/CFHeading.svelte';
    import TopSpacer from '$lib/Components/TopSpacer.svelte';
    import { page } from '$app/stores';
    import { pb } from '$lib/Services/PocketbaseWrapper';
    import ClimateFriday from '$lib/Components/ClimateFriday.svelte';
    import CFSectionHeading from '$lib/Components/CFSectionHeading.svelte';
    import { onMount } from 'svelte';

    let fragment = ""
    let myRegion


    $: fragment = $page.url.hash
    onMount(async () => {
        let rec = await pb.collection("users").getOne(pb.authStore.model?.id);
        myRegion = rec.region;
    })
</script>
<TopSpacer/>
<CfHeading>Ergebnisse</CfHeading>

<div class="my-4 font-bold text-lg ">

    <a href="#code">Meine PLZ</a>

    <span class="text-neutral-500" href="#state">Umgebung</span>
    <span class="text-neutral-500" href="#all">Alle</span>

</div>

{#if fragment === "#state"}
    <div>
        <table>
            <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Bundesland</th>
                <th scope="col">Teilnahmen</th>

            </tr>
            </thead>
            <tbody>
            <tr>
                <th scope="row">1</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">2</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">3</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">4</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">5</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">6</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">7</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">8</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">9</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">10</th>
                <td>28231</td>
                <td>100</td>
            </tr>
            </tbody>
        </table>
    </div>
{:else if fragment === "#all"}
    <div>
        <table>
            <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">PLZ</th>
                <th scope="col">Teilnahmen</th>

            </tr>
            </thead>
            <tbody>
            <tr>
                <th scope="row">1</th>
                <td>Bremen</td>
                <td>100</td>
            </tr>

            <tr>
                <th scope="row">2</th>
                <td>Hamburg</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">3</th>
                <td>Niedersachsen</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">4</th>
                <td>Nordrhein-Westfalen</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">5</th>
                <td>Rheinland-Pfalz</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">6</th>
                <td>Saarland</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">7</th>
                <td>Baden-Württemberg</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">8</th>
                <td>Mecklemburg-Vorpommern</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">9</th>
                <td>Berlin</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">10</th>
                <td>Brandenburg</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">11</th>
                <td>Sachsen-Anhalt</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">12</th>
                <td>Sachsen</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">13</th>
                <td>Thüringen</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">14</th>
                <td>Hessen</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">15</th>
                <td>Bayern</td>
                <td>100</td>
            </tr>
            <tr>
                <th scope="row">16</th>
                <td>Schleswig-Holstein</td>
                <td>100</td>
            </tr>
            </tbody>
        </table>
    </div>
{:else}
    <div>
        <CFSectionHeading>Meine PLZ</CFSectionHeading>
        {#if myRegion}
            {#await pb.collection('user_count').getOne(myRegion) then count}
               In deiner Postleitzahl ({myRegion}) haben {count.count} Menschen beim <ClimateFriday/> mitgemacht.
            {/await}
            {:else}
            Bitte gib eine Postleitzahl ein um zu sehen wievieke Menschen aus deiner Umgebung beim <ClimateFriday/> mitgemacht haben.
        {/if}
    </div>
{/if}
<span class="text-sm text-neutral-700 my-10 italic">Diese Seite ist noch sehr unfertig, hier könnt ihr später Details über die, hoffentlich, bremen- und dann bundesweite Nutzung der App finden.</span>
<br />
