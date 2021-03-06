package osquery

import (
	"errors"
	"fmt"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/spf13/viper"
)

type osqueryClient struct {
	client *osquery.ExtensionManagerClient
}

// getConfigParam parses the config file and returns the value of the specified parameter.
func getConfigParam(param string) string {
	viper.SetConfigName("osquery.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error when reading config file: %w \n", err))
	}

	return viper.GetString(param)
}

// getOsqueryClient prepares a connection to the osquery daemon.
func getOsqueryClient() *osqueryClient {
	client, err := osquery.NewClient(getConfigParam("osquery_socket"), 10*time.Second)
	if err != nil {
		panic(fmt.Errorf("Fatal error getting osquery socket: %w \n", err))
		return nil
	}
	return &osqueryClient{
		client: client,
	}
}

// execQuery is a general function for executing OSQuery queries.
func (c *osqueryClient) execQuery(sql string) ([]map[string]string, error) {
	response, err := c.client.Query(sql)
	if err != nil {
		return nil, err
	}
	return response.Response, nil
}

// getTableInfo is a shared function to get the first result of a row from a specified table.
func (c *osqueryClient) getTableInfo(table, column string) (string, error) {
	sql := "SELECT " + column + " FROM " + table + " LIMIT 1;"
	tmp, err := c.execQuery(sql)
	if err != nil {
		return "", err
	}
	if len(tmp) == 0 {
		return "", errors.New("no results found")
	}
	if val, exists := tmp[0][column]; exists {
		return val, nil
	}
	return "", errors.New("incorrect param or table specified")
}

// getOsInfo based on a parameter. The parameter is one of the rows of the
// os_version table of OSQuery: https://www.osquery.io/schema/5.1.0/#os_version
func (c *osqueryClient) getOsInfo(param string) (string, error) {
	return c.getTableInfo("os_version", param)
}

// getKernelInfo based on a parameter. The parameter is one of the rows of the
// kernel_info table of OSQuery: https://www.osquery.io/schema/5.1.0/#kernel_info
func (c *osqueryClient) getKernelInfo(param string) (string, error) {
	return c.getTableInfo("kernel_info", param)
}

// getProcessID gets the ID of the process which communicates on an open socket, specified by
// the source and destination IP addresses and ports. If no such socket is found,
// returns an empty string and an error.
func (c *osqueryClient) getProcessID(srcIp, dstIp, srcPort, dstPort string) (string, error) {
	var sql string

	if srcPort == "" || dstPort == "" {
		sql = "SELECT pid FROM process_open_sockets WHERE " +
			"(local_address='" + srcIp + "' AND " +
			"remote_address='" + dstIp + "') OR " +
			"(local_address='" + dstIp + "' AND " +
			"remote_address='" + srcIp + "') LIMIT 1;\r\n"
	} else {
		sql = "SELECT pid FROM process_open_sockets WHERE (" +
			"local_address='" + srcIp + "' AND " +
			"remote_address='" + dstIp + "' AND " +
			"local_port='" + srcPort + "' AND " +
			"remote_port='" + dstPort + "') OR (" +
			"local_address='" + dstIp + "' AND " +
			"remote_address='" + srcIp + "' AND " +
			"local_port='" + dstPort + "' AND " +
			"remote_port='" + srcPort + "') LIMIT 1;\r\n"
	}

	tmp, err := c.execQuery(sql)
	if err != nil {
		return "", err
	}
	if len(tmp) == 0 {
		return "", errors.New("socket not found")
	}
	return tmp[0]["pid"], nil
}

// getProcessName gets the name of the process based on the process ID. When no process
// with the specified PID is found, return empty string and error.
func (c *osqueryClient) getProcessName(pid string) (string, error) {
	sql := "SELECT name FROM processes WHERE pid='" + pid + "' LIMIT 1;"
	tmp, err := c.execQuery(sql)
	if err != nil {
		return "", err
	}
	if len(tmp) == 0 {
		return "", errors.New("process not found")
	}
	return tmp[0]["name"], nil
}
