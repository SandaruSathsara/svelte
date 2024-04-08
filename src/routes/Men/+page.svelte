<script>
  
  import { addItemToCart } from '../Cart/cart'; 

  
  let items = [
    { id: 1, name: 'Nike Air Jordan 1 Low', image: '../src/assets/Men Shoe.jpeg', count: 1 },
    { id: 2, name: 'Nike Dunk Low Retro', image: '../src/assets/Nike Dunk Low Retro.png', count: 1 }
  ];

  
  // @ts-ignore
  function addToCart(item) {
    addItemToCart(item);
  }

  
  // @ts-ignore
  function increaseCount(item) {
    item.count++;
    items = [...items];
  }

  
  // @ts-ignore
  function decreaseCount(item) {
    if (item.count > 1) { 
      item.count--;
      items = [...items];
    }
  }
</script>

<div class="container">
  {#each items as item (item.id)}
    <div class="card">
      <img src="{item.image}" alt="{item.name}" style="width: 300px; height: 300px;" />
      <p class="name">{item.name}</p>
  
      <div class="counter">
        <button on:click={() => decreaseCount(item)}>-</button>
        <span>{item.count}</span>
        <button on:click={() => increaseCount(item)}>+</button>
      </div>
  
      <button class="cart-button" on:click={() => addToCart(item)}>Add to Cart</button>
    </div>
  {/each}
</div>



<style>
.container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  height: 100vh;
  
}

.card {
  width: 50%;
  margin-bottom: 20px;
  text-align: center; 
  margin-left: 50px;
  margin-right: 50px;
}

.card p {
  margin-top: 10px;
}

.name {
  margin-top: 10px; 
  font-weight: bold; 
  font-size: 24px;
  font-family: sans-serif;
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
