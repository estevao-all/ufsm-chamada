import { API_BASE_URL } from "./routes";

export interface RequestOptions {
    method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
    path: string;
    json?: Record<string, any>;
}

export interface APIErrorResponse {
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

    if (response.headers.get("Content-Type")?.includes("application/json")) {
        return (text ? JSON.parse(text) : null) as T;
    }

    return text as T;
}

const DWR_SECURITY_PREFIX = "throw 'allowScriptTagRemoting is false.';";
let latestDwrReply: any = null;

window.dwr = {
    _: [{
        engine: {
            remote: {
                handleCallback: function (_: string, __: string, reply: any) {
                    latestDwrReply = reply;
                }
            }
        }
    }]
}

export function evaluateDwrReply<T = any>(dwrReply: string) {
    eval(dwrReply.replace(DWR_SECURITY_PREFIX, "") + "\n//# sourceURL=latestDwrReply.js");
    return latestDwrReply as T;
}
