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

    .table-wrapper :global(th) {
        text-align: left;
        padding: 0.5rem 1rem 0.5rem 0.25rem;
        border-bottom: 1px solid var(--color-border);
    }

    .table-wrapper :global(td) {
        padding: 0.5rem 1rem 0.5rem 0;
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

    .table-wrapper :global(td button) {
        all: unset;
        cursor: pointer;
        display: inline-flex;
        align-items: center;
        color: var(--color-text);
    }

    .table-wrapper :global(td button:hover) {
        color: var(--color-primary-hover);
    }

    .table-wrapper :global(td:first-child) {
        text-align: center;
    }

</style>
