package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestPlaceOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := PlaceOrderRequest{}
	resp, err := ts.VortexApiClient.PlaceOrder(ctx, request)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while placing order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestModifyOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := ModifyOrderRequest{}
	resp, err := ts.VortexApiClient.ModifyOrder(ctx, request, ExchangeTypesNSEEQUITY, "NXAAE00002K3")
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while modifying order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestCancelOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.CancelOrder(ctx, "NXAAE00002K3")
	if err != nil {
		t.Errorf("Error while cancelling order. %v", err)
		return
	}
	if resp.Data.OrderID != "NXAAE00002K3" {
		t.Errorf("Error while cancelling order. %s", "order id is not same")
	}
}

func (ts *TestSuite) TestOrderBook(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Orders(ctx)
	if err != nil {
		t.Errorf("Error while fetching order book. %v", err)
		return
	}
	if len(resp.Orders) == 0 {
		t.Errorf("Errorwhile fetching order book. %s", "order book is empty")
	}
}

func (ts *TestSuite) TestOrderHistory(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.OrderHistory(ctx, "test")
	if err != nil {
		t.Errorf("Error while fetching order history. %v", err)
		return
	}
	if len(resp.Data) == 0 {
		t.Errorf("Error while fetching order history. %s", "order history is empty")
	}
}

func (ts *TestSuite) TestMultipleOrderCancel(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := MultipleOrderCancelRequest{}
	resp, err := ts.VortexApiClient.CancelMultipleRegularOrders(ctx, request)
	if err != nil {
		t.Errorf("Error while cancelling order: %v", err.Error())
		return
	}
	if len(resp.Data) == 0 {
		t.Errorf("Error while cancelling order: %s", "order history is empty")
	}
}

func (ts *ErrorTestSuite) TestPlaceOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := PlaceOrderRequest{}
	_, err := ts.VortexApiClient.PlaceOrder(ctx, request)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestModifyOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := ModifyOrderRequest{}
	_, err := ts.VortexApiClient.ModifyOrder(ctx, request, ExchangeTypesNSEEQUITY, "NXAAE00002K3")
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestCancelOrder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.CancelOrder(ctx, "NXAAE00002K3")
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestOrderBook(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Orders(ctx)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestOrderHistory(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.OrderHistory(ctx, "test")
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestMultipleOrderCancel(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := MultipleOrderCancelRequest{}
	_, err := ts.VortexApiClient.CancelMultipleRegularOrders(ctx, request)
	checkError429(t, err)
}
