<script>
  import { queryStore, gql, getContextClient } from "@urql/svelte";

  const GET_USERS = gql`
    query {
      mensneakers {
        id
        name
        price
        description
      }
    }
  `;
  // Fetch data using the queryStore
  const todos = queryStore({
    client: getContextClient(),
    query: GET_USERS,
  });
</script>

{#if $todos.fetching}
  <p>Loading...</p>
{:else if $todos.error}
  <p>Oh no... {$todos.error?.message}</p>
{:else}
   <pre>{JSON.stringify($todos.data, null, 2)}</pre>
{/if}
