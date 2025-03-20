<script>
    let {name} = $props();
</script>

<style>
    div {
        display: flex;
        flex-direction: row;
        gap: 5px;
    }

    p {
        margin-left: 1.5rem;
    }

    p, a {
        font: var(--sk-font-body-small);
        color: var(--sk-fg-2);
        text-decoration: none;
    }

    a:hover {
        text-decoration: underline;
    }
</style>

<div>
    <p>∟</p>
    <a href="/notes">{name}</a>
</div>
