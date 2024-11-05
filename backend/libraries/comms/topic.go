package comms

import "strings"

type TopicOption func(string) (string, error)

const REDIS_WILDCARD = "*"
const DEFAULT_DELIMITER = "." // redis channel delimiter
const REDIS_KEY_DELIMITER = "."

func WithKey(topic string) (string, error) {
	keyTopic := strings.ReplaceAll(topic, DEFAULT_DELIMITER, REDIS_KEY_DELIMITER)
	return keyTopic, nil
}

func PufEthConversionRateTopic(chain string, realtime bool, opts ...TopicOption) (string, error) {
	var err error
	var topic string

	var quality string
	if realtime {
		quality = "latest"
	} else {
		quality = "historical"
	}
	levels := []string{"puffer", "rate", chain, quality}
	for i, part := range levels {
		if part != "" {
			continue
		}

		levels[i] = REDIS_WILDCARD
	}
	topic = strings.Join(levels, DEFAULT_DELIMITER)

	for _, opt := range opts {
		topic, err = opt(topic)
		if err != nil {
			return "", err
		}
	}

	return topic, nil
}
