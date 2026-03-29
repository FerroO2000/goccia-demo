package internal

type PingEvent struct {
	SrcIP uint32 `json:"src_ip"`
	DstIP uint32 `json:"dst_ip"`
	ID    uint16 `json:"id"`
	Seq   uint16 `json:"seq"`
}
