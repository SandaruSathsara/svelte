<script>
// @ts-nocheck


    import { onMount } from 'svelte';
    
    
    // @ts-ignore
    let wshoes = [];
    
    let errorMessage = '';
    
    onMount(async () => {
        try {
            
            const response = await fetch('http://localhost:8080/womensneakers');
            
            
            if (!response.ok) {
                throw new Error('Failed to fetch data');
            }
            
            
            wshoes = await response.json();
        } catch (error) {
            
            // @ts-ignore
            errorMessage = `Error: ${error.message}`;
            console.error(error);
        }
    });
    </script>
    
    <h1>Women Sneakers</h1>
    
    <div class="wshoes">
        {#if errorMessage}
            
            <p>{errorMessage}</p>
        {:else if wshoes.length > 0}
            
            {#each wshoes as post (post.id)}
                <div class="post">
                    <h2>
                        <span class="name">{post.name}</span>
                        <span class="price">${post.price}</span>
                    </h2>
                    <p>{post.description}</p>
                    <img src={post.image_url} alt=""/>
                    <hr />
                </div>
            {/each}
        {:else}
            
            <p>Loading...</p>
        {/if}
    </div>
    
    <style>
        .wshoes {
            padding-bottom: 50px;
        }

        h1 {

            text-align: center;
            font-style: italic;
            text-decoration: underline;
        }

        .name {
             margin-right: 50px; 
        }

        .price{

            color: greenyellow;
            font-weight: bolder;
        }

       


        
    </style>
    