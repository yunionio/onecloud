package netlink

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// ioctl for statistics.
const (
	// ETHTOOL_GSSET_INFO gets string set info
	ETHTOOL_GSSET_INFO = 0x00000037
	// SIOCETHTOOL is Ethtool interface
	SIOCETHTOOL = 0x8946
	// ETHTOOL_GSTRINGS gets specified string set
	ETHTOOL_GSTRINGS = 0x0000001b
	// ETHTOOL_GSTATS gets NIC-specific statistics
	ETHTOOL_GSTATS = 0x0000001d
)

// string set id.
const (
	// ETH_SS_TEST is self-test result names, for use with %ETHTOOL_TEST
	ETH_SS_TEST = iota
	// ETH_SS_STATS statistic names, for use with %ETHTOOL_GSTATS
	ETH_SS_STATS
	// ETH_SS_PRIV_FLAGS are driver private flag names
	ETH_SS_PRIV_FLAGS
	// _ETH_SS_NTUPLE_FILTERS is deprecated
	_ETH_SS_NTUPLE_FILTERS
	// ETH_SS_FEATURES are device feature names
	ETH_SS_FEATURES
	// ETH_SS_RSS_HASH_FUNCS is RSS hush function names
	ETH_SS_RSS_HASH_FUNCS
)

// IfreqSubordinate is a struct for ioctl bond manipulation syscalls.
// It is used to assign subordinate to bond interface with Name.
type IfreqSubordinate struct {
	Name  [unix.IFNAMSIZ]byte
	Subordinate [unix.IFNAMSIZ]byte
}

// Ifreq is a struct for ioctl ethernet manipulation syscalls.
type Ifreq struct {
	Name [unix.IFNAMSIZ]byte
	Data uintptr
}

// ethtoolSset is a string set information
type ethtoolSset struct {
	cmd      uint32
	reserved uint32
	mask     uint64
	data     [1]uint32
}

// ethtoolGstrings is string set for data tagging
type ethtoolGstrings struct {
	cmd       uint32
	stringSet uint32
	length    uint32
	data      [32]byte
}

type ethtoolStats struct {
	cmd    uint32
	nStats uint32
	data   [1]uint64
}

// newIocltSubordinateReq returns filled IfreqSubordinate with proper interface names
// It is used by ioctl to assign subordinate to bond main
func newIocltSubordinateReq(subordinate, main string) *IfreqSubordinate {
	ifreq := &IfreqSubordinate{}
	copy(ifreq.Name[:unix.IFNAMSIZ-1], main)
	copy(ifreq.Subordinate[:unix.IFNAMSIZ-1], subordinate)
	return ifreq
}

// newIocltStringSetReq creates request to get interface string set
func newIocltStringSetReq(linkName string) (*Ifreq, *ethtoolSset) {
	e := &ethtoolSset{
		cmd:  ETHTOOL_GSSET_INFO,
		mask: 1 << ETH_SS_STATS,
	}

	ifreq := &Ifreq{Data: uintptr(unsafe.Pointer(e))}
	copy(ifreq.Name[:unix.IFNAMSIZ-1], linkName)
	return ifreq, e
}

// getSocketUDP returns file descriptor to new UDP socket
// It is used for communication with ioctl interface.
func getSocketUDP() (int, error) {
	return syscall.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
}
