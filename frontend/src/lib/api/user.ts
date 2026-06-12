import { evaluateDwrReply, request } from "./utils";
import * as Routes from "./routes";
import { rawScheduleToParsed } from "./discipline_utils";

export async function login(username: string, password: string) {
    return request({
        method: "POST",
        path: Routes.USER_LOGIN,
        json: { username, password }
    });
}

export interface UserInfo {
    name: string;
}

export async function getUserInfo() {
    return request<UserInfo>({
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

export interface Student {
    id: string;
    name: string;
    enrollmentId: string;
}

export interface DisciplineClass {
    disciplineId: string;
    disciplineName: string;
    className: string;
    defaultLessonStartTime: string;
    students: Student[];
}

export async function getDisciplineClass(classId: string) {
    return request<DisciplineClass>({
        method: "GET",
        path: Routes.USER_DISCIPLINE_CLASS(classId)
    });
}

export interface StudentPresence {
    studentId: string;
    status: boolean;
}

export interface SaveLessonRequest {
    disciplineId: string;
    startTime: string;
    hourAmount: string;
    type: string;
    noteText: string;
    remoteLesson: boolean;
    coil: boolean;
    studentPresences: StudentPresence[];
}

export async function saveLesson(classId: string, saveLessonRequest: SaveLessonRequest) {
    return request({
        method: "POST",
        path: Routes.USER_SAVE_LESSON(classId),
        json: saveLessonRequest
    });
}
