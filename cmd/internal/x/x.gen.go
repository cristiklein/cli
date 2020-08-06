// Code generated by openapi-cli-generator. DO NOT EDIT.
// See https://github.com/danielgtaylor/openapi-cli-generator

package x

import (
	"strings"

	"github.com/danielgtaylor/openapi-cli-generator/cli"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gentleman.v2"
)

var xSubcommand bool

func xServers() []map[string]string {
	return []map[string]string{

		map[string]string{
			"description": "",
			"url":         "{protocol}://{environment}-{zone}.exoscale.com/v2.alpha",
		},
	}
}

// XListInstanceTypes list-instance-types
func XListInstanceTypes(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "list-instance-types"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/instance-type"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XGetInstanceType get-instance-type
func XGetInstanceType(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "get-instance-type"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/instance-type/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XListLoadBalancers list-load-balancers
func XListLoadBalancers(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "list-load-balancers"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XCreateLoadBalancer create-load-balancer
func XCreateLoadBalancer(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "create-load-balancer"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer"

	req := cli.Client.Post().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XDeleteLoadBalancer delete-load-balancer
func XDeleteLoadBalancer(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "delete-load-balancer"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Delete().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XGetLoadBalancer get-load-balancer
func XGetLoadBalancer(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "get-load-balancer"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XUpdateLoadBalancer update-load-balancer
func XUpdateLoadBalancer(paramId string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "update-load-balancer"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Put().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XAddServiceToLoadBalancer add-service-to-load-balancer
func XAddServiceToLoadBalancer(paramId string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "add-service-to-load-balancer"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}/service"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Post().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XGetLoadBalancerService get-load-balancer-service
func XGetLoadBalancerService(paramId string, paramServiceId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "get-load-balancer-service"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}/service/{service-id}"
	url = strings.Replace(url, "{id}", paramId, 1)
	url = strings.Replace(url, "{service-id}", paramServiceId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XUpdateLoadBalancerService update-load-balancer-service
func XUpdateLoadBalancerService(paramId string, paramServiceId string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "update-load-balancer-service"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}/service/{service-id}"
	url = strings.Replace(url, "{id}", paramId, 1)
	url = strings.Replace(url, "{service-id}", paramServiceId, 1)

	req := cli.Client.Put().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XDeleteLoadBalancerService delete-load-balancer-service
func XDeleteLoadBalancerService(paramId string, paramServiceId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "delete-load-balancer-service"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/load-balancer/{id}/service/{service-id}"
	url = strings.Replace(url, "{id}", paramId, 1)
	url = strings.Replace(url, "{service-id}", paramServiceId, 1)

	req := cli.Client.Delete().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XListOperations list-operations
func XListOperations(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "list-operations"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/operation"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XGetOperation get-operation
func XGetOperation(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "get-operation"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/operation/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// XListZones list-zones
func XListZones(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "list-zones"
	if xSubcommand {
		handlerPath = "x " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = xServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/zone"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

func xRegister(subcommand bool) {
	root := cli.Root

	if subcommand {
		root = &cobra.Command{
			Use:   "x",
			Short: "Exoscale Public API",
			Long:  cli.Markdown("Infrastructure automation API, allowing programmatic access to all Exoscale products and services."),
		}
		xSubcommand = true
	} else {
		cli.Root.Short = "Exoscale Public API"
		cli.Root.Long = cli.Markdown("Infrastructure automation API, allowing programmatic access to all Exoscale products and services.")
	}

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "list-instance-types",
			Short:   "list-instance-types",
			Long:    cli.Markdown("List Compute instance types"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XListInstanceTypes(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "get-instance-type id",
			Short:   "get-instance-type",
			Long:    cli.Markdown("Show instance type details"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XGetInstanceType(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "list-load-balancers",
			Short:   "list-load-balancers",
			Long:    cli.Markdown("List load balancers"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XListLoadBalancers(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "create-load-balancer",
			Short:   "create-load-balancer",
			Long:    cli.Markdown("Create a new load balancer\n## Request Schema (application/json)\n\nproperties:\n  created-at:\n    format: date-time\n    readOnly: true\n    type: string\n  description:\n    type: string\n  id:\n    format: uuid\n    readOnly: true\n    type: string\n  ip:\n    format: ipv4\n    readOnly: true\n    type: string\n  name:\n    type: string\n  services:\n    items:\n      $ref: '#/components/schemas/load-balancer-service'\n    type: array\n  state:\n    enum:\n    - creating\n    - deleting\n    - running\n    - error\n    readOnly: true\n    type: string\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[0:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := XCreateLoadBalancer(params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "delete-load-balancer id",
			Short:   "delete-load-balancer",
			Long:    cli.Markdown("Delete a load balancer"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XDeleteLoadBalancer(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "get-load-balancer id",
			Short:   "get-load-balancer",
			Long:    cli.Markdown("Show load balancer details"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XGetLoadBalancer(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "update-load-balancer id",
			Short:   "update-load-balancer",
			Long:    cli.Markdown("Update a load balancer\n## Request Schema (application/json)\n\nproperties:\n  description:\n    type: string\n  name:\n    type: string\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[1:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := XUpdateLoadBalancer(args[0], params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "add-service-to-load-balancer id",
			Short:   "add-service-to-load-balancer",
			Long:    cli.Markdown("Add a single service to a load balancer\n## Request Schema (application/json)\n\nproperties:\n  description:\n    type: string\n  healthcheck:\n    $ref: '#/components/schemas/healthcheck'\n  healthcheck-status:\n    items:\n      $ref: '#/components/schemas/load-balancer-server-status'\n    readOnly: true\n    type: array\n  id:\n    format: uuid\n    readOnly: true\n    type: string\n  instance-pool:\n    $ref: '#/components/schemas/resource'\n  name:\n    type: string\n  port:\n    format: int64\n    type: integer\n  protocol:\n    enum:\n    - tcp\n    - udp\n    type: string\n  state:\n    enum:\n    - creating\n    - deleting\n    - running\n    - updating\n    - error\n    readOnly: true\n    type: string\n  strategy:\n    enum:\n    - round-robin\n    - source-hash\n    type: string\n  target-port:\n    format: int64\n    type: integer\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[1:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := XAddServiceToLoadBalancer(args[0], params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "get-load-balancer-service id service-id",
			Short:   "get-load-balancer-service",
			Long:    cli.Markdown("Show load balancer service details"),
			Example: examples,
			Args:    cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XGetLoadBalancerService(args[0], args[1], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "update-load-balancer-service id service-id",
			Short:   "update-load-balancer-service",
			Long:    cli.Markdown("Update a load balancer service name\n## Request Schema (application/json)\n\nproperties:\n  description:\n    type: string\n  healthcheck:\n    $ref: '#/components/schemas/healthcheck'\n  name:\n    type: string\n  port:\n    format: int64\n    type: integer\n  protocol:\n    enum:\n    - tcp\n    - udp\n    type: string\n  strategy:\n    enum:\n    - round-robin\n    - source-hash\n    type: string\n  target-port:\n    format: int64\n    type: integer\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[2:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := XUpdateLoadBalancerService(args[0], args[1], params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "delete-load-balancer-service id service-id",
			Short:   "delete-load-balancer-service",
			Long:    cli.Markdown("Delete a single load balancer service from a load balancer"),
			Example: examples,
			Args:    cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XDeleteLoadBalancerService(args[0], args[1], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "list-operations",
			Short:   "list-operations",
			Long:    cli.Markdown("Lookup recent operation statuses"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XListOperations(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "get-operation id",
			Short:   "get-operation",
			Long:    cli.Markdown("Lookup operation status by ID"),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XGetOperation(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "list-zones",
			Short:   "list-zones",
			Long:    cli.Markdown("List zones"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := XListZones(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

}