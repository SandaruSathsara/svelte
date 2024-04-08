<script>

import { page } from "$app/stores";

</script>


<div class = "topnav">

    <a href="/">
        <img src="../src/assets/nike.png" alt="" class="logo-icon" />
      </a>


    <div class="nav-links">
        <a href="/" class:active={$page.url.pathname === '/'} >Home</a> 
        <a href="/Men" class:active={$page.url.pathname === '/Men'} >Men</a>  
        <a href="/Women" class:active={$page.url.pathname === '/Women'} >Women</a>  
        <a href="/About" class:active={$page.url.pathname === '/About'} >About</a>  
    </div>


    <div class="search-container">
        <input type="text" placeholder="Search..." class="search-bar">
        <img src="../src/assets/search.png" alt="Search" class="search-icon">
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
        width: 25px; 
        height: 25px;
        margin-right: 16px; 
        text-align: center;
        
    }

    .logo-icon{
        width: 50px; 
        height: 20px;
        margin-right: 16px; 
        text-align: center;
        
    }

    .search-container {
    position: relative;
  }

  .search-bar {
    padding-left: 40px;
    border-radius: 20px; 
    height: 25px; 
    border: 1px solid #ccc; 
    font-size: 16px; 
  }

  .search-icon {
    width: 20px;
    height: 20px;
    position: absolute;
    left: 8px;
    top: 50%;
    transform: translateY(-50%);
    cursor: pointer;
  }

  .search-bar:focus {
    outline: none;
  }
  
    @media screen and (max-width: 768px) {
        .topnav {
            flex-direction: column; 
            align-items: flex-start; 
        }

        .nav-links {
            flex-direction: column;
            align-items: center;
        }

        .topnav a {
            padding: 10px;
        }

        .cart-icon {
            margin-top: 10px;
        }
    }
    

</style>