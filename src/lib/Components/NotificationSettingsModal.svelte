<script>
	import CfButton from './CFButton.svelte';
	import CfSectionHeading from './CFSectionHeading.svelte';

	const submit = () => {
		console.log('sent');
	};

	const markParticipateWithoutNotifications = () => {};
</script>

<CfSectionHeading>Bevor es richtig los Geht: Benachrichtigungen einrichten</CfSectionHeading>
<div class="flex flex-col gap-4 py-8">
	<div>
		Wir wollen dich auf dem Laufenden halten, damit du keine #ClimateFriday Aktionen verpasst. Gib
		uns bitte dein Einverständnis. Danke!
	</div>

	<form>
		<fieldset class="flex flex-col gap-4">
			<legend class="sr-only"> Benachrichtigungseinstellungen </legend>
			<div>
				<label for="push_notifications"> Push Benachrichtigungen aktivieren </label>
				<input id="push_notifications" type="checkbox" />
				<label for="push_notifications" class="col-span-2 text-sm text-neutral-900 text-opacity-90">
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Vitae reprehenderit quia iure.
				</label>
			</div>
			<div>
				<label for="email_notification"> E-Mail Benachrichtigungen aktivieren </label>
				<input id="email_notification" type="checkbox" />
				<label for="email_notification" class="col-span-2 text-sm text-neutral-900 text-opacity-90">
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Vitae reprehenderit quia iure.
				</label>
			</div>
		</fieldset>
	</form>

	<button
		class="text-right text-sm text-neutral-700 underline pt-4"
		on:click={markParticipateWithoutNotifications}>Neee, ach lass mal!</button
	>

	<!-- Wie machen wir die Will-ich-nicht-Option? Einfach über Speichern ohne das etwas angekreuzt wurde?
	Ich fände es besser, wenn eine aktive Ablehnung erforderlich ist. Also zwei Tickboxen. Einen
	Togglebutton der auf an voreingestellt ist, ist ja nicht erlaubt, oder? -->
</div>
<div class="flex flex-row justify-end pt-8">
	<CfButton rounded onclick={submit}>Speichern</CfButton>
</div>
