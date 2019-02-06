#!/usr/bin/env bash

test_description="Test ping over QUIC command"

. lib/test-lib.sh

test_init_ipfs

# start iptb + wait for peering
test_expect_success 'init iptb' '
  iptb init -n 2 --bootstrap=none --port=0
'

test_expect_success "enable QUIC experiment" '
  ipfsi 0 config --json Experimental.QUIC true &&
  ipfsi 1 config --json Experimental.QUIC true
'

addr1='"[\"/ip4/127.0.0.1/udp/0/quic/\"]"'
addr2='"[\"/ip4/127.0.0.1/udp/0/quic/\"]"'
test_expect_success "add QUIC swarm addresses" '
  ipfsi 0 config --json Addresses.Swarm '$addr1' &&
  ipfsi 1 config --json Addresses.Swarm '$addr2'
'

startup_cluster 2

test_expect_success 'peer ids' '
  PEERID_0=$(iptb get id 0) &&
  PEERID_1=$(iptb get id 1)
'

test_expect_success "test ping other" '
  ipfsi 0 ping -n2 -- "$PEERID_1" &&
  ipfsi 1 ping -n2 -- "$PEERID_0"
'

test_expect_success "test ping self" '
  test_must_fail ipfsi 0 ping -n2 -- "$PEERID_0" &&
  test_must_fail ipfsi 1 ping -n2 -- "$PEERID_1"
'

test_expect_success "test ping 0" '
  test_must_fail ipfsi 0 ping -n0 -- "$PEERID_1" &&
  test_must_fail ipfsi 1 ping -n0 -- "$PEERID_0"
'

test_expect_success 'stop iptb' '
  iptb stop
'

test_done
