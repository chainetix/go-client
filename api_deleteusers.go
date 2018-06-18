package chainetix

func (client *Client) DeleteUsers() error {
	src := map[string]interface{}{
	}
	_, err := client.Delete(
		"https://multichain.chainetix.com/api/users",
		src,
		nil,
		client.headers,
	)
	if err != nil {
		return err
	}
	return nil
}