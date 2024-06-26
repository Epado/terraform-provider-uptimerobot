package provider

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	uptimerobotapi "github.com/Epado/terraform-provider-uptimerobot/internal/provider/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestUptimeRobotDataResourceMonitor_http_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http monitor"
	var Type = "http"
	var URL = "https://google.com"
	var URL2 = "https://yahoo.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL2),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_keyword_monitor(t *testing.T) {
	var FriendlyName = "TF Test: keyword"
	var Type = "keyword"
	var URL = "https://google.com"
	var KeywordType = "not exists"
	var KeywordType2 = "exists"
	var KeywordValue = "yahoo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_value", KeywordValue),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType2, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType2),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http port monitor"
	var Type = "port"
	var URL = "google.com"
	var URL2 = "yahoo.com"
	var SubType = "http"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL, SubType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL2, SubType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL2),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: custom port monitor"
	var Type = "port"
	var URL = "google.com"
	var SubType = "custom"
	var Port = 8080
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					sub_type      = "%s"
					port          = %d
				}
				`, FriendlyName, Type, URL, SubType, Port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "port", fmt.Sprintf(`%d`, Port)),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_ignore_ssl_errors(t *testing.T) {
	var FriendlyName = "TF Test:  custom ignore ssl errors"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name     = "%s"
					type              = "%s"
					url               = "%s"
					ignore_ssl_errors = true
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "ignore_ssl_errors", "true"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}
func TestUptimeRobotDataResourceMonitor_custom_alert_contact_threshold_and_recurrence(t *testing.T) {
	var FriendlyName = "TF Test: custom alert contact threshold & recurrence"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "SRE Team"
					type          = "e-mail"
					value         = "sre@Epado.com"
				}
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = uptimerobot_alert_contact.test.id
						threshold  = 0
						recurrence = 0
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.recurrence", "0"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_alert_contacts(t *testing.T) {
	var FriendlyName = "TF Test: custom alert contacts"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test1" {
					friendly_name = "Test 1"
					type          = "e-mail"
					value         = "test1@example.com"
				}

				resource "uptimerobot_alert_contact" "test2" {
					friendly_name = "Test 2"
					type          = "e-mail"
					value         = "test2@example.com"
				}

				resource "uptimerobot_alert_contact" "test3" {
					friendly_name = "Test 3"
					type          = "e-mail"
					value         = "test3@example.com"
				}

				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = uptimerobot_alert_contact.test1.id
						threshold  = 0
						recurrence = 0
					}
					alert_contact {
						id         = uptimerobot_alert_contact.test2.id
						threshold  = 0
						recurrence = 0
					}
					alert_contact {
						id         = uptimerobot_alert_contact.test3.id
						threshold  = 0
						recurrence = 0
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "3"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.recurrence", "0"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.1.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.1.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.1.recurrence", "0"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.2.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.2.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.2.recurrence", "0"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_http_headers(t *testing.T) {
	var FriendlyName = "TF Test:  custom http headers"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					custom_http_headers = {
						// Accept-Language = "en"
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "custom_http_headers.%", "0"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_http_status_codes(t *testing.T) {
	var FriendlyName = "TF Test:  custom http status codes"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					custom_http_headers = {
						// Accept-Language = "en"
					}
					custom_http_statuses = "404:1"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "custom_http_headers.%", "0"),
					//resource.TestMatchResourceAttr("uptimerobot_monitor.test", "custom_http_statuses", regexp.MustCompile(".*404:1.*")),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "custom_http_statuses", "404:1_200:0_201:0_202:0_203:0_204:0_205:0_206:0_207:0_208:0_226:0_300:0_301:0_302:0_303:0_304:0_305:0_306:0_307:0_308:0_400:0_401:0_402:0_403:0_405:0_406:0_407:0_408:0_409:0_410:0_411:0_412:0_413:0_414:0_415:0_416:0_417:0_418:0_421:0_422:0_423:0_424:0_426:0_428:0_429:0_430:0_431:0_440:0_449:0_450:0_451:0_495:0_496:0_497:0_499:0_500:0_501:0_502:0_503:0_504:0_505:0_508:0_509:0_510:0_511:0_520:0_521:0_522:0_523:0_524:0_525:0_526:0_527:0_530:0_598:0_599:0"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_ping_monitor(t *testing.T) {
	var FriendlyName = "TF Test: ping monitor"
	var Type = "ping"
	var URL = "1.1.1.1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_interval(t *testing.T) {
	var FriendlyName = "TF Test: custom interval"
	var Type = "ping"
	var URL = "1.1.1.1"
	var Interval = 300
	var Interval2 = 360
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval)),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval2)),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_method(t *testing.T) {
	var FriendlyName = "TF Test: http method monitor"
	var Type = "http"
	var URL = "https://httpbin.org/post"
	var Method = "POST"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_method    = "%s"
				}
				`, FriendlyName, Type, URL, Method),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_method", Method),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_auth_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http auth monitor"
	var Type = "http"
	var Username = "tester"
	var Password = "secret"
	var AuthType = "basic"
	var AuthType2 = "digest"
	var URL = fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", Username, Password)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType2),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_default_alert_contact(t *testing.T) {
	var FriendlyName = "TF Test: using the default alert contact"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`

				data "uptimerobot_account" "account" {}

				data "uptimerobot_alert_contact" "default" {
				friendly_name = data.uptimerobot_account.account.email
				}

				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = data.uptimerobot_alert_contact.default.id
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_method issue
				// ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMonitorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_monitor" {
			continue
		}

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = client.GetMonitor(id)

		if err == nil {
			return fmt.Errorf("Monitor still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
