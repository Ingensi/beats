package network

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)
var netService NETService
var oldNetRaw NETRaw
var newNetRaw NETRaw


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

	value:= netService.getRxDroppedPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t, float64(100),value)
}
func TestGetRxPacketsPerSeconde( t *testing.T){
	setTime()
	oldNetRaw.RxPackets= 140
	newNetRaw.RxPackets= 240

	value:= netService.getRxPacketsPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t, float64(50),value)
}
func TestGetRxErrorsPerSeconde( t *testing.T){
	setTime()
	oldNetRaw.RxErrors=0
	newNetRaw.RxErrors=0

	value:= netService.getRxErrorsPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t,float64(0),value)
}

func TestGetTxBytesPerSecond(t *testing.T){
	setTime()
	oldNetRaw.TxPackets=10
	newNetRaw.TxPackets=0

	value:= netService.getTxBytesPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t, float64(0),value)
}
func TestGetTxDroppedPerSeconde(t *testing.T){
	setTime()
	oldNetRaw.TxDropped=95
	newNetRaw.TxDropped=195

	value:= netService.getTxDroppedPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t,float64(50),value)
}
func TestGetTxPacketsPerSeconde(t *testing.T){
	setTime()
	oldNetRaw.TxPackets=951
	newNetRaw.TxPackets=1951

	value:=netService.getTxPacketsPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t,float64(500),value)
}
func TestGetTxErrorsPerSecond(t *testing.T){
	setTime()
	oldNetRaw.TxErrors=995
	newNetRaw.TxErrors=1995

	value:=netService.getTxErrorsPerSecond(&newNetRaw,&oldNetRaw)
	assert.Equal(t, float64(500),value)
}

func setTime(){
	oldNetRaw.Time= time.Now()
	newNetRaw.Time= oldNetRaw.Time.Add(time.Duration(2000000000))
}
