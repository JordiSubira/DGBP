package main

import (
	"bytes"
	"fmt"
	"log"
	"flag"
	"os/exec"
	"io/ioutil"
	"net"
	"encoding/binary"
	//"os"
	//"strings"
)

const createUserNodeFileName = "createUser.js"
const deleteUserNodeFileName = "deleteUser.js"
const queryUserNodeFileName = "queryUser.js"
const queryUsersByMspNodeFileName = "queryByMSP.js"

const createDptNodeFileName = "createDepartment.js"
const deleteDptNodeFileName = "deleteDepartment.js"
const queryDptNodeFileName = "queryDepartment.js"

const createPolicyFileName = "createPolicy.js"
const deletePolicyFileName = "deletePolicy.js"
const queryPolicyFileName = "queryPolicy.js"

const createResourceFileName = "createResource.js"
const deleteResourceFileName = "deleteResource.js"
const queryResourceFileName = "queryResource.js"

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}


func main() {

	//var args []string

	//args = os.Args[1:]

	//User flags
	//userNamePtr := flag.String("usrName", "", "string user name")
	userEidPtr := flag.String("usrEid", "", "string user EID")
	userPkiPtr := flag.String("usrPki","","string public key RFC 1...")
	userDptPtr := flag.String("usrDpt","","string department name")
	qUserMspPtr := flag.String("qUserMsp","","string MSP name (just for querys)")

	//Dep flags
	dptNamePtr := flag.String("dptName","","string department name")
	qDptMspPtr := flag.String("qDptMsp","","string MSP name (just for querys)")


	//Policy flags
	//dstDptNamePtr := flag.String("dstDpt","","string destination department")
	dstResEidPtr := flag.String("dstResEid","","string destination department")
	fromDptnamePtr := flag.String("fromDpt","","string from department")
	fromMspNamePtr := flag.String("fromMsp","","string from MSP")
	fromUserEid := flag.String("fromUserPki","","string from user")
	qDstMspPtr := flag.String("qDstMsp","","string MSP name (just for querys)")

	//Res Flags
	resNamePtr := flag.String("resEid","","string resource eid")
	qResMspPtr := flag.String("qResMsp","","string MSP name (just for querys)")


    flag.Parse()

    //fmt.Println(flag.Args())

	if len(flag.Args()) != 1{
		fmt.Println("Usage: expected just one action")
		return
	}

	action := flag.Args()[0]
	//fmt.Println("tail:", action)

	cmd := exec.Command("node")
	cmd.Dir = "/home/jordi/Fabric1.1/myapp"
	var out bytes.Buffer
	var err2 bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err2

	switch action{

		case "createUser":
			if *userPkiPtr== "" || *userEidPtr == "" || *userDptPtr == ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			content, err := ioutil.ReadFile(*userPkiPtr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File contents: %s", string(content))

			cmd.Args = []string{"node",createUserNodeFileName,string(content),*userEidPtr,*userDptPtr}

		case "create1000User":
			if *userPkiPtr== "" || *userDptPtr == ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			content, err := ioutil.ReadFile(*userPkiPtr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File contents: %s", string(content))

			userIpfromInt:= int2ip(2)	

			fmt.Printf("Ip: %s", userIpfromInt.String())

			for i := 20; i < 999; i++{

				cmd = exec.Command("node")
				cmd.Dir = "/home/jordi/Fabric1.1/myapp"
				cmd.Stdout = &out
				cmd.Stderr = &err2

				userIpfromInt = int2ip(uint32(i))
				fmt.Printf("Ip: %s", userIpfromInt.String())

				cmd.Args = []string{"node",createUserNodeFileName,string(content), userIpfromInt.String(), *userDptPtr}

				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("stdOut: %s\n", out.String())
				fmt.Printf("stdErr: %s\n", err2.String())

			}

			cmd = exec.Command("node")
			cmd.Dir = "/home/jordi/Fabric1.1/myapp"
			cmd.Stdout = &out
			cmd.Stderr = &err2

			
			userIpfromInt = int2ip(1000)
			cmd.Args = []string{"node",createUserNodeFileName,string(content), userIpfromInt.String(), *userDptPtr}

		case "deleteUser":
			if *userEidPtr== "" {
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			/*content, err := ioutil.ReadFile(*userPkiPtr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File contents: %s", string(content))*/
			cmd.Args = []string{"node",deleteUserNodeFileName, *userEidPtr}


		case "queryUser":
			if *userEidPtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			/*content, err := ioutil.ReadFile(*userEidPtr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File contents: %s", string(content))*/
			cmd.Args = []string{"node",queryUserNodeFileName,*userEidPtr}

		case "queryMspUsers":
			if *qUserMspPtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",queryUsersByMspNodeFileName,*qUserMspPtr}


		case "createDpt":
			if *dptNamePtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",createDptNodeFileName, *dptNamePtr}


		case "deleteDpt":
			if *dptNamePtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",deleteDptNodeFileName, *dptNamePtr}

		case "queryDpt":
			if *dptNamePtr== "" || *qDptMspPtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",queryDptNodeFileName,*qDptMspPtr, *dptNamePtr}

		case "createPolicy":
			if *dstResEidPtr== "" || *fromUserEid == ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",createPolicyFileName,*dstResEidPtr,*fromUserEid}

		case "deletePolicy":
			if *dstResEidPtr == "" || *fromUserEid == ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",deletePolicyFileName,*dstResEidPtr,*fromUserEid}

		case "queryPolicy":
			if *dstResEidPtr == "" || *fromUserEid == ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",queryPolicyFileName, *dstResEidPtr,*fromUserEid}

		case "createResource":
			if *resNamePtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",createResourceFileName, *resNamePtr}


		case "deleteResource":
			if *resNamePtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node", deleteResourceFileName, *resNamePtr}

		case "queryResource":
			if *resNamePtr== "" || *qResMspPtr== ""{
				fmt.Println("Usage: with " + action + " bad needed args")
				return
			}
			cmd.Args = []string{"node",queryResourceFileName, *qResMspPtr, *resNamePtr}

		default:
			fmt.Println("Usage: action is not recognized")
			return

	}

	_ = fromDptnamePtr
	_ = fromMspNamePtr
	_ = qDstMspPtr

	/*for i:=0 ; i<len(args);i++{
		cmd.Args = append(cmd.Args,args[i])
	}*/
	fmt.Println(cmd.Args)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stdOut: %s\n", out.String())
	fmt.Printf("stdErr: %s\n", err2.String())
}