package enums

type EventTypes string

const (
	LOG_IN                 EventTypes = "sn.user.login"
	LOG_IN_CHALLENGE       EventTypes = "sn.user.login.challenge"
	LOG_IN_FAILURE         EventTypes = "sn.user.login.failure"
	LOG_OUT                EventTypes = "sn.user.logout"
	SIGN_UP                EventTypes = "sn.user.signup"
	AUTH_CHALLENGE         EventTypes = "sn.user.auth.challenge"
	AUTH_CHALLENGE_SUCCESS EventTypes = "sn.user.auth.challenge.success"
	AUTH_CHALLENGE_FAILURE EventTypes = "sn.user.auth.challenge.failure"
	TWO_FACTOR_DISABLE     EventTypes = "sn.user.2fa.disable"
	EMAIL_UPDATE           EventTypes = "sn.user.email.update"
	PASSWORD_REST          EventTypes = "sn.user.password.reset"
	PASSWORD_REST_SUCCESS  EventTypes = "sn.user.password.reset.success"
	PASSWORD_UPDATE        EventTypes = "sn.user.password.update"
	PASSWORD_REST_FAILURE  EventTypes = "sn.user.password.reset.failure"
	USER_INVITE            EventTypes = "sn.user.invite"
	ROLE_UPDATE            EventTypes = "sn.user.role.update"
	PROFILE_UPDATE         EventTypes = "sn.user.profile.update"
	PAGE_VIEW              EventTypes = "sn.user.page.view"
	VERIFY_EVENT           EventTypes = "sn.verify"
)
