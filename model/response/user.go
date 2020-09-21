package response

// IDName 返回ID和nickname
type IDName struct {
	ID       uint   `json:"id"`
	NickName string `json:"nickname"`
}
