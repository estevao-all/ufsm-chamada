<script lang="ts">
    import type { HTMLButtonAttributes } from "svelte/elements";

    interface Props extends HTMLButtonAttributes {
        loading?: boolean;
        error?: boolean;
        variant?: "primary" | "ghost" | "icon";
        size?: "sm" | "md" | "lg";
    }

    let {
        variant = "primary",
        size = "md",
        disabled = false,
        loading = false,
        error = false,
        children,
        ...restProps
    }: Props = $props();
</script>

<button
    disabled={disabled || loading}
    data-variant={variant}
    data-size={size}
    data-loading={loading}
    data-error={error}
    {...restProps}
>
    <span class="spinner"></span>
    <span class="content">
        {@render children?.()}
    </span>
</button>

<style>
    .content {
        display: inline-flex;
        align-items: center;
        gap: 0.25rem;
    }

    button {
        position: relative;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        font-weight: bold;
        cursor: pointer;
        border-radius: 6px;
        border: 1px solid transparent;
        transition: background-color 0.15s ease, border-color 0.15s ease, color 0.15s ease, filter 0.2s ease;
    }

    button[data-size="sm"] { padding: 0.3em 0.75em; font-size: 0.85rem; }
    button[data-size="md"] { padding: 0.6em 1.2em; font-size: 1rem; }
    button[data-size="lg"] { padding: 0.75em 1.5em; font-size: 1.15rem; }

    button[data-variant="primary"] {
        width: 100%;
        background-color: var(--color-primary);
        border-color: var(--color-border);
        color: var(--color-primary-contrast);
    }
    button[data-variant="primary"]:hover:not(:disabled) {
        background-color: var(--color-primary-hover);
    }

    /* Ghost */
    button[data-variant="ghost"] {
        background-color: transparent;
        border-color: transparent;
        color: inherit;
    }
    button[data-variant="ghost"]:hover:not(:disabled) {
        color: var(--color-primary-hover);
    }

    /* Icon */
    button[data-variant="icon"] {
        background-color: transparent;
        border-color: transparent;
        color: inherit;
        padding: 0.25em;
        border-radius: 4px;
    }
    button[data-variant="icon"]:hover:not(:disabled) {
        color: var(--color-primary-hover);
    }

    /* Error */
    button[data-error="true"] {
        background-color: var(--color-error);
        border-color: var(--color-error);
        color: var(--color-primary-contrast);
    }
    button[data-error="true"]:hover:not(:disabled) {
        background-color: var(--color-error-hover);
    }

    button:disabled {
        filter: grayscale(0.5);
    }

    .spinner {
        visibility: hidden;
        position: absolute;
        width: 1em;
        height: 1em;
        border: 2px solid color-mix(in srgb, currentColor, transparent 60%);
        border-top-color: currentColor;
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
