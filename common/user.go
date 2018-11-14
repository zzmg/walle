package common

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type QyToken struct {
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type QyUser struct {
	ErrCode  int64        `json:"errcode"`
	ErrMsg   string       `json:"errmsg"`
	UserList []QyUserItem `json:"userlist"`
}
type QyUserItem struct {
	UserId     string  `json:"user_id"`
	Name       string  `json:"name"`
	Department []int64 `json:"department"`
	Position   string  `json:"position"`
	Mobile     string  `json:"mobile"`
	Gender     string  `json:"gender"`
	Email      string  `json:"email"`
	Avatar     string  `json:"avatar"`
	Status     int     `json:"status"`
	Enable     int     `json:"enable"`
	IsLeader   int     `json:"is_leader"`
	Extattr    struct {
		attrs []string
	} `json:"extattr"`
	HideMobile  int    `json:"hide_mobile"`
	EnglishName string `json:"english_name"`
	Telephone   string `json:"telephone"`
	Order       []int  `json:"order"`
	QrCode      string `json:"qr_code"`
}

type GitlabUser struct {
	Id               int64       `json:"id"`
	Name             string      `json:"name"`
	Username         string      `json:"username"`
	State            string      `json:"state"`
	AvatarUrl        string      `json:"avatar_url"`
	WebUrl           string      `json:"web_url"`
	CreatedAt        string      `json:"created_at"`
	Bio              string      `json:"bio"`
	Location         string      `json:"location"`
	Skype            string      `json:"skype"`
	linkedin         string      `json:"linkedin"`
	twitter          string      `json:"twitter"`
	WebsiteUrl       string      `json:"website_url"`
	Organization     string      `json:"organization"`
	LastSignInAt     string      `json:"last_sign_in_at"`
	ConfirmedAt      string      `json:"confirmed_at"`
	LastActivityOn   string      `json:"last_activity_on"`
	Email            string      `json:"email"`
	ThemeId          interface{} `json:"theme_id"`
	ColorSchemeId    int64       `json:"color_scheme_id"`
	ProjectsLimit    int64       `json:"projects_limit"`
	CurrentSignInAt  string      `json:"current_sign_in_at"`
	Identities       []string    `json:"identities"`
	CanCreateGroup   bool        `json:"can_create_group"`
	CanCreateProject bool        `json:"can_create_project"`
	TwoFactorEnabled bool        `json:"two_factor_enabled"`
	External         bool        `json:"external"`
	IsAdmin          bool        `json:"is_admin"`
}
