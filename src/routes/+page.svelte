<script>

import {Router, Link, Route } from 'svelte-routing';    
let flipped = false;

</script>


<h1>Welcome </h1>
<p>Visit <a href="https://www.nike.com/">Nike,Inc.</a></p>

<div class="container">
    <h1>New Arrivals</h1>
    <button
        class="card"
        class:flipped={flipped}
        on:click={() => flipped = !flipped}
    >
        <div class="front">
            <!-- svelte-ignore a11y-img-redundant-alt -->
            <img src="../src/assets/Front.png" alt="Front Image" style="max-width: 100%; max-height: 100%;"/>
        </div>
        <div class="back">
            <!-- svelte-ignore a11y-img-redundant-alt -->
            <img src="../src/assets/Full.jpeg" alt="Back Image" style="max-width: 100%; max-height: 100%;"/>
        </div>
    </button>
</div>

<style>

.container {
		display: flex;
		flex-direction: column;
		gap: 1em;
		height: 100%;
		align-items: center;
		justify-content: center;
		perspective: 100vh;
	}

	.card {
		position: relative;
		aspect-ratio: 2.5 / 3.5;
		font-size: min(1vh, 0.25rem);
		height: 80em;
		background: var(--bg-1);
		border-radius: 2em;
		transform: rotateY(180deg);
		transition: transform 0.4s;
		transform-style: preserve-3d;
		padding: 0;
		user-select: none;
		cursor: pointer;
	}

	.card.flipped {
		transform: rotateY(0);
	}

	.front, .back {
		display: flex;
		align-items: center;
		justify-content: center;
		position: absolute;
		width: 100%;
		height: 100%;
		left: 0;
		top: 0;
		backface-visibility: hidden;
		border-radius: 2em;
		border: 1px solid var(--fg-2);
		box-sizing: border-box;
		padding: 2em;
        
	}

	.front {
		background: url(./svelte-logo.svg) no-repeat 5em 5em, url(./svelte-logo.svg) no-repeat calc(100% - 5em) calc(100% - 5em);
		background-size: 8em 8em, 8em 8em;
	}

	.back {
		transform: rotateY(180deg);
	}

	

</style>





