package entity

// GraylogJSON represent json object receive from Graylog notification, see reference at https://docs.graylog.org/en/3.3/pages/alerts.html
type GraylogJSON struct {
	EventDefinitionID          string `json:"event_definition_id"`
	EventDefinitionType        string `json:"event_definition_type"`
	EventDefinitionTitle       string `json:"event_definition_title"`
	EventDefinitionDescription string `json:"event_definition_description"`
	JobDefinitionID            string `json:"job_definition_id"`
	JobTriggerID               string `json:"job_trigger_id"`
	Event                      struct {
		ID                  string                 `json:"id"`
		EventDefinitionID   string                 `json:"event_definition_id"`
		EventDefinitionType string                 `json:"event_definition_type"`
		OriginContext       string                 `json:"origin_context"`
		Timestamp           string                 `json:"timestamp"`
		TimestampProcessing string                 `json:"timestamp_processing"`
		TimerangeStart      string                 `json:"timerange_start"`
		TimerangeEnd        string                 `json:"timerange_end"`
		Streams             []string               `json:"streams"`
		SourceStreams       []string               `json:"source_streams"`
		Alert               bool                   `json:"alert"`
		Message             string                 `json:"message"`
		Source              string                 `json:"source"`
		KeyTuple            []string               `json:"key_tuple"`
		Key                 string                 `json:"key"`
		Priority            int64                  `json:"priority"`
		Fields              map[string]interface{} `json:"fields"`
	} `json:"event"`
	Backlog []struct {
		ID        string                 `json:"id"`
		Index     string                 `json:"index"`
		Source    string                 `json:"source"`
		Message   string                 `json:"message"`
		Timestamp string                 `json:"timestamp"`
		SteamIDs  []string               `json:"stream_ids"`
		Fields    map[string]interface{} `json:"fields"`
	} `json:"backlog"`
	// Backlog []interface{} `json:"backlog"`
}
