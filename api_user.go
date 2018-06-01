package chainetix

func (client *Client) User(user string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/user/" + user + "",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}