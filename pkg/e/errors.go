package e

type Errors struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func New(errno int) *Errors {
	return &Errors{Errmsg: GetMsg(errno), Errno: errno}
}
