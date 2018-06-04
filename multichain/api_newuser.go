package chainetix

func (client *Client) NewUser(groups []interface {}) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"groups": groups,
	}
	_, err := client.Post(
		"https://multichain.blokhub.io/api/user",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}