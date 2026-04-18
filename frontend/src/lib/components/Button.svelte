<script lang="ts">
    export let isDisabled = false;
    export let isLoading = false;
</script>

<button
    disabled={isDisabled || isLoading}
    class:isLoading
    on:click
    type="button"
    {...$$restProps}
>
    {#if isLoading}
        <span class="spinner"></span>
    {:else}
        <slot />
    {/if}
</button>

<style>
    button {
        width: 100%;

        background-color: var(--color-primary);
        color: var(--color-primary-contrast);
        border: 1px solid var(--color-border);
        border-radius: 6px;

        display: inline-flex;
        align-items: center;
        justify-content: center;
        padding: 0.6em 1.2em;
        font-size: 1rem;
        font-weight: bold;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    button:hover:not(:disabled) {
        background-color: var(--color-primary-hover);
        border-color: color-mix(in srgb, var(--color-border), transparent 50%);
    }

    button:disabled {
        filter: grayscale(0.5);
    }

    .spinner {
        width: 16px;
        height: 16px;
        border: 2px solid rgba(255, 255, 255, 0.3);
        border-top-color: var(--color-primary-contrast);
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }
</style>
