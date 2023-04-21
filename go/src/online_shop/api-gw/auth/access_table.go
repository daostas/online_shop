package auth

var accessTable = map[string][]string{
	"POST/auth/register/user":    {"ROLE_CLIENT", "ANONYMOUS"},
	"POST/auth/register/admin":   {"ROLE_MAIN_ADMIN"},
	"POST/auth/login/user":       {"ROLE_CLIENT", "ANONYMOUS"},
	"POST/auth/login/admin":      {"ROLE_MAIN_ADMIN", "ROLE_ADMIN", "ANONYMOUS"},
	"POST/admin/register/groups": {"ROLE_MAIN_ADMIN", "ROLE_ADMIN"},
	"POST/admin/get/list/groups": {"ROLE_MAIN_ADMIN", "ROLE_ADMIN", "ANONYMOUS"},
}

func isAnonymous(method string, path string) bool {
	if access, ok := accessTable[method+path]; ok {
		for _, r := range access {
			if r == "ANONYMOUS" {
				return true
			}
		}
		return false
	}
	return true
}

func IsAccessGranted(method string, path string, user *ClaimsUser) bool {
	if access, ok := accessTable[method+path]; ok {
		for _, r := range access {
			if user.Role == r {
				return true
			}
		}
		return false
	}
	return true
}
