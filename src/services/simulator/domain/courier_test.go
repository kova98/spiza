package domain_test

import (
	"encoding/json"
	"github.com/kova98/spiza/services/simulator/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"os"
	"testing"
	"time"
)

type MockBus struct{ mock.Mock }

func (m MockBus) Publish(topic string, msg interface{}) {
	m.Called(topic, msg)
}

var (
	BusMock = new(MockBus)
	// TODO: Use a mock logger, test logs
	Logger = log.New(os.Stdout, "test", log.LstdFlags)
)

func NewTestCourier(loc domain.Location) *domain.Courier {
	return domain.NewCourier(1, "Test Courier", loc, BusMock, Logger)
}

func TestAssignToOrder(t *testing.T) {
	courier := domain.Courier{}

	orderID := int64(123)
	courier.AssignToOrder(orderID)

	assert.Equal(t, orderID, courier.CurrentOrderId, "CurrentOrderId should match the assigned order ID")
}

func TestTravel(t *testing.T) {
	courier := NewTestCourier(domain.Location{})
	orderId := int64(1)
	path := []domain.Location{
		{Lat: 0.1, Lng: 0.1},
		{Lat: 0.2, Lng: 0.2},
	}
	BusMock.On("Publish", mock.AnythingOfType("string"), mock.AnythingOfType("Location")).Times(len(path))

	courier.Travel(orderId, path)

	BusMock.AssertExpectations(t)
	assert.Equal(t, path[len(path)-1], courier.Loc, "The courier's final location should be the last location in the path")
}

func TestPickUpOrder(t *testing.T) {
	courier := NewTestCourier(domain.Location{})
	orderId := int64(1)
	BusMock.On("Publish", "order/1", mock.AnythingOfType("[]uint8")).Run(func(args mock.Arguments) {
		var msg domain.OrderUpdated
		json.Unmarshal(args.Get(1).([]byte), &msg)
		assert.Equal(t, domain.OrderStatusPickedUp, msg.Status, "The order status should be picked up")
		assert.NotZero(t, msg.DeliveryTime, "The delivery time should not be zero") // TODO: revisit once proper time calculation is implemented
	})

	courier.PickUpOrder(orderId)

	BusMock.AssertExpectations(t)
}

func TestCompleteOrder(t *testing.T) {
	courier := NewTestCourier(domain.Location{})
	orderId := int64(1)
	msg := domain.OrderUpdated{
		Status:       domain.OrderStatusDelivered,
		DeliveryTime: time.Now().UTC().Add(15 * time.Minute),
	}
	BusMock.On("Publish", "order/1", mock.AnythingOfType("[]uint8")).Run(func(args mock.Arguments) {
		var actual domain.OrderUpdated
		json.Unmarshal(args.Get(1).([]byte), &actual)
		assert.Equal(t, domain.OrderStatusDelivered, actual.Status, "The order status should be delivered")
		assert.Equal(t, msg.DeliveryTime, actual.DeliveryTime, "The delivery time should match the one in the message")
	})

	courier.CompleteOrder(orderId, msg)

	BusMock.AssertExpectations(t)
}
