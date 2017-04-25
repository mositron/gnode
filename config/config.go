package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	Name       string
	Host       string
	Location   string
	PortWeb    int
	PortMaster int
	PortCam    int
	Ip         string
	IpMaster         string
	//is_master   bool
)

func LoadConf() {
	//f2, _ := ioutil.ReadFile(f)
	//var dat map[string]interface{}
	//if err := json.Unmarshal(f2, &dat); err != nil {
	//	panic(err)
	//}
	host, err := os.Hostname()
	if err == nil {
		Host = host
	}
  Name = "GNode/0.1"
  IpMaster := os.Getenv("GNODE_IP")
  PortMaster := os.Getenv("GNODE_PORT")
  PortWeb := os.Getenv("GNODE_WEB")
  Ip, _ = server.ExternalIP()

}
