package chainetix

func (client *Client) Users() ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/users",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}