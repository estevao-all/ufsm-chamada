import { navigate } from "../router";
import { APIError } from "./api/utils";

export function isLoggedIn() {
    return document.cookie.split(";").some(cookie => cookie.trimStart().startsWith("JSESSIONIDSSO="));
}

export async function guardAuthenticatedRequest<T>(requestPromise: Promise<T>) {
    return requestPromise.catch(error => {
        if (error instanceof APIError && error.status === 401) {
            navigate("/login");
        }

        throw error;
    });
}
