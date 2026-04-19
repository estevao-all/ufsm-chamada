import { request } from "./utils";
import * as Routes from "./routes";

export async function userLogin(username: string, password: string) {
    return await request({
        method: "POST",
        path: Routes.USER_LOGIN,
        json: { username, password }
    });
}
