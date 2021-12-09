package f5g

func (c *client) GTM() *GTM {
	return &GTM{c: c}
}
