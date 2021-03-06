/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"flag"
	"fmt"
	"github.com/nebulaim/telegramd/baselib/mysql_client"
	"github.com/BurntSushi/toml"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/proto/zproto"
	"github.com/nebulaim/telegramd/baselib/redis_client"
)

var (
	confPath string
	Conf     *syncConfig
)

type syncConfig struct {
	ServerId      int32 // 服务器ID
	Redis         []redis_client.RedisConfig
	Mysql         []mysql_client.MySQLConfig
	Server        *grpc_util.RPCServerConfig
	SessionClient *zproto.ZProtoClientConfig
}

func (c *syncConfig) String() string {
	return fmt.Sprintf("{server_id: %d, redis: %v. mysql: %v, server: %v, sessionClient: %v}",
		c.ServerId,
		c.Redis,
		c.Mysql,
		c.Server,
		c.SessionClient)
}

//type rpcServerConfig struct {
//	Addr string
//}
//
//type syncConfig struct {
//	Server        *rpcServerConfig
//	Discovery     service_discovery.ServiceDiscoveryServerConfig
//	Redis         []redis_client.RedisConfig
//	Mysql         []mysql_client.MySQLConfig
//	SessionClient *ClientConfig
//}
//
//type ClientConfig struct {
//	Name      string
//	ProtoName string
//	AddrList  []string
//	EtcdAddrs []string
//	Balancer  string
//}

func init() {
	flag.StringVar(&confPath, "conf", "./sync.toml", "config path")
}

func InitializeConfig() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		err = fmt.Errorf("decode file %s error: %v", confPath, err)
	}
	return
}
