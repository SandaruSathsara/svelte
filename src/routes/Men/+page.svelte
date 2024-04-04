<script>

  let count = 0;

  function increaseCount() {
    count++;
  }

  function decreaseCount() {
    if (count > 0) {
      count--;
    }
  }

  function cart() {

    console.log("Item is successfully added to the cart!");
  }

</script>




<div class="container">
    
    <div class="card">
        <img src="../src/assets/Men Shoe.jpeg" alt="Nike Air Jordan 1 Low" style="width: 300px; height: 300px;" />
      </div>
  
    
    <div class="counter">
      <button on:click={decreaseCount}>-</button>
      <span>{count}</span>
      <button on:click={increaseCount}>+</button>
    </div>
  
    
    <button class="cart-button" on:click={cart}>Add to Cart</button>
  </div>

  <style>

.container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh; 
  }

  .card {
    
    margin-bottom: 20px;
  }

  .counter {
    
    margin-bottom: 20px;
  }

  .counter button {
    font-size: 20px;
    padding: 5px 10px;
    cursor: pointer;
  }

  .counter span {
    font-size: 20px;
    margin: 0 10px;
  }

  .cart-button {
    font-size: 20px;
    padding: 10px 20px;
    cursor: pointer;
    background-color: blue; 
    color: white; 
    border-radius: 10px; 
    border: none; 
    transition: background-color 0.3s ease; 
  }

  .cart-button:hover {
    background-color: #007bff; 
  }

  .cart-button:active {
    background-color: #0056b3; 
  }

  </style>