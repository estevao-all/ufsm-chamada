const SAFE_CHARACTERS = new Set([":", "?", "@"]);
function encode(strings: TemplateStringsArray, ...args: Array<string | number>) {
    return strings.reduce((acc, str, i) => {
        acc += str;

        if (args[i] !== undefined && args[i] !== null) {
            acc += Array.from(
                String(args[i]),
                char => SAFE_CHARACTERS.has(char)
                    ? char
                    : (decodeURIComponent(char) === char
                        ? encodeURIComponent(char)
                        : char)
            ).join("");
        }

        return acc;
    }, "");
}

export const USER_LOGIN = "/user/login";
export const USER_INFO = "/user/info";
