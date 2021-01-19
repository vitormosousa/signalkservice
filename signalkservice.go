package signalkservice

import (
	"github.com/paypal/gatt"
)

func NewSignalkService() *gatt.Service {
	lv := byte(100)
	//n := 0
	//s := gatt.NewService(gatt.MustParseUUID("03689026-5a14-11eb-ae93-0242ac130002"))
	s := gatt.NewService(gatt.UUID16(0x181a))
	// inovo
	c := s.AddCharacteristic(gatt.UUID16(0x2A73))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write([]byte{lv})
			lv--
		})
	c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{4, 1, 39, 173, 1, 0, 0})
	// fnovo

	// i-rem
	/* 	s.AddCharacteristic(gatt.MustParseUUID("03689260-5a14-11eb-ae93-0242ac130002")).HandleReadFunc(
	   		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
	   			fmt.Fprintf(rsp, "count: %d", n)
	   			n++
	   		})

	   	s.AddCharacteristic(gatt.MustParseUUID("03689350-5a14-11eb-ae93-0242ac130002")).HandleWriteFunc(
	   		func(r gatt.Request, data []byte) (status byte) {
	   			log.Println("Wrote:", string(data))
	   			return gatt.StatusSuccess
	   		})

	   	s.AddCharacteristic(gatt.MustParseUUID("03689418-5a14-11eb-ae93-0242ac130002")).HandleNotifyFunc(
	   		func(r gatt.Request, n gatt.Notifier) {
	   			cnt := 0
	   			for !n.Done() {
	   				fmt.Fprintf(n, "Count: %d", cnt)
	   				cnt++
	   				time.Sleep(time.Second)
	   			}
	   		}) */

	// f-rem
	return s
}
