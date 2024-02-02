<script lang="ts">
	import { goto } from "$app/navigation";
  import { participationCount, streakCount } from '$lib/Services/UserState';
	import CfButton from "./CFButton.svelte";
	import ClimateFriday from "./ClimateFriday.svelte";
  import { onMount } from 'svelte';
  import { type Award, getNewAward } from '$lib/Services/awards';


  let newAwardP: Award = getNewAward($participationCount, "p");
  let newAwardS: Award = getNewAward($streakCount, "s");

  onMount(async () => {
    console.log($participationCount, $streakCount);
  });

</script>


<div class="flex flex-col gap-8">
    <div class="flex flex-col">
        <span class="text-center font-serif text-2xl font-semibold w-full">Super</span>
        <div>
            <div>
                Danke, dass du diese Woche mit deinem Banner mitgemacht hast.
            </div>
            <div>
                Es war deine {$participationCount}. Teilnahme und dein aktueller Streak ist {$streakCount} Wochen lang. Super!
            </div>
            {#if newAwardP}
            <div>
                Herzlichen Glückwunsch! Mit jetzt {$participationCount} Teilnahmen bist du {newAwardP.title}
            </div>
            {/if}
            {#if newAwardS}
                <div>
                    Herzlichen Glückwunsch! Mit jetzt {$streakCount} Teilnahmen bist du {newAwardS.title}
                </div>
            {/if}
        </div>
    </div>
    <div class="flex flex-col">
        <span class="text-center font-serif text-2xl font-semibold w-full">Teilen</span>
        <div>
            Lass andere wissen, dass Du beim <ClimateFriday/> mitgemacht hast
        </div>
    </div>
    <div class="flex flex-row w-full justify-end">
        <CfButton rounded onclick={() => {
            goto("/app/history")
        }}>Zu meinem Verlauf</CfButton>    
    </div>
</div>