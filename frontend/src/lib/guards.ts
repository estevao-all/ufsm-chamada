export function isLoggedIn(): boolean {
    return document.cookie.split(';').some(cookie => cookie.trimStart().startsWith('JSESSIONIDSSO='));
}
