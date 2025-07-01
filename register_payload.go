package main

type InstancePayload struct {
	Instance RegisterPayload `json:"instance"`
}

type RegisterPayload struct {
	HostName       string         `json:"hostName"`
	App            string         `json:"app"`
	IpAddr         string         `json:"ipAddr"`
	VipAddress     string         `json: "vipAddress"`
	Status         string         `json:"status"`
	Port           Port           `json:"port"`
	DataCenterInfo DataCenterInfo `json:"dataCenterInfo,omitempty"`
}

type Port struct {
	Port    int    `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}
