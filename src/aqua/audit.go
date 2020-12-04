package aqua

// Audit a generic struct to store all messages from Aqua
type Audit struct {
	Action            string   `json:"action,omitempty"` // exec, login, policy.failure, create, accept, start, Scan,
	Adjective         string   `json:"adjective,omitempty"`
	ApplicationScopes []string `json:"application_scopes,omitempty"`
	Category          string   `json:"category,omitempty"` // file, os, network, secret, CVE, container, host, kubernetes.enforcer, image, KubernetesAssurancePolicy, Integration
	Command           string   `json:"command,omitempty"`
	Container         string   `json:"container,omitempty"`
	Containerid       string   `json:"containerid,omitempty"`
	Control           string   `json:"control,omitempty"`
	CreateTime        int      `json:"create_time,omitempty"`
	Critical          int      `json:"critical,omitempty"`
	Data              string   `json:"data,omitempty"`
	Date              int      `json:"date,omitempty"`
	Description       string   `json:"description,omitempty"`
	DtaSkipped        bool     `json:"dta_skipped,omitempty"`
	Euid              string   `json:"euid,omitempty"`
	Euser             string   `json:"euser,omitempty"`
	High              int      `json:"high,omitempty"`
	Host              string   `json:"host,omitempty"`
	Hostgroup         string   `json:"hostgroup,omitempty"`
	Hostid            string   `json:"hostid,omitempty"`
	Hostip            string   `json:"hostip,omitempty"`
	ID                int      `json:"id,omitempty"`
	Image             string   `json:"image,omitempty"`
	Imagehash         string   `json:"imagehash,omitempty"`
	Imageid           string   `json:"imageid,omitempty"`
	K8SCluster        string   `json:"k8s_cluster,omitempty"`
	Level             string   `json:"level,omitempty"` // success, block, detect, alert
	Low               int      `json:"low,omitempty"`
	Medium            int      `json:"medium,omitempty"`
	Pid               int      `json:"pid,omitempty"`
	Poddeployment     string   `json:"poddeployment,omitempty"`
	Podname           string   `json:"podname,omitempty"`
	Podnamespace      string   `json:"podnamespace,omitempty"`
	Podtype           string   `json:"podtype,omitempty"`
	Process           string   `json:"process,omitempty"`
	Reason            string   `json:"reason,omitempty"`
	Registry          string   `json:"registry,omitempty"`
	Resource          string   `json:"resource,omitempty"`
	ResourceDigest    string   `json:"resource_digest,omitempty"`
	ResourceName      string   `json:"resource_name,omitempty"`
	ResourceType      string   `json:"resoure_type,omitempty"`
	Result            int      `json:"result,omitempty"` // 2 = block, 4 = alert, 1 = success, 3 = detect?
	Rule              string   `json:"rule,omitempty"`
	RuleType          string   `json:"rule_type,omitempty"`
	Secret            string   `json:"secret,omitempty"`
	SourceAddress     string   `json:"source_address,omitempty"`
	Subtype           string   `json:"subtype,omitempty"`
	Time              int      `json:"time,omitempty"`
	Type              string   `json:"type,omitempty"` // alert, administration
	UID               string   `json:"uid,omitempty"`
	User              string   `json:"user,omitempty"`
	VMGroup           string   `json:"vm_group,omitempty"`
	VMID              string   `json:"vm_id,omitempty"`
	VMLocation        string   `json:"vm_location,omitempty"`
	VMName            string   `json:"vm_name,omitempty"`
}

// Data is an alert substructure of Audits
type Data struct {
	AdhocScanRegistry string   `json:"adhoc_scan_registry,omitempty"`
	Adjective         string   `json:"adjective,omitempty"`
	Blocking          bool     `json:"blocking,omitempty"`
	BuildNumber       string   `json:"build_number,omitempty"`
	BuildPipeline     string   `json:"build_pipeline,omitempty"`
	Controls          []string `json:"controls,omitempty"`
	Critical          int      `json:"critical,omitempty"`
	Description       string   `json:"description,omitempty"`
	DtaSeverityScore  string   `json:"dta_severity_score,omitempty"`
	DtaSkipped        bool     `json:"dta_skipped,omitempty"`
	High              int      `json:"high,omitempty"`
	Host              string   `json:"host,omitempty"`
	Low               int      `json:"low,omitempty"`
	Malware           int      `json:"malware,omitempty"`
	Medium            int      `json:"medium,omitempty"`
	Pending           bool     `json:"pending,omitempty"`
	PolicyID          int      `json:"policy_id,omitempty"`
	PolicyName        string   `json:"policy_name,omitempty"`
	Registry          string   `json:"registry,omitempty"`
	Repository        string   `json:"repository,omitempty"`
	ScanDuration      int      `json:"scan_duration,omitempty"`
	ScanEndTime       int      `json:"scan_end_time,omitempty"`
	ScanStartTime     int      `json:"scan_start_time,omitempty"`
	SensitiveData     int      `json:"sensitive_data,omitempty"`
	SourceIP          string   `json:"source_ip,omitempty"`
	Time              int      `json:"time,omitempty"`
	Type              string   `json:"type,omitempty"`
}
