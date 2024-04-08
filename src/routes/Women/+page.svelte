<script>
// @ts-nocheck


    import { onMount } from 'svelte';
    
    // @ts-ignore
    let posts = [];

    onMount (async () => {

        const response = await fetch ('https://api.escuelajs.co/api/v1/products')
        posts = await response.json()
    
    })

  </script>


  
<h1>Women</h1>

<div class="posts">
    {#each posts as post (post.id) }
    <h2>{post.title}. $ {post.price}</h2>
    <p>{post.description}</p>
    <img src={post.images} alt=""/>
    <hr />
    {:else}
    <p>Loading...</p>
    {/each}
</div>

<style>
    .posts {
        padding-bottom: 50px;
    }
</style>



