package models

type ColorResponse struct {
	Uid string
}

func (c *ColorResponse) GetUID() string {
	return c.Uid
}
