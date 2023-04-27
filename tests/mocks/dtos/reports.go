package mocks

import (
	"testing"

	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetCreateReportValidateDTOMock returns a slice of BatchTest for CreateReportDTO.
func GetCreateReportValidateDTOMock(t *testing.T, createReportData reports.CreateReportDTO) []tests.BatchTest {
	createReportDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				assert.NoError(t, createReportData.Validate())
			},
		},
		{
			OnRun: func() {
				createReportData.Reason = ""
				assert.EqualError(t, createReportData.Validate(), "reason_field_required.")

				createReportData.Reason = "foobar"
				assert.EqualError(t, createReportData.Validate(), "reason_field_invalid.")
			},
		},
		{
			OnRun: func() {
				createReportData.Reason = "hate speech or symbols"
				assert.NoError(t, createReportData.Validate())
			},
		},
		{
			OnRun: func() {
				createReportData.Type = ""
				assert.EqualError(t, createReportData.Validate(), "type_field_required.")

				createReportData.Type = "block"
				assert.EqualError(t, createReportData.Validate(), "type_field_invalid.")
			},
		},
		{
			OnRun: func() {
				createReportData.Type = "user"
				assert.NoError(t, createReportData.Validate())

				createReportData.Type = "question"
				assert.NoError(t, createReportData.Validate())
			},
		},
	}

	return createReportDataTest
}
