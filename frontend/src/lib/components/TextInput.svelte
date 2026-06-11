<script lang="ts">
    import type { HTMLInputAttributes, HTMLTextareaAttributes } from "svelte/elements";

    interface Props extends HTMLInputAttributes {
        label?: string;
        value?: string;
        multiline?: boolean;
    }

    let {
        label,
        id = crypto.randomUUID(),
        value = $bindable(),
        multiline = false,
        ...restProps
    }: Props = $props();
</script>

<div class="input-container">
    {#if label}
        <label for={id}>{label}</label>
    {/if}

    {#if multiline}
        <textarea {id} bind:value={value} {...restProps as HTMLTextareaAttributes}></textarea>
    {:else}
        <input {id} bind:value={value} {...restProps} />
    {/if}
</div>

<style>
    .input-container {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    label {
        font-size: 1.25rem;
        font-weight: bold;
    }

    input, textarea {
        width: 100%;
        padding: 0.25em;
        font-size: 1.25rem;
        border: 1px solid var(--color-border);
        border-radius: 4px;
    }

    textarea {
        resize: vertical;
        min-height: 6rem;
    }

    input:focus, textarea:focus {
        outline: none;
        border-color: var(--color-border-dark);
        box-shadow: 0 0 4px color-mix(in srgb, var(--color-border-dark), transparent 50%);
    }
</style>
