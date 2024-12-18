package ikuai

import (
	"github.com/duke-git/lancet/v2/slice"
)

const (
	monitorIfaceFuncName = "monitor_iface"
)

type ifaceStream struct {
	Interface   string `json:"interface"`
	Comment     string `json:"comment"`
	IpAddr      string `json:"ip_addr"`
	ConnectNum  string `json:"connect_num"`
	Upload      int    `json:"upload"`
	Download    int    `json:"download"`
	TotalUp     int64  `json:"total_up"`
	TotalDown   int64  `json:"total_down"`
	UpDropped   int    `json:"updropped"`
	DownDropped int    `json:"downdropped"`
	UpPacked    int    `json:"uppacked"`
	DownPacked  int    `json:"downpacked"`
}

type MonitorIfaceType string

const (
	MonitorIfaceTypeCheck  MonitorIfaceType = "iface_check"
	MonitorIfaceTypeStream                  = "iface_stream"
)

type ShowMonitorIfaceResult struct {
	IfaceStream []ifaceStream `json:"iface_stream"`
}

func ShowMonitorIface(types ...MonitorIfaceType) (*ShowMonitorIfaceResult, error) {
	Type := slice.Join(types, ",")
	res, err := CallAction[ShowMonitorIfaceResult, map[string]string](monitorIfaceFuncName, "show", map[string]string{
		"TYPE": Type,
	})
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
