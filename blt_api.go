package mysql

import (
	"database/sql/driver"
	"errors"
)

// ExtendConn creates an extended connection.
func ExtendConn(conn driver.Conn) (*ExtendedConn, error) {
	if conn == nil {
		return nil, errors.New("Connection is nil")
	}
	mc, ok := conn.(*mysqlConn)
	if !ok || mc == nil {
		return nil, errors.New("Invalid connection")
	}

	return &ExtendedConn{mc}, nil
}

// ExtendedConn provides access to internal packet functions.
type ExtendedConn struct {
	*mysqlConn
}

// Exec ...
func (c *ExtendedConn) Exec(query string) error {
	return c.exec(query)
}

// ReadPacket reads a packet from a given connection.
func (c *ExtendedConn) ReadPacket() ([]byte, error) {
	return c.readPacket()
}

// WritePacket writes a packet to a given connection.
func (c *ExtendedConn) WritePacket(p []byte) error {
	return c.writePacket(p)
}

// ReadResultOK ...
func (c *ExtendedConn) ReadResultOK() error {
	return c.readResultOK()
}

// HandleErrorPacket ...
func (c *ExtendedConn) HandleErrorPacket(data []byte) error {
	return c.handleErrorPacket(data)
}

// ResetSequence ...
func (c *ExtendedConn) ResetSequence() {
	c.sequence = 0
}
