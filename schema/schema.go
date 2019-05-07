package schema


import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/auth"
)


func BuildSchema() graphql.Schema {
	queryFields := graphql.Fields {
		"feerates": &graphql.Field {
			Type: feeRateEstimateType,
			Description: "Return feerate estimates, either satoshi-per-kw ({style} perkw) or satoshi-per-kb ({style} perkb).",
			Args: graphql.FieldConfigArgument {
				"style": &graphql.ArgumentConfig {
					Type: graphql.String,
					Description: "either perkw for satoshi-per-kw or perkb for satoshi-per-kb",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_feerates, authLevels, p)
			},
		},

		"getinfo": &graphql.Field {
			Type:  nodeinfoType,
			Description: "Get my node info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_getinfo, authLevels, p)
			},
		},

		"getroute": &graphql.Field {
			Type: graphql.NewList(routeHopType),
			Description: "get route",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "Id for getroute query. '' is all nodes.",
				},
				"msatoshis": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "msatoshis to send",
				},
				"riskfactor": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Float),
					Description: "risk factor",
				},
				"cltv": &graphql.ArgumentConfig {
					Type: graphql.Int,
					DefaultValue: 9,
					Description: "cltv (default is 9)",
				},
				"fromid": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "route from this id rather than current node",
				},
				"fuzzpercent": &graphql.ArgumentConfig {
					Type: graphql.Float,
					DefaultValue: 5.0,
					Description: "fuzz percent (default is 5.0)",
				},
				"exclude": &graphql.ArgumentConfig {
					Type: graphql.NewList(graphql.String),
					DefaultValue: make([]string, 0),
					Description: "channels to exclude from route",
				},
				"maxhops": &graphql.ArgumentConfig {
					Type: graphql.Int,
					DefaultValue: 20,
					Description: "max hops",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_getroute, authLevels, p)
			},
		},
		"listchannels": &graphql.Field {
			Type: graphql.NewList(channelType),
			Description: "List channels",
			Args: graphql.FieldConfigArgument {
				"scid": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "short channel id",
				},
				"source": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "source id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listchannels, authLevels, p)
			},
		},

		"listinvoices": &graphql.Field {
			Type: graphql.NewList(invoiceType),
			Description: "List invoices",
			Args: graphql.FieldConfigArgument {
				"label": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "list invoices. Opional label argument",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listinvoices, authLevels, p)
			},
		},

		"listnodes": &graphql.Field {
			Type: graphql.NewList(nodeType),
			Description: "Get a list of all nodes seen in network though channels and node announcement messages",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					DefaultValue: "",
					Description: "Id for listnodes query. '' is all nodes.",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listnodes, authLevels, p)
			},
		},

                "listpeers": &graphql.Field {
			Type:  graphql.NewList(peerType),
			Description: "List peers",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "Id for listpeers query. '' is all peers.",
				},
				"level": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "choose level of logs",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listpeers, authLevels, p)
			},
		},

                "waitanyinvoice": &graphql.Field {
			Type:  completedInvoiceType,
			Description: "wait for an invoice to be paid",
			Args: graphql.FieldConfigArgument {
				"lastpay_index": &graphql.ArgumentConfig {
					Type: graphql.Int,
					DefaultValue: 0,
					Description: "wait for an invoice to be paid after this index if this index is supplied as a param/arg",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_waitanyinvoice, authLevels, p)
			},
		},

                "waitinvoice": &graphql.Field {
			Type:  completedInvoiceType,
			Description: "wait for an invoice specified by a label to be paid",
			Args: graphql.FieldConfigArgument {
				"label": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "wait for the invoice with this label to be paid",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_waitinvoice, authLevels, p)
			},
		},
	}

	mutationFields := graphql.Fields {
		"connect": &graphql.Field {
			Type: graphql.String,
			Description: "Connect to {id} at {host}:{port}",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "id of peer",
				},
				"host": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "address of peer. It can be tor, ipv4 or ipv6",
				},
				"port": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
					Description: "port of peer",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Peers}
				return auth.AuthWrapper(r_connect, authLevels, p)
			},
		},
		"delinvoice": &graphql.Field {
			Type: graphql.String,
			Description: "delete invoice with label and status as non-optional params",
			Args: graphql.FieldConfigArgument {
				"label": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "label of invoice",
				},
				"status": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "status of invoice to be deleted",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Invoices}
				return auth.AuthWrapper(r_delinvoice, authLevels, p)
			},
		},
		"invoice": &graphql.Field {
			Type: invoiceType,
			Description: "Create new lightning payment invoice",
			Args: graphql.FieldConfigArgument {
				"msatoshis": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "msatoshis to send",
				},
				"label": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "list invoices. Opional label argument",
				},
				"description": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "invoice description",
				},
				//optional args
				"expiry": &graphql.ArgumentConfig {
					Type: graphql.Int,
					DefaultValue: 3600, //1 hour
					Description: "number of seconds the invoice is valid for",
				},
				"fallbacks": &graphql.ArgumentConfig {
					Type: graphql.NewList(graphql.String),
					DefaultValue: make([]string, 0),
					Description: "list invoices. Opional label argument",
				},
				"preimage": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "the preimage of the htlc chain",
				},
				"exposeprivatechannels": &graphql.ArgumentConfig {
					Type: graphql.Boolean,
					DefaultValue: false,
					Description: "exposing channels that are not public to all nodes in the lightning network",
				},

			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Invoices}
				return auth.AuthWrapper(r_invoice, authLevels, p)
			},
		},
		"pay": &graphql.Field {
			Type: paymentSuccessType,
			Description: "Pay via bolt11 as argument",
			Args: graphql.FieldConfigArgument {
				"bolt11": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String), //non null means that argument is required
					Description: "full bolt11 invoice to pay to",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Funds}
				return auth.AuthWrapper(r_pay, authLevels, p)
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
	mutations := graphql.ObjectConfig{Name: "Mutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(mutations)}
	schema, _ := graphql.NewSchema(schemaConfig)
        return schema
}




