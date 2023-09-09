package htmx


type Headers struct {
	HXRequest     string   `json:"HX-Request"`
	HXTrigger     string `json:"HX-Trigger"`
	HXTriggerName string `json:"HX-Trigger-Name"`
	HXTarget      string `json:"HX-Target"`
	HXCurrentUrl  string `json:"HX-Current-URL"`
}
