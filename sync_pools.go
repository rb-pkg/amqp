package amqp

import (
	"sync"
	"time"
)

var (
	poolPublishing      sync.Pool
	poolBasicPublishing sync.Pool
	poolDelivery        sync.Pool
)

func (c *Publishing) Reset() {
	c.Headers = nil
	c.ContentType = ""
	c.ContentEncoding = ""
	c.DeliveryMode = 0
	c.Priority = 0
	c.CorrelationId = ""
	c.ReplyTo = ""
	c.Expiration = ""
	c.MessageId = ""
	c.Timestamp = time.Time{}
	c.Type = ""
	c.UserId = ""
	c.AppId = ""
	c.Body = []byte{}
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

func (msg *basicPublish) Reset() {
	msg.Exchange = ""
	msg.RoutingKey = ""
	msg.Mandatory = false
	msg.Immediate = false
	msg.Properties = properties{}
	msg.Body = []byte{}
}

// acquireBasicPublish returns an empty basicPublish instance from basic Publish pool.
//
// The returned basicPublish instance may be passed to releaseBasicPublishing when it is
// no longer needed. This allows basicPublish recycling, reduces GC pressure
// and usually improves performance.
func acquireBasicPublish() *basicPublish {
	v := poolBasicPublishing.Get()
	if v == nil {
		return &basicPublish{}
	}
	return v.(*basicPublish)
}

// releaseBasicPublishing returns pub acquired via acquireBasicPublishing to basicPublish pool.
//
// It is forbidden accessing bPub and/or its' members after returning it to request pool.
func releaseBasicPublishing(bPub *basicPublish) {
	if bPub != nil {
		bPub.Reset()
	}
	poolBasicPublishing.Put(bPub)
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
