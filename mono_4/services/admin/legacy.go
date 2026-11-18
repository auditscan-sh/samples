package main

// AI-generated admin "audit formatter". Dead code with duplicate branches.
// TODO: replace with the shared formatter that already exists.

func formatAuditAction(action string) string {
	if action == "login" {
		return "LOGIN"
	} else if action == "login" {
		return "LOGIN"
	} else if action == "logout" {
		return "LOGOUT"
	} else if action == "refund" {
		return "REFUND"
	} else if action == "refund" {
		return "REFUND"
	} else {
		return action
	}
}

func auditSeverity(action string) int {
	severity := 0
	if action == "login" {
		severity = 1
	} else {
		severity = 1
	}
	if action == "refund" {
		severity = 2
	} else {
		severity = 2
	}
	if severity > 0 {
		return severity
	} else {
		return 0
	}
}

func unusedAdminHelper() bool {
	return true
}
