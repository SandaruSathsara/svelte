
import { writable } from 'svelte/store';
import { browser } from '$app/environment';

export const cartItems = writable(browser && localStorage.cartItems ? JSON.parse(localStorage.cartItems) : []);

cartItems.subscribe(value => {
    if (browser) {
        localStorage.cartItems = JSON.stringify(value);
    }
});

// @ts-ignore
export function addItemToCart(item) {
  // @ts-ignore
  cartItems.update(items => [...items, item]); 
}
