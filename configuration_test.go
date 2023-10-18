package treblle

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
	"testing"
)

func TestLoggerImplementation(t *testing.T) {
	var buf bytes.Buffer
	want := "treblle: testing logger implementation"
	tests := map[string]struct {
		logger Logger
	}{
		"log":    {logger: log.New(&buf, "", log.Lshortfile)},
		"logrus": {logger: &logrus.Logger{Out: &buf}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			Configure(Configuration{Logger: tc.logger})
			Config.Logger.Print(want)

			if !strings.Contains(buf.String(), want) {
				t.Fatalf("expected: %v, got: %v", want, buf.String())
			}

			buf.Reset()
		})
	}
}
