<script lang="ts">
	import { goto } from '$app/navigation';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';

	let email: string;
	let password: string;

	const login = async () => {
		const authData = await pb.collection('users').authWithPassword(email, password);
		if (await authData.token) {
			goto('/app/');
		}
	};
</script>

<form class="flex flex-col mx-auto">
	<fieldset>
		<legend>Login</legend>

		<div class="form-control w-full max-w-xs">
			<label class="label" for="email">
				<span class="label-text">Email</span>
			</label>
			<input
				type="text"
				placeholder="Email"
				class="input input-bordered w-full max-w-xs"
				id="email"
				bind:value={email}
			/>
		</div>

		<div class="form-control w-full max-w-xs">
			<label class="label" for="password">
				<span class="label-text">Passwort</span>
			</label>
			<input
				placeholder=""
				class="input input-bordered w-full max-w-xs"
				type="password"
				id="password"
				bind:value={password}
			/>
		</div>

		<div class="form-control w-full max-w-xs py-4 gap-4">
			<button
				class="btn btn-primary bg-primary-200 hover:bg-primary-300 text-neutral-900"
				on:click={() => login()}>Absenden</button
			>
			<a class="btn btn-ghost" href="/app/auth/withEmail/signup"> Account anlegen </a>
		</div>
	</fieldset>
</form>

<hr />

<div class="flex flex-col mx-auto max-w-xs w-full py-4 gap-4 px-6">
	<button class="btn btn-primary" on:click={() => login()}>Mit Apple anmelden</button>
	<button class="btn btn-primary" on:click={() => login()}>Mit Google anmelden</button>
</div>
