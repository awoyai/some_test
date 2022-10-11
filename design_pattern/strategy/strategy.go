package strategy

import (
	"fmt"
	"time"
)

type Action interface {
	TransPort()
}

type Car struct {
}

func (c *Car) TransPort() {
	fmt.Println("开汽车")
}

type Plane struct {
}

func (c *Plane) TransPort() {
	fmt.Println("坐飞机")
}

type Tran struct {
}

func (c *Tran) TransPort() {
	fmt.Println("坐高铁")
}

func GoToShenZhen(wanTime time.Duration) Action {
	switch {
	case wanTime.Minutes() > 5*time.Hour.Minutes() && wanTime.Minutes() < 7*time.Hour.Minutes():
		return new(Tran)
	case wanTime.Minutes() <= 5*time.Hour.Minutes() && wanTime.Minutes() >= 1*time.Hour.Minutes():
		return new(Plane)
	default:
		return new(Car)
	}
}
