package main

import (
	"encoding/xml"
	"fmt"
	"net"
	"os"
	"strings"
)

type xmlIpc struct {
	XMLName                 xml.Name `xml:"ProbeMatch"`
	Uuid                    string   `xml:"Uuid"`
	Types                   string   `xml:"Types"`
	DeviceType              int      `xml:"DeviceType"`
	DeviceDescription       string   `xml:"DeviceDescription"`
	DeviceSN                string   `xml:"DeviceSN"`
	CommandPort             int      `xml:"CommandPort"`
	HttpPort                int8     `xml:"HttpPort"`
	EHomeVer                string   `xml:"EHomeVer"`
	IPv4Address             string   `xml:"IPv4Address"`
	IPv4SubnetMask          string   `xml:"IPv4SubnetMask"`
	IPv4Gateway             string   `xml:"IPv4Gateway"`
	IPv6Address             string   `xml:"IPv6Address"`
	IPv6Gateway             string   `xml:"IPv6Gateway"`
	IPv6MaskLen             string   `xml:"IPv6MaskLen"`
	DHCP                    bool     `xml:"DHCP"`
	AnalogChannelNum        int8     `xml:"AnalogChannelNum"`
	DigitalChannelNum       int8     `xml:"DigitalChannelNum"`
	SoftwareVersion         string   `xml:"SoftwareVersion"`
	DSPVersion              string   `xml:"DSPVersion"`
	Encrypt                 bool     `xml:"Encrypt"`
	Salt                    string   `xml:"Salt"`
	BootTime                string   `xml:"BootTime"`
	DiskNumber              int8     `xml:"DiskNumber"`
	OEMInfo                 string   `xml:"OEMInfo"`
	Activated               bool     `xml:"Activated"`
	PasswordResetAbility    bool     `xml:"PasswordResetAbility"`
	ResetAbility            bool     `xml:"ResetAbility"`
	SyncIPCPassword         bool     `xml:"SyncIPCPassword"`
	PasswordResetModeSecond bool     `xml:"PasswordResetModeSecond"`
	DeviceLock              bool     `xml:"DeviceLock"`
	DHCPAbility             bool     `xml:"DHCPAbility"`
	SupportGUID             bool     `xml:"SupportGUID"`
	SupportSecurityQuestion bool     `xml:"SupportSecurityQuestion"`
	SupportIPv6             string   `xml:"supportIPv6"`
	SupportModifyIPv6       string   `xml:"supportModifyIPv6"`
	SupportHCPlatform       bool     `xml:"SupportHCPlatform"`
	//xml???????????????????????? strconv.ParseBool: parsing "flase": invalid syntax ????????????"flase"????????????bool?????????????????????????????????string
	HCPlatformEnable         string `xml:"HCPlatformEnable"`
	IsModifyVerificationCode bool   `xml:"IsModifyVerificationCode"`
	SupportMailBox           bool   `xml:"SupportMailBox"`
	SupportEzvizUnbind       string `xml:"supportEzvizUnbind"`
}

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(ipAddr()),
		Port: 37020,
	})
	if err != nil {
		fmt.Println("??????????????????, err:", err)
		return
	}
	defer func(listen *net.UDPConn) {
		err := listen.Close()
		if err != nil {
			fmt.Println("err:", err)
		}
	}(listen)
	fmt.Println("??????????????????!")
	for {
		var data [2048]byte
		n, _, err := listen.ReadFromUDP(data[:]) // ????????????
		if err != nil {
			fmt.Println("??????udp????????????, err:", err)
			continue
		}
		xmlData := data[:n]
		//fmt.Println("xml??????:", string(xmlData))
		var ipc xmlIpc
		err = xml.Unmarshal(xmlData, &ipc)
		if err != nil {
			fmt.Println("xml????????????????????????", err.Error())
		} else {
			fmt.Println(ipc.DeviceSN, ipc.IPv4Address)
		}
	}
}

// ????????????IP??????
func ipAddr() string {
	//var ipAddr
	var ipAddr string
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("????????????????????????: ", err)
		os.Exit(-1)
	}
	for _, inter := range interfaces {
		flags := inter.Flags.String()
		if strings.Contains(flags, "up") && strings.Contains(flags, "broadcast") {
			addrs, err := inter.Addrs()
			if err != nil {
				print("you")
			}
			for i := 0; i < len(addrs); i++ {
				addr := addrs[i].String()
				ip := addr[len(addr)-2:]
				if ip == "24" {
					ipAddr = addr[:len(addr)-3]
				}
			}
		}
	}
	return ipAddr
}
