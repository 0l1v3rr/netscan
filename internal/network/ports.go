package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var Ports []int = []int{7, 19, 20, 21, 22, 23, 24, 25, 26, 30, 32, 33, 37, 42, 43, 49, 53, 70, 79, 85, 88, 90, 99, 100, 106, 109, 111, 113, 119, 125, 135, 139, 143, 144, 146, 161, 163, 179, 199, 211, 212, 222, 254, 256, 259, 264, 280, 301, 306, 311, 340, 366, 389, 406, 407, 416, 417, 425, 427, 443, 445, 458, 464, 465, 481, 497, 500, 512, 515, 524, 541, 543, 545, 548, 554, 555, 563, 587, 593, 616, 617, 625, 631, 636, 646, 648, 666, 668, 683, 687, 691, 700, 705, 711, 714, 720, 722, 726, 749, 765, 777, 783, 787, 800, 801, 808, 843, 873, 880, 888, 898, 900, 903, 911, 912, 981, 987, 990, 992, 993, 995, 999, 1002, 1007, 1009, 1011, 1021, 1100, 1102, 1104, 1108, 1110, 1114, 1117, 1119, 1121, 1124, 1126, 1130, 1132, 1137, 1138, 1141, 1145, 1147, 1149, 1151, 1152, 1154, 1163, 1166, 1169, 1174, 1175, 1183, 1185, 1187, 1192, 1198, 1199, 1201, 1213, 1216, 1218, 1233, 1234, 1236, 1244, 1247, 1248, 1259, 1271, 1272, 1277, 1287, 1296, 1300, 1301, 1309, 1311, 1322, 1328, 1334, 1352, 1417, 1433, 1434, 1443, 1455, 1461, 1494, 1500, 1501, 1503, 1521, 1524, 1533, 1556, 1580, 1583, 1594, 1600, 1641, 1658, 1666, 1687, 1688, 1700, 1717, 1721, 1723, 1755, 1761, 1782, 1783, 1801, 1805, 1812, 1839, 1840, 1862, 1864, 1875, 1900, 1914, 1935, 1947, 1971, 1972, 1974, 1984, 1998, 2010, 2013, 2020, 2022, 2030, 2033, 2035, 2038, 2040, 2043, 2045, 2049, 2065, 2068, 2099, 2100, 2103, 2105, 2107, 2111, 2119, 2121, 2126, 2135, 2144, 2160, 2161, 2170, 2179, 2190, 2191, 2196, 2200, 2222, 2251, 2260, 2288, 2301, 2323, 2366, 2381, 2383, 2393, 2394, 2399, 2401, 2492, 2500, 2522, 2525, 2557, 2601, 2602, 2604, 2605, 2607, 2608, 2638, 2701, 2702, 2710, 2717, 2718, 2725, 2800, 2809, 2811, 2869, 2875, 2909, 2910, 2920, 2967, 2968, 2998, 3000, 3001, 3003, 3005, 3007, 3011, 3013, 3017, 3030, 3031, 3052, 3071, 3077, 3128, 3168, 3211, 3221, 3260, 3261, 3268, 3269, 3283, 3300, 3301, 3306, 3322, 3325, 3333, 3351, 3367, 3369, 3372, 3389, 3390, 3404, 3476, 3493, 3517, 3527, 3546, 3551, 3580, 3659, 3689, 3690, 3703, 3737, 3766, 3784, 3800, 3801, 3809, 3814, 3826, 3828, 3851, 3869, 3871, 3878, 3880, 3889, 3905, 3914, 3918, 3920, 3945, 3971, 3986, 3995, 3998, 4000, 4006, 4045, 4111, 4125, 4126, 4129, 4224, 4242, 4279, 4321, 4343, 4443, 4446, 4449, 4550, 4567, 4662, 4848, 4899, 4900, 4998, 5000, 5004, 5009, 5030, 5033, 5050, 5051, 5054, 5060, 5061, 5080, 5087, 5100, 5102, 5120, 5190, 5200, 5214, 5221, 5222, 5225, 5226, 5269, 5280, 5298, 5357, 5405, 5414, 5431, 5432, 5440, 5500, 5510, 5544, 5550, 5555, 5560, 5566, 5631, 5633, 5666, 5678, 5679, 5718, 5730, 5800, 5802, 5810, 5811, 5815, 5822, 5825, 5850, 5859, 5862, 5877, 5900, 5904, 5906, 5907, 5910, 5911, 5915, 5922, 5925, 5950, 5952, 5959, 5963, 5987, 5989, 5998, 6007, 6009, 6025, 6059, 6100, 6101, 6106, 6112, 6123, 6129, 6156, 6346, 6389, 6502, 6510, 6543, 6547, 6565, 6567, 6580, 6646, 6666, 6669, 6689, 6692, 6699, 6779, 6788, 6789, 6792, 6839, 6881, 6901, 6969, 7000, 7002, 7004, 7007, 7019, 7025, 7070, 7100, 7103, 7106, 7200, 7201, 7402, 7435, 7443, 7496, 7512, 7625, 7627, 7676, 7741, 7777, 7778, 7800, 7911, 7920, 7921, 7937, 7938, 7999, 8002, 8007, 8011, 8021, 8022, 8031, 8042, 8045, 8080, 8090, 8093, 8099, 8100, 8180, 8181, 8192, 8194, 8200, 8222, 8254, 8290, 8292, 8300, 8333, 8383, 8400, 8402, 8443, 8500, 8600, 8649, 8651, 8652, 8654, 8701, 8800, 8873, 8888, 8899, 8994, 9000, 9003, 9009, 9011, 9040, 9050, 9071, 9080, 9081, 9090, 9091, 9099, 9103, 9110, 9111, 9200, 9207, 9220, 9290, 9415, 9418, 9485, 9500, 9502, 9503, 9535, 9575, 9593, 9595, 9618, 9666, 9876, 9878, 9898, 9900, 9917, 9929, 9943, 9944, 9968, 9998, 10004, 10009, 10010, 10012, 10024, 10025, 10082, 10180, 10215, 10243, 10566, 10616, 10617, 10621, 10626, 10628, 10629, 10778, 11110, 11111, 11967, 12000, 12174, 12265, 12345, 13456, 13722, 13782, 13783, 14000, 14238, 14441, 14442, 15000, 15002, 15004, 15660, 15742, 16000, 16001, 16012, 16016, 16018, 16080, 16113, 16992, 16993, 17877, 17988, 18040, 18101, 18988, 19101, 19283, 19315, 19350, 19780, 19801, 19842, 20000, 20005, 20031, 20221, 20222, 20828, 21571, 22939, 23502, 24444, 24800, 25734, 25735, 26214, 27000, 27352, 27353, 27355, 27356, 27715, 28201, 30000, 30718, 30951, 31038, 31337, 32768, 32785, 33354, 33899, 34571, 34573, 35500, 38292, 40193, 40911, 41511, 42510, 44176, 44442, 44443, 44501, 45100, 48080, 49152, 49161, 49163, 49165, 49167, 49175, 49176, 49400, 49999, 50003, 50006, 50300, 50389, 50500, 50636, 50800, 51103, 51493, 52673, 52822, 52848, 52869, 54045, 54328, 55055, 55056, 55555, 55600, 56737, 56738, 57294, 57797, 58080, 60020, 60443, 61532, 61900, 62078, 63331, 64623, 64680, 65000, 65129, 65389, 280, 4567, 7001, 8008, 9080}

