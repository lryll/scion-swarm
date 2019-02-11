//package to manage everything related to the test and meta information about scion-swarm
package scionutils

import (
	"fmt"
	"github.com/scionproto/scion/go/lib/snet"

	"github.com/ethereum/go-ethereum/log"
)

var TestRunScionSwarm = false
var sciondPath =""

func SetupSCION(SCIONAddr, sciond, pathDictionary string) error {

	if pathDictionary !=""{
		err := LoadPathDirectory(pathDictionary)
			if err != nil{
			log.Error(err.Error())
		}else{
			fmt.Println("PathDirectory loaded successfully from file.")
			PrintPathDirectory()
		}
	}

	if SCIONAddr != "" {
		scionAddr, err := snet.AddrFromString(SCIONAddr)
		if err != nil {
			fmt.Println("SCIONADDR: " + SCIONAddr)
			return fmt.Errorf("error in decoding scion addr. %v", err)
		}

		dispatcherAddr := "/run/shm/dispatcher/default.sock"

		var sciondAddr string
		if sciond != "" {
			sciondAddr = sciond
		} else {
			sciondAddr = "/run/shm/sciond/default.sock"
		}

		sciondPath=sciondAddr //for path selection later on (dirty)

		//init not only for udp but complete snet,
		//udp is the first established connection in geth, swarm though
		return snet.Init(scionAddr.IA, sciondAddr, dispatcherAddr)

	} else {
		return fmt.Errorf("Scion address needs to be specified with -scion)")
	}
}
