export type ToastType = "success" | "failure" | "info";

export interface Toast {
    id: string;
    message: string;
    type: ToastType;
}

let toasts = $state<Toast[]>([]);

export function getToasts() {
    return toasts;
}

export function toast(message: string, type: ToastType = "info", duration = 3000) {
    const id = crypto.randomUUID();
    toasts.push({ id, message, type });

    setTimeout(() => {
        toasts = toasts.filter(t => t.id !== id);
    }, duration);
}

export const success = (message: string) => toast(message, "success");
export const failure = (message: string) => toast(message, "failure");
