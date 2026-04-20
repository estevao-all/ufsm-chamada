<script lang="ts">
    import { isLoggedIn } from "../lib/guards";
    import { navigate } from "../router";

    if (isLoggedIn()) {
        navigate("/user/dashboard");
    } else {
        navigate("/login");
    }
</script>
