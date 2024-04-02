<script>
    import { page } from "$app/stores";
    
 export let data;
    let windowWidth = window.innerWidth;
    let clicked = false;

    function handleResize(){
        windowWidth = window.innerWidth;
        if(windowWidth >= 1024)clicked = false;
    }

    const handleClick = ()=> clicked = !clicked;
    const resetClick = ()=> clicked = false;

</script>

<svelte:window on:resize={handleResize} />
<nav class={clicked? 'navbar-clicked' : 'navbar'}>
    <div class="menu-icon">
        <div class="icon-wrapper">
          <button on:click={handleClick}>
            <i class={clicked? 'fa fa-times' : 'fas fa-bars'} />
          </button>
        </div>

        {#if data.logoLink}
        <a class="logo-link-text" href={data.linkUrl}><img class='logo' src={data.logoSrc} alt={data.altText}/>{data.optionalLinkText? data.optionalLinkText : ''} </a>
      
        {/if}
        </div>
        </nav>







<style>   

button{
      background-color: rgb(255, 255, 255,0);
    }
    .navbar {   
        position: fixed;
        top: 0;
        width: 100%;  
        height: 5rem; 
        background-color: #000;   
        color: #FFF;   z-index: 999;  
        display: flex;   
        flex-direction: row;
        justify-content: space-around; 
        align-items: center; 
        transition: all 0.2s ease-out;
    } 

    navbar-clicked{
        position: fixed;
        top: 0;
        width: 100%;  
        background-color: #000;   
        color: #FFF;   z-index: 999;  
        display: flex;   
        flex-direction: row;
        justify-content: space-around; 
        align-items: center; 
        transition: all 0.2s ease-out;
        height:5em;
    }  
    
   

</style>
