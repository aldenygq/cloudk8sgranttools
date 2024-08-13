package server 
import (
	"fmt"
	"os"
	cs20151215  "github.com/alibabacloud-go/cs-20151215/v5/client"
	openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util  "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
  )
func StartAckGrant(uid,role_name,role_type,cluster string,is_custom bool) {
	var (
		err error
		client *cs20151215.Client = &cs20151215.Client{}
	)
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
	  // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
	  AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
	  // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	  AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/CS
	config.Endpoint = tea.String("cs.cn-qingdao.aliyuncs.com")
	client, err = cs20151215.NewClient(config)
	if err != nil {
		fmt.Sprintf("初始化阿里云ack client失败,失败原因:%v\n",err)
		os.Exit(-1)
	}
	if role_type == "all-clusters" {
		cluster = ""
	}
	body0 := &cs20151215.GrantPermissionsRequestBody{
		IsCustom: tea.Bool(is_custom),
		IsRamRole: tea.Bool(true),
		RoleType: tea.String(role_type),
		RoleName: tea.String(role_name),
		Cluster: tea.String(cluster),
	  }
	grantPermissionsRequest := &cs20151215.GrantPermissionsRequest{
		Body: []*cs20151215.GrantPermissionsRequestBody{body0},
	}
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	// 复制代码运行请自行打印 API 的返回值
	_, err = client.GrantPermissionsWithOptions(tea.String(uid), grantPermissionsRequest, headers, runtime)
	if err != nil {
		fmt.Sprintf("阿里云集群授权失败，失败原因:%v\n",err)
		os.Exit(-1)
	}
	fmt.Sprintf("授权完成，授权信息如下:\n\tuid:%v\n\t,role_type:%v\n\trole_name:%v\n\t",uid,role_type,role_name)
	os.Exit(0)
}