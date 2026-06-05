package misc

// TODO Replace with slog? Or something better?
import "github.com/go-kit/kit/log"

func LogMessage(logger log.Logger, message string) {
	logger.Log("message", message)
}
