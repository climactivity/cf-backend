import { Capacitor } from "@capacitor/core";
import PocketBase from "pocketbase"; 
import { writable } from "svelte/store";

let pocketbaseUrl; 

if (Capacitor.isNativePlatform()) {
    pocketbaseUrl = "https://app.climatefriday.de"
} else {
    pocketbaseUrl = import.meta.env.VITE_PB_BASE_URL

}


const pb = new PocketBase(pocketbaseUrl); 
pb.autoCancellation(false); //this messes up image loading

const currentUser = writable(pb.authStore.model)

pb.authStore.onChange(
    () => currentUser.set(pb.authStore.model)
)

export {
    pb, currentUser
}
