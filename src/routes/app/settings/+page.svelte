<script lang="ts">
	import { goto } from '$app/navigation';
	import CfHeading from '$lib/Components/CFHeading.svelte';
	import TopSpacer from '$lib/Components/TopSpacer.svelte';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import { onMount } from 'svelte';
	import CfSectionHeading from '$lib/Components/CFSectionHeading.svelte';
	import Modal from '$lib/Components/Modal.svelte';
	import EditSaveTextField from '$lib/Components/EditSaveTextField.svelte';
	import { browser } from '$app/environment';
	import { Capacitor } from '@capacitor/core';

	let notification_email = false;
	let notification_push = false;

	onMount(async () => {
		notification_email = pb.authStore.model?.notification_email;
		notification_push = pb.authStore.model?.notification_push;
	});

	let updateSuccess = false;

	const saveNotificationSettings = async () => {
		console.log('sent');
		console.log(pb.authStore.model?.id);
		await pb.collection('users').update(pb.authStore.model?.id,
			{
				'notification_email': notification_email,
				'notification_push': notification_push,
				'notification_setup': true

			}
		);
		updateSuccess = true;
		//modal.hide()
	};


	const deleteAccountRequest = () => {

	};

	const changePasswordRequest = async (value: string) => {
		await pb.collection('users').requestPasswordReset(value);
	};

	const deleteAccount = async () => {
		await pb.collection('demo').delete(pb.authStore.model?.id);
		pb.authStore.clear();
		goto('/');
	};

	const savePlz = async (value: string) => {
		let rec = await pb.collection('users').update(pb.authStore.model?.id, { region: value });
		return rec.region;
	};

	const saveUser = async (value: string) => {
		await pb.collection('users').requestEmailChange(value);
		return value;
	};

	const savePass = async (value: string) => {

		return value;
	};

	let userReccord;
	onMount(async () => {
		userReccord = pb.collection('users').getOne(pb.authStore.model?.id);
	});

	let deleteAccountModal: Modal;
	let changePasswordModal: Modal;

</script>

<TopSpacer />

<CfHeading>Einstellungen</CfHeading>
<div>

	<div>
		{#if browser}
			{#await userReccord}
				<span class="loading loading-dots text-primary-500 text-center text-4xl" />

			{:then user}
				<div>
					<form>
						<fieldset class="flex flex-col gap-4">
							<legend class="sr-only"> Benachrichtigungseinstellungen</legend>
							{#if Capacitor.isNativePlatform()}
								<div>
									<label for="push_notifications"> Push Benachrichtigungen aktivieren </label>
									<input id="push_notifications" type="checkbox" bind:checked={notification_push}
												 on:click={saveNotificationSettings} />
									<label for="push_notifications" class="col-span-2 text-sm text-neutral-900 text-opacity-90">

									</label>
								</div>
							{/if}
							<div>
								<label for="email_notification"> E-Mail Benachrichtigungen aktivieren </label>
								<input id="email_notification" type="checkbox" bind:checked={notification_email}
											 on:click={saveNotificationSettings} />
								<label for="email_notification" class="col-span-2 text-sm text-neutral-900 text-opacity-90">

								</label>
							</div>
						</fieldset>
						<!--						<button on:click={saveNotificationSettings}>
													Speichern
												</button>
						-->          </form>
				</div>


				<div>
					<form>
						<fieldset>
						</fieldset>
					</form>
				</div>

				<div>
					<form>
						<EditSaveTextField name="Postleitzahl" id="plz" placeholder="Postleitzahl" value={user?.region}
															 update={savePlz} />
						<EditSaveTextField name="EMail" id="email" placeholder="Email" value={user?.email}
															 update={saveUser} successMessage="Bitte bestätige deine neue Email!" />


						<!--	<EditSaveTextField name="Passwort" id="plz" placeholder="Neues Passwort" isPass value={"********"}
																 update={savePass} /> -->
						<div>
							<button class="bg-neutral-100 rounded-md py-4 px-8 font-bold underline my-4"
											on:click={changePasswordModal.show()}>Passwort ändern
							</button>
						</div>
					</form>
				</div>

				<div class="h-8" />


				<div>
					<button class="text-error-500 font-bold underline my-10" on:click={deleteAccountModal.show()}>Account
						löschen
					</button>
				</div>
			{/await}
		{/if}
	</div>
</div>


<!--
Abmelden confirm dialog
-->


<Modal bind:this={changePasswordModal}>
	<CfSectionHeading>Passwort ändern</CfSectionHeading>
	<div class="flex flex-col gap-4 py-8">
		<div>
			<button class="my-10 underline" on:click={() => changePasswordRequest($currentUser?.email)}>Klicke hier um dein
				Passwort zu ändern
			</button>
		</div>
	</div>
</Modal>


<Modal bind:this={deleteAccountModal}>
	<CfSectionHeading>Account löschen</CfSectionHeading>
	<div class="flex flex-col gap-4 py-8">
		<div>
			Schade, dass du gehst :(
			<button class="my-10 underline text-error-700" on:click={() => deleteAccount()}>Klicke hier um deinen Account zu
				endgültig zu löschen
			</button>
		</div>
	</div>
</Modal>
