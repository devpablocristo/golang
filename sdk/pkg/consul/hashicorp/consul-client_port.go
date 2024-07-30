package cslhash

import "github.com/hashicorp/consul/api"

type ConsulClientPort interface {
	Client() *api.Client
}