func ScanPort(hostname string, port int, dialtime int) bool {
	address := fmt.Sprintf("%v:%v", hostname, port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(dialtime)*time.Second)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(60 * time.Second)
			ScanPort(hostname, port, dialtime)
		} else {
			return false
		}
	}
	defer conn.Close()

	return true
}

func ScanPorts(ports []int, hostname string, dialtime int, c func(port int, isOpen bool)) {
	for _, port := range ports {
		c(port, ScanPort(hostname, port, dialtime))
	}
}

func ParsePortString(s string) ([]int, error) {
	splitted := strings.Split(s, ",")
	var ports []int

	for _, p := range splitted {
		n, err := strconv.Atoi(p)
		if err != nil {
			return []int{}, err
		}
		ports = append(ports, n)
	}

	return ports, nil
}

func PortService(port int) string {
	switch port {
	case 7:
		return "echo"
	case 19:
		return "chargen"
	case 20:
		return "ftp"
	case 21:
		return "ftp"
	case 22:
		return "ssh"
	case 23:
		return "telnet"
	case 25:
		return "smtp"
	case 43:
		return "whois"
	case 49:
		return "tacacs"
	case 53:
		return "dns"
	case 67:
		return "dhcp/bootp"
	case 68:
		return "dhcp/bootp"
	case 69:
		return "tftp"
	case 70:
		return "gopher"
	case 79:
		return "finger"
	case 80:
		return "http"
	case 88:
		return "kerberos"
	case 102:
		return "ms exchange"
	case 110:
		return "pop3"
	case 111:
		return "rpcbind"
	case 113:
		return "ident"
	case 119:
		return "nntp"
	case 123:
		return "ntp"
	case 135:
		return "microsoft-rpc"
	case 139:
		return "netbios-ssn"
	case 143:
		return "imap4"
	case 161:
		return "snmp"
	case 162:
		return "snmp"
	case 177:
		return "xdmcp"
	case 179:
		return "bgp"
	case 201:
		return "appletalk"
	case 264:
		return "bgmp"
	case 318:
		return "tsp"
	case 319:
		return "ldap"
	case 411:
		return "direct connect"
	case 412:
		return "direct connect"
	case 443:
		return "https"
	case 445:
		return "microsoft-ds"
	case 464:
		return "kerberos"
	case 465:
		return "smtp over ssl"
	case 497:
		return "retrospect"
	case 500:
		return "isakmp"
	case 512:
		return "rexec"
	case 513:
		return "login"
	case 514:
		return "syslong"
	case 515:
		return "ldp/lpr"
	case 520:
		return "rip"
	case 521:
		return "ripng (ipv6)"
	case 540:
		return "uucp"
	case 554:
		return "rtsp"
	case 546:
		return "dhcpv6"
	case 547:
		return "dhcpv6"
	case 560:
		return "rmonitor"
	case 563:
		return "nntp over ssl"
	case 587:
		return "smtp"
	case 591:
		return "filemaker"
	case 593:
		return "microsoft dcom"
	case 631:
		return "internet printing"
	case 636:
		return "dlap over ssl"
	case 639:
		return "msdp (pim)"
	case 646:
		return "ldp (mpls)"
	case 691:
		return "ms exchange"
	case 860:
		return "iscsi"
	case 873:
		return "rsync"
	case 902:
		return "vmware server"
	case 989:
		return "ftp over ssl"
	case 990:
		return "ftp over ssl"
	case 993:
		return "imap4 over ssl"
	case 995:
		return "pop3 over ssl"
	case 1025:
		return "microsoft rpc"
	case 1080:
		return "socks proxy"
	case 1194:
		return "openvpn"
	case 1214:
		return "kazaa"
	case 1241:
		return "nessus"
	case 1311:
		return "deil openmanage"
	case 1337:
		return "waste"
	case 1433:
		return "microsoft sql"
	case 1434:
		return "microsoft sql"
	case 1512:
		return "wins"
	case 1589:
		return "cisco vqp"
	case 1701:
		return "l2tp"
	case 1723:
		return "ms pptp"
	case 1725:
		return "steam"
	case 1741:
		return "ciscoworks 2000"
	case 1755:
		return "ms media server"
	case 1812:
		return "radius"
	case 1813:
		return "radius"
	case 1863:
		return "msn"
	case 1985:
		return "cisco hsrp"
	case 2000:
		return "cisco scco"
	case 2002:
		return "cisco asc"
	case 2049:
		return "nfs"
	case 2082:
		return "cpanel"
	case 2083:
		return "cpanel ssl"
	case 2086:
		return "whm"
	case 2087:
		return "whm ssl"
	case 2095:
		return "webmail"
	case 2096:
		return "webmail ssl"
	case 2100:
		return "oracle xdb"
	case 2121:
		return "ccproxy-ftp"
	case 2222:
		return "directadmin"
	case 2302:
		return "halo"
	case 2483:
		return "oracle db"
	case 2484:
		return "oracle db"
	case 2745:
		return "bagle.h"
	case 2967:
		return "symantec av"
	case 3050:
		return "interbase db"
	case 3074:
		return "xbox live"
	case 3124:
		return "http proxy"
	case 3127:
		return "mydoom"
	case 3128:
		return "http proxy"
	case 3222:
		return "glbp"
	case 3260:
		return "iscsi target"
	case 3306:
		return "mysql"
	case 3389:
		return "terminal server"
	case 3689:
		return "itunes"
	case 3690:
		return "subversion"
	case 3724:
		return "world of warcraft"
	case 3784:
		return "ventrilo"
	case 3785:
		return "ventrilo"
	case 4333:
		return "msql"
	case 4444:
		return "blaster"
	case 4664:
		return "google desktop"
	case 4672:
		return "emule"
	case 4899:
		return "radmin"
	case 5000:
		return "upnp"
	case 5001:
		return "slingbox"
	case 5004:
		return "rtp"
	case 5005:
		return "rtp"
	case 5050:
		return "yahoo messenger"
	case 5060:
		return "sip"
	case 5190:
		return "aim/icq"
	case 5222:
		return "xmpp/jabber"
	case 5223:
		return "xmpp/jabber"
	case 5432:
		return "postgresql"
	case 5500:
		return "vnc server"
	case 5554:
		return "sasser"
	case 5800:
		return "vnc over http"
	case 5900:
		return "vnc server"
	case 6000:
		return "x11"
	case 6001:
		return "x11"
	case 6112:
		return "battle.net"
	case 6129:
		return "dameware"
	case 6257:
		return "winmx"
	case 6346:
		return "gnutella"
	case 6347:
		return "gnutella"
	case 6500:
		return "gamespy arcade"
	case 6566:
		return "sane"
	case 6588:
		return "analogx"
	case 6665:
		return "irc"
	case 6679:
		return "irc over ssl"
	case 6699:
		return "napster"
	case 6881:
		return "bittorrent"
	case 6999:
		return "bittorrent"
	case 6891:
		return "windows live"
	case 6970:
		return "quicktime"
	case 7212:
		return "ghostsurf"
	case 7648:
		return "cu-seeme"
	case 7649:
		return "cu-seeme"
	case 8000:
		return "internet radio"
	case 8080:
		return "http proxy"
	case 8086:
		return "kaspersky av"
	case 8087:
		return "kaspersky av"
	case 8118:
		return "privoxy"
	case 8200:
		return "vmware server"
	case 8500:
		return "adobe coldfusion"
	case 8767:
		return "teamspeak"
	case 8866:
		return "bagle.b"
	case 9100:
		return "hp jetdirect"
	case 9101:
		return "bacula"
	case 9102:
		return "bacula"
	case 9103:
		return "bacula"
	case 9119:
		return "mxit"
	case 9418:
		return "git"
	case 9800:
		return "webdav"
	case 9898:
		return "dabber"
	case 9988:
		return "rbot/spybot"
	case 9999:
		return "webmin"
	case 10000:
		return "webmin"
	case 10113:
		return "netlq"
	case 11371:
		return "openpgp"
	case 12035:
		return "second life"
	case 12036:
		return "second life"
	case 12345:
		return "netbus"
	case 13720:
		return "netbackup"
	case 13721:
		return "netbackup"
	case 14567:
		return "battlefield"
	case 15118:
		return "dipnet/oddbob"
	case 19226:
		return "adminsecure"
	case 19638:
		return "ensim"
	case 20000:
		return "usermin"
	case 24800:
		return "synergy"
	case 25999:
		return "xfine"
	case 27015:
		return "half-life"
	case 27374:
		return "sub7"
	case 28960:
		return "call of duty"
	case 31337:
		return "back orifice"
	case 33434:
		return "traceroute"
	}

	return "unknown"
}
