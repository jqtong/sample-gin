package middleware

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"os"
	"time"
)

// TraceInfo trace info struct
type TraceInfo struct {
	TraceID string
	SpanID  string
}

var clientIP string

// AppendTraceID append trace id on request
func AppendTraceID() gin.HandlerFunc {

	return func(c *gin.Context) {
		//通过ip设置trance
		clientIP = c.ClientIP()

		traceContext := newTrace()
		if traceID := c.Request.Header.Get("com-header-rid"); traceID != "" {
			traceContext.TraceID = traceID
		}

		c.Set("trace", traceContext.TraceID)

		c.Next()
	}
}

func newTrace() *TraceInfo {
	trace := &TraceInfo{}
	trace.TraceID = getTraceID()
	return trace
}

func getTraceID() (traceID string) {
	return calcTraceID(clientIP)
}

func calcTraceID(ip string) (traceID string) {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	pid := os.Getpid()

	b := bytes.Buffer{}
	netIP := net.ParseIP(ip)

	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}

	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%04x", pid&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))
	b.WriteString("b0") // 末两位标记来源,b0为go

	return b.String()
}
