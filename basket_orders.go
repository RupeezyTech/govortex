package govortex

import (
	"context"
	"fmt"
	"time"
)

func (v *VortexApi) GetBaskets(ctx context.Context) (*Baskets, error) {
	var resp Baskets
	_, err := v.doJson(ctx, "GET", URIBaskets, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) CreateBasket(ctx context.Context, request *CreateBasketRequest) (*BasketOrder, error) {
	var resp BasketOrder
	_, err := v.doJson(ctx, "POST", URIBaskets, request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (v *VortexApi) DeleteBasket(ctx context.Context, basketID int) (*ModifyBasketResponse, error) {
	var resp ModifyBasketResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URIBasket, basketID), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (v *VortexApi) UpdateBasket(ctx context.Context, basketId int, request *ModifyBasketRequest) (*ModifyBasketResponse, error) {
	req := modifyBasketOrder{}
	var resp ModifyBasketResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URIBasket, basketId), req, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// {
//     "status": "success",
//     "message": "",
//     "response": [
//         {
//             "id": 1,
//             "clientCode": "DEMOUSER",
//             "basketName": "Test Basket",
//             "basketOrders": [
//                 {
//                     "id": 1,
//                     "orderPosition": 0,
//                     "token": 2885,
//                     "marketSegmentId": 1,
//                     "stockData": {
//                         "scripToken": 2885,
//                         "marketSegmentId": 1,
//                         "symbol": "RELIANCE",
//                         "securityDesc": "RELIANCE INDUSTRIES LTD",
//                         "instrumentName": "EQUITIES",
//                         "series": "EQ",
//                         "eligibility": 1,
//                         "lotSize": 1,
//                         "tick": 10,
//                         "industry": "Refineries",
//                         "sectorName": "Energy",
//                         "hasOptionChain": true,
//                         "isinCode": "INE002A01018",
//                         "has_icon": true
//                     },
//                     "buySell": 1,
//                     "productType": "3",
//                     "orderType": "2",
//                     "quantity": 1,
//                     "price": 1450.8,
//                     "disclosedQuantity": 0,
//                     "validity": "1",
//                     "validityDays": 0,
//                     "isAmo": false,
//                     "orderIdentifier": "",
//                     "orderId": "",
//                     "triggerPrice": 0,
//                     "createdAt": "2025-06-24T11:45:41.879769Z",
//                     "updatedAt": "2025-06-24T11:45:41.879769Z",
//                     "tag_ids": [],
//                     "metadata": {}
//                 }
//             ],
//             "ordersCount": 1,
//             "createdAt": "2025-06-24T11:45:12.917014Z",
//             "updatedAt": "2025-06-24T11:45:41.879431Z",
//             "lastExecutedAt": "0001-01-01T00:00:00Z",
//             "alertId": null,
//             "alert": null,
//             "isExecuted": false,
//             "tag_ids": []
//         },
//         {
//             "id": 1,
//             "clientCode": "DEMOUSER",
//             "basketName": "My First Basket",
//             "basketOrders": [],
//             "ordersCount": 0,
//             "createdAt": "2024-12-18T09:26:09.2368Z",
//             "updatedAt": "2024-12-18T09:26:09.2368Z",
//             "lastExecutedAt": "0001-01-01T00:00:00Z",
//             "alertId": null,
//             "alert": null,
//             "isExecuted": false,
//             "tag_ids": null
//         }
//     ]
// }

type Baskets struct {
	Status   string   `json:"status"`
	Message  string   `json:"message"`
	Response []Basket `json:"response"`
}

type Basket struct {
	ID             int           `json:"id"`
	ClientCode     string        `json:"clientCode"`
	BasketName     string        `json:"basketName"`
	BasketOrders   []BasketOrder `json:"basketOrders"`
	OrdersCount    int           `json:"ordersCount"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	LastExecutedAt time.Time     `json:"lastExecutedAt"`
	AlertId        *int          `json:"alertId"`
	Alert          *Alert        `json:"alert"`
	IsExecuted     bool          `json:"isExecuted"`
	TagIds         []int         `json:"tag_ids"`
}
type BasketOrder struct {
	ID                int                    `json:"id"`
	OrderPosition     int                    `json:"orderPosition"`
	Token             int                    `json:"token"`
	MarketSegmentId   int                    `json:"marketSegmentId"`
	BuySell           int                    `json:"buySell"`
	ProductType       string                 `json:"productType"`
	OrderType         string                 `json:"orderType"`
	Quantity          int                    `json:"quantity"`
	Price             float64                `json:"price"`
	DisclosedQuantity int                    `json:"disclosedQuantity"`
	Validity          string                 `json:"validity"`
	ValidityDays      int                    `json:"validityDays"`
	IsAmo             bool                   `json:"isAmo"`
	OrderIdentifier   string                 `json:"orderIdentifier"`
	OrderId           string                 `json:"orderId"`
	TriggerPrice      float64                `json:"triggerPrice"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	TagIds            []int                  `json:"tag_ids"`
	Metadata          map[string]interface{} `json:"metadata"`
}

type ModifyBasketResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

type CreateBasketRequest struct {
	Name string `json:"name"`
}

type ModifyBasketRequest struct {
	Name   string
	Orders []ModifyBasketOrder
}

type ModifyBasketOrder struct {
	Token             int
	Exchange          ExchangeTypes
	TransactionType   TransactionTypes
	ProductType       ProductTypes
	Variety           VarietyTypes
	Quantity          int
	Price             float64
	DisclosedQuantity int
	Validity          ValidityTypes
	ValidityDays      int
	IsAmo             bool
	OrderIdentifier   string
	TriggerPrice      float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	TagIds            []int
	Metadata          map[string]interface{}
}

type modifyBasketRequest struct {
	Name   string              `json:"basketName"`
	Orders []modifyBasketOrder `json:"orders"`
}

//	{
//	            "buySell": 1,
//	            "disclosedQuantity": 0,
//	            "isAmo": false,
//	            "marketSegmentId": 1,
//	            "orderIdentifier": "Your identifier",
//	            "orderType": "2",
//	            "price": 1450.8,
//	            "productType": "3",
//	            "quantity": 1,
//	            "tag_ids": [],
//	            "token": 2885,
//	            "triggerPrice": 0,
//	            "validity": "1",
//	            "validityDays": 0
//	        },
type modifyBasketOrder struct {
	BuySell           int                    `json:"buySell"`
	DisclosedQuantity int                    `json:"disclosedQuantity"`
	IsAmo             bool                   `json:"isAmo"`
	MarketSegmentId   int                    `json:"marketSegmentId"`
	OrderIdentifier   string                 `json:"orderIdentifier"`
	OrderType         string                 `json:"orderType"`
	Price             float64                `json:"price"`
	ProductType       string                 `json:"productType"`
	Quantity          int                    `json:"quantity"`
	TagIds            []int                  `json:"tag_ids"`
	Token             int                    `json:"token"`
	TriggerPrice      float64                `json:"triggerPrice"`
	Validity          string                 `json:"validity"`
	ValidityDays      int                    `json:"validityDays"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	Metadata          map[string]interface{} `json:"metadata"`
}
