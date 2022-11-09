package simplehttp

import (
	"fmt"
	"testing"

	"github.com/uber-company/gonmap/simplenet"
)

func TestGet(t *testing.T) {
	fmt.Println("Initializing PortRingPool...")
	simplenet.InitPortRingPool(1000, 2000)
	fmt.Println("PortRingPool Initialized")
	re, err := Get("https://www.baidu.com")
	fmt.Println(re, "err", err)
}
