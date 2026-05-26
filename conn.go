// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbio

import (
	"time"
)

// ConnType is used to identify different types of Conn.
type ConnType = int8

const (
	// ConnTypeTCP represents TCP Conn.
	ConnTypeTCP ConnType = iota + 1
	// ConnTypeUDPServer represents UDP Conn used as a listener.
	ConnTypeUDPServer
	// ConnTypeUDPClientFromRead represents UDP connection that
	// is sending data to our UDP Server from peer.
	ConnTypeUDPClientFromRead
	// ConnTypeUDPClientFromDial represents UDP Conn that is sending
	// data to other UDP Server from ourselves.
	ConnTypeUDPClientFromDial
	// ConnTypeUnix represents Unix Conn.
	ConnTypeUnix
)

// Type .
//
//go:norace
func (c *Conn) Type() ConnType {
	_ = "STUB: not implemented"

	// IsTCP returns whether this Conn is a TCP Conn.
	//
	//go:norace
	return *new(ConnType)
}

func (c *Conn) IsTCP() bool { _ = "STUB: not implemented"; return false }

// IsUDP returns whether this Conn is a UDP Conn.
//
//go:norace
func (c *Conn) IsUDP() bool { _ = "STUB: not implemented"; return false }

// IsUnix  returns whether this Conn is a Unix Conn.
//
//go:norace
func (c *Conn) IsUnix() bool { _ = "STUB: not implemented"; return false }

// Session returns user session.
//
//go:norace
func (c *Conn) Session() interface{} {
	_ = "STUB: not implemented"

	// SetSession sets user session.
	//
	//go:norace
	return nil
}

func (c *Conn) SetSession(session interface{}) { _ = "STUB: not implemented"; return }

// OnData registers Conn's data handler.
// Notice:
//  1. The data readed by the poller is not handled by this Conn's data
//     handler by default.
//  2. The data readed by the poller is handled by nbio.Engine's data
//     handler which is registered by nbio.Engine.OnData by default.
//  3. This Conn's data handler is used to customize your implementation,
//     you can set different data handler for different Conns,
//     and call Conn's data handler in nbio.Engine's data handler.
//     For example:
//     engine.OnData(func(c *nbio.Conn, data byte){
//     c.DataHandler()(c, data)
//     })
//     conn1.OnData(yourDatahandler1)
//     conn2.OnData(yourDatahandler2)
//
//go:norace
func (c *Conn) OnData(h func(conn *Conn, data []byte)) {
	_ = "STUB: not implemented"

	// DataHandler returns Conn's data handler.
	//
	//go:norace
	return
}

func (c *Conn) DataHandler() func(conn *Conn, data []byte) { _ = "STUB: not implemented"; return nil }

// Dial calls net.Dial to make a net.Conn and convert it to *nbio.Conn.
//
//go:norace
func Dial(network string, address string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dial calls net.DialTimeout to make a net.Conn and convert it to *nbio.Conn.
//
//go:norace
func DialTimeout(network string, address string, timeout time.Duration) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lock .
//
//go:norace
func (c *Conn) Lock() {
	_ = "STUB: not implemented"

	// Unlock .
	//
	//go:norace
	return
}

func (c *Conn) Unlock() {
	_ = "STUB: not implemented"

	// IsClosed returns whether the Conn is closed.
	//
	//go:norace
	return
}

func (c *Conn) IsClosed() (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// ExecuteLen returns the length of the Conn's job list.
		//
		//go:norace
		nil
}

func (c *Conn) ExecuteLen() int { _ = "STUB: not implemented"; return 0 }

// Execute is used to run the job.
//
// How it works:
// If the job is the head/first of the Conn's job list, it will call the
// nbio.Engine.Execute to run all the jobs in the job list that include:
//  1. This job
//  2. New jobs that are pushed to the back of the list before this job
//     is done.
//  3. nbio.Engine.Execute returns until there's no more jobs in the job
//     list.
//
// Else if the job is not the head/first of the job list, it will push the
// job to the back of the job list and wait to be called.
// This guarantees there's at most one flow or goroutine running job/jobs
// for each Conn.
// This guarantees all the jobs are executed in order.
//
// Notice:
//  1. The job wouldn't run or pushed to the back of the job list if the
//     connection is closed.
//  2. nbio.Engine.Execute is handled by a goroutine pool by default, users
//     can customize it.
//
//go:norace
func (c *Conn) Execute(job func()) bool { _ = "STUB: not implemented"; return false }

// If there's no job running, run Engine.Execute to run this job
// and new jobs appended before this head job is done.

// MustExecute implements a similar function as Execute did,
// but will still execute or push the job to the
// back of the job list no matter whether Conn has been closed,
// it guarantees the job to be executed.
// This is used to handle the close event in nbio/nbhttp.
//
//go:norace
func (c *Conn) MustExecute(job func()) { _ = "STUB: not implemented"; return }

// If there's no job running, run Engine.Execute to run this job
// and new jobs appended before this head job is done.

//go:norace
func (c *Conn) execute(job func()) { _ = "STUB: not implemented"; return }

// set nil to release the job and gc

// reuse the slice

// get next job

// set nil to release the job and gc
