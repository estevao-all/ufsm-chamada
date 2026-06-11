<script lang="ts">
    import type { HTMLInputAttributes } from "svelte/elements";

    interface Props extends HTMLInputAttributes {
        label?: string;
        value?: string;
    }

    let {
        label,
        id = crypto.randomUUID(),
        value = $bindable(),
        type = "datetime-local",
        ...restProps
    }: Props = $props();
</script>

<div class="field">
    {#if label}
        <label for={id}>{label}</label>
    {/if}
    <input {id} {type} bind:value {...restProps} />
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

    input {
        padding: 0.25em 0.5em;
        font-size: 1rem;
        border: 1px solid var(--color-border);
        border-radius: 4px;
        background-color: var(--color-background);
        color: var(--color-text-main);
    }

    input:focus {
        outline: none;
        border-color: var(--color-border-dark);
        box-shadow: 0 0 4px color-mix(in srgb, var(--color-border-dark), transparent 50%);
    }
</style>
