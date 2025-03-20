<script>
    import Tracker from "$lib/tracker/tracker.svelte";
</script>

<style>
  #page-title {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  p {
    font: var(--sk-font-ui-small);
    color: var(--sk-fg-4);
  }

  input {
    font: var(--sk-font-h1);
    color: var(--sk-fg-1);
    border: none;
    background-color: transparent;
    width: auto;
  }
</style>

<div id="page-title">
    <p>TASK TRACKING</p>
    <input type="text" placeholder="Tasks" value="Tasks" />
</div>

<Tracker />