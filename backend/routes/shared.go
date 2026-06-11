package routes

var UFSM_PORTAL_HOST = "portal.ufsm.br"
var UFSM_PORTAL_BASE_URL = "https://portal.ufsm.br"
var UFSM_PORTAL_TYPE = "docente"
var UFSM_PORTAL_DWR_PLAINCALL_BASE_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/dwr/call/plaincall"

var UFSM_PORTAL_LOGIN_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/j_security_check"
var UFSM_PORTAL_INDEX_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/index.html"
var UFSM_PORTAL_MAIN_MENU_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/mainMenu.html"

var UFSM_PORTAL_SCHEDULE_URL = UFSM_PORTAL_DWR_PLAINCALL_BASE_URL + "/gradeHorariosAjaxService.horarios.dwr"
var UFSM_PORTAL_CLASSES_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/turma/turma.html?action=list"

var UFSM_PORTAL_CLASS_VIEW_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/turma/turma.html?action=view&id="
var UFSM_PORTAL_CLASS_FORM_URL = UFSM_PORTAL_BASE_URL + "/" + UFSM_PORTAL_TYPE + "/diario/form.html?turma="
