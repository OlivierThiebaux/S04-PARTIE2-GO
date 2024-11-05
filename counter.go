package main

type Counter struct {
	Value int
}
func (c *Counter) Increment() {
	c.Value = c.Value +1
}