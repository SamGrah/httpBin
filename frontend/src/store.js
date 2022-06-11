import { writable } from 'svelte/store';

export const mainComponentToDisplay = writable('Intro');
export const scrollEnabled = writable(false);