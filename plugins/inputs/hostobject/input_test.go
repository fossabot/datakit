// +build !noconntrack

package hostobject

import (
	"fmt"
	"testing"

	"github.com/prometheus/procfs"
)

/*
type conntrackCollector struct {
	current       *prometheus.Desc
	limit         *prometheus.Desc
	found         *prometheus.Desc
	invalid       *prometheus.Desc
	ignore        *prometheus.Desc
	insert        *prometheus.Desc
	insertFailed  *prometheus.Desc
	drop          *prometheus.Desc
	earlyDrop     *prometheus.Desc
	searchRestart *prometheus.Desc
}

type conntrackStatistics struct {
	found         uint64 // Number of searched entries which were successful
	invalid       uint64 // Number of packets seen which can not be tracked
	ignore        uint64 // Number of packets seen which are already connected to a conntrack entry
	insert        uint64 // Number of entries inserted into the list
	insertFailed  uint64 // Number of entries for which list insertion was attempted but failed (happens if the same entry is already present)
	drop          uint64 // Number of packets dropped due to conntrack failure. Either new conntrack entry allocation failed, or protocol helper dropped the packet
	earlyDrop     uint64 // Number of dropped conntrack entries to make room for new ones, if maximum table size was reached
	searchRestart uint64 // Number of conntrack table lookups which had to be restarted due to hashtable resizes
}

// NewConntrackCollector returns a new Collector exposing conntrack stats.
func NewConntrackCollector() (*conntrackCollector, error) {
	return &conntrackCollector{}, nil
}

func (c *conntrackCollector) Update(ch chan<- prometheus.Metric) error {
	value, err := readUintFromFile(procFilePath("sys/net/netfilter/nf_conntrack_count"))
	if err != nil {
		return err
	}
	ch <- prometheus.MustNewConstMetric(
		c.current, prometheus.GaugeValue, float64(value))

	value, err = readUintFromFile(procFilePath("sys/net/netfilter/nf_conntrack_max"))
	if err != nil {
		return err
	}
	ch <- prometheus.MustNewConstMetric(
		c.limit, prometheus.GaugeValue, float64(value))

	conntrackStats, err := getConntrackStatistics()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.found, prometheus.GaugeValue, float64(conntrackStats.found))
	ch <- prometheus.MustNewConstMetric(
		c.invalid, prometheus.GaugeValue, float64(conntrackStats.invalid))
	ch <- prometheus.MustNewConstMetric(
		c.ignore, prometheus.GaugeValue, float64(conntrackStats.ignore))
	ch <- prometheus.MustNewConstMetric(
		c.insert, prometheus.GaugeValue, float64(conntrackStats.insert))
	ch <- prometheus.MustNewConstMetric(
		c.insertFailed, prometheus.GaugeValue, float64(conntrackStats.insertFailed))
	ch <- prometheus.MustNewConstMetric(
		c.drop, prometheus.GaugeValue, float64(conntrackStats.drop))
	ch <- prometheus.MustNewConstMetric(
		c.earlyDrop, prometheus.GaugeValue, float64(conntrackStats.earlyDrop))
	ch <- prometheus.MustNewConstMetric(
		c.searchRestart, prometheus.GaugeValue, float64(conntrackStats.searchRestart))
	return nil
}

func getConntrackStatistics() (*conntrackStatistics, error) {
	c := conntrackStatistics{}

	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, fmt.Errorf("failed to open procfs: %w", err)
	}

	connStats, err := fs.ConntrackStat()
	if err != nil {
		return nil, err
	}

	for _, connStat := range connStats {
		c.found += connStat.Found
		c.invalid += connStat.Invalid
		c.ignore += connStat.Ignore
		c.insert += connStat.Insert
		c.insertFailed += connStat.InsertFailed
		c.drop += connStat.Drop
		c.earlyDrop += connStat.EarlyDrop
		c.searchRestart += connStat.SearchRestart
	}

	return &c, nil
}
*/
func TestConn(t *testing.T) {
	value, err := readUintFromFile(procFilePath("netfilter/nf_conntrack_count"))
	fmt.Println(value, err)

	fs, err := procfs.NewFS("/tmp")
	if err != nil {
		fmt.Println(err)
	}

	connStats, err := fs.ConntrackStat()
	if err != nil {
		fmt.Println(err)
	}

	for _, connStat := range connStats {
		fmt.Println(connStat)
	}
}
