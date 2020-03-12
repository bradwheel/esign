// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package folders implements the DocuSign SDK
// category Folders.
//
// Use the Folders category to manage envelopes in your folders.
//
// You can list the folder contents and move envelopes between folders.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/v2/reference/Folders
// Usage example:
//
//   import (
//       "github.com/bradwheel/esign"
//       "github.com/bradwheel/esign/v2/folders"
//       "github.com/bradwheel/esign/v2/model"
//   )
//   ...
//   foldersService := folders.New(esignCredential)
package folders // import "github.com/bradwheel/esign/v2/folders"

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/bradwheel/esign"
	"github.com/bradwheel/esign/v2/model"
)

// Service implements DocuSign Folders Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a folders service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// List gets a list of the folders for the account.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/folders/folders/list
//
// SDK Method Folders::list
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "folders",
		QueryOpts:  make(url.Values),
	}
}

// ListOp implements DocuSign API SDK Folders::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.FoldersResponse, error) {
	var res *model.FoldersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Include reserved for DocuSign.
func (op *ListOp) Include(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("include", val)
	}
	return op
}

// StartPosition reserved for DocuSign.
func (op *ListOp) StartPosition(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Template specifies the items that are returned. Valid values are:
//
// * include - The folder list will return normal folders plus template folders.
// * only - Only the list of template folders are returned.
func (op *ListOp) Template(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("template", val)
	}
	return op
}

// UserFilter reserved for DocuSign.
func (op *ListOp) UserFilter(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("user_filter", val)
	}
	return op
}

// ListItems gets a list of the envelopes in the specified folder.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/folders/folders/listitems
//
// SDK Method Folders::listItems
func (s *Service) ListItems(folderID string) *ListItemsOp {
	return &ListItemsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"folders", folderID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// ListItemsOp implements DocuSign API SDK Folders::listItems
type ListItemsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListItemsOp) Do(ctx context.Context) (*model.FolderItemsResponse, error) {
	var res *model.FolderItemsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate only return items on or after this date. If no value is provided, the default search is the previous 30 days.
func (op *ListItemsOp) FromDate(val time.Time) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// OwnerEmail is the email of the folder owner.
func (op *ListItemsOp) OwnerEmail(val string) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("owner_email", val)
	}
	return op
}

// OwnerName is the name of the folder owner.
func (op *ListItemsOp) OwnerName(val string) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("owner_name", val)
	}
	return op
}

// SearchText is the search text used to search the items of the envelope. The search looks at recipient names and emails, envelope custom fields, sender name, and subject.
func (op *ListItemsOp) SearchText(val string) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("search_text", val)
	}
	return op
}

// StartPosition is the position of the folder items to return. This is used for repeated calls, when the number of envelopes returned is too much for one return (calls return 100 envelopes at a time). The default value is 0.
func (op *ListItemsOp) StartPosition(val int) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Status is a comma-separated list of current envelope statuses to included in the response. Possible values are:
//
// * completed
// * created
// * declined
// * deleted
// * delivered
// * processing
// * sent
// * signed
// * timedout
// * voided
//
// The `any` value is equivalent to any status.
func (op *ListItemsOp) Status(val string) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// ToDate only return items up to this date. If no value is provided, the default search is to the current date.
func (op *ListItemsOp) ToDate(val time.Time) *ListItemsOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// MoveEnvelopes moves an envelope from its current folder to the specified folder.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/folders/folders/moveenvelopes
//
// SDK Method Folders::moveEnvelopes
func (s *Service) MoveEnvelopes(folderID string, foldersRequest *model.FoldersRequest) *MoveEnvelopesOp {
	return &MoveEnvelopesOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"folders", folderID}, "/"),
		Payload:    foldersRequest,
		QueryOpts:  make(url.Values),
	}
}

// MoveEnvelopesOp implements DocuSign API SDK Folders::moveEnvelopes
type MoveEnvelopesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *MoveEnvelopesOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// Search gets a list of envelopes in folders matching the specified criteria.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/folders/folders/search
//
// SDK Method Folders::search
func (s *Service) Search(searchFolderID string) *SearchOp {
	return &SearchOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"search_folders", searchFolderID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// SearchOp implements DocuSign API SDK Folders::search
type SearchOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *SearchOp) Do(ctx context.Context) (*model.FolderItemResponse, error) {
	var res *model.FolderItemResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// All specifies that all envelopes that match the criteria are returned.
func (op *SearchOp) All() *SearchOp {
	if op != nil {
		op.QueryOpts.Set("all", "true")
	}
	return op
}

// Count specifies the number of records returned in the cache. The number must be greater than 0 and less than or equal to 100.
func (op *SearchOp) Count(val int) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// FromDate specifies the start of the date range to return. If no value is provided, the default search is the previous 30 days.
func (op *SearchOp) FromDate(val time.Time) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// IncludeRecipients when set to **true**, the recipient information is returned in the response.
func (op *SearchOp) IncludeRecipients() *SearchOp {
	if op != nil {
		op.QueryOpts.Set("include_recipients", "true")
	}
	return op
}

// Order specifies the order in which the list is returned. Valid values are: `asc` for ascending order, and `desc` for descending order.
func (op *SearchOp) Order(val string) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("order", val)
	}
	return op
}

// OrderBy specifies the property used to sort the list. Valid values are: `action_required`, `created`, `completed`, `sent`, `signer_list`, `status`, or `subject`.
func (op *SearchOp) OrderBy(val string) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("order_by", val)
	}
	return op
}

// StartPosition specifies the the starting location in the result set of the items that are returned.
func (op *SearchOp) StartPosition(val int) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// ToDate specifies the end of the date range to return.
func (op *SearchOp) ToDate(val time.Time) *SearchOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}
