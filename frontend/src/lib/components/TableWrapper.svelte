<script lang="ts">
    import type { Snippet } from "svelte";

    let { children }: { children: Snippet } = $props();
</script>

<div class="table-wrapper">
    {@render children()}
</div>

<style>
    .table-wrapper {
        overflow-x: auto;
    }

    .table-wrapper :global(table) {
        width: 100%;
        border-collapse: collapse;
    }

    .table-wrapper :global(th),
    .table-wrapper :global(td) {
        padding: 0.5rem 1rem 0.5rem 0.25rem;
    }

    .table-wrapper :global(th) {
        text-align: left;
        border-bottom: 1px solid var(--color-border);
    }

    .table-wrapper :global(td) {
        font-size: 0.9rem;
        color: var(--color-text-main);
    }

    .table-wrapper :global(tbody tr:nth-child(odd) td) {
        background-color: color-mix(in srgb, var(--color-primary) 10%, var(--color-background));
    }

    .table-wrapper :global(tbody tr:nth-child(even) td) {
        background-color: color-mix(in srgb, var(--color-primary) 3%, var(--color-background));
    }

    .table-wrapper :global(tbody tr:hover td) {
        background-color: color-mix(in srgb, var(--color-primary) 18%, var(--color-background));
    }

    .table-wrapper :global(.column-fit) {
        width: 1%;
        white-space: nowrap;
    }

    .table-wrapper :global(.column-fit-center) {
        width: 1%;
        white-space: nowrap;
        text-align: center;
    }
</style>
