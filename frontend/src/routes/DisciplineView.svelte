<script lang="ts">
    import { failure, success } from "../lib/api/svelte/toasts.svelte";
    import { getDisciplineClass, saveLesson } from "../lib/api/user";
    import Button from "../lib/components/Button.svelte";
    import DateInput from "../lib/components/DateInput.svelte";
    import BackIcon from "../lib/components/icons/BackIcon.svelte";
    import Select from "../lib/components/Select.svelte";
    import TableWrapper from "../lib/components/TableWrapper.svelte";
    import TextInput from "../lib/components/TextInput.svelte";
    import { guardAuthenticatedRequest } from "../lib/guards";
    import { dateInputTimeToPortalTime, portalTimeToDateInputTime } from "../lib/utils";
    import { navigate, route } from "../router";

    const params = route.getParams("/user/disciplines/:classId");
    let disciplineId: string;

    let lessonStartTime = $state("");
    let lessonHourAmount = $state("2");
    let lessonType = $state("2");
    let lessonIsRemoteLession = $state(false);
    let lessionIsCoil = $state(false);
    let lessonNoteText = $state("");

    let allPresencesChecked = $state(false);
    let studentsPresences = $state<Record<string, boolean>>({});

    function navigateToDashboard() {
        navigate("/user/dashboard");
    }

    function toggleAllPresences() {
        allPresencesChecked = !allPresencesChecked;
        for (const studentId in studentsPresences) {
            studentsPresences[studentId] = allPresencesChecked;
        }
    }

    async function handleSaveLesson() {
        const sanitizedLessonNoteText = lessonNoteText.trim();

        if (sanitizedLessonNoteText === "") {
            failure("As notas da aula não podem estar vazias");
            return;
        }

        try {
            await saveLesson(params.classId, {
                disciplineId: disciplineId,
                startTime: dateInputTimeToPortalTime(lessonStartTime),
                hourAmount: lessonHourAmount,
                type: lessonType,
                noteText: sanitizedLessonNoteText,
                remoteLesson: lessonIsRemoteLession,
                coil: lessionIsCoil,
                studentPresences: Object.entries(studentsPresences).map(([studentId, status]) => ({
                    studentId,
                    status
                }))
            });

            success("Aula adicionada com sucesso");
        } catch (err) {
            failure("Houve um erro ao adicionar a aula");
            throw err;
        }
    }

    const disciplineClassPromise = guardAuthenticatedRequest(getDisciplineClass(params.classId));
    disciplineClassPromise.then(disciplineClass => {
        disciplineId = disciplineClass.disciplineId;
        lessonStartTime = portalTimeToDateInputTime(disciplineClass.defaultLessonStartTime);

        studentsPresences = Object.fromEntries(
            disciplineClass.students.map(student => [student.id, false])
        );
    });
</script>

<div class="main-container">
    <div class="discipline-header">
        {#await disciplineClassPromise}
            <h2>Carregando...</h2>
        {:then disciplineClass}
            <div class="discipline-info">
                <h2>{disciplineClass.disciplineName}</h2>
                <h3>{disciplineClass.className}</h3>
            </div>
        {/await}

        <Button variant="ghost" title="Voltar" onclick={navigateToDashboard}>
            <BackIcon />
            <span>Voltar</span>
        </Button>
    </div>

    <h2>Dados da aula</h2>
    <div class="lesson-data-container">
        <DateInput label="Início" bind:value={lessonStartTime} />

        <Select label="Horas de aula" bind:value={lessonHourAmount}>
            <option value="1">1</option>
            <option value="2">2</option>
            <option value="3">3</option>
            <option value="4">4</option>
            <option value="5">5</option>
            <option value="6">6</option>
            <option value="7">7</option>
            <option value="8">8</option>
        </Select>

        <Select label="Tipo de aula" bind:value={lessonType}>
            <option value="1">Aula Prática</option>
            <option value="2">Aula Teórica</option>
            <option value="3">Prova</option>
            <option value="4">Revisão</option>
            <option value="5">Exame</option>
            <option value="6">Avaliação</option>
            <option value="7">Reavaliação</option>
            <option value="8">Teórica-Prática</option>
            <option value="9">Atividades Complementares</option>
            <option value="10">Extensionista (Prática)</option>
            <option value="11">Extensionista (Teórica)</option>
        </Select>

        <label class="lesson-data-checkbox">
            <span>Modalidade EAD</span>
            <input type="checkbox" bind:checked={lessonIsRemoteLession} />
        </label>

        <label class="lesson-data-checkbox">
            <span>COIL</span>
            <input type="checkbox" bind:checked={lessionIsCoil} />
        </label>
    </div>

    <TextInput label="Notas da aula" multiline bind:value={lessonNoteText}/>

    {#await disciplineClassPromise then disciplineClass}
        <div class="presences-header">
            <h2>Presenças</h2>
            <label class="toggle-all-presences">
                <input
                    type="checkbox"
                    checked={allPresencesChecked}
                    onchange={toggleAllPresences}
                />
                <span>{allPresencesChecked ? "Desmarcar presença para todos" : "Marcar presença para todos"}</span>
            </label>
        </div>

        <TableWrapper>
            <table>
                <thead>
                    <tr>
                        <th class="column-fit-center">Presente</th>
                        <th class="column-fit">Matrícula</th>
                        <th>Nome</th>
                    </tr>
                </thead>
                <tbody>
                    {#each disciplineClass.students as student (student.id)}
                        <tr>
                            <td class="column-fit-center">
                                <input
                                    type="checkbox"
                                    bind:checked={studentsPresences[student.id]}
                                />
                            </td>
                            <td>{student.enrollmentId}</td>
                            <td>{student.name}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </TableWrapper>

        <div class="actions">
            <Button variant="primary" onclick={handleSaveLesson}>Salvar</Button>
        </div>
    {/await}
</div>

<style>
    .main-container {
        min-height: 100vh;
        padding: 10vh 0;
        width: min(80rem, 90vw);
        margin: 0 auto;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .discipline-header {
        padding-bottom: 1rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 2px solid var(--color-primary);
    }

    .discipline-info {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .lesson-data-container {
        display: flex;
        flex-wrap: wrap;
        gap: 1.5rem;
    }

    .lesson-data-checkbox {
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 0.5rem;
        cursor: pointer;
    }

    .lesson-data-checkbox span {
        font-size: 1rem;
        font-weight: 600;
    }

    .presences-header {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .toggle-all-presences {
        display: flex;
        align-items: center;
        gap: 0.25rem;
        user-select: none;
    }

    :global(input[type="checkbox"]) {
        width: 1rem;
        height: 1rem;
        cursor: pointer;
    }

    .actions {
        max-width: 8rem;
    }
</style>
