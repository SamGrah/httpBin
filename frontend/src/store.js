import { writable } from 'svelte/store';

export const subPageToDisplay = writable('create-bin');
export const scrollEnabled = writable(false);