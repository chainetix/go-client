package chainetix

func (client *Client) GetStreams() ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/streams",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}