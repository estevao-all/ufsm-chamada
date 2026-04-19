export function isLoggedIn() {
    return document.cookie.split(";").some(cookie => cookie.trimStart().startsWith("JSESSIONIDSSO="));
}
