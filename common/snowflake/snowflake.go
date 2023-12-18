package snowflake

import (
	"errors"
	"math/rand"
	"net"
	"time"

	"github.com/sony/sonyflake"
)

var defaultSonyflake *sonyflake.Sonyflake

func init() {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2021-04-09 00:00:00")
	setting := sonyflake.Settings{ //使用默认的内网ip转换为machineID
		StartTime: startTime,
		MachineID: machineID,
		//CheckMachineID: checkMachineID,
	}

	defaultSonyflake = sonyflake.NewSonyflake(setting)
	if defaultSonyflake == nil {
		panic("init sonyflake failed")
	}
}

func NextUID() uint64 {
	id, _ := defaultSonyflake.NextID()
	return id
}

// 优先使用内网ip后16位，其次使用随机值，如果服务实例数很大，可以考虑改用zookeeper或者redis生成唯一机器id
func machineID() (uint16, error) {
	r, err := lower16BitPrivateIP()
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		r = uint16(rand.Intn(65536))
	}

	return r, nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
