<script>
    let description = '';
    let showText = false; 

    // @ts-ignore
    function handleSubmit(event) {
        event.preventDefault();
    }

    // @ts-ignore
    function handleInput(event) {
        
        description = event.target.value.replace(/[^a-zA-Z]/g, '');
    }

    function toggleTextVisibility() {
        showText = !showText; 
    }
</script>

<h1>Form</h1>

<form method="POST" on:submit={handleSubmit}>
    <label>
        Type your name:
        <input
            name="description"
            autocomplete="off"
            required
            bind:value={description}
            on:input={handleInput}
        />
    </label>
</form>

<button on:click={toggleTextVisibility}>Submit</button> 

{#if showText}
    <p>{description}</p> 
{/if}
