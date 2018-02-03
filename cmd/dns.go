// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var dnsDomain string
var dnsZonefile string
var dnsResource string
var dnsName string

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Manage DNS",
	Long:  `Set and obtain DNS information held in Ethereum`,
}

// Reverse map of TypeToString in DNS package
var stringToType = map[string]uint16{}

func initStringToTypeMap() {
	for k, v := range dns.TypeToString {
		stringToType[v] = k
	}
}

func init() {
	RootCmd.AddCommand(dnsCmd)
	initStringToTypeMap()
}

func dnsFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&dnsDomain, "domain", "", "Domain against which to operate (e.g. wealdtech.eth)")
	cmd.Flags().StringVar(&dnsZonefile, "zonefile", "", "Path to DNS zone file")
	cmd.Flags().StringVar(&dnsResource, "resource", "", "The resource (A, NS, CNAME etc.)")
	cmd.Flags().StringVar(&dnsName, "name", "", "The name for the resource (end with \".\" for fully-qualified domain, otherwise domain will be added)")
}
