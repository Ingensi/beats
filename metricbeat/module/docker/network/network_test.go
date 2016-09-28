package network

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)
var netService NETService
var oldNetRaw NETRaw
var newNetRaw NETRaw

func setTime(){
	oldNetRaw.Time= time.Now()
	newNetRaw.Time= oldNetRaw.Time.Add(time.Duration(2000000000))
}
func TestGetRxBytesPerSecond( t *testing.T){
	setTime()
	// set old & new rxBytes
	oldNetRaw.RxBytes= 20
	newNetRaw.RxBytes=120
	//WHEN
	value := netService.getRxBytesPerSecond(&newNetRaw,&oldNetRaw)
	//THEN
	assert.Equal(t, float64(50),value)
}
func TestGetRxDroppedPerSeconde(t * testing.T){
	setTime()
	oldNetRaw.RxDropped = 40
	newNetRaw.RxDropped= 240
	//WHEN
	value:= netService.getRxDroppedPerSecond(&newNetRaw,&oldNetRaw)
	//THEN
	assert.Equal(t, float64(100),value)
}
