package facex

type Config struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	GroupId   string `json:"group_id"`
	Timeout   int    `json:"timeout"`
}
