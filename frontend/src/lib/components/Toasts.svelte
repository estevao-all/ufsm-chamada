<script lang="ts">
    import { getToasts } from "../api/svelte/toasts.svelte";

    const toasts = $derived(getToasts());
</script>

<div class="toasts-container">
    {#each toasts as toast (toast.id)}
        <div class="toast toast-{toast.type}">
            {toast.message}
        </div>
    {/each}
</div>

<style>
    .toasts-container {
        position: fixed;
        top: 1.5rem;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        z-index: 1000;
    }

    .toast {
        padding: 0.75rem 1.25rem;
        border-radius: 6px;
        font-size: 0.95rem;
        font-weight: 500;
        color: var(--color-primary-contrast);
        box-shadow: 0 4px 12px color-mix(in srgb, black 20%, transparent);
        animation: slide-in 0.2s ease;
    }

    @keyframes slide-in {
        from {
            opacity: 0;
            transform: translateY(-1rem);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .toast-success {
        background-color: var(--color-success);
    }

    .toast-failure {
        background-color: var(--color-error);
    }

    .toast-info {
        background-color: var(--color-primary);
    }
</style>
