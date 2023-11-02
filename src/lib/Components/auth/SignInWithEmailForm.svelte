<script lang="ts">
	import { goto } from '$app/navigation';
	import { currentUser, pb } from '$lib/Services/PocketbaseWrapper';
	import type { ValidationErrors } from '$lib/types/pb';
	import CfButton from '../CFButton.svelte';
	import logo from '$lib/Images/an_logo.png';
	import { fly } from 'svelte/transition';
	let identity: string;
	let password: string;

	let code: number;
	let message: string;
	let errors: ValidationErrors;

	const login = async () => {
		try {
			const authData = await pb.collection('users').authWithPassword(identity, password);
			if (await authData.token) {
				goto('/app/');
			}
		} catch (e) {
			const { data } = e as { data: any };

			code = data.code;
			message = data.message;
			errors = data.data;

			console.log(data);
		}
	};
</script>

<div class="flex flex-col items-center mx-auto pb-8 gap-4">
	<img class="w-24 h-24" alt="ClimateFriday logo" src={logo} />
	<h1 class="font-serif text-center font-bold text-2xl text-neutral-900">Anmelden</h1>
	<span class="font-sans text-center text-neutral-800"
		>Noch nicht dabei? <a
			href="/app/auth/withEmail/signup"
			class="text-primary-500 hover:text-primary-700">Account erstellen</a
		></span
	>
</div>
<form
	class="flex flex-col mx-auto bg-neutral-50 shadow-md px-8 py-4 rounded-md"
	in:fly={{ y: 100 }}
>
	<fieldset>
		{#if message}
			<div class="text-center italic text-error-500 w-[16rem]">
				{message}
			</div>
		{/if}

		<div class="form-control w-[16rem]">
			<label class="label" for="email">
				<span class="label-text">Email</span>
			</label>
			<input
				type="text"
				placeholder="Email"
				class="input input-bordered w-[16rem]"
				id="email"
				bind:value={identity}
			/>
			{#if errors && errors.identity}
				<label class="label" for="email">
					<span class="text-xs text-error-500 italic">*{errors.identity.message}</span>
				</label>
			{/if}
		</div>

		<div class="form-control w-[16rem]">
			<label class="label" for="password">
				<span class="label-text">Passwort</span>
			</label>
			<input
				placeholder=""
				class="input input-bordered w-[16rem]"
				type="password"
				id="password"
				bind:value={password}
			/>
			{#if errors && errors.password}
				<label class="label" for="email">
					<span class="text-xs text-error-500 italic">*{errors.password.message}</span>
				</label>
			{/if}
		</div>

		<div class="form-control w-[16rem] py-4 gap-4">
			<CfButton onclick={() => login()}>Absenden</CfButton>
			<!-- <CfButton
				onclick={(e) => {
					e.preventDefault();
					goto('/app/auth/withEmail/signup');
				}}>Account anlegen</CfButton
			> -->
		</div>
	</fieldset>
	<hr />
	<div class="flex flex-col mx-auto w-[16rem] py-4 gap-4">
		<CfButton secondary onclick={() => goto('/app/auth/withApple/')}>Mit Apple anmelden</CfButton>
		<CfButton secondary onclick={() => goto('/app/auth/withGoogle/')}>Mit Google anmelden</CfButton>
	</div>
</form>
