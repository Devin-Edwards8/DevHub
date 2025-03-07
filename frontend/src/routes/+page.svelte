<script>
    import { Client, setContextClient, cacheExchange, fetchExchange, queryStore, getContextClient, gql } from '@urql/svelte';
  
    const client = new Client({
      url: 'http://localhost:8080/query',
      exchanges: [cacheExchange, fetchExchange],
    });
  
    setContextClient(client);

    const tasks = queryStore({
    client: getContextClient(),
    query: gql`
      query {
        tasks {
          id,
          status
        }
      }
    `,
  });
</script>

{#if $tasks.fetching}
<p>Loading...</p>
{:else if $tasks.error}
<p>Oh no... {$tasks.error.message}</p>
{:else}
<ul>
  {#each $tasks.data.tasks as task}
  <li>{task.id}</li>
  <li>{task.status}</li>
  {/each}
</ul>
{/if}
