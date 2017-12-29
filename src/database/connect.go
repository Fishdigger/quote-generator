package database

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"gopkg.in/mgo.v2"
)

//OpenSession ... returns an open session with the database
func OpenSession() *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs: []string{os.Getenv("DB_PREFIX0"),
			os.Getenv("DB_PREFIX1"),
			os.Getenv("DB_PREFIX1")},
		Database: os.Getenv("DB_AUTHDB"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}

	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Println("Problems connecting to database!!!!", err)
		panic(err)
	}

	return session
}

//CloseSession ... closes the session with the db
func CloseSession(session *mgo.Session) {
	session.Close()
}
