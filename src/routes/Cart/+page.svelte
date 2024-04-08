<script>
  import { cartItems } from './cart.js'; 

  // @ts-ignore
  function removeItem(item) {

    // @ts-ignore
    cartItems.update(items => items.filter(i => i !== item));
    
  }
  
</script>

<div class="container">
  <h1>Cart</h1>
  <div class="cart">
      {#each $cartItems as item}
          <div class="cart-item">
              <div class="item-details">
                  <img src="{item.image}" alt="{item.name}" class="item-image" />
                  <div class="text">
                      <p class="item-name">{item.name}</p>
                      <p class="item-count">Count: {item.count}</p>
                  </div>
                  <button class="delete-button" on:click={() => removeItem(item)}>Delete</button>
              </div>
          </div>
      {/each}
  </div>
</div>

<style>
  
.container {
  padding-bottom: 100px; 
}

h1 {
  margin-top: 20px; 
  margin-left: 50px;
}

.cart {
  margin-bottom: 50px; 
}

.cart-item {
  margin-bottom: 20px;
  margin-left: 50px;
}

.item-details {
  display: flex;
  align-items: center;
}

.item-image {
  width: 180px; 
  height: 180px;
  margin-right: 20px; 
}

.cart-item img {
  margin-right: 20px; 
}

.text {
  flex-grow: 1; 
}

.text p {
  margin: 0;
}

.item-name {
  font-size: 24px; 
  font-weight: bold;
  margin-bottom: 10px; 
}

.item-count {
  font-size: 20px; 
  margin: 0;
}

.delete-button {
  font-size: 16px;
  padding: 10px 10px;
  background-color: red;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin-left: auto;
}

</style>


