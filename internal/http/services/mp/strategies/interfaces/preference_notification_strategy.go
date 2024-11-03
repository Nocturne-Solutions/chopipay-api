package interfaces

import (
	"chopipay/internal/http/services/mp/notifications/merchantorder"
	notiPayment "chopipay/internal/http/services/mp/notifications/payment"
	"chopipay/internal/models/entities"
)

type PreferenceNotificationStrategy interface {
    Process() error
}

type NotificationProcessor struct {
    PaymentStrategy PreferenceNotificationStrategy
    Strategies map[string]PreferenceNotificationStrategy
}

func (p *NotificationProcessor) loadStrategies(notificationId int, 
                                               credentials *entities.PersonalCredentials) {
    
        p.Strategies = map[string]PreferenceNotificationStrategy{
        "merchant_order": &merchantorder.Processor{
            NotificationId: notificationId,
            Credentials: credentials,
        },
        "payment": &notiPayment.Processor{
            NotificationId: notificationId,
            Credentials: credentials,
        },
    }
}

func (p *NotificationProcessor) Process(strategy string, 
                                        notificationId int, 
                                        credentials *entities.PersonalCredentials) error {
    if p.Strategies == nil {
        p.loadStrategies(notificationId, credentials)
    }

    if strategy, ok := p.Strategies[strategy]; ok {
        return strategy.Process()
    }

    return nil
}