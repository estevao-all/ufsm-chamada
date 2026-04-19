<script lang="ts">
    import { getUserInfo } from '../lib/api/user';
    import { guardAuthenticatedRequest } from "../lib/guards";

    function getGreeting() {
        const hour = new Date().getHours();
        if (hour < 12) return "Bom dia";
        if (hour < 18) return "Boa tarde";
        return "Boa noite";
    }

    const userInfoPromise = guardAuthenticatedRequest(getUserInfo());
</script>

{#await userInfoPromise}
    <h1>Loading...</h1>
{:then userInfo}
    <div class="content">
        <h1>{getGreeting()}, {userInfo.name}</h1>
    </div>
{:catch}
    <h1>Failed to load user info</h1>
{/await}
