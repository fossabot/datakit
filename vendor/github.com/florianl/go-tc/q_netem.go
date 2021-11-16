package tc

import (
	"fmt"

	"github.com/mdlayher/netlink"
)

const (
	tcaNetemUnspec = iota
	tcaNetemCorr
	tcaNetemDelayDist
	tcaNetemReorder
	tcaNetemCorrupt
	tcaNetemLoss
	tcaNetemRate
	tcaNetemEcn
	tcaNetemRate64
	tcaNetemPad
	tcaNetemLatency64
	tcaNetemJitter64
	tcaNetemSlot
	tcaNetemSlotDist
)

// Netem contains attributes of the netem discipline
type Netem struct {
	Qopt      NetemQopt
	Corr      *NetemCorr
	Reorder   *NetemReorder
	Corrupt   *NetemCorrupt
	Rate      *NetemRate
	Ecn       *uint32
	Rate64    *uint64
	Latency64 *int64
	Jitter64  *int64
	Slot      *NetemSlot
}

// NetemQopt from include/uapi/linux/pkt_sched.h
type NetemQopt struct {
	Latency   uint32
	Limit     uint32
	Loss      uint32
	Gap       uint32
	Duplicate uint32
	Jitter    uint32
}

// NetemCorr from include/uapi/linux/pkt_sched.h
type NetemCorr struct {
	Delay uint32
	Loss  uint32
	Dup   uint32
}

// NetemReorder from include/uapi/linux/pkt_sched.h
type NetemReorder struct {
	Probability uint32
	Correlation uint32
}

// NetemCorrupt from include/uapi/linux/pkt_sched.h
type NetemCorrupt struct {
	Probability uint32
	Correlation uint32
}

// NetemRate from include/uapi/linux/pkt_sched.h
type NetemRate struct {
	Rate           uint32
	PacketOverhead int32
	CellSize       int32
	CellOverhead   int32
}

// NetemSlot from include/uapi/linux/pkt_sched.h
type NetemSlot struct {
	MinDelay   int64
	MaxDelay   int64
	MaxPackets int32
	MaxBytes   int32
	DistDelay  int64
	DistJitter int64
}

// unmarshalNetem parses the Netem-encoded data and stores the result in the value pointed to by info.
func unmarshalNetem(data []byte, info *Netem) error {
	qopt := NetemQopt{}
	if err := unmarshalStruct(data, &qopt); err != nil {
		return err
	}
	info.Qopt = qopt

	// continue decoding attributes after the size of the NetemQopt struct
	ad, err := netlink.NewAttributeDecoder(data[24:])
	if err != nil {
		return err
	}
	ad.ByteOrder = nativeEndian
	for ad.Next() {
		switch ad.Type() {
		case tcaNetemCorr:
			tmp := &NetemCorr{}
			if err := unmarshalStruct(ad.Bytes(), tmp); err != nil {
				return err
			}
			info.Corr = tmp
		case tcaNetemReorder:
			tmp := &NetemReorder{}
			if err := unmarshalStruct(ad.Bytes(), tmp); err != nil {
				return err
			}
			info.Reorder = tmp
		case tcaNetemCorrupt:
			tmp := &NetemCorrupt{}
			if err := unmarshalStruct(ad.Bytes(), tmp); err != nil {
				return err
			}
			info.Corrupt = tmp
		case tcaNetemRate:
			tmp := &NetemRate{}
			if err := unmarshalStruct(ad.Bytes(), tmp); err != nil {
				return err
			}
			info.Rate = tmp
		case tcaNetemEcn:
			tmp := ad.Uint32()
			info.Ecn = &tmp
		case tcaNetemRate64:
			tmp := ad.Uint64()
			info.Rate64 = &tmp
		case tcaNetemLatency64:
			var val int64
			if err := unmarshalNetlinkAttribute(ad.Bytes(), &val); err != nil {
				return err
			}
			info.Latency64 = &val
		case tcaNetemJitter64:
			var val int64
			if err := unmarshalNetlinkAttribute(ad.Bytes(), &val); err != nil {
				return err
			}
			info.Jitter64 = &val
		case tcaNetemSlot:
			tmp := &NetemSlot{}
			if err := unmarshalStruct(ad.Bytes(), tmp); err != nil {
				return err
			}
			info.Slot = tmp
		default:
			return fmt.Errorf("unmarshalNetem()\t%d\n\t%v", ad.Type(), ad.Bytes())
		}
	}
	return nil
}

// marshalNetem returns the binary encoding of Qfq
func marshalNetem(info *Netem) ([]byte, error) {
	options := []tcOption{}

	if info == nil {
		return []byte{}, fmt.Errorf("Netem: %w", ErrNoArg)
	}

	if info.Corr != nil {
		data, err := marshalStruct(info.Corr)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaNetemCorr, Data: data})
	}
	if info.Reorder != nil {
		data, err := marshalStruct(info.Reorder)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaNetemReorder, Data: data})
	}
	if info.Corrupt != nil {
		data, err := marshalStruct(info.Corrupt)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaNetemCorrupt, Data: data})
	}
	if info.Rate != nil {
		data, err := marshalStruct(info.Rate)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaNetemRate, Data: data})
	}
	if info.Ecn != nil {
		options = append(options, tcOption{Interpretation: vtUint32, Type: tcaNetemEcn, Data: *info.Ecn})
	}
	if info.Rate64 != nil {
		options = append(options, tcOption{Interpretation: vtUint64, Type: tcaNetemRate64, Data: *info.Rate64})
	}
	if info.Latency64 != nil {
		options = append(options, tcOption{Interpretation: vtInt64, Type: tcaNetemLatency64, Data: *info.Latency64})
	}
	if info.Jitter64 != nil {
		options = append(options, tcOption{Interpretation: vtInt64, Type: tcaNetemJitter64, Data: *info.Jitter64})
	}
	if info.Slot != nil {
		data, err := marshalStruct(info.Slot)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaNetemSlot, Data: data})
	}

	data, err := marshalAttributes(options)

	var qoptErr error
	var qoptData []byte
	if qoptData, qoptErr = marshalStruct(info.Qopt); qoptErr != nil {
		return []byte{}, err
	}

	return append(qoptData[:], data[:]...), err
}
