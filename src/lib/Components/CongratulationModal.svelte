<script>
	import { goto } from "$app/navigation";
	import { participationCount } from "$lib/Services/UserState";
	import CfButton from "./CFButton.svelte";
	import ClimateFriday from "./ClimateFriday.svelte";

</script>


<div class="flex flex-col gap-8">
    <div class="flex flex-col">
        <span class="text-center font-serif text-2xl font-semibold w-full">Super</span>
        <div>
            <div>
                Danke, dass du diese Woche mit deinem Banner mitgemacht hast.
            </div>
            <div>
                Es war deine {$participationCount}. Teilnahme und dein aktueller Streak ist 6 Wochen lang. Super
            </div>
            <div>
                Herzlichen Glückwunsch! Mit jetzt {$participationCount} Teilnahmen bist du Klimaschutz-Lehrling
            </div>
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