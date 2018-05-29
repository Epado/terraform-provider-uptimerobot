package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type UptimeRobotConfig struct {
	apiKey string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_alert_contact": resourceAlertContact(),
			"uptimerobot_monitor":       resourceMonitor(),
		},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := UptimeRobotConfig{
				apiKey: r.Get("api_key").(string),
			}
			return config, nil
		},
	}
}