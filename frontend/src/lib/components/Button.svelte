<script lang="ts">
    import type { HTMLButtonAttributes } from 'svelte/elements';

    interface Props extends HTMLButtonAttributes {
        loading?: boolean;
    }

    let { disabled = false, loading = false, children, ...restProps }: Props = $props();
</script>

<button
    disabled={disabled || loading}
    data-loading={loading}
    {...restProps}
>
    <span class="spinner"></span>
    <span class="content">
        {@render children?.()}
    </span>
</button>

<style>
    button {
        visibility: visible;
        position: relative;

        width: 100%;

        background-color: var(--color-primary);
        color: var(--color-primary-contrast);
        border: 1px solid var(--color-border);
        border-radius: 6px;

        display: inline-flex;
        align-items: center;
        justify-content: center;
        padding: 0.6em 1.2em;
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
        visibility: hidden;
        position: absolute;

        width: 1em;
        height: 1em;

        border: 2px solid color-mix(in srgb, var(--color-primary-contrast), transparent 50%);
        border-top-color: var(--color-primary-contrast);
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }

    button[data-loading="true"] .spinner {
        visibility: visible;
    }

    button[data-loading="true"] .content {
        visibility: hidden;
    }
</style>
