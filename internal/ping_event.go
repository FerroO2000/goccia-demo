package internal

import "net"

type PingEvent struct {
	SrcIP uint32 `json:"src_ip"`
	DstIP uint32 `json:"dst_ip"`
	ID    uint16 `json:"id"`
	Seq   uint16 `json:"seq"`
}

func (pe *PingEvent) getIP(ip uint32) net.IP {
	return net.IPv4(
		byte(ip),
		byte(ip>>8),
		byte(ip>>16),
		byte(ip>>24),
	)
}

func (pe *PingEvent) GetSrcIP() net.IP {
	return pe.getIP(pe.SrcIP)
}

func (pe *PingEvent) GetDstIP() net.IP {
	return pe.getIP(pe.DstIP)
}
