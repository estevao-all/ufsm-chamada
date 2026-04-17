<script lang="ts">
    import TextInput from "../lib/components/TextInput.svelte";
    import Button from "../lib/components/Button.svelte";

    let username = "";
    let password = "";
    let isLoading = false;

    async function handleLogin() {
        const sanitizedUsername = username.trim();
        const sanitizedPassword = password.trim();

        if (sanitizedUsername === "" || sanitizedUsername === "") {
            return;
        }

        isLoading = true;

        const response = await fetch("/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username: sanitizedUsername, password: sanitizedPassword }),
        }).catch(() => null);

        isLoading = false;
    }

    function handleKeyDown(event: KeyboardEvent) {
        switch (event.key) {
            case "Enter": {
                handleLogin();
                break;
            }
        }
    }
</script>

<div class="main-container">
    <div class="login-container">
        <TextInput
            name="username" autocomplete="username"
            label="Usuário" bind:value={username}
            on:keydown={handleKeyDown}
        />
        <TextInput
            name="password" autocomplete="current-password"
            label="Senha" type="password"
            bind:value={password} on:keydown={handleKeyDown}
        />
        <Button on:click={handleLogin} {isLoading}>Login</Button>
    </div>
</div>

<style>
  .main-container {
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
