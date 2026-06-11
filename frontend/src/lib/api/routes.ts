const SAFE_CHARACTERS = new Set([":", "?", "@"]);
function encode(strings: TemplateStringsArray, ...args: Array<string | number>) {
    return strings.reduce((acc, str, i) => {
        acc += str;

        if (args[i] != null) {
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

export const API_BASE_URL = "/api";
export const USER_LOGIN = "/user/login";
export const USER_INFO = "/user/info";

export const USER_TEACHER_SCHEDULE = "/user/teacher-schedule";
export const USER_DISCIPLINE_CLASS = (classId: string) => encode`/user/disciplines/${classId}`;
