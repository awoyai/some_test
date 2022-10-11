package strategy

import (
	"testing"
	"time"
)

func TestTransPort(t *testing.T)  {
	GoToShenZhen(12*time.Hour).TransPort()
}
