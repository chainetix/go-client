package chainetix

func (client *Client) DeleteUsers() error {
	src := map[string]interface{}{
	}
	_, err := client.Delete(
		"https://multichain.blokhub.io/api/users",
		src,
		nil,
		client.headers,
	)
	if err != nil {
		return err
	}
	return nil
}