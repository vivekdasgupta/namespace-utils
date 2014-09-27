/*
 nsview.go :
		Program to view various namespaces currently running in the system

 Build : go build nsview.go

 Copyright (C) 2014 Red Hat Inc.
 Author Vivek Dasgupta <vivek.dasgupta@gmail.com> | <vdasgupt@redhat.com>

 This program is free software; you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation; either version 2 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program; if not, write to the Free Software
 Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 
*/



package main

import	(
	"fmt"
	"os"
	"strings"
	"strconv"
	"log"
)

type Debug bool

func (dbg Debug) Printf(msg string, val ...interface{}) {
	if dbg {
		fmt.Printf(msg, val...)
	}
}

var dbg Debug


func getNamespaceId(procPidStr string, nsType string) (string, int) {

	nsLink, err := os.Readlink("/proc/" + procPidStr + "/ns/" + nsType)
        if err != nil {
                log.Fatal(err)
	}
        dbg.Printf("\nProc PID1 ipc ns link : %s\n", nsLink)


	nsNameId := strings.FieldsFunc(nsLink, func (r rune) bool {
		return r == '[' || r == ']'
	})

	nsIdStr := nsNameId[1]
	dbg.Printf("\nProc PID1 ipc ns id str: %s\n", nsNameId[1])

	nsId, err := strconv.Atoi(nsNameId[1])
        if err != nil {
                log.Fatal(err)
        }


	dbg.Printf("\nProc PID1 ipc ns id int: %d\n", nsId)

	return nsIdStr,nsId

}


func getProcPidList() []string {

        procDir, err := os.Open("/proc") 
        if err != nil {
                log.Fatal(err)
        }


        procContent, err := procDir.Readdirnames(0)
        if err != nil {
                log.Fatal(err)
        }

        procContentLen := len(procContent)

        dbg.Printf("\nProc list : %s \n Length : %d\n", procContent, procContentLen)

        procIndex := findIndex(procContent, "1")
        if (procIndex == -1) {
                log.Fatal()
        }

        dbg.Printf("\nProc Pid 1 index : %d\n", procIndex)

        procPidList := procContent[procIndex:procContentLen]

	dbg.Printf("\nProc PID list : %s\n", procPidList)

	return procPidList

}

func findIndex(strSlice []string, val string) int {

	for i, v := range strSlice {
		if (v == val) {
			return i
		}
	}
	return -1
}


func main() {

	dbg = false

	procPidList := getProcPidList()

//	fmt.Printf("\nMainfunc Proc PID1 ipc ns id int: %d\n", getNamespaceId(procPidList[0], "ipc"))

	// find host namespace ids , create a slice for host namespace group

	hostNsGroup := make([]string, 1)

	hostNsGroup[0] = "1"

	ipcIdList := make([]string, 1)
	mntIdList := make([]string, 1)
	netIdList := make([]string, 1)
	pidIdList := make([]string, 1)
	utsIdList := make([]string, 1)

	ipcIdList[0],_ = getNamespaceId(procPidList[0], "ipc")
	mntIdList[0],_ = getNamespaceId(procPidList[0], "mnt")
	netIdList[0],_ = getNamespaceId(procPidList[0], "net")
	pidIdList[0],_ = getNamespaceId(procPidList[0], "pid")
	utsIdList[0],_ = getNamespaceId(procPidList[0], "uts")


	for _, pid := range procPidList {

		// for range of namespaces
/*
		if ( (getNamespaceId(pid,"ipc") == getNamespaceId(nsLeader,"ipc")) &&
		    (getNamespaceId(pid,"mnt") == getNamespaceId(nsLeader,"mnt")) &&
		    (getNamespaceId(pid,"net") == getNamespaceId(nsLeader,"net")) &&
		    (getNamespaceId(pid,"pid") == getNamespaceId(nsLeader,"pid")) &&
		    (getNamespaceId(pid,"uts") == getNamespaceId(nsLeader,"uts")) ) 
*/


	// ipc

		ipcIdStr,_ := getNamespaceId(pid,"ipc")
		ipcIdIndex := findIndex(ipcIdList, ipcIdStr)
		if (ipcIdIndex == -1) {
			ipcIdList = append(ipcIdList, ipcIdStr)
		}

	// mnt

		mntIdStr,_ := getNamespaceId(pid,"mnt")
		mntIdIndex := findIndex(mntIdList, mntIdStr)
		if (mntIdIndex == -1) {
			mntIdList = append(mntIdList, mntIdStr)
		}

	// net

		netIdStr,_ := getNamespaceId(pid,"net")
		netIdIndex := findIndex(netIdList, netIdStr)
		if (netIdIndex == -1) {
			netIdList = append(netIdList, netIdStr)
		}

	// pid

		pidIdStr,_ := getNamespaceId(pid,"pid")
		pidIdIndex := findIndex(pidIdList, pidIdStr)
		if (pidIdIndex == -1) {
			pidIdList = append(pidIdList, pidIdStr)
		}

	// uts

		utsIdStr,_ := getNamespaceId(pid,"uts")
		utsIdIndex := findIndex(utsIdList, utsIdStr)
		if (utsIdIndex == -1) {
			utsIdList = append(utsIdList, utsIdStr)
		}


	}   // for range PidList


	fmt.Printf("\nIPC namespaces %s\n", ipcIdList)
	fmt.Printf("\nMNT namespaces %s\n", mntIdList)
	fmt.Printf("\nNET namespaces %s\n", netIdList)
	fmt.Printf("\nPID namespaces %s\n", pidIdList)
	fmt.Printf("\nUTS namespaces %s\n", utsIdList)

	// TODO: for each process check namespace id and add it to host list if matches

	// TODO: create another slice for each new namespace and add matching processes

}
