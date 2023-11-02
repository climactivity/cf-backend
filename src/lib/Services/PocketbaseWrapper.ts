import PocketBase from "pocketbase"; 
import { writable } from "svelte/store";

const pb = new PocketBase("http://127.0.0.1:8090"); 
// pb.autoCancellation(false); //this messes up image loading


const currentUser = writable(pb.authStore.model)

pb.authStore.onChange(
    () => currentUser.set(pb.authStore.model)
)

export {
    pb, currentUser
}
