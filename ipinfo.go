package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func toBinary(num int) string {
	n := strconv.FormatInt(int64(num), 2)

	binary := "00000000"[len(string(n)):] + string(n)

	return binary
}

func convertIPToBinary(ip string) string {
	segments := strings.Split(ip, ".")

	binary := []string{}

	for _, m := range segments {
		n, _ := strconv.Atoi(m)
		binary = append(binary, toBinary(n))
	}

	return strings.Join(binary, "")
}

func getRequiredBitsForNumber(num int) int {
	ret := -1

	for i := 1; i < 8; i++ {
		if powInt(2, i) >= num {
			ret = i
			break
		}
	}

	return ret
}

func binaryIPToDecimal(binary string) string {
	re := regexp.MustCompile(`.{1,8}`)
	segments := re.FindAllStringSubmatch(binary, -1)

	decimal := []string{}

	for _, m := range segments {
		segNum, _ := strconv.ParseInt(m[0], 2, 64)
		decimal = append(decimal, strconv.FormatInt(segNum, 10))
	}

	return strings.Join(decimal, ".")
}

func getSubnetMask(addr string, numSubnets int) string {
	num := getRequiredBitsForNumber(numSubnets)

	subnet := ""

	foundZeros := false

	for i := 0; i < len(addr); i += 8 {
		seg := addr[i : i+8]

		if seg == "00000000" && foundZeros == false {
			foundZeros = true
			subnet += strings.Repeat("1", num) + strings.Repeat("0", 8-num)
		} else if seg == "00000000" && foundZeros == true {
			subnet += strings.Repeat("0", 8)
		} else {
			subnet += strings.Repeat("1", 8)
		}
	}

	if !foundZeros {
		subnet = ""

		for i := 0; i < len(addr); i++ {
			if i < numSubnets {
				subnet += "1"
			} else {
				subnet += "0"
			}
		}
	}

	return subnet
}

func replaceAt(str string, index int, ch string) string {
	return str[:index-1] + ch + str[index:]
}

func getFirstAvailable(address string, subnet string) string {
	numTrailingZeros := len(subnet[len(subnet)-len(strings.Split(subnet, "0"))+1:])

	addr := address

	addr = replaceAt(addr, len(addr)-numTrailingZeros, "1")
	addr = replaceAt(addr, len(addr), "1")

	return addr
}

func getMaxHosts(subnet string) int {
	numTrailingZeros := len(subnet[len(subnet)-len(strings.Split(subnet, "0"))+1:])

	return powInt(2, numTrailingZeros) - 2
}

type extraInfo struct {
	broadcastAddress string
	networkAddress   string
}

func getExtraInfo(binaryIP string, binarySubnet string) extraInfo {
	networkAddress := ""
	broadcastAddress := ""

	for i := 0; i < len(binaryIP); i++ {
		if binarySubnet[i] == '1' {
			networkAddress += string(binaryIP[i])
		} else {
			networkAddress += "0"
		}

		if binarySubnet[i] == '0' {
			broadcastAddress += "1"
		} else {
			broadcastAddress += string(binaryIP[i])
		}
	}

	return extraInfo{broadcastAddress: binaryIPToDecimal(broadcastAddress), networkAddress: binaryIPToDecimal(networkAddress)}
}

func getInfo(ip string, subnets int) {
	binaryIP := convertIPToBinary(ip)
	subnetMask := getSubnetMask(binaryIP, subnets)
	firstAvail := getFirstAvailable(binaryIP, subnetMask)
	maxHosts := getMaxHosts(subnetMask)
	info := getExtraInfo(binaryIP, subnetMask)

	fmt.Printf("Subnet mask: %s\n", binaryIPToDecimal(subnetMask))
	fmt.Printf("First available host: %s\n", binaryIPToDecimal(firstAvail))

	fmt.Printf("Max hosts: %d\n", maxHosts)

	fmt.Printf("Broadcast address: %s\n", info.broadcastAddress)
	fmt.Printf("Network address: %s\n", info.networkAddress)
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<IP>", "<SUBNETS>")
		return
	}

	ip := args[0]
	subnets, _ := strconv.Atoi(args[1])

	getInfo(ip, subnets)
}
