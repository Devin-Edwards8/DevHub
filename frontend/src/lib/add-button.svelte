<script>
    let {title} = $props();
</script>

<style>
  .add-button {
    width: 2.5rem;
    border-radius: 3px;
    text-align: center;
    cursor: pointer;
  }

  .add-button:hover {
    background-color: var(--sk-bg-4);
  }

  .add-button:hover p {
    color: var(--sk-fg-1);
  }
</style>

<div class="add-button" title={title}>
    <p>+</p>
</div>