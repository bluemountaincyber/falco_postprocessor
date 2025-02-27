package processors

type FalcoEvent struct {
	Time         string                 `json:"time"`
	HostName     string                 `json:"hostname"`
	Rule         string                 `json:"rule"`
	OutputFields map[string]interface{} `json:"output_fields"`
}

// ProcessData processes the Falco event data and modifies the event data in place.
//
// The input to this function is a pointer to the FalcoEvent struct.
//
// An expected usage might be:
//
//	if err := ProcessData(&event); err != nil {
func ProcessData(data *FalcoEvent) error {
	if data.Rule == "DNS Query Logging" {
		hostName, err := RetrieveDNSQueryHost(data.OutputFields["evt.arg.data"].(string))
		if err != nil {
			return err
		}
		data.OutputFields["dns_query"] = hostName
	}
	if data.Rule == "Metadata Access" {
		hostName, err := RetrieveMetadataAccessPath(data.OutputFields["evt.arg.data"].(string))
		if err != nil {
			return err
		}
		data.OutputFields["metadata_path"] = hostName
	}
	delete(data.OutputFields, "evt.time")
	return nil
}
