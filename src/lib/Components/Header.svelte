<script>



</script>



<div class = "topnav">

    
    
    <div class="nav-links">
        <a href="/" class="active">Home</a> 
        <a href="/Men">Men</a>  
        <a href="/Women">Women</a>  
        <a href="/About">About</a>  
    </div>

</div>


<style>

    .topnav{

        background-color:  #333;
        overflow: hidden;
        display: flex;
        justify-content: center;
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


    @media screen and (max-width: 768px) {
        .topnav {
            flex-direction: column; 
            align-items: flex-start; 
        }

    }

</style>