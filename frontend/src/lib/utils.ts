 // DD/MM/YYYY HH:MM => YYYY-MM-DDTHH:MM
export function portalTimeToDateInputTime(time: string | undefined): string {
    if (!time) {
        return "";
    }

    const [datePart, timePart] = time.split(" ");
    const [day, month, year] = datePart.split("/");
    return `${year}-${month}-${day}T${timePart}`;
}

// YYYY-MM-DDTHH:MM => DD/MM/YYYY HH:MM
export function dateInputTimeToPortalTime(time: string): string {
    if (!time) {
        return "";
    }

    const [datePart, timePart] = time.split("T");
    const [year, month, day] = datePart.split("-");
    return `${day}/${month}/${year} ${timePart}`;
}
