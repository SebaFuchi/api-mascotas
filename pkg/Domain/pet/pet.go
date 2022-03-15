package pet

type Pet struct {
	Token      string `json:"token"`
	Ownertoken string `json:"ownertoken"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
}
