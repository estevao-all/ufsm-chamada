import { request } from "./utils";
import * as Routes from "./routes";

export async function login(username: string, password: string) {
    return await request({
        method: "POST",
        path: Routes.USER_LOGIN,
        json: { username, password }
    });
}

interface UserInfo {
    name: string;
}

export async function getUserInfo() {
    return await request<UserInfo>({
        method: "GET",
        path: Routes.USER_INFO
    });
}
