package soap

type Logger interface {
	LogRequest(ctx interface{}, uuid string, messageType string, message string) error
}
