// Factory for DNS providers
package dns

import (
	"fmt"

	"github.com/nicky-dev/lego/acmev2"
	"github.com/nicky-dev/lego/providers/dns/auroradns"
	"github.com/nicky-dev/lego/providers/dns/azure"
	"github.com/nicky-dev/lego/providers/dns/bluecat"
	"github.com/nicky-dev/lego/providers/dns/cloudflare"
	"github.com/nicky-dev/lego/providers/dns/cloudxns"
	"github.com/nicky-dev/lego/providers/dns/digitalocean"
	"github.com/nicky-dev/lego/providers/dns/dnsimple"
	"github.com/nicky-dev/lego/providers/dns/dnsmadeeasy"
	"github.com/nicky-dev/lego/providers/dns/dnspod"
	"github.com/nicky-dev/lego/providers/dns/duckdns"
	"github.com/nicky-dev/lego/providers/dns/dyn"
	"github.com/nicky-dev/lego/providers/dns/dynu"
	"github.com/nicky-dev/lego/providers/dns/exec"
	"github.com/nicky-dev/lego/providers/dns/exoscale"
	"github.com/nicky-dev/lego/providers/dns/fastdns"
	"github.com/nicky-dev/lego/providers/dns/gandi"
	"github.com/nicky-dev/lego/providers/dns/gandiv5"
	"github.com/nicky-dev/lego/providers/dns/glesys"
	"github.com/nicky-dev/lego/providers/dns/godaddy"
	"github.com/nicky-dev/lego/providers/dns/googlecloud"
	"github.com/nicky-dev/lego/providers/dns/lightsail"
	"github.com/nicky-dev/lego/providers/dns/linode"
	"github.com/nicky-dev/lego/providers/dns/namecheap"
	"github.com/nicky-dev/lego/providers/dns/namedotcom"
	"github.com/nicky-dev/lego/providers/dns/ns1"
	"github.com/nicky-dev/lego/providers/dns/otc"
	"github.com/nicky-dev/lego/providers/dns/ovh"
	"github.com/nicky-dev/lego/providers/dns/pdns"
	"github.com/nicky-dev/lego/providers/dns/rackspace"
	"github.com/nicky-dev/lego/providers/dns/rfc2136"
	"github.com/nicky-dev/lego/providers/dns/route53"
	"github.com/nicky-dev/lego/providers/dns/vultr"
)

func NewDNSChallengeProviderByName(name string) (acmev2.ChallengeProvider, error) {
	var err error
	var provider acmev2.ChallengeProvider
	switch name {
	case "azure":
		provider, err = azure.NewDNSProvider()
	case "auroradns":
		provider, err = auroradns.NewDNSProvider()
	case "bluecat":
		provider, err = bluecat.NewDNSProvider()
	case "cloudflare":
		provider, err = cloudflare.NewDNSProvider()
	case "cloudxns":
		provider, err = cloudxns.NewDNSProvider()
	case "digitalocean":
		provider, err = digitalocean.NewDNSProvider()
	case "dnsimple":
		provider, err = dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		provider, err = dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		provider, err = dnspod.NewDNSProvider()
	case "duckdns":
		provider, err = duckdns.NewDNSProvider()
	case "dyn":
		provider, err = dyn.NewDNSProvider()
	case "dynu":
		provider, err = dynu.NewDNSProvider()
	case "fastdns":
		provider, err = fastdns.NewDNSProvider()
	case "exoscale":
		provider, err = exoscale.NewDNSProvider()
	case "gandi":
		provider, err = gandi.NewDNSProvider()
	case "gandiv5":
		provider, err = gandiv5.NewDNSProvider()
	case "glesys":
		provider, err = glesys.NewDNSProvider()
	case "gcloud":
		provider, err = googlecloud.NewDNSProvider()
	case "godaddy":
		provider, err = godaddy.NewDNSProvider()
	case "lightsail":
		provider, err = lightsail.NewDNSProvider()
	case "linode":
		provider, err = linode.NewDNSProvider()
	case "manual":
		provider, err = acmev2.NewDNSProviderManual()
	case "namecheap":
		provider, err = namecheap.NewDNSProvider()
	case "namedotcom":
		provider, err = namedotcom.NewDNSProvider()
	case "rackspace":
		provider, err = rackspace.NewDNSProvider()
	case "route53":
		provider, err = route53.NewDNSProvider()
	case "rfc2136":
		provider, err = rfc2136.NewDNSProvider()
	case "vultr":
		provider, err = vultr.NewDNSProvider()
	case "ovh":
		provider, err = ovh.NewDNSProvider()
	case "pdns":
		provider, err = pdns.NewDNSProvider()
	case "ns1":
		provider, err = ns1.NewDNSProvider()
	case "otc":
		provider, err = otc.NewDNSProvider()
	case "exec":
		provider, err = exec.NewDNSProvider()
	default:
		err = fmt.Errorf("Unrecognised DNS provider: %s", name)
	}
	return provider, err
}
