<script>
    import { Client, setContextClient, cacheExchange, fetchExchange, queryStore, getContextClient, gql } from '@urql/svelte';
    import EnumCell from './enum-cell.svelte';
    import AddButton from '$lib/add-button.svelte';
  
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
          title,
          status,
          priority
        }
      }
    `,
  });
</script>

<style>
  .tracker-title {
    font: var(--sk-font-h2);
    color: var(--sk-fg-2);
    margin-top: 7rem;
    border: none;
    background-color: transparent;
    width: 60%;
    margin: 7rem 0 3px 3px;
    padding: 2px;
  }

  .tracker-title:focus {
    border: 1px solid rgba(200, 200, 200, 0.5);
    border-radius: 5px;
    outline: none;
    padding: 1px;
  }

  /* Scrollable Table */

  .task-tracker {
    display: flex;
    flex-direction: row;
    border-radius: 5px;
    overflow: hidden;
    border: 1px solid rgba(200, 200, 200, 0.5);
  }

  .scroll-table {
    width: 100%;
    max-width: calc(100% - 30rem);
    overflow-x: scroll;
    overflow-y: visible;
  }

  /* Table Structure */

  td, th {
    height: 4rem;
    min-width: 20rem;
    padding: 0;
  }

  .task-column {
    width: 30rem;
  }

  /* Table Border Style */

  table {
    background-color: var(--sk-bg-0);
    border-collapse: collapse;
    overflow: visible;
  }

  table th {
    border-top: none;
  }

  th:last-child, td:last-child {
    border-right: none;
  }

  tr:last-child td {
    border-bottom: none;
  }

  .task-column td, .task-column th {
    border-right: none;
    border-left: none;
  }

  th, td {
    border: 1px solid rgba(200, 200, 200, 0.5);
  }

  /* Table Fonts */

  th {
    font: var(--sk-font-ui-medium);
  }

  td, input {
    font: var(--sk-font-ui-medium);
  }

  /* Cell Styles */

  input {
    width: 95%;
    background-color: transparent;
    outline: none;
    border: none;
  }

  .task-column input {
    margin-left: 5px;
  }

  .fill-cell {
    width: 100%;
    max-width: 100%;
    min-width: 10rem;
    text-align: left;
    padding-left: 5px;
  }

  .add-row:hover input {
    outline: 1px solid var(--sk-fg-4);
  }
</style>

{#if $tasks.fetching}
<p></p>
{:else if $tasks.error}
<p>Oh no... {$tasks.error.message}</p>
{:else}
<input class="tracker-title" placeholder="My Projects" value="My Projects" maxlength="30"/>
<div class="task-tracker">
    <table class="task-column">
        <thead>
        <tr>
            <th>Task</th>
        </tr>
        </thead>
        <tbody>
        {#each $tasks.data.tasks as task}
        <tr>
            <td><input value={task.title}/></td>
        </tr>
        {/each}
        <tr>
          <td class="add-row"><input placeholder="+ Add task" /></td>
        </tr>
        </tbody>
    </table>
    <div class="scroll-table">
    <table>
        <thead>
            <tr>
                <th>Status</th>
                <th>Priority</th>
                <th class="fill-cell"><AddButton title="Add Column"/></th>
            </tr>
        </thead>
        <tbody>
            {#each $tasks.data.tasks as task}
            <tr>
                <td><EnumCell value={task.status} /></td>
                <td class="priority-cell">{task.priority}</td>
                <td class="fill-cell"></td>
            </tr>
            {/each}
            <tr>
              <td style="border-left: none;"></td>
            </tr>
        </tbody>
    </table>
    </div>
</div>
{/if}
