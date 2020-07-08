package enums

type EventTypesEnum struct {
	LogIn                string
	LogInChallenge       string
	LogInFailure         string
	LogOut               string
	SignUp               string
	AuthChallenge        string
	AuthChallengeSuccess string
	AuthChallengeFailure string
	TwoFactorDisable     string
	EmailUpdate          string
	PasswordRest         string
	PasswordRestSuccess  string
	PasswordUpdate       string
	PasswordRestFailure  string
	UserInvite           string
	RoleUpdate           string
	ProfileUpdate        string
	PageView             string
	Verify               string
}

var EventTypes = EventTypesEnum{
	LogIn:                "sn.user.login",
	LogInChallenge:       "sn.user.login.challenge",
	LogInFailure:         "sn.user.login.failure",
	LogOut:               "sn.user.logout",
	SignUp:               "sn.user.signup",
	AuthChallenge:        "sn.user.auth.challenge",
	AuthChallengeSuccess: "sn.user.auth.challenge.success",
	AuthChallengeFailure: "sn.user.auth.challenge.failure",
	TwoFactorDisable:     "sn.user.2fa.disable",
	EmailUpdate:          "sn.user.email.update",
	PasswordRest:         "sn.user.password.reset",
	PasswordRestSuccess:  "sn.user.password.reset.success",
	PasswordUpdate:       "sn.user.password.update",
	PasswordRestFailure:  "sn.user.password.reset.failure",
	UserInvite:           "sn.user.invite",
	RoleUpdate:           "sn.user.role.update",
	ProfileUpdate:        "sn.user.profile.update",
	PageView:             "sn.user.page.view",
	Verify:               "sn.verify",
}
