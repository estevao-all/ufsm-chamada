<script lang="ts">
    import TextInput from "../lib/components/TextInput.svelte";
    import Button from "../lib/components/Button.svelte";
    import { login } from "../lib/api/user";
    import { APIError } from "../lib/api/utils";
    import { navigate } from "../router";

    let username = $state("");
    let password = $state("");
    let loading = $state(false);

    let invalidCredentials = $state(false);
    let invalidCredentialsTimeoutId: number | undefined;

    async function handleLogin(event: SubmitEvent) {
        event.preventDefault();

        invalidCredentials = false;
        clearTimeout(invalidCredentialsTimeoutId);

        const sanitizedUsername = username.trim();
        const sanitizedPassword = password.trim();

        if (sanitizedUsername === "" || sanitizedPassword === "") {
            return;
        }

        loading = true;

        try {
            await login(sanitizedUsername, sanitizedPassword);
            navigate("/user/dashboard");
        } catch (error) {
            if (error instanceof APIError) {
                if (error.status !== 401) {
                    return;
                }

                invalidCredentials = true;
                invalidCredentialsTimeoutId = setTimeout(() => invalidCredentials = false, 1500);
            }
        } finally {
            loading = false;
        }
    }
</script>

<div class="content-container">
    <form class="login-container" onsubmit={handleLogin}>
        <TextInput
            name="username" autocomplete="username"
            label="Usuário"
            bind:value={username}
        />
        <TextInput
            name="password" autocomplete="current-password"
            type="password" label="Senha"
            bind:value={password}
        />
        <Button type="submit" {loading} error={invalidCredentials}>Login</Button>
    </form>
</div>

<style>
  .content-container {
    min-height: 100vh;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .login-container {
    width: min(25rem, 90vw);

    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
</style>
