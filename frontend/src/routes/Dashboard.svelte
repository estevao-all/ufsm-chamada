<script lang="ts">
    import { getDisciplinesFromSchedule } from "../lib/api/discipline_utils";
    import { getTeacherSchedule, getUserInfo } from "../lib/api/user";
    import Button from "../lib/components/Button.svelte";
    import EditSquareIcon from "../lib/components/icons/EditSquareIcon.svelte";
    import LogoutIcon from "../lib/components/icons/LogoutIcon.svelte";
    import TableWrapper from "../lib/components/TableWrapper.svelte";
    import { guardAuthenticatedRequest } from "../lib/guards";
    import { navigate } from "../router";

    function getGreeting() {
        const hour = new Date().getHours();
        if (hour < 12) return "Bom dia";
        if (hour < 18) return "Boa tarde";
        return "Boa noite";
    }

    function handleLogout() {
        document.cookie.split(";").forEach(cookie => {
            const name = cookie.split("=")[0].trim();
            document.cookie = `${name}=; path=/; Max-Age=0`;
        });

        navigate("/login");
    }

    function editDisciplineClass(classId: number) {
        navigate("/user/disciplines/:classId", {
            params: { classId: String(classId) }
        });
    }

    const userInfoPromise = guardAuthenticatedRequest(getUserInfo());
    const teacherSchedulePromise = guardAuthenticatedRequest(getTeacherSchedule());
</script>

<div class="main-container">
    <div class="user-greeting-container">
        {#await userInfoPromise then userInfo}
            <h1>{getGreeting()}, {userInfo.name}</h1>
            <Button variant="ghost" title="Sair" onclick={handleLogout}>
                <LogoutIcon />
                <span>Sair</span>
            </Button>
        {/await}
    </div>
    {#await teacherSchedulePromise then teacherSchedule}
        <h2>Suas disciplinas:</h2>
        <TableWrapper>
            <table>
                <thead>
                    <tr>
                        <th class="column-fit-center">Ações</th>
                        <th>Nome</th>
                        <th>Turma</th>
                    </tr>
                </thead>
                <tbody>
                    {#each getDisciplinesFromSchedule(teacherSchedule) as discipline (discipline.class.id)}
                        <tr>
                            <td class="column-fit-center">
                                <Button
                                    variant="icon"
                                    title="Editar"
                                    onclick={() => editDisciplineClass(discipline.class.id)}
                                >
                                    <EditSquareIcon />
                                </Button>
                            </td>
                            <td>{discipline.name}</td>
                            <td>{discipline.class.name}</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </TableWrapper>
    {/await}
</div>

<style>
    .main-container {
        min-height: 100vh;
        padding: 10vh 0;
        width: min(60rem, 90vw);
        margin: 0 auto;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .user-greeting-container {
        padding-bottom: 1rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
        border-bottom: 2px solid var(--color-primary);
    }
</style>
