package aqua

type Audit struct {
	Action         string `json:"action"`
	Category       string `json:"category"`
	Command        string `json:"command"`
	Container      string `json:"container"`
	Containerid    string `json:"containerid"`
	Control        string `json:"control"`
	Euid           string `json:"euid"`
	Euser          string `json:"euser"`
	Host           string `json:"host"`
	Hostgroup      string `json:"hostgroup"`
	Hostid         string `json:"hostid"`
	Hostip         string `json:"hostip"`
	Image          string `json:"image"`
	Imageid        string `json:"imageid"`
	K8SCluster     string `json:"k8s_cluster"`
	Level          string `json:"level"`
	Pid            int    `json:"pid"`
	Poddeployment  string `json:"poddeployment"`
	Podname        string `json:"podname"`
	Podnamespace   string `json:"podnamespace"`
	Podtype        string `json:"podtype"`
	Process        string `json:"process"`
	Reason         string `json:"reason"`
	Resource       string `json:"resource"`
	ResourceDigest string `json:"resource_digest"`
	Result         int    `json:"result"`
	Rule           string `json:"rule"`
	RuleType       string `json:"rule_type"`
	Subtype        string `json:"subtype"`
	Time           int    `json:"time"`
	UID            string `json:"uid"`
	User           string `json:"user"`
	VMGroup        string `json:"vm_group"`
	VMID           string `json:"vm_id"`
	VMLocation     string `json:"vm_location"`
	VMName         string `json:"vm_name"`
}
