package reports

import (
	"errors"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// AlreadySent is a function that receives a boolean indicating if a report has already been sent for the same content.
// If the boolean is true, it returns an error indicating that the report can't be sent again.
// If the boolean is false, it returns nil indicating that the report can be sent.
func AlreadySent(alreadySent bool) error {
	if alreadySent {
		return errors.New(pkgErrors.CANT_REPORT_ALREADY_SENT)
	}

	return nil
}

// IsReportingYourself checks whether a user is trying to report themselves or not.
// It receives the authenticated user ID and the ID of the user to whom the report will be sent.
// If the IDs are the same, it returns an error indicating that the user cannot report themselves.
// Otherwise, it returns nil, indicating that the user can proceed with reporting the other user.
func IsReportingYourself(authenticatedUserID toolkitEntities.ID, sendTo toolkitEntities.ID) error {
	if sendTo == authenticatedUserID {
		return errors.New(pkgErrors.CANT_REPORT_YOURSELF)
	}

	return nil
}

// ReportExists validates whether a report exists or not based on its ID.
func ReportExists(r *Report) error {
	if toolkitEntities.IsZeroID(r.ID) {
		return errors.New(pkgErrors.REPORT_NOT_FOUND)
	}

	return nil
}

// CanUserDeleteReport validates whether the user who is trying to delete the report is the report owner.
func CanUserDeleteReport(r *Report, authenticatedUserID toolkitEntities.ID) error {
	if r.SentBy != authenticatedUserID {
		return errors.New(pkgErrors.CANT_DELETE_REPORT_NOT_SENT_BY_YOU)
	}

	return nil
}

// CanViewReport validates whether the authenticated user is authorized to view the report.
func CanViewReport(r *Report, authenticatedUserID toolkitEntities.ID) error {
	if r.SentBy != authenticatedUserID {
		return errors.New(pkgErrors.REPORT_NOT_AUTHORIZED)
	}

	return nil
}