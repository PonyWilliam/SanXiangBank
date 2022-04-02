package main

import (
	"context"
)

// PublishImp servant implementation
type PublishImp struct {
}

// Init servant init
func (imp *PublishImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *PublishImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *PublishImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *PublishImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
