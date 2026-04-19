const API_BASE_URL = "/api";

interface RequestOptions {
    method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
    path: string;
    json?: Record<string, any>;
}

interface APIErrorResponse {
    error: string;
}

export class APIError extends Error {
    constructor(public message: string, public status: number) {
        super(message);
    }
}

export async function request<T = null>(options: RequestOptions) {
    const { method, path, json } = options;

    const response = await fetch(API_BASE_URL + path, {
        method,
        headers: { "Content-Type": "application/json" },
        body: json != null ? JSON.stringify(json) : undefined
    });

    if (!response.ok) {
        const error: APIErrorResponse = await response.json().catch(() => ({ error: "Unknown Error" }));
        throw new APIError(error.error ?? "Unknown Error", response.status);
    }

    const text = await response.text();
    return (text ? JSON.parse(text) : null) as T;
}
