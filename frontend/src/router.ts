import { createRouter } from 'sv-router';
import { isLoggedIn } from './lib/guards';
import Login from "./routes/Login.svelte";

export const { p, navigate, isActive, route } = createRouter({
    '/login':  Login,
    '/user': {
        "/dashboard": () => import('./routes/Dashboard.svelte'),
        hooks: {
            beforeLoad() {
                if (!isLoggedIn()) {
                    throw navigate('/login');
                }
            }
        }
    }
});
