package amqp

import (
	"sync"
	"time"
)

var (
	poolPublishing sync.Pool
	poolDelivery   sync.Pool
)

func (p *Publishing) Reset() {
	p.Headers = nil
	p.ContentType = ""
	p.ContentEncoding = ""
	p.DeliveryMode = 0
	p.Priority = 0
	p.CorrelationId = ""
	p.ReplyTo = ""
	p.Expiration = ""
	p.MessageId = ""
	p.Timestamp = time.Time{}
	p.Type = ""
	p.UserId = ""
	p.AppId = ""
	p.Body = []byte{}
}

// AcquirePublishing returns an empty Publishing instance from publishing pool.
//
// The returned Publishing instance may be passed to ReleasePublishing when it is
// no longer needed. This allows Publishing recycling, reduces GC pressure
// and usually improves performance.
func AcquirePublishing() *Publishing {
	v := poolPublishing.Get()
	if v == nil {
		return &Publishing{}
	}
	return v.(*Publishing)
}

// ReleasePublishing returns pub acquired via AcquirePublishing to publishing pool.
//
// It is forbidden accessing pub and/or its' members after returning it to request pool.
func ReleasePublishing(pub *Publishing) {
	if pub != nil {
		pub.Reset()
	}
	poolPublishing.Put(pub)
}

func (d *Delivery) Reset() {
	d.Acknowledger = nil
	d.Headers = nil
	d.ContentType = ""
	d.ContentEncoding = ""
	d.DeliveryMode = 0
	d.Priority = 0
	d.CorrelationId = ""
	d.ReplyTo = ""
	d.Expiration = ""
	d.MessageId = ""
	d.Timestamp = time.Time{}
	d.Type = ""
	d.UserId = ""
	d.AppId = ""
	d.ConsumerTag = ""
	d.MessageCount = 0
	d.DeliveryTag = 0
	d.Redelivered = false
	d.Exchange = ""
	d.RoutingKey = ""
	d.Body = []byte{}
}

// AcquireDelivery returns an empty Request instance from delivery pool.
//
// The returned Delivery instance may be passed to ReleaseDelivery when it is
// no longer needed. This allows Delivery recycling, reduces GC pressure
// and usually improves performance.
func AcquireDelivery() *Delivery {
	v := poolDelivery.Get()
	if v == nil {
		return &Delivery{}
	}
	return v.(*Delivery)
}

// ReleaseDelivery returns del acquired via AcquireDelivery to delivery pool.
//
// It is forbidden accessing del and/or its' members after returning it to request pool.
func ReleaseDelivery(del *Delivery) {
	if del != nil {
		del.Reset()
	}
	poolDelivery.Put(del)
}
