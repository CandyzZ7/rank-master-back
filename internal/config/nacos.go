package config

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"

	"rank-master-back/infrastructure/pkg/xlogger"
)

type (
	BootstrapConfig struct {
		NacosConfig NacosConfig
	}

	ListenConfig func(data string)

	NacosServerConfig struct {
		IpAddr string
		Port   uint64
	}

	NacosClientConfig struct {
		NamespaceId         string
		TimeoutMs           uint64
		NotLoadCacheAtStart bool
		LogDir              string
		LogLevel            string
		UserName            string
		Password            string
		AppendToStdout      bool
	}

	NacosConfig struct {
		ServerConfigs []NacosServerConfig
		ClientConfig  NacosClientConfig
		DataId        string
		Group         string
	}
)

func (n *NacosConfig) Discovery(c Config) {
	sc, cc := n.buildConfig()
	opts := nacos.NewNacosConfig(c.Name, fmt.Sprintf("%s:%d", c.Host, c.Port), sc, &cc)
	err := nacos.RegisterService(opts)
	if err != nil {
		panic(err)
	}
}

func (n *NacosConfig) InitConfig(listenConfigCallback ListenConfig) string {
	// nacos server
	logc.SetLevel(3)
	logger.SetLogger(xlogger.DefaultNacosLogger)
	sc, cc := n.buildConfig()

	pa := vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
	configClient, err := clients.NewConfigClient(pa)
	if err != nil {
		panic(err)
	}
	// 获取配置中心内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
	})
	if err != nil {
		panic(err)
	}
	// 设置配置监听
	if err = configClient.ListenConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
		OnChange: func(namespace, group, dataId, data string) {
			// 配置文件产生变化就会触发
			if len(data) == 0 {
				logx.Errorf("listen nacos data nil error ,  namespace : %s，group : %s , dataId : %s , data : %s")
				return
			}
			listenConfigCallback(data)
		},
	}); err != nil {
		panic(err)
	}

	if len(content) == 0 {
		panic("read nacos config  content err , content is nil")
	}

	return content
}

func (n *NacosConfig) buildConfig() ([]constant.ServerConfig, constant.ClientConfig) {
	var sc []constant.ServerConfig
	if len(n.ServerConfigs) == 0 {
		panic("nacos server no set")
	}
	for _, serveConfig := range n.ServerConfigs {
		sc = append(sc, constant.ServerConfig{
			Port:   serveConfig.Port,
			IpAddr: serveConfig.IpAddr,
		},
		)
	}

	// nacos client
	cc := constant.ClientConfig{
		NamespaceId:         n.ClientConfig.NamespaceId,
		TimeoutMs:           n.ClientConfig.TimeoutMs,
		NotLoadCacheAtStart: n.ClientConfig.NotLoadCacheAtStart,
		LogDir:              n.ClientConfig.LogDir,
		LogLevel:            n.ClientConfig.LogLevel,
		Username:            n.ClientConfig.UserName,
		Password:            n.ClientConfig.Password,
		AppendToStdout:      n.ClientConfig.AppendToStdout,
	}
	return sc, cc
}

func LoadNacosConfig(configFile *string) Config {
	var bootstrapConfig BootstrapConfig
	conf.MustLoad(*configFile, &bootstrapConfig)
	// 解析业务配置
	var c Config
	nacosConfig := bootstrapConfig.NacosConfig
	serviceConfigContent := nacosConfig.InitConfig(func(data string) {
		err := conf.LoadFromYamlBytes([]byte(data), &c)
		if err != nil {
			panic(err)
		}
	})
	err := conf.LoadFromYamlBytes([]byte(serviceConfigContent), &c)
	if err != nil {
		panic(err)
	}

	// 注册到nacos
	nacosConfig.Discovery(c)
	return c
}
