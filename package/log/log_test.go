package log

import (
	"bytes"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type MockLogger struct {
	buffer bytes.Buffer
}

func (l *MockLogger) Write(p []byte) (n int, err error) {
	return l.buffer.Write(p)
}

func (l *MockLogger) GetOutput() string {
	return l.buffer.String()
}

func TestNewLog(t *testing.T) {
	log := NewLog()

	// Ensure that the returned log object is not nil
	require.NotNil(t, log)

	// Ensure that the log object is of the correct type
	_, ok := log.(*option)
	require.True(t, ok)

	// Assert the log level and formatter
	option := log.(*option)
	assert.Equal(t, logrus.DebugLevel, option.logger.GetLevel())
	assert.IsType(t, &prefixed.TextFormatter{}, option.logger.Formatter)
}

func TestInfo(t *testing.T) {
	// Create a mock logger
	mockLogger := &MockLogger{}

	// Create a log object with the mock logger
	log := &option{
		logger: logrus.New(),
	}
	log.logger.Out = mockLogger

	// Call the Info function
	log.Info("This is an info message", nil)

	// Assert the output in the mock logger
	expectedOutput := `level=info msg="This is an info message"`
	assert.Contains(t, mockLogger.GetOutput(), expectedOutput)
}

func TestWarning(t *testing.T) {
	// Create a mock logger
	mockLogger := &MockLogger{}

	// Create a log object with the mock logger
	log := &option{
		logger: logrus.New(),
	}
	log.logger.Out = mockLogger

	// Call the Warning function with an error
	err := errors.New("test error")
	log.Warning("This is a warning message", err, nil)

	// Assert the output in the mock logger
	expectedOutput := `level=warning msg="This is a warning message | Error: test error"`
	assert.Contains(t, mockLogger.GetOutput(), expectedOutput)
}

func TestWarningNoMessage(t *testing.T) {
	// Create a mock logger
	mockLogger := &MockLogger{}

	// Create a log object with the mock logger
	log := &option{
		logger: logrus.New(),
	}
	log.logger.Out = mockLogger

	// Call the Warning function with an error
	err := errors.New("test error")
	log.Warning("", err, nil)

	// Assert the output in the mock logger
	expectedOutput := `level=warning msg="Error: test error"`
	assert.Contains(t, mockLogger.GetOutput(), expectedOutput)
}

func TestError(t *testing.T) {
	// Create a mock logger
	mockLogger := &MockLogger{}

	// Create a log object with the mock logger
	log := &option{
		logger: logrus.New(),
	}
	log.logger.Out = mockLogger

	// Call the Error function with an error
	err := errors.New("test error")
	log.Error(err, "This is an error message", nil)

	// Assert the output in the mock logger
	expectedOutput := `level=error msg="This is an error message | Error: test error"`
	assert.Contains(t, mockLogger.GetOutput(), expectedOutput)
}

func TestErrorNoMessage(t *testing.T) {
	// Create a mock logger
	mockLogger := &MockLogger{}

	// Create a log object with the mock logger
	log := &option{
		logger: logrus.New(),
	}
	log.logger.Out = mockLogger

	// Call the Error function with an error
	err := errors.New("test error")
	log.Error(err, "", nil)

	// Assert the output in the mock logger
	expectedOutput := `level=error msg="Error: test error"`
	assert.Contains(t, mockLogger.GetOutput(), expectedOutput)
}
