<script>
    import Footer from '$lib/Components/Footer.svelte';
    import Header from '$lib/Components/Header.svelte';

    import { Client, cacheExchange, fetchExchange,setContextClient } from '@urql/svelte';

const client = new Client({
url: 'http://localhost:8080/query',
exchanges: [cacheExchange, fetchExchange],



});

setContextClient(client);

</script>

<Header />

<main>
    
<slot></slot>
</main>
<Footer />