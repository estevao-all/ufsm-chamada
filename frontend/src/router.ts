import { createRouter } from "sv-router";
import Login from "./routes/Login.svelte";
import RedirectUnknownRoute from "./routes/RedirectUnknownRoute.svelte";
import { isLoggedIn } from "./lib/guards";

export const { p, navigate, isActive, route } = createRouter({
    "*": RedirectUnknownRoute,
    "/login": Login,
    "/user": {
        "/dashboard": () => import("./routes/Dashboard.svelte"),
        "/disciplines/:classId": () => import("./routes/DisciplineEditor.svelte"),
        hooks: {
            beforeLoad() {
                if (!isLoggedIn()) {
                    throw navigate("/login", { replace: true });
                }
            }
        }
    }
});
