package ibdiagnet2

import (
	"fmt"
	"infiniband_exporter/global"
	"infiniband_exporter/log"
	"infiniband_exporter/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	pmLabels      = util.GetFieldNames(Pm{})
	linkDownGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_link_down_counter",
			Help: "link_down_counter",
		},
		pmLabels,
	)

	linkErrorRecoveryGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_link_error_recovery_counter",
			Help: "link_error_recovery_counter",
		},
		pmLabels,
	)
	symbolErrorCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_symbol_error_counter",
			Help: "symbol_error_counter",
		},
		pmLabels,
	)
	portRcvRemotePhysicalErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_remote_physical_errors",
			Help: "port_rcv_remote_physical_errors",
		},
		pmLabels,
	)

	portRcvErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_errors",
			Help: "port_rcv_errors",
		},
		pmLabels,
	)

	portXmitDiscard = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_discard",
			Help: "port_xmit_discard",
		},
		pmLabels,
	)

	portRcvSwitchRelayErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_switch_relay_errors",
			Help: "port_rcv_switch_relay_errors",
		},
		pmLabels,
	)

	excessiveBufferErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_excessive_buffer_errors",
			Help: "excessive_buffer_errors",
		},
		pmLabels,
	)

	localLinkIntegrityErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_local_link_integrity_errors",
			Help: "local_link_integrity_errors",
		},
		pmLabels,
	)

	portRcvConstraintErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_constraint_errors",
			Help: "port_rcv_constraint_errors",
		},
		pmLabels,
	)

	portXmitConstraintErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_constraint_errors",
			Help: "port_xmit_constraint_errors",
		},
		pmLabels,
	)

	vl15Dropped = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_vl15_dropped",
			Help: "vl15_dropped",
		},
		pmLabels,
	)
	portXmitData = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_data",
			Help: "port_xmit_data",
		},
		pmLabels,
	)
	portRcvData = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_data",
			Help: "port_rcv_data",
		},
		pmLabels,
	)
	portXmitPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_pkts",
			Help: "port_xmit_pkts",
		},
		pmLabels,
	)
	portRcvPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_pkts",
			Help: "port_rcv_pkts",
		},
		pmLabels,
	)
	portXmitWait = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_wait",
			Help: "port_xmit_wait",
		},
		pmLabels,
	)
	portXmitDataExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_data_extended",
			Help: "port_xmit_data_extended",
		},
		pmLabels,
	)
	portRcvDataExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_data_extended",
			Help: "port_rcv_data_extended",
		},
		pmLabels,
	)
	portXmitPktsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_pkts_extended",
			Help: "port_xmit_pkts_extended",
		},
		pmLabels,
	)
	portRcvPktsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_pkts_extended",
			Help: "port_rcv_pkts_extended",
		},
		pmLabels,
	)
	portUnicastXmitPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_unicast_xmit_pkts",
			Help: "port_unicast_xmit_pkts",
		},
		pmLabels,
	)
	portUnicastRcvPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_unicast_rcv_pkts",
			Help: "port_unicast_rcv_pkts",
		},
		pmLabels,
	)
	portMulticastXmitPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_multicast_xmit_pkts",
			Help: "port_multicast_xmit_pkts",
		},
		pmLabels,
	)
	portMulticastRcvPkts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_multicast_rcv_pkts",
			Help: "port_multicast_rcv_pkts",
		},
		pmLabels,
	)
	symbolErrorCounterExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_symbol_error_counter_extended",
			Help: "symbol_error_counter_extended",
		},
		pmLabels,
	)
	linkErrorRecoveryCounterExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_link_error_recovery_counter_extended",
			Help: "link_error_recovery_counter_extended",
		},
		pmLabels,
	)
	linkDownedCounterExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_link_downed_counter_extended",
			Help: "link_downed_counter_extended",
		},
		pmLabels,
	)
	portRcvErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_errors_extended",
			Help: "port_rcv_errors_extended",
		},
		pmLabels,
	)
	portRcvRemotePhysicalErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_remote_physical_errors_extended",
			Help: "port_rcv_remote_physical_errors_extended",
		},
		pmLabels,
	)
	portRcvSwitchRelayErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_switch_relay_errors_extended",
			Help: "port_rcv_switch_relay_errors_extended",
		},
		pmLabels,
	)
	portXmitDiscardsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_discards_extended",
			Help: "port_xmit_discards_extended",
		},
		pmLabels,
	)
	portXmitConstraintErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_constraint_errors_extended",
			Help: "port_xmit_constraint_errors_extended",
		},
		pmLabels,
	)
	portRcvConstraintErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_rcv_constraint_errors_extended",
			Help: "port_rcv_constraint_errors_extended",
		},
		pmLabels,
	)
	localLinkIntegrityErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_local_link_integrity_errors_extended",
			Help: "local_link_integrity_errors_extended",
		},
		pmLabels,
	)
	excessiveBufferOverrunErrorsExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_excessive_buffer_overrun_errors_extended",
			Help: "excessive_buffer_overrun_errors_extended",
		},
		pmLabels,
	)
	vl15DroppedExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_vl15_dropped_extended",
			Help: "vl15_dropped_extended",
		},
		pmLabels,
	)
	portXmitWaitExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_xmit_wait_extended",
			Help: "port_xmit_wait_extended",
		},
		pmLabels,
	)
	qp1DroppedExtended = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_qp1_dropped_extended",
			Help: "qp1_dropped_extended",
		},
		pmLabels,
	)
	retransmissionPerSec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_retransmission_per_sec",
			Help: "retransmission_per_sec",
		},
		pmLabels,
	)
	maxRetransmissionRate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_max_retransmission_rate",
			Help: "max_retransmission_rate",
		},
		pmLabels,
	)
	portLocalPhysicalErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_local_physical_errors",
			Help: "port_local_physical_errors",
		},
		pmLabels,
	)
	portMalformedPacketErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_malformed_packet_errors",
			Help: "port_malformed_packet_errors",
		},
		pmLabels,
	)
	portBufferOverrunErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_buffer_overrun_errors",
			Help: "port_buffer_overrun_errors",
		},
		pmLabels,
	)
	portDlidMappingErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_dlid_mapping_errors",
			Help: "port_dlid_mapping_errors",
		},
		pmLabels,
	)
	portVlMappingErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_vl_mapping_errors",
			Help: "port_vl_mapping_errors",
		},
		pmLabels,
	)
	portLoopingErrors = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_looping_errors",
			Help: "port_looping_errors",
		},
		pmLabels,
	)
	portInactiveDiscards = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_inactive_discards",
			Help: "port_inactive_discards",
		},
		pmLabels,
	)
	portNeighborMtuDiscards = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_neighbor_mtu_discards",
			Help: "port_neighbor_mtu_discards",
		},
		pmLabels,
	)
	portSwLifetimeLimitDiscards = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_sw_lifetime_limit_discards",
			Help: "port_sw_lifetime_limit_discards",
		},
		pmLabels,
	)
	portSwHoqLifetimeLimitDiscards = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "infiniband_port_sw_hoq_lifetime_limit_discards",
			Help: "port_sw_hoq_lifetime_limit_discards",
		},
		pmLabels,
	)
)

