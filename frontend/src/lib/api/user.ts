import { evaluateDwrReply, request } from "./utils";
import * as Routes from "./routes";
import { rawScheduleToParsed } from "./discipline_utils";

export async function login(username: string, password: string) {
    return await request({
        method: "POST",
        path: Routes.USER_LOGIN,
        json: { username, password }
    });
}

export interface UserInfo {
    name: string;
}

export async function getUserInfo() {
    return await request<UserInfo>({
        method: "GET",
        path: Routes.USER_INFO
    });
}

export async function getTeacherSchedule() {
    const rawTeacherScheduleDwrReply = await request<string>({
        method: "GET",
        path: Routes.USER_TEACHER_SCHEDULE
    });

    return rawScheduleToParsed(evaluateDwrReply(rawTeacherScheduleDwrReply));
}
