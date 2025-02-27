package main

type FalcoEvent struct {
	Time         string                 `json:"time"`
	HostName     string                 `json:"hostname"`
	Rule         string                 `json:"rule"`
	OutputFields map[string]interface{} `json:"output_fields"`
}

// processData processes the Falco event data and modifies the event data in place.
//
// The input to this function is a pointer to the FalcoEvent struct.
//
// An expected usage might be:
//
//	if err := processData(&event); err != nil {
func processData(data *FalcoEvent) error {
	if data.Rule == "DNS Query Logging" {
		hostName, err := retrieveDNSQueryHost(data.OutputFields["evt.arg.data"].(string))
		if err != nil {
			return err
		}
		data.OutputFields["dns_query"] = hostName
	}
	delete(data.OutputFields, "evt.time")
	return nil
}