type Pmer interface {
	UpdateMetrics()
}

type LinkPm struct {
	FilePath string
}

type Pm struct {
	//port       string
	//lid        string
	//guid       string
	//device     string
	//name       string
	remoteGuid string
	remoteName string
	remotePort string
	component  string
	localGuid  string
	localName  string
	localPort  string
}

func init() {
	metrics := []prometheus.Collector{
		linkDownGauge,
		linkErrorRecoveryGauge,
		symbolErrorCounter,
		portRcvRemotePhysicalErrors,
		portRcvErrors,
		portXmitDiscard,
		portRcvSwitchRelayErrors,
		excessiveBufferErrors,
		localLinkIntegrityErrors,
		portRcvConstraintErrors,
		portXmitConstraintErrors,
		vl15Dropped,
		portXmitData,
		portRcvData,
		portXmitPkts,
		portRcvPkts,
		portXmitWait,
		portXmitDataExtended,
		portRcvDataExtended,
		portXmitPktsExtended,
		portRcvPktsExtended,
		portUnicastXmitPkts,
		portUnicastRcvPkts,
		portMulticastXmitPkts,
		portMulticastRcvPkts,
		symbolErrorCounterExtended,
		linkErrorRecoveryCounterExtended,
		linkDownedCounterExtended,
		portRcvErrorsExtended,
		portRcvRemotePhysicalErrorsExtended,
		portRcvSwitchRelayErrorsExtended,
		portXmitDiscardsExtended,
		portXmitConstraintErrorsExtended,
		portRcvConstraintErrorsExtended,
		localLinkIntegrityErrorsExtended,
		excessiveBufferOverrunErrorsExtended,
		vl15DroppedExtended,
		portXmitWaitExtended,
		qp1DroppedExtended,
		retransmissionPerSec,
		maxRetransmissionRate,
		portLocalPhysicalErrors,
		portMalformedPacketErrors,
		portBufferOverrunErrors,
		portDlidMappingErrors,
		portVlMappingErrors,
		portLoopingErrors,
		portInactiveDiscards,
		portNeighborMtuDiscards,
		portSwLifetimeLimitDiscards,
		portSwHoqLifetimeLimitDiscards,
	}
	for _, metric := range metrics {
		prometheus.MustRegister(metric)
	}
}

