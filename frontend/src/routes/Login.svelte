<script lang="ts">
    import TextInput from "../lib/components/TextInput.svelte";
    import Button from "../lib/components/Button.svelte";

    let username = $state("");
    let password = $state("");
    let loading = $state(false);

    async function handleLogin(event: SubmitEvent) {
        event.preventDefault();

        const sanitizedUsername = username.trim();
        const sanitizedPassword = password.trim();

        if (sanitizedUsername === "" || sanitizedUsername === "") {
            return;
        }

        loading = true;

        const response = await fetch("/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username: sanitizedUsername, password: sanitizedPassword }),
        }).catch(() => null);

        loading = false;
    }
</script>

<div class="content">
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
        <Button type="submit" {loading}>Login</Button>
    </form>
</div>

<style>
  .content {
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
