import { writable } from "svelte/store";

export const store = writable<State>({ contacts: [] });
