<script>

import { page } from "$app/stores";

</script>


<div class = "topnav">

    <div class="nav-links">
        <a href="/" class:active={$page.url.pathname === '/'} >Home</a> 
        <a href="/Men" class:active={$page.url.pathname === '/Men'} >Men</a>  
        <a href="/Women" class:active={$page.url.pathname === '/Women'} >Women</a>  
        <a href="/About" class:active={$page.url.pathname === '/About'} >About</a>  
        

    </div>

    <a href="/Cart" class="cart-icon">
        <img src="../src/assets/cart.png" alt="" />
    </a>

</div>


<style>

    .topnav{

        background-color:  #333;
        overflow: hidden;
        display: flex;
        justify-content: space-between; 
        align-items: center; 
        padding: 0 20px; 
        
    }

    .nav-links {
        display: flex;
        justify-content: center; 
        flex-grow: 1; 
    }

    .topnav a {

        float: left;
        display: block;
        color: #f2f2f2;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 20px;
    }

    .topnav a:hover{

        background-color: #ddd;
        color: black;
    }

    .topnav a.active{

        background-color: greenyellow;
        color: white;
    }

    .cart-icon img {
        width: 24px; 
        height: 24px;
        margin-right: 16px; 
        text-align: center;
        
    }


    @media screen and (max-width: 768px) {
        .topnav {
            flex-direction: column; 
            align-items: flex-start; 
        }

    }

    

</style>