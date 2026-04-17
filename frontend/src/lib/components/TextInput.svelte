<script lang="ts">
    import type { FullAutoFill } from "svelte/elements";

    export let label: string | undefined = undefined;
    export let name: string | undefined = undefined;
    export let autocomplete: FullAutoFill | undefined = undefined;
    export let type: "text" | "password" = "text";
    export let value = "";

    let id: string = crypto.randomUUID();
</script>

<div class="input-container">
    {#if label}
        <label for={id}>{label}</label>
    {/if}

    <input
        {id}
        {name}
        {autocomplete}
        {type}
        on:keydown
        bind:value={value}
    />
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
