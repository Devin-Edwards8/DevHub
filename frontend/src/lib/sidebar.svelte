<script>
    import Folder from "./file-system/folder.svelte";
    import NoteLink from "./file-system/note-link.svelte";
    import AddButton from "./add-button.svelte";
</script>

<style>
    nav {
        position: absolute;
        top: var(--header-height);
        left: 0;
        width: calc(var(--sb-width) - var(--sb-padding-left));
        background-color: var(--sk-bg-0);
        height: calc(100vh - var(--header-height) - var(--sb-padding-top));
        padding-left: var(--sb-padding-left);
        padding-top: var(--sb-padding-top);
    }

    .sidebar-title {
        margin-top: 30px;
        width: calc(var(--sb-width) - var(--sb-padding-left) - 1rem);
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }

    ul {
        list-style-type: none;
        padding: 0;
        margin: 0;
    }

    a {
        font: var(--sk-font-body-small);
        color: var(--sk-fg-2);
        text-decoration: none;
    }

    a:hover {
        text-decoration: underline;
    }
</style>

<nav>
    <div class="sidebar-title">
        <h3>Task Tracking</h3>
        <AddButton title="Add a Tracker" />
    </div>
    <ul>
        <li><a href="/">Work</a></li>
        <li><a href="/">Personal</a></li>
    </ul>


    <div class="sidebar-title">
        <h3>Note Taking</h3>
        <AddButton title="Add a Note" />
    </div>
    <ul>
        <li><Folder name="Folder 1" /></li>
        <li><NoteLink name="Note 1"/></li>
        <li><Folder name="Folder 2" /></li>
        <li><NoteLink name="Note 1"/></li>
        <li><NoteLink name="Note 2"/></li>
    </ul>
</nav>