<script lang="ts">
    import { getDisciplinesFromSchedule } from "../lib/api/discipline_utils";
    import { getTeacherSchedule, getUserInfo } from "../lib/api/user";
    import EditSquareIcon from "../lib/components/icons/EditSquare.svelte";
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
        document.cookie.split(';').forEach(cookie => {
            const name = cookie.split('=')[0].trim();
            document.cookie = `${name}=; path=/; Max-Age=0`;
        });

        navigate("/login");
    }

    function editDisciplineClass(classId: number) {
        console.log(`Edit class with id ${classId}`);
    }

    const userInfoPromise = guardAuthenticatedRequest(getUserInfo());
    const teacherSchedulePromise = guardAuthenticatedRequest(getTeacherSchedule());
</script>

<div class="content-container">
    <div class="dashboard-container">
        <div class="user-greeting-container">
            {#await userInfoPromise then userInfo}
                <h1>{getGreeting()}, {userInfo.name}</h1>
                <button title="Sair" class="logout-container" onclick={handleLogout}>
                    <LogoutIcon />
                    <h4>Sair</h4>
                </button>
            {/await}
        </div>
        {#await teacherSchedulePromise then teacherSchedule}
            <h2>Suas disciplinas:</h2>
            <TableWrapper>
                <table>
                    <thead>
                        <tr>
                            <th>Ações</th>
                            <th>Nome</th>
                            <th>Turma</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each getDisciplinesFromSchedule(teacherSchedule) as discipline}
                            <tr>
                                <td>
                                    <button title="Editar" onclick={() => editDisciplineClass(discipline.class.id)}>
                                        <EditSquareIcon />
                                    </button>
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
</div>

<style>
    .content-container {
        min-height: 100vh;
        padding: 10vh 0;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .dashboard-container {
        width: min(60rem, 90vw);
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

    .logout-container {
        display: flex;
        align-items: center;
        gap: 0.1rem;
    }

    .logout-container:hover {
        color: var(--color-primary-hover);
    }
</style>
