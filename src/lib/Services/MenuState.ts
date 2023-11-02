import { writable } from "svelte/store";

export const menuOpen = writable(false);

menuOpen.subscribe(v => console.log("Menu open:", v))