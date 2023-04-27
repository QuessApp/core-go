package dtos

import (
	"testing"

	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetCreateReportValidateDTOBatches returns a slice of BatchTest for CreateReportDTO.
func GetCreateReportValidateDTOBatches(t *testing.T, createReportData reports.CreateReportDTO) []tests.BatchTest {

	return []tests.BatchTest{
		{
			OnRun: func() {
				assert.NoError(t, createReportData.Validate())
			},
		},
		{
			OnRun: func() {
				createReportData.Reason = ""
				assert.ErrorContains(t, createReportData.Validate(), "reason_field_required")

				createReportData.Reason = "foobar"
				assert.ErrorContains(t, createReportData.Validate(), "reason_field_invalid")

				createReportData.Reason = "hate speech or symbols"
				assert.NoError(t, createReportData.Validate())
			},
		},
		{
			OnRun: func() {
				createReportData.Type = ""
				assert.ErrorContains(t, createReportData.Validate(), "type_field_required")

				createReportData.Type = "block"
				assert.ErrorContains(t, createReportData.Validate(), "type_field_invalid")

				createReportData.Type = "user"
				assert.NoError(t, createReportData.Validate())

				createReportData.Type = "question"
				assert.NoError(t, createReportData.Validate())
			},
		},
	}
}