func (p *LinkPm) UpdateMetrics() {
	blocks, err := util.GetContent(p.FilePath, `(?m)Port=(\d+)\sLid=(\w+)\sGUID=(\w{18})\sDevice=(\d+)\sPort\sName=(.*)`)
	if err != nil {
		log.GetLogger().Error("Get pm content error")
		return
	}
	for _, block := range *blocks {
		switchCaExpr := `Port=(\d+)\sLid=(\w+)\sGUID=(\w{18})\sDevice=(\d+)\sPort\sName=(.*)`
		switchCaMatch, err := regexp.Compile(switchCaExpr)
		if err != nil {
			log.GetLogger().Error("Switch or ca error compiling regex")
			break
		}
		subSwitchCaMatch := switchCaMatch.FindStringSubmatch(block)
		guid := subSwitchCaMatch[3]
		port := subSwitchCaMatch[1]
		name := subSwitchCaMatch[5]
		var remoteGuid, remoteName, remotePort, localGuid, localName, localPort string
		component := global.ComponentCa
		_, exists := util.GetKeysFromCache(guid)
		if exists {
			component = global.ComponentSw
			linkMap, exists := util.GetValueFromCache(fmt.Sprintf("%s_%s", guid, port))
			if exists {
				remoteGuid = guid
				remoteName = linkMap["remoteName"]
				remotePort = port
				localGuid = linkMap["localGuid"]
				localName = linkMap["localName"]
				localPort = linkMap["localPort"]
			}
		} else {
			for _, linkMap := range util.Cache {
				remoteGuid = guid
				remotePort = port
				if linkMap["localGuid"] == guid {
					remoteName = linkMap["remoteName"]
					localGuid = linkMap["localGuid"]
					localName = linkMap["localName"]
					localPort = linkMap["localPort"]
					break
				} else {
					remoteName = name
				}
			}
		}
		pm := Pm{
			remoteGuid: remoteGuid,
			remoteName: remoteName,
			remotePort: remotePort,
			component:  component,
			localGuid:  localGuid,
			localName:  localName,
			localPort:  localPort,
		}
		getValue := func(regexStr string) (value float64) {
			re := regexp.MustCompile(regexStr)
			match := re.FindStringSubmatch(block)
			if match == nil {
				return 0
			}
			metricValue := match[1]
			if metricValue == "NA" {
				return 0
			}
			numberStr := strings.Replace(metricValue, "0x", "", -1)
			numberStr = strings.Replace(numberStr, "0X", "", -1)
			dec, err := strconv.ParseInt(numberStr, 16, 64)
			if err != nil {
				log.GetLogger().Error(fmt.Sprintf("Parse error:%s", err))
				//TODO
				return 0
			}
			return float64(dec)
		}
		labelValues := []string{pm.remoteGuid, pm.remoteName, pm.remotePort, pm.component, pm.localGuid, pm.localName, pm.localPort}
		linkDownGauge.
			WithLabelValues(labelValues...).
			Set(getValue(`link_down_counter=(\w+)`))
		linkErrorRecoveryGauge.
			WithLabelValues(labelValues...).
			Set(getValue(`link_error_recovery_counter=(\w+)`))
		symbolErrorCounter.
			WithLabelValues(labelValues...).
			Set(getValue(`symbol_error_counter=(\w+)`))
		portRcvRemotePhysicalErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_remote_physical_errors=(\w+)`))
		portRcvErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_errors=(\w+)`))
		portXmitDiscard.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_discard=(\w+)`))
		portRcvSwitchRelayErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_switch_relay_errors=(\w+)`))
		excessiveBufferErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`excessive_buffer_errors=(\w+)`))
		localLinkIntegrityErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`local_link_integrity_errors=(\w+)`))
		portRcvConstraintErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_constraint_errors=(\w+)`))
		portXmitConstraintErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_constraint_errors=(\w+)`))
		vl15Dropped.
			WithLabelValues(labelValues...).
			Set(getValue(`vl15_dropped=(\w+)`))
		portXmitData.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_data=(\w+)`))
		portRcvData.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_data=(\w+)`))
		portXmitPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_pkts=(\w+)`))
		portRcvPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_pkts=(\w+)`))
		portXmitWait.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_wait=(\w+)`))
		portXmitDataExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_data_extended=(\w+)`))
		portRcvDataExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_data_extended=(\w+)`))
		portXmitPktsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_pkts_extended=(\w+)`))
		portRcvPktsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_pkts_extended=(\w+)`))
		portUnicastXmitPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_unicast_xmit_pkts=(\w+)`))
		portUnicastRcvPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_unicast_rcv_pkts=(\w+)`))
		portMulticastXmitPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_multicast_xmit_pkts=(\w+)`))
		portMulticastRcvPkts.
			WithLabelValues(labelValues...).
			Set(getValue(`port_multicast_rcv_pkts=(\w+)`))
		symbolErrorCounterExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`symbol_error_counter_extended=(\w+)`))
		linkErrorRecoveryCounterExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`link_error_recovery_counter_extended=(\w+)`))
		linkDownedCounterExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`link_downed_counter_extended=(\w+)`))
		portRcvErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_errors_extended=(\w+)`))
		portRcvRemotePhysicalErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_remote_physical_errors_extended=(\w+)`))
		portRcvSwitchRelayErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_switch_relay_errors_extended=(\w+)`))
		portXmitDiscardsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_discards_extended=(\w+)`))
		portXmitConstraintErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_constraint_errors_extended=(\w+)`))
		portRcvConstraintErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_rcv_constraint_errors_extended=(\w+)`))
		localLinkIntegrityErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`local_link_integrity_errors_extended=(\w+)`))
		excessiveBufferOverrunErrorsExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`excessive_buffer_overrun_errors_extended=(\w+)`))
		vl15DroppedExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`vl15_dropped_extended=(\w+)`))
		portXmitWaitExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`port_xmit_wait_extended=(\w+)`))
		qp1DroppedExtended.
			WithLabelValues(labelValues...).
			Set(getValue(`qp1_dropped_extended=(\w+)`))
		retransmissionPerSec.
			WithLabelValues(labelValues...).
			Set(getValue(`retransmission_per_sec=(\w+)`))
		maxRetransmissionRate.
			WithLabelValues(labelValues...).
			Set(getValue(`max_retransmission_rate=(\w+)`))
		portLocalPhysicalErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_local_physical_errors=(\w+)`))
		portMalformedPacketErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_malformed_packet_errors=(\w+)`))
		portBufferOverrunErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_buffer_overrun_errors=(\w+)`))
		portDlidMappingErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_dlid_mapping_errors=(\w+)`))
		portVlMappingErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_vl_mapping_errors=(\w+)`))
		portLoopingErrors.
			WithLabelValues(labelValues...).
			Set(getValue(`port_looping_errors=(\w+)`))
		portInactiveDiscards.
			WithLabelValues(labelValues...).
			Set(getValue(`port_inactive_discards=(\w+)`))
		portNeighborMtuDiscards.
			WithLabelValues(labelValues...).
			Set(getValue(`port_neighbor_mtu_discards=(\w+)`))
		portSwLifetimeLimitDiscards.
			WithLabelValues(labelValues...).
			Set(getValue(`port_sw_lifetime_limit_discards=(\w+)`))
		portSwHoqLifetimeLimitDiscards.
			WithLabelValues(labelValues...).
			Set(getValue(`port_sw_hoq_lifetime_limit_discards=(\w+)`))
	}
}
