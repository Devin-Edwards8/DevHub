<script>
    const { value } = $props();
    const options = ["Not started", "Working on it", "Stuck", "Done"];
</script>

<style>
    .enum-fill {
        position: relative;
        width: 100%;
        height: 100%;
        background-color: var(--sk-bg-accent);
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
    }

    .enum-fill:hover {
        background-color: color-mix(in srgb, var(--sk-bg-accent) 90%, black 10%);
    }

    p {
        color: var(--sk-fg-1);
        font: var(--sk-font-ui-medium)
    }

    /* .enum-selector {
        position: absolute;
        display: flex;
        flex-direction: column;

        top: 100%;
        z-index: 99;
        background-color: var(--sk-bg-1);
    } */
</style>


<div class="enum-fill">
    <p>{value}</p>

    <!-- <div class="enum-selector">
        {#each options as option}
        <div class="enum-option">
            <p>{option}</p>
        </div>
        {/each}
    </div> -->
</div>
