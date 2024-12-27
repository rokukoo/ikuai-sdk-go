package ikuai

import "github.com/duke-git/lancet/v2/slice"

const (
	dNatFuncName = "dnat"
)

type DNatParam struct {
	Id int `json:"id"`
	AddDNatParam
}

type AddDNatParam struct {
	Enabled   bool   `json:"enabled"`
	LanAddr   string `json:"lan_addr"`
	LanPort   string `json:"lan_port"`
	WanPort   string `json:"wan_port"`
	Interface string `json:"interface"`
	Protocol  string `json:"protocol"`
	SrcAddr   string `json:"src_addr"`
	Comment   string `json:"comment"`
}

const (
	ProtocolTcp = "tcp"
	ProtocolUdp = "udp"
	ProtocolAll = "tcp+udp"
)

func GetAddDNatParam(lanAddr string, lanPort string, wanPort string, comment string) *AddDNatParam {
	return &AddDNatParam{
		Enabled:   true,
		LanAddr:   lanAddr,
		LanPort:   lanPort,
		WanPort:   wanPort,
		Interface: "all",
		Protocol:  ProtocolAll,
		SrcAddr:   "",
		Comment:   comment,
	}
}

// AddDNat Add DNat
func AddDNat(params *AddDNatParam) (int, error) {
	res, err := CallAction[any, AddDNatParam](dNatFuncName, "add", *params)
	if err != nil {
		return -1, err
	}
	return *res.RowId, nil
}

// DelDNat Del DNat
func DelDNat(ids ...int) (bool, error) {
	id := slice.Join(ids, ",")
	_, err := CallAction[any, map[string]string](dNatFuncName, "del", map[string]string{"id": id})
	if err != nil {
		return false, err
	}
	return true, nil
}

type EditDNatParam struct {
	DNatParam
}

func GetEditDNatParam(params DNatParam) *EditDNatParam {
	return &EditDNatParam{
		DNatParam: params,
	}
}

// EditDNat Edit DNat
func EditDNat(params *EditDNatParam) (bool, error) {
	_, err := CallAction[any, EditDNatParam](dNatFuncName, "add", *params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DownDNat Down DNat
func DownDNat(ids ...int) (bool, error) {
	id := slice.Join(ids, ",")
	_, err := CallAction[any, map[string]string](dNatFuncName, "down", map[string]string{"id": id})
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpDNat Up DNat
func UpDNat(ids ...int) (bool, error) {
	id := slice.Join(ids, ",")
	_, err := CallAction[any, map[string]string](dNatFuncName, "up", map[string]string{"id": id})
	if err != nil {
		return false, err
	}
	return true, nil
}

type ShowDNatParam struct {
	Type     string `json:"TYPE"`
	Limit    string `json:"limit"`
	OrderBy  string `json:"ORDER_BY"`
	Order    string `json:"ORDER"`
	Finds    string `json:"FINDS"`
	Keywords string `json:"KEYWORDS"`
	Filter1  string `json:"FILTER1"`
	Filter2  string `json:"FILTER2"`
	Filter3  string `json:"FILTER3"`
	Filter4  string `json:"FILTER4"`
	Filter5  string `json:"FILTER5"`
}

func GetShowDNatParam(keywords string) *ShowDNatParam {
	return &ShowDNatParam{
		Type:     "total,data",
		Finds:    "lan_addr,lan_port,wan_port,comment",
		Keywords: keywords,
	}
}

// ShowDNat Show DNat
func ShowDNat(params *ShowDNatParam) ([]DNatParam, error) {
	type result struct {
		Data  []DNatParam `json:"data"`
		Total int         `json:"total"`
	}
	res, err := CallAction[result, ShowDNatParam](dNatFuncName, "show", *params)
	if err != nil {
		return nil, err
	}
	return res.Data.Data, nil
}
