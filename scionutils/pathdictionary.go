package scionutils

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
	"github.com/scionproto/scion/go/lib/spath"
	"os"
	"strings"
	"time"
)

//SCION
//predefined paths as strings[19-ffaa:1:b5 1>46 19-ffaa:0:1303 1>5 19-ffaa:0:1301 1>2 16-ffaa:0:1001 5>3 16-ffaa:0:1004 1>2 18-ffaa:0:1201 6>1 18-ffaa:0:1202 34>1 18-ffaa:1:123]
//This is a workaround since there doesn't seem to be a method to get a FWDPath or host info from a string
//Therefore we just compare the string representation
//This is parsed from file if flag SCIONPathDictionaryLocationFlag is set and points to a valid file
// <key, value> == <remoteAS,Path> (name of map is the <key, value> == <remoteAS,Path> (name of map is the localAS
//IF local == 19-ffaa:1:b5
var IAPathDirectory = map[string]map[string]string{}

func SetPath(local *snet.Addr, remote *snet.Addr) (string, error) {

	localIA := strings.Split(local.String(), ",")[0] // local.String() is of format "19-ffaa:1:b5,[10.0.8.32]:0"
	remoteIA := strings.Split(remote.String(), ",")[0]

	if localIA == remoteIA {
		return "", errors.New(fmt.Sprintf("There is no path when the local AS equals the remote AS", localIA, remoteIA))
	}

	if currentPathRegister := IAPathDirectory[localIA]; currentPathRegister != nil {

		fmt.Printf("Found register for %s\n", localIA)

		chosenPath := currentPathRegister[remoteIA]

		if chosenPath == "" {
			return "", errors.New(fmt.Sprintf("No predefined path found from %s to %s", localIA, remoteIA))
		} else {

			fmt.Printf("Use path %s\n", chosenPath)

			pathReply, err := getPathReplyEntryFromString(chosenPath, *local, *remote)

			if err != nil {
				return "", err
			} else {
				setPathFromReplyEntry(remote, pathReply)
				fmt.Printf("PATH CHOSEN: \n %s \n", pathReply.Path.String())
				return pathReply.Path.String(), nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("CurrentPathRegister empty for %s", localIA))
}

//SCION
//modified from bwtestclient.go:main():328
func setPathFromReplyEntry(remoteAddr *snet.Addr, pathEntry *sciond.PathReplyEntry) {

	remoteAddr.Path = spath.New(pathEntry.Path.FwdPath)
	remoteAddr.Path.InitOffsets()
	remoteAddr.NextHop , _ = pathEntry.HostInfo.Overlay()

}

//SCION
//This is a workaround since there doesn't seem to be a method to get a PathReplyEntry or host info from a string
//Therefore we just compare the string representation
func getPathReplyEntryFromString(wantedPath string , local snet.Addr, remote snet.Addr ) (*sciond.PathReplyEntry, error) {
	maxPaths := 99
	sd := sciond.NewService(sciondPath , true) //can be called potentially another way without global var
	sdConn, err := sd.ConnectTimeout(time.Second*2)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect to SCIOND: %v\n", err))
	}
	reply, err := sdConn.Paths(context.Background(),remote.IA, local.IA, uint16(maxPaths), sciond.PathReqFlags{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve paths from SCIOND: %v\n", err))
	}

	//fmt.Printf("Found %d paths \n", len(reply.Entries))

	for	i, path := range reply.Entries{

		currentPath := path.Path.String() //especially this operation is highly dependent on string representation of the path entry ofc.

		//fmt.Printf("Retrieved Path [%d]: %s \n",i,currentPath)

		if currentPath	== wantedPath{
			fmt.Printf("Found path: %s", path.Path.String())
			return &reply.Entries[i], nil
		}
	}
	return nil, errors.New("Couldn't set predefined path.")
}


func LoadPathDirectory(filePath string) error {
	pathDirectoryFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer pathDirectoryFile.Close()
	scanner := bufio.NewScanner(pathDirectoryFile)
	var curLocalIAMap map[string]string
	for scanner.Scan() {

		line := fmt.Sprintf(scanner.Text())

		//localIA
		if !strings.HasPrefix(line, ">") {
			curLocalIAMap = map[string]string{}
			IAPathDirectory[line] = curLocalIAMap
		} else { //new remote IA with path
			remoteIA, path := string(strings.Split(line, ";")[0][1:]), strings.Split(line, ";")[1]
			curLocalIAMap[remoteIA] = path
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PrintPathDirectory() {
	for key := range IAPathDirectory {

		fmt.Printf("%s \n", key)

		for remotAI, remoteEntry := range IAPathDirectory[key] {

			fmt.Printf("%s; %s \n", remotAI, remoteEntry)
		}
	}
}


