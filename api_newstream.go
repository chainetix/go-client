package chainetix

func (client *Client) NewStream(name string, secret string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"name": name,
		"secret": secret,
	}
	_, err := client.Post(
		"https://multichain.blokhub.io/api/stream",
		src,
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}