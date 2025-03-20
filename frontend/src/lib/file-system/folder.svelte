<script lang="ts">
    let { name }: { name: string } = $props();

    let arrowDown = $state(false);

    function toggleArrow() {
        arrowDown = ! arrowDown;
    }
</script>

<style>
    button {
        display: flex;
        flex-direction: row;
        gap: 5px;
        cursor: default;
        border: none;
        background-color: transparent;
    }

    button p {
        font: var(--sk-font-body-small);
    }

    button:hover p {
        color: var(--sk-fg-1);
    }

    button:hover .down-arrow {
        font-weight: bold;
    }

    .down-arrow {
        transform: rotate(90deg);
    }
</style>

<button onclick={() => toggleArrow()}>
    <p class="down-arrow">›</p>
    <p>{name}</p>
</button>