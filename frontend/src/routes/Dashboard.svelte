<script lang="ts">
    import { getTeacherDisciplines, getUserInfo } from "../lib/api/user";
    import LogoutIcon from "../lib/components/icons/LogoutIcon.svelte";
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

    const userInfoPromise = guardAuthenticatedRequest(getUserInfo());
    const teacherDisplicinesPromise = guardAuthenticatedRequest(getTeacherDisciplines());
</script>

<div class="content-container">
    <div class="dashboard-container">
        <div class="user-greeting-container">
            {#await userInfoPromise then userInfo}
                <h1>{getGreeting()}, {userInfo.name}</h1>
                <button class="logout-container" onclick={handleLogout}>
                    <LogoutIcon />
                    <h4>Sair</h4>
                </button>

            {/await}
        </div>
        <h2>Suas disciplinas:</h2>
        {#await teacherDisplicinesPromise then teacherDisciplines}
            <div class="disciplines-container">
                {#each teacherDisciplines.disciplines as discipline}
                    <p>{discipline.name}</p>
                {/each}
            </div>
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
        width: min(50vw, 90vw);

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

    .disciplines-container {
        padding: 0.5rem;

        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }
</style>
