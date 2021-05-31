package kubernetes

import (
	"github.com/influxdata/telegraf/plugins/common/tls"
	"github.com/stretchr/testify/assert"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	"testing"
)

func TestInitCfg(t *testing.T) {
	i := &Input{
		Tags:              make(map[string]string),
		URL:               "https://172.16.2.41:6443",
		BearerTokenString: "eyJhbGciOiJSUzI1NiIsImtpZCI6InFWNzd1LTNDNEdEd0FlTjdPQzF1NXBGVnYxU2JrTlVJQ3RUUnZlbXRGZ1EifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4ta3F4NzUiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjM5ZmQxOTQ4LTY5YTAtNDZlZi1hZjc3LWYxYzUwMmFmZDdiMiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.f4oPuQ0fuY1jZI7o7CeGr-FtfQbnYlxzphtZAeKo31HjQAG5ynl4rYLRt1PK7lpCoMiMrAw5xDSMlG2DN9bTF3OYQJbfC4Mq3olPGxHHjxoTSotrfGrMK779NZ_JzRw6OQ9mKEgG8vadFpd4nGRi4KuD-7w8ysOzm_j6Z78eVTxhKrOuU11a6WEUh_LGnJSNLjAdN8xKqim90qcWy5jvdYl2s9N2tRPvkSJ22xwJ9Icts0HHZfvAywG7Rb69WyN13ct37N1_bICwjVrWuONyXOgNSiV7JvUFI2ZFpKfpDrDhpGRwwmVCR5a8BjP0S1kNjjckK9ma4ubYyvLIDS86Xw",
		ClientConfig: tls.ClientConfig{
			TLSCA: "/Users/liushaobo/go/src/gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kubernetes/pki/cacert.pem",
		},
	}

	err := i.initCfg()
	t.Log("error ---->", err)

	// 通过 ServerVersion 方法来获取版本号
	versionInfo, err := i.client.ServerVersion()
	if err != nil {
		assert.Error(t, err, "")
	}

	t.Log("version ==>", versionInfo.String())
}

func TestMain(t *testing.T) {
	// t.Run("config", func(t *testing.T) {
	// 	i := &Input{
	// 		Tags:           make(map[string]string),
	// 		KubeConfigPath: "/Users/liushaobo/.kube/config",
	// 	}

	// 	err := i.initCfg()
	// 	t.Log("error ---->", err)

	// 	i.Collect()

	// 	for _, obj := range i.collectCache {
	// 		point, err := obj.LineProto()
	// 		if err != nil {
	// 			t.Log("error ->", err)
	// 		} else {
	// 			t.Log("point ->", point.String())
	// 		}
	// 	}
	// })

	t.Run("bear token", func(t *testing.T) {
		i := &Input{
			Tags:              make(map[string]string),
			URL:               "https://172.16.2.41:6443",
			BearerTokenString: "eyJhbGciOiJSUzI1NiIsImtpZCI6InFWNzd1LTNDNEdEd0FlTjdPQzF1NXBGVnYxU2JrTlVJQ3RUUnZlbXRGZ1EifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4ta3F4NzUiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjM5ZmQxOTQ4LTY5YTAtNDZlZi1hZjc3LWYxYzUwMmFmZDdiMiIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.f4oPuQ0fuY1jZI7o7CeGr-FtfQbnYlxzphtZAeKo31HjQAG5ynl4rYLRt1PK7lpCoMiMrAw5xDSMlG2DN9bTF3OYQJbfC4Mq3olPGxHHjxoTSotrfGrMK779NZ_JzRw6OQ9mKEgG8vadFpd4nGRi4KuD-7w8ysOzm_j6Z78eVTxhKrOuU11a6WEUh_LGnJSNLjAdN8xKqim90qcWy5jvdYl2s9N2tRPvkSJ22xwJ9Icts0HHZfvAywG7Rb69WyN13ct37N1_bICwjVrWuONyXOgNSiV7JvUFI2ZFpKfpDrDhpGRwwmVCR5a8BjP0S1kNjjckK9ma4ubYyvLIDS86Xw",
			ClientConfig: tls.ClientConfig{
				TLSCA: "/Users/liushaobo/go/src/gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kubernetes/pki/cacert.pem",
			},
		}

		err := i.initCfg()
		t.Log("error ---->", err)

		i.Collect()

		t.Log("======>", i.collectCache)

		for _, obj := range i.collectCache {
			point, err := obj.LineProto()
			if err != nil {
				t.Log("error ->", err)
			} else {
				t.Log("point ->", point.String())
			}
		}
	})
}

