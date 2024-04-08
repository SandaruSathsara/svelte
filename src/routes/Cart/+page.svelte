
<script>
  import { cartItems } from './cart.js'; 
</script>

<div class="cart">
  <h1>Cart</h1>
  {#each $cartItems as item}
    <div class="cart-item">
      <img src="{item.image}" alt="{item.name}" style="width: 100px; height: 100px;" />
      <p>{item.name}</p>
      <p>Count: {item.count}</p>
    </div>
  {/each}
</div>
