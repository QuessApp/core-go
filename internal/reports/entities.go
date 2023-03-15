package reports

import (
	"time"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// Report is a model for each report in app.
type Report struct {
	ID        toolkitEntities.ID `json:"id" bson:"_id"`
	Type      string             `json:"type" bson:"type"`
	Reason    string             `json:"reason" bson:"reason"`
	SendTo    any                `json:"sendTo" bson:"sendTo"`
	SentBy    toolkitEntities.ID `json:"sentBy" bson:"sentBy"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

// PaginatedReports is a model for paginated reports in app.
type PaginatedReports struct {
	Reports    *[]Report `json:"reports"`
	TotalCount int64     `json:"totalCount"`
}
