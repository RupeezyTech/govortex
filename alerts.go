package govortex

import (
	"context"
	"fmt"
)

func (v *VortexApi) GetAlerts(ctx context.Context) (*AlertBook, error) {
	var resp AlertBook
	_, err := v.doJson(ctx, "GET", URIAlerts, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) CreateAlert(ctx context.Context, request *CreateAlertRequest) (*CreateAlertResponse, error) {
	req := createAlertRequest{
		BasketID:        request.BasketID,
		Condition:       request.Condition,
		MarketSegmentId: exchangeToMarketSegmentId[request.Exchange],
		Name:            request.Name,
		Note:            request.Note,
		Property:        request.Property,
		Token:           request.Token,
		TriggerValue:    request.TriggerValue,
	}
	var resp CreateAlertResponse
	_, err := v.doJson(ctx, "POST", URIAlerts, req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (v *VortexApi) DeleteAlert(ctx context.Context, alertID int) (*ModifyAlertResponse, error) {
	var resp ModifyAlertResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIAlert, alertID), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (v *VortexApi) UpdateAlert(ctx context.Context, alertID int, request *UpdateAlertRequest) (*ModifyAlertResponse, error) {
	req := updateAlertRequest{
		BasketID:        request.BasketID,
		Condition:       request.Condition,
		MarketSegmentId: exchangeToMarketSegmentId[request.Exchange],
		Name:            request.Name,
		Note:            request.Note,
		Property:        request.Property,
		Token:           request.Token,
		TriggerValue:    request.TriggerValue,
	}
	var resp ModifyAlertResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIAlert, alertID), req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type Alert struct {
	DeletedAt       *string        `json:"DeletedAt"`
	ClientCode      string         `json:"clientCode"`
	Condition       AlertCondition `json:"condition"`
	CreatedAt       string         `json:"createdAt"`
	Expiry          string         `json:"expiry"`
	ID              int            `json:"id"`
	IsExecuted      bool           `json:"isExecuted"`
	LastExecutedAt  string         `json:"lastExecutedAt"`
	MarketSegmentID int            `json:"marketSegmentId"`
	Name            string         `json:"name"`
	Note            string         `json:"note"`
	Property        AlertProperty  `json:"property"`
	Source          string         `json:"source"`
	Status          string         `json:"status"`
	Token           int            `json:"token"`
	TriggerValue    float64        `json:"triggerValue"`
	UpdatedAt       string         `json:"updatedAt"`
}

type AlertBook struct {
	Data    []Alert `json:"data"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}

type CreateAlertRequest struct {
	BasketID     *int
	Condition    AlertCondition
	Exchange     ExchangeTypes
	Name         string
	Note         string
	Property     AlertProperty
	Token        int
	TriggerValue float64
}

type createAlertRequest struct {
	BasketID        *int            `json:"basketId,omitempty"`
	Condition       AlertCondition  `json:"condition"`
	MarketSegmentId marketSegmentId `json:"marketSegmentId"`
	Name            string          `json:"name"`
	Note            string          `json:"note,omitempty"`
	Property        AlertProperty   `json:"property"`
	Token           int             `json:"token"`
	TriggerValue    float64         `json:"triggerValue"`
}

type CreateAlertResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    AlertId `json:"data"`
}

type AlertId struct {
	AlertId int `json:"alertId"`
}

type ModifyAlertResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

type UpdateAlertRequest struct {
	BasketID     *int
	Condition    AlertCondition
	Exchange     ExchangeTypes
	Name         string
	Note         string
	Property     AlertProperty
	Token        int
	TriggerValue float64
}

type updateAlertRequest struct {
	BasketID        *int            `json:"basketId,omitempty"`
	Condition       AlertCondition  `json:"condition"`
	MarketSegmentId marketSegmentId `json:"marketSegmentId"`
	Name            string          `json:"name"`
	Note            string          `json:"note,omitempty"`
	Property        AlertProperty   `json:"property"`
	Token           int             `json:"token"`
	TriggerValue    float64         `json:"triggerValue"`
}