func TestServiceAccount(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		i := &Input{
			Tags:           make(map[string]string),
			KubeConfigPath: "/Users/liushaobo/.kube/config",
		}

		err := i.initCfg()
		t.Log("error ---->", err)

		i.Collect()

		for _, obj := range i.collectCache {
			point, err := obj.LineProto()
			if err != nil {
				t.Log("error ->", err)
			} else {
				t.Log("point ->", point.String())
			}
		}
	})

	t.Run("bear token", func(t *testing.T) {
		i := &Input{
			Tags:              make(map[string]string),
			URL:               "https://172.16.2.41:6443",
			BearerTokenString: "eyJhbGciOiJSUzI1NiIsImtpZCI6InFWNzd1LTNDNEdEd0FlTjdPQzF1NXBGVnYxU2JrTlVJQ3RUUnZlbXRGZ1EifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkYXRha2l0LW1vbml0b3IiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGF0YWtpdC1tb25pdG9yLXRva2VuLWo5OWJkIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRhdGFraXQtbW9uaXRvciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6Ijg4NjI5NmMzLTQ0MGYtNGVjYy1iZjI2LTQ1Y2MwMTI3NGJmZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkYXRha2l0LW1vbml0b3I6ZGF0YWtpdC1tb25pdG9yIn0.BTd-SscJy5OlDykLnEIHAFVD1WN16bIMp5lxklXHMS1m2wVCGE4ttls2vMTdqZWThlSq78iTdnXaaG3pZsCaT7NLxjlON_uoUUcjq4V8vgeoiy-iinIgfvnUyOxTY84hvSSfJEM2DmVZsRV6cgk9sVbvLu4MEhfQN1g0N2pX2ZpIL8GdPjzIw6lxSrv1Zh4GqvOwHcHxS5ZL8JM-ahkta_XwKqKhxmMxuX0L62HJP9Df4hYzbOYT80R4-a0twEGqTF4Lu7wWAHwT3Y7gEGXHdH30LTq2KV_P1uS9ykAtjfrMIiOqBUgSqnC8h04N02hb1P_lSWBjD5xOZtnfjI8AkQ",
			ClientConfig: tls.ClientConfig{
				TLSCA: "/Users/liushaobo/go/src/gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/kubernetes/pki/cacert.pem",
			},
		}

		err := i.initCfg()
		t.Log("error ---->", err)

		i.Collect()

		t.Log("======>", i.collectCache)

		for _, obj := range i.collectCache {
			point, err := obj.LineProto()
			if err != nil {
				t.Log("error ->", err)
			} else {
				t.Log("point ->", point.String())
			}
		}
	})
}

func TestLoadCfg(t *testing.T) {
	arr, err := config.LoadInputConfigFile("./cfg.conf", func() inputs.Input {
		return &Input{}
	})

	if err != nil {
		t.Fatalf("%s", err)
	}

	kube := arr[0].(*Input)

	t.Log("url ---->", kube.URL)
	t.Log("token ---->", kube.BearerTokenString)
	t.Log("ca ---->", kube.TLSCA)
}

func TestRun(t *testing.T) {
	arr, err := config.LoadInputConfigFile("./cfg.conf", func() inputs.Input {
		return &Input{}
	})

	if err != nil {
		t.Fatalf("%s", err)
	}

	datakit.OutputFile = "./res.dat"

	kube := arr[0].(*Input)
	kube.Run()
}
