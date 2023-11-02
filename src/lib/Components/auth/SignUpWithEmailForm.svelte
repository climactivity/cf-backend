<script lang="ts">
	import { goto } from '$app/navigation';
	import { pb } from '$lib/Services/PocketbaseWrapper';

	let errors: { [key: string]: any } = {};

	let email: string;
	let password: string;
	let passwordConfirm: string;
	let privacy_policy: boolean;

	const signup = async () => {
		if (password !== passwordConfirm) {
			return;
		}
		if (!privacy_policy) {
			return;
		}
		const result = await pb.collection('users').create({
			email,
			password,
			passwordConfirm
		});
		if (!result) {
			errors.form = result;
		}
		const authData = await pb.collection('users').authWithPassword(email, password);
		if (await authData.token) {
			goto('/app/');
		}
	};
</script>

<form class="flex flex-col mx-auto">
	<label>
		Email
		<input type="text" id="email" bind:value={email} />
	</label>
	<label>
		Passwort
		<input type="password" id="password" bind:value={password} />
	</label>
	<label>
		Passwort bestätigen
		<input type="password" id="passwordConfirm" bind:value={passwordConfirm} />
	</label>
	<label>
		Privacy thing
		<input type="checkbox" id="privacy_policy" bind:checked={privacy_policy} />
	</label>
	<button on:click={() => signup()}>Absenden</button>
</form>

<style lang="scss">
	input {
		border: 1px solid black;
		display: flex;
	}
</style>
