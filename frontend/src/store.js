import { writable } from 'svelte/store';

export const scrollEnabled = writable(false);
export const binRequestHistory = writable([])