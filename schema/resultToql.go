package schema


import (
        "github.com/niftynei/glightning/glightning"
        "strconv"
)



func peerToql(peer glightning.Peer) Peer_ql {
        var ql Peer_ql
	ql.Id = peer.Id
	ql.Connected = peer.Connected
	ql.NetAddresses = peer.NetAddresses
	ql.GlobalFeatures = peer.GlobalFeatures
	ql.LocalFeatures = peer.LocalFeatures
	for _, channel := range peer.Channels {
	        ql.Channels = append(ql.Channels, peerChannelToql(channel))
	}
	ql.Logs = peer.Logs
	return ql
}

func peerChannelToql(peerChannel glightning.PeerChannel) PeerChannel_ql {
	var ql PeerChannel_ql
	ql.State = peerChannel.State
	ql.ScratchTxId = peerChannel.ScratchTxId
	ql.Owner = peerChannel.Owner
	ql.ShortChannelId = peerChannel.ShortChannelId
	ql.ChannelDirection = peerChannel.ChannelDirection
	ql.ChannelId = peerChannel.ChannelId
	ql.FundingTxId = peerChannel.FundingTxId
	ql.Funding = peerChannel.Funding
	ql.Status = peerChannel.Status
	ql.Private = peerChannel.Private
	for key, val := range peerChannel.FundingAllocations {
		ql.FundingAllocations = append(ql.FundingAllocations, FundingAllocations_ql{key, strconv.FormatUint(val, 10)})
	}
	ql.MilliSatoshiToUs = strconv.FormatUint(peerChannel.MilliSatoshiToUs, 10)
	ql.MilliSatoshiToUsMin = strconv.FormatUint(peerChannel.MilliSatoshiToUsMin, 10)
	ql.MilliSatoshiToUsMax = strconv.FormatUint(peerChannel.MilliSatoshiToUsMax, 10)
	ql.MilliSatoshiTotal = strconv.FormatUint(peerChannel.MilliSatoshiTotal, 10)
	ql.DustLimitSatoshi = strconv.FormatUint(peerChannel.DustLimitSatoshi, 10)
	ql.MaxHtlcValueInFlightMilliSatoshi = strconv.FormatUint(peerChannel.MaxHtlcValueInFlightMilliSatoshi, 10)
	ql.TheirChannelReserveSatoshi = strconv.FormatUint(peerChannel.TheirChannelReserveSatoshi, 10)
	ql.OurChannelReserveSatoshi = strconv.FormatUint(peerChannel.OurChannelReserveSatoshi, 10)
	ql.SpendableMilliSatoshi = strconv.FormatUint(peerChannel.SpendableMilliSatoshi, 10)
	ql.HtlcMinMilliSatoshi = strconv.FormatUint(peerChannel.HtlcMinMilliSatoshi, 10)
	ql.TheirToSelfDelay = peerChannel.TheirToSelfDelay
	ql.OurToSelfDelay = peerChannel.OurToSelfDelay
	ql.MaxAcceptedHtlcs = peerChannel.MaxAcceptedHtlcs
	ql.InPaymentsOffered = strconv.FormatUint(peerChannel.InPaymentsOffered, 10)
	ql.InMilliSatoshiOffered = strconv.FormatUint(peerChannel.InMilliSatoshiOffered, 10)
	ql.InPaymentsFulfilled = strconv.FormatUint(peerChannel.InPaymentsFulfilled, 10)
	ql.InMilliSatoshiFulfilled = strconv.FormatUint(peerChannel.InMilliSatoshiFulfilled, 10)
	ql.OutPaymentsOffered = strconv.FormatUint(peerChannel.OutPaymentsOffered, 10)
	ql.OutMilliSatoshiOffered = strconv.FormatUint(peerChannel.OutMilliSatoshiOffered, 10)
	ql.OutPaymentsFulfilled = strconv.FormatUint(peerChannel.OutPaymentsFulfilled, 10)
	ql.OutMilliSatoshiFulfilled = strconv.FormatUint(peerChannel.OutMilliSatoshiFulfilled, 10)
	for _, htlc := range peerChannel.Htlcs {
		ql.Htlcs = append(ql.Htlcs, htlcToql(htlc))
	}
	return ql
}

func htlcToql(htlc *glightning.Htlc) Htlc_ql {
	var ql Htlc_ql
	ql.Direction = htlc.Direction
	ql.Id = strconv.FormatUint(htlc.Id, 10)
	ql.MilliSatoshi = strconv.FormatUint(htlc.MilliSatoshi, 10)
	ql.Expiry = strconv.FormatUint(htlc.Expiry, 10)
	ql.PaymentHash = htlc.PaymentHash
	ql.State = htlc.State
	return ql
}


func invoiceToql(invoice glightning.Invoice) Invoice_ql {
	var ql Invoice_ql
	ql.PaymentHash = invoice.PaymentHash
	ql.ExpiresAt = strconv.FormatUint(invoice.ExpiresAt, 10)
	ql.Bolt11 = invoice.Bolt11
	ql.WarningOffline = invoice.WarningOffline
	ql.WarningCapacity = invoice.WarningCapacity
	ql.Label = ql.Label
	ql.Status = ql.Status
	ql.Description = ql.Description
	return ql
}


func nodeToNodeInfo(nodeinfo *glightning.NodeInfo, nodeinfo_ql *NodeInfo_ql) {
	nodeinfo_ql.Id = nodeinfo.Id
	nodeinfo_ql.Alias = nodeinfo.Alias
	nodeinfo_ql.Color = nodeinfo.Color
	nodeinfo_ql.PeerCount = nodeinfo.PeerCount
	nodeinfo_ql.PendingChannelCount = nodeinfo.PendingChannelCount
        nodeinfo_ql.ActiveChannelCount = nodeinfo.ActiveChannelCount
	nodeinfo_ql.InactiveChannelCount = nodeinfo.InactiveChannelCount
	nodeinfo_ql.Addresses = nodeinfo.Addresses
        nodeinfo_ql.Binding = nodeinfo.Binding
	nodeinfo_ql.Version = nodeinfo.Version
	nodeinfo_ql.Blockheight = nodeinfo.Blockheight
	nodeinfo_ql.Network = nodeinfo.Network
	nodeinfo_ql.FeesCollectedMilliSatoshis = strconv.FormatUint(nodeinfo.FeesCollectedMilliSatoshis, 10)
}
