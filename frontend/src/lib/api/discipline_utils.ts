export enum RawWeekDay {
    Sunday =  1,
    Monday = 2,
    Tuesday = 3,
    Wednesday = 4,
    Thursday = 5,
    Friday = 6,
    Saturday = 7,
}

export interface RawTeacherSchedule {
    ano: number;
    horarios: Record<"segunda" | "terca" | "quarta" | "quinta" | "sexta" | "sabado", RawTeacherDayLesson[]>;
}

export interface RawTeacherDayLesson {
    id: number;
    tipo: string;
    diaSemana: RawWeekDay;
    disciplinas: string[];
    fim: Date;
    inicio: Date;
    turmas: RawTeacherClass[]
}

export interface RawTeacherClass {
    id: number;
    nome: string;
}

export enum WeekDay {
    Sunday,
    Monday,
    Tuesday,
    Wednesday,
    Thursday,
    Friday,
    Saturday,
}

export interface TeacherSchedule {
    year: number;
    weekDaysLessons: Record<WeekDay, TeacherDayLesson[]>;
}

export interface TeacherDayLesson {
    disciplineId: number;
    disciplinesNames: string[];
    type: string;
    weekDay: WeekDay,
    startTime: Date;
    endTime: Date;
    classes: TeacherClass[]
}

export interface TeacherClass {
    id: number;
    name: string;
}

function rawDayLessonToParsed(rawDayLesson: RawTeacherDayLesson): TeacherDayLesson {
    return {
        disciplineId: rawDayLesson.id,
        disciplinesNames: rawDayLesson.disciplinas,
        type: rawDayLesson.tipo,
        weekDay: rawDayLesson.diaSemana - 1,
        startTime: rawDayLesson.inicio,
        endTime: rawDayLesson.fim,
        classes: rawDayLesson.turmas.map(rawTeacherClass => ({
            id: rawTeacherClass.id,
            name: rawTeacherClass.nome,
        })),
    }
}

export function rawScheduleToParsed(rawSchedule: RawTeacherSchedule): TeacherSchedule {
    const { ano, horarios } = rawSchedule;
    return {
        year: ano,
        weekDaysLessons: {
            [WeekDay.Sunday]: [],
            [WeekDay.Monday]: horarios.segunda.map(rawDayLessonToParsed),
            [WeekDay.Tuesday]: horarios.terca.map(rawDayLessonToParsed),
            [WeekDay.Wednesday]: horarios.quarta.map(rawDayLessonToParsed),
            [WeekDay.Thursday]: horarios.quinta.map(rawDayLessonToParsed),
            [WeekDay.Friday]: horarios.sexta.map(rawDayLessonToParsed),
            [WeekDay.Saturday]: horarios.sabado.map(rawDayLessonToParsed),
        },
    }
}

export interface TeacherDiscipline {
    name: string;
    class: TeacherClass;
}

export function getDisciplinesFromSchedule(teacherYearSchedule: TeacherSchedule): TeacherDiscipline[] {
    const disciplines: TeacherDiscipline[] = [];

    for (const weekDayLessons of Object.values(teacherYearSchedule.weekDaysLessons)) {
        for (const lesson of weekDayLessons) {
            if (disciplines.some(discipline => discipline.name === lesson.disciplinesNames[0]
                && discipline.class.id === lesson.classes[0].id)) {
                continue;
            }

            disciplines.push({
                name: lesson.disciplinesNames[0],
                class: lesson.classes[0],
            });
        }
    }

    return disciplines;
}
