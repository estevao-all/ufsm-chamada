<script lang="ts">
    import { getDisciplineClass } from "../lib/api/user";
    import { guardAuthenticatedRequest } from "../lib/guards";
	import { route } from "../router";

	const params = route.getParams("/user/disciplines/:classId");
    const disciplineClassPromise =
        guardAuthenticatedRequest(getDisciplineClass(params.classId));
</script>

{#await disciplineClassPromise then disciplineClass}
    {#each disciplineClass.students as student}
        <p>{student.name}</p>
    {/each}
{/await}
