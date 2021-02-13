package golang

func checksum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])<<8 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var ans uint16 = uint16(^sum)
	return ans
}
