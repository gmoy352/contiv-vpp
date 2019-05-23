/*REMEMBER TO USE THIS MODEL WHEN BRINGING OVER THE PLUG-IN*/

package model

import (
	"github.com/contiv/vpp/plugins/ksr/model/ksrkey"
)

/*const ModuleName = "bgp"

var (
	ModelBgpGlobal = models.Register(&GlobalConf{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "global",
	})
	ModelBgpPeer = models.Register(&PeerConf{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "peers",
	}, models.WithNameTemplate("{{.Name}}"))
)

//given the peer name, this function will return the key
func PeerKey(name string) string {
	return models.Key(&PeerConf{
		Name: name,
	})
}*/



// Keyword defines the keyword identifying NodeConfig data.
const Keyword = "bgpconfig"

// KeyPrefix return prefix where all node configs are persisted.
func KeyPrefix() string {
	return ksrkey.KsrK8sPrefix + "/" + Keyword + "/"
}

// Key returns the key for configuration of a given node.
func Key(conf string) string {
	return KeyPrefix() + conf
}


/*func GlobalKey(name string) string {
	return models.Key(&GlobalConf{
		Name: name,
	})
}*/
