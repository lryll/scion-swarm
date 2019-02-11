package scionutils

import (
	"fmt"
	"os"
	"os/user"
	"sync"
	"sync/atomic"
	"time"
)

//metric/stats related files
//
// Custom chunk stats currently disabled!
//

var statsReceiveChunkDeliveryMsg *stats
var statsRetrieveRequestMsg *stats

var timeFormat = "2006-01-02 15:04:05.000"

var CustomStatsLogFilePath = "/home/ubuntu/go/src/github.com/ethereum/go-ethereum/~scion/stats/chunkStatsNormalSwarm.log"

type stats struct {
	firstChunkReceived time.Time
	lastChunkReceived  time.Time
	timer              *time.Timer
	finished           bool
	name               string
	numberChunks       int64
	selfID             string
	peerID             string
	sync.Mutex
	//peerName string //own peer name for log
}

func (s *stats) Received() {

	if s.firstChunkReceived.IsZero() {
		s.firstChunkReceived = time.Now()
		s.timer = time.NewTimer(7 * time.Second)
		atomic.AddInt64(&s.numberChunks, 1)
		go func() {
			<-s.timer.C
			fmt.Printf("[%s|%s local %s ; remote %s] Duration: %.6fs Chunks: %d \n", time.Now().Format(timeFormat), s.name,  s.selfID, s.peerID, s.lastChunkReceived.Sub(s.firstChunkReceived).Seconds(), s.numberChunks)
			s.WriteLogToFile()
			s.finished = true
		}()
	} else {
		s.Lock()
		s.lastChunkReceived = time.Now()
		s.timer.Reset(10 * time.Second)
		s.numberChunks++
		//atomic.AddInt64(&s.numberChunks, 1)
		s.Unlock()
	}
}

//write stats to logfile, create file if it doesn't exist
func (s *stats) WriteLogToFile() {

	user, err := user.Current()

	fmt.Printf("UserName %s \n", user.Name)
	//if user.Name != "ubuntu" || user.Name != "Ubuntu" { //somehow in the scion-VM the user name gets "Ubuntu" with upper case, why is that the case?
    //	customStatsLogFilePath = "/home/"+user.Name+"/go/src/github.com/ethereum/go-ethereum/~scion/stats/chunkStatsNormalSwarm.log"
	//}

	f, err := os.OpenFile(CustomStatsLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat,err := f.Stat()
	if stat.Size() == 0 {
		header := fmt.Sprintf("Timestamp;Event;self;peer;TimeInS;NumberOfChunks\n")
		if _, err = f.WriteString(header); err != nil {
			panic(err)
		}
	}

	//otherwise there will be a difference for this case
	if s.numberChunks == 1 {
		s.lastChunkReceived = s.firstChunkReceived
	}

	//msg := fmt.Sprintf("%s;%s local %s ; remote %s] Duration: %.6fs Chunks: %d \n", time.Now().Format(timeFormat), s.name,  s.selfID, s.peerID, s.lastChunkReceived.Sub(s.firstChunkReceived).Seconds(), s.numberChunks)
	msg := fmt.Sprintf("%s;%s;%s;%s;%.6f;%d\n", time.Now().Format(timeFormat), s.name,  s.selfID, s.peerID, s.lastChunkReceived.Sub(s.firstChunkReceived).Seconds(), s.numberChunks)

	if _, err = f.WriteString(msg); err != nil {
		panic(err)
	}
}


func CustomStatsReceiveChunkDelivery(selfID, peerID string) {
	if statsReceiveChunkDeliveryMsg == nil {
		statsReceiveChunkDeliveryMsg = &stats{time.Time{}, time.Time{}, nil, false, "ReceiveChunks", 0, selfID, peerID, sync.Mutex{}}
		statsReceiveChunkDeliveryMsg.Received()
	} else if !statsReceiveChunkDeliveryMsg.finished {
		statsReceiveChunkDeliveryMsg.Received()
	} else if statsReceiveChunkDeliveryMsg.finished { //reset for next run
		statsReceiveChunkDeliveryMsg = &stats{time.Time{}, time.Time{}, nil, false, "ReceiveChunks", 0, selfID, peerID,sync.Mutex{}}
		statsReceiveChunkDeliveryMsg.Received()
	}
}

func CustomStatsRetrieveRequestMsgCount(selfID, peerID string) {
	if statsRetrieveRequestMsg == nil {
		statsRetrieveRequestMsg = &stats{time.Time{}, time.Time{}, nil, false, "DeliverChunks", 0, selfID, peerID,sync.Mutex{}}
		statsRetrieveRequestMsg.Received()
	} else if !statsRetrieveRequestMsg.finished {
		statsRetrieveRequestMsg.Received()
	} else if statsRetrieveRequestMsg.finished { //reset for next run
		statsRetrieveRequestMsg = &stats{time.Time{}, time.Time{}, nil, false, "DeliverChunks", 0, selfID, peerID,sync.Mutex{}}
		statsRetrieveRequestMsg.Received()
	}
}
