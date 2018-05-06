package dynu

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	dynuLiveTest     bool
	dynuCustomerName string
	dynuUserName     string
	dynuPassword     string
	dynuDomain       string
)

func init() {
	dynuCustomerName = os.Getenv("DYNU_CUSTOMER_NAME")
	dynuUserName = os.Getenv("DYNU_USER_NAME")
	dynuPassword = os.Getenv("DYNU_PASSWORD")
	dynuDomain = os.Getenv("DYNU_DOMAIN")
	if len(dynuCustomerName) > 0 && len(dynuUserName) > 0 && len(dynuPassword) > 0 && len(dynuDomain) > 0 {
		dynuLiveTest = true
	}
}

func TestLiveDynPresent(t *testing.T) {
	if !dynuLiveTest {
		t.Skip("skipping live test")
	}

	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.Present(dynuDomain, "", "123d==")
	assert.NoError(t, err)
}

func TestLiveDynCleanUp(t *testing.T) {
	if !dynuLiveTest {
		t.Skip("skipping live test")
	}

	time.Sleep(time.Second * 1)

	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.CleanUp(dynuDomain, "", "123d==")
	assert.NoError(t, err)
}
