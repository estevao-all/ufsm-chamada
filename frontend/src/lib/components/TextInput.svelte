<script lang="ts">
    import type { HTMLInputAttributes } from "svelte/elements";

    interface Props extends HTMLInputAttributes {
        label?: string;
        value?: string;
    }

    let { label, id = crypto.randomUUID(), value = $bindable(), ...restProps }: Props = $props();
</script>

<div class="input-container">
    {#if label}
        <label for={id}>{label}</label>
    {/if}

    <input {id} bind:value={value} {...restProps} />
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

    input {
        width: 100%;
        padding: 0.25em;

        font-size: 1.25rem;
        border: 1px solid var(--color-border);
        border-radius: 4px;
    }

    input:focus {
        outline: none;
        border-color: var(--color-border-dark);
        box-shadow: 0 0 4px color-mix(in srgb, var(--color-border-dark), transparent 50%);
    }
</style>
