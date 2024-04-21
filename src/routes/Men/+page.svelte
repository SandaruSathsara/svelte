<script>
  import { addItemToCart } from '../Cart/cart'; 
  
  let items = [
    { id: 1, name: 'Nike Air Jordan 1 Low', image: '../src/assets/Men Shoe.jpeg', count: 1 },
    { id: 2, name: 'Nike Dunk Low Retro', image: '../src/assets/Nike Dunk Low Retro.png', count: 1 }
  ];
  
  // @ts-ignore
  function addToCart(item) {
    addItemToCart(item);
    alert(`${item.name} is successfully added to the cart!`);
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

<h1>Men Sneakers</h1>

<div class="container">
  {#each items as item (item.id)}
    <div class="card">
      <img src="{item.image}" alt="{item.name}" />
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
  flex-wrap: wrap;
  justify-content: center;
  align-items: flex-start;
  padding: 20px;
}

.card {
  flex: 0 0 calc(50% - 40px);
  margin: 20px;
  text-align: center; 
}

.card img {
  width: 40%;
  height: auto;
}

.name {
  margin: 10px 0;
  font-weight: bold; 
  font-size: 18px;
}

.counter {
  margin-bottom: 10px;
}

.counter button {
  font-size: 18px;
  padding: 5px 10px;
  cursor: pointer;
}

.counter span {
  font-size: 18px;
  margin: 0 10px;
}

.cart-button {
  font-size: 18px;
  padding: 10px 20px;
  cursor: pointer;
  background-color: black;
  color: white;
  border-radius: 10px;
  border: none;
  transition: background-color 0.3s ease;
}

.cart-button:hover {
  background-color: greenyellow;
}

.cart-button:active {
  background-color: greenyellow;
}

h1{

  text-align: center;
  font-style: italic;
  text-decoration: underline;
}

@media screen and (max-width: 768px) {
  .card {
    flex-basis: calc(50% - 20px);
  }
}

</style>
