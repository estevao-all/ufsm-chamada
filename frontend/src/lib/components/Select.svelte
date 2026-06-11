<script lang="ts">
    import type { HTMLSelectAttributes } from "svelte/elements";

    interface Props extends HTMLSelectAttributes {
        label?: string;
        value?: string;
    }

    let {
        label,
        id = crypto.randomUUID(),
        value = $bindable(),
        children,
        ...restProps
    }: Props = $props();
</script>

<div class="field">
    {#if label}
        <label for={id}>{label}</label>
    {/if}
    <select {id} bind:value {...restProps}>
        {@render children?.()}
    </select>
</div>

<style>
    .field {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    label {
        font-size: 1rem;
        font-weight: 600;
    }

    select {
        padding: 0.25em 0.5em;
        font-size: 1rem;
        border: 1px solid var(--color-border);
        border-radius: 4px;
        background-color: var(--color-background);
        color: var(--color-text-main);
        cursor: pointer;
    }

    select:focus {
        outline: none;
        border-color: var(--color-border-dark);
        box-shadow: 0 0 4px color-mix(in srgb, var(--color-border-dark), transparent 50%);
    }
</style>
