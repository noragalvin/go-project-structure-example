package custom_types

// {"key_id":9,"user_id":"123","consumer_key":"ck_6e769f426e3434a286b0e3071542fc38b971e539","consumer_secret":"cs_72b6fdbd59cae6c0fe06f1480fa38e39d2400fc8","key_permissions":"read_write"}
type AuthCallBackBody struct {
	KeyID          int    `json:"key_id"`
	UserId         string `json:"user_id"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	KeyPermissions string `json:"key_permissions"`
}
