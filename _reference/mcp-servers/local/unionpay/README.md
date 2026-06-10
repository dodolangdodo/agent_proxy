## 1. 简介

> <https://www.npmjs.com/package/unionpay-mcp-server>

unionpay-mcp-server是银联基于MCP协议为AI智能体提供的支付工具（MCP Server），各类支持MCP协议的智能体应用均可安全、便捷地接入银联支付能力。以下是一个行程规划智能体为用户提供行程规划和酒店预定的示例：
传统方式下，用户需要自主查询酒店、对比酒店价格，并完成预订和支付。 智能体模式下，用户只需在智能体应用中输入出行需求（如出发地、目的地、出行时间、预算偏好等），智能体即可自动分析需求，并推荐最佳酒店选择方案，用户确认后，智能体通过银联MCP Server生成支付订单，待用户支付后，智能体完成线上酒店预订操作，并将订单信息同步给用户。整个流程无需人工反复查询和操作，高效便捷。

## 2. 使用和配置

使用前，需先注册成为银联网络商户，并开通业务权限，获取商户私钥。之后，即可在主流的支持MCP的客户端应用上使用银联MCP Server中的各类支付工具。

### 在 Cursor 中使用

在 Cursor 项目中的 `.cursor/mcp.json` 加入如下配置：

```json
{
  "mcpServers": {
    "unionpay-mcp-server": {
 "command": "npx",
 "args": [
  "-y",
  "unionpay-mcp-server"
 ],
 "env": {
  "UP_ACQ_INS_CODE": "<机构号，收单机构接入必填>",
  "UP_ACCESS_TYPE": "<必填，0-商户直连|1-收单机构接入|2-平台商户接入>",
  "UP_MER_ID": "<商户号，商户直连接入必填>",
  "UP_TR_ID": "<trId, 签约支付类交易必填>",
  "UP_TOKEN_TYPE": "<token类型，使用签约支付工具时必填，01-标记申请>",
  "UP_FRONT_URL": "<前台跳转地址，选填>",
  "UP_FRONT_FAIL_URL": "<失败跳转地址，选填>",
  "UP_BACK_URL": "<后台通知地址，选填>",
  "UP_SIGN_CERT_PATH": "<机构/商户签名证书绝对路径,必填>",
  "UP_SIGN_CERT_PWD": "<签名证书密码,必填>",
  "UP_VALIDATE_CERT_DIR": "<验签证书绝对路径,必填>",
  "UP_NEED_ENCRYPT": "<敏感信息是否加密，0-不加密，1-加密,可选，默认不加密>",
  "UP_ENCRYPT_CERT_PATH": "<如需加密，加密证书绝对路径，可选>",
  "UP_DECRYPT_CERT_PATH": "<如需加解密，解密证书绝对路径，可选>",
  "UP_ENCRYPT_CERT_PWD": "<如需加解密，解密证书密码，可选>",
  "UP_LOG_DIR": "日志打印地址，选填，默认在HOME目录打印日志",
  "UP_URL": "<银联交易地址, 选填，可填写生产或验证环境PM地址，不填默认生产地址>",
  "UP_TIME_OUT": "<交易地址超时时间（单位毫秒）, 选填，默认5000毫秒>",
  "UP_AVAILABLE_TOOLS": "<可用工具列表，选填，默认all,可按需配置工具名称，英文逗号分隔，如create-contract-order-unionpay-payment,create-contract-unionpay-payment>"
 }
    },
    "其他工具": { 
      "...": "..."
    }
  }
}
```

### 在 Cline 中使用

在 Cline 设置中找到 `cline_mcp_settings.json` 配置文件，并加入如下配置：

```json
{
  "mcpServers": {
    "unionpay-mcp-server": {
      "command": "npx",
      "args": ["-y", "unionpay-mcp-server"],
      "env": {
  "UP_ACQ_INS_CODE": "<机构号，收单机构接入必填>",
  "UP_ACCESS_TYPE": "<必填，0-商户直连|1-收单机构接入|2-平台商户接入>",
  "UP_MER_ID": "<商户号，商户直连接入必填>",
  "UP_TR_ID": "<trId, 签约支付类交易必填>",
  "UP_TOKEN_TYPE": "<token类型，使用签约支付工具时必填，01-标记申请>",
  "UP_FRONT_URL": "<前台跳转地址，选填>",
  "UP_FRONT_FAIL_URL": "<失败跳转地址，选填>",
  "UP_BACK_URL": "<后台通知地址，选填>",
  "UP_SIGN_CERT_PATH": "<机构/商户签名证书绝对路径,必填>",
  "UP_SIGN_CERT_PWD": "<签名证书密码,必填>",
  "UP_VALIDATE_CERT_DIR": "<验签证书绝对路径,必填>",
  "UP_NEED_ENCRYPT": "<敏感信息是否加密，0-不加密，1-加密,可选，默认不加密>",
  "UP_ENCRYPT_CERT_PATH": "<如需加密，加密证书绝对路径，可选>",
  "UP_DECRYPT_CERT_PATH": "<如需加解密，解密证书绝对路径，可选>",
  "UP_ENCRYPT_CERT_PWD": "<如需加解密，解密证书密码，可选>",
  "UP_LOG_DIR": "日志打印地址，选填，默认在HOME目录打印日志",
  "UP_URL": "<银联交易地址, 选填，可填写生产或验证环境PM地址，不填默认生产地址>",
  "UP_TIME_OUT": "<交易地址超时时间（单位毫秒）, 选填，默认5000毫秒>",
  "UP_AVAILABLE_TOOLS": "<可用工具列表，选填，默认all,可按需配置工具名称，英文逗号分隔，如create-contract-order-unionpay-payment,create-contract-unionpay-payment>"
      },
      "disable": false,
      "autoApprove": []
    },
    "其他工具": { 
      "...": "..."
    }
  }
}
```

### 在其他 MCP Client中使用

在其它支持MCP的客户端中，通过合理配置 Server 进程启动方式`npx -y unionpay-mcp-server` ，并按下文介绍设置环境参数即可使用。

### 所有参数

银联MCP Server通过环境变量接收参数。

```json
 {
  "UP_ACQ_INS_CODE": "<机构号，收单机构接入必填>",
  "UP_ACCESS_TYPE": "<必填，0-商户直连|1-收单机构接入|2-平台商户接入>",
  "UP_MER_ID": "<商户号，商户直连接入必填>",
  "UP_TR_ID": "<trId, 签约支付类交易必填>",
  "UP_TOKEN_TYPE": "<token类型，使用签约支付工具时必填，01-标记申请>",
  "UP_FRONT_URL": "<前台跳转地址，选填>",
  "UP_FRONT_FAIL_URL": "<失败跳转地址，选填>",
  "UP_BACK_URL": "<后台通知地址，选填>",
  "UP_SIGN_CERT_PATH": "<机构/商户签名证书绝对路径,必填>",
  "UP_SIGN_CERT_PWD": "<签名证书密码,必填>",
  "UP_VALIDATE_CERT_DIR": "<验签证书绝对路径,必填>",
  "UP_NEED_ENCRYPT": "<敏感信息是否加密，0-不加密，1-加密,可选，默认不加密>",
  "UP_ENCRYPT_CERT_PATH": "<如需加密，加密证书绝对路径，可选>",
  "UP_DECRYPT_CERT_PATH": "<如需加解密，解密证书绝对路径，可选>",
  "UP_ENCRYPT_CERT_PWD": "<如需加解密，解密证书密码，可选>",
  "UP_LOG_DIR": "日志打印地址，选填，默认在HOME目录打印日志",
  "UP_URL": "<银联交易地址, 选填，可填写生产或验证环境PM地址，不填默认生产地址>",
  "UP_TIME_OUT": "<交易地址超时时间（单位毫秒）, 选填，默认5000毫秒>",
  "UP_AVAILABLE_TOOLS": "<可用工具列表，选填，默认all,可按需配置工具名称，英文逗号分隔，如create-contract-order-unionpay-payment,create-contract-unionpay-payment>"
}
```

MCP Server配置：

 "UP_LOG_DIR": "日志打印地址，选填，默认在HOME目录打印日志"

## 3. 使用 MCP Inspector 调试

开发人员可使用 MCP Inspector 来调试和了解银联 MCP Server 的各项功能，具体操作如下：

1.通过export设置各环境变量；

2.执行 npx -y @modelcontextprotocol/inspector && npx -y unionpay-mcp-serve，启动 MCP Inspector；

3.在 MCP Inspector WebUI 中进行调试。

## 4. 支持的能力

下表列出了本版本MCP Server可用的支付工具，本版本提供能力具体使用方法参考银联开放平台签约支付产品（ <https://open.unionpay.com/tjweb/acproduct/list?apiSvcId=3301> ）

| 名称 | 描述 | 参数 | 输出 |
|:-----:|:----:|:----:|:----:|
| `create-contract-order-unionpay-payment` | 创建一笔签约支付订单，并返回授权签约链接。 | - orderId: 交易订单号,格式:8至40位字母数字 <br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- certifTp: 证件类型,格式:2位数字<br>- certifId: 证件号码,格式:1至20位字母数字<br>- customerNm: 用户姓名,格式:1至120字母数字<br>- phoneNo: 手机号,格式:1至20位手机号<br>- riskRateInfo: 风险信息域的JSON字符串格式 | -  code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- contractUrl: 签约url |
| `create-contract-unionpay-payment` | 发起签约交易，并返回签约信息,该交易是签约下单的后续交易，是支付的前序交易，且签约交易只需做一次，可以实现多次支付。 | - orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br> - origOrderId: 签约下单交易请求的订单号orderId<br>- origTxnTime: 签约下单交易应答的txnTime<br>- tokenType: token类型,格式:2位数字<br> | code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- tokenInfo: 签约成功后返回。token：签约协议ID， tokenEnd：标记失效时间  <br>- cardContractInfo: 银行卡签约信息<br>- protocolFlag: 电子协议签约标识|
| `contract-pay-sms` | 创建一笔支付短信，当需要在支付前做短信验证时调用此接口| - orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- token: 签约交易返回的token<br>- txnAmt: 交易金额<br>- currencyCode: 交易币种,默认156人民币 | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- contractUrl: 签约url|
| `pay-contract-order-unionpay-payment` | 创建一笔签约支付订单，并返回用户支付结果 | orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- certifTp: 证件类型,格式:2位数字<br>- certifId: 证件号码,格式:1至20位字母数字<br>- customerNm: 用户姓名,格式:1至120字母数字<br>- phoneNo: 手机号,格式:1至20位手机号<br>- riskRateInfo: 风险信息域的JSON字符串格式 <br>- currencyCode: 交易币种,默认156人民币 <br>- token:  签约交易返回的token  | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- orderId: 支付订单ID|
| `refund-contract-order-unionpay-payment` | 创建一笔退货订单，并返回退货结果| - orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br> - origOrderId: 支付交易请求的订单号orderId<br>- origTxnTime: 交易应答的txnTime<br> - txnAmt: 需要退货的金额<br> | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- orderId: 支付订单ID |
| `query-unionpay-payment` | 发起查询交易，支持签约交易查询、支付交易查询、退款等查询类交易| - orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- txnAmt: 退货金额 <br>- transStatus: 交易状态<br>- tokenInfo: token域信息<br>- cardContractInfo: 银行卡签约信息 <br>- origBizMethod: 查询订单对应的原始方法<br>- origTn: 查询订单对应的原始订单号  |
| `cancel-contract-order-unionpay-payment` | 创建一笔解约订单，并返回解约结果,该交易是签约的反向交易 | - orderId: 交易订单号,格式:8至40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- token: 签约交易返回的token | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- orderId: 支付订单ID |
| `apply-unionpay-qrCode` | 申请消费二维码,返回qrCode | - orderId: 交易订单号,格式:8~40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- payTimeout: 二维码有效时间 <br>- txnAmt: 交易金额,单位元| - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- qrCode: 返回的二维码链接 |
| `refund-unionpay-qrCode` | 对之前二维码交易进行退货，仅30天之内交易可退货，多次退货累计退货金额不超过原始交易金额 | - orderId: 交易订单号,格式:8~40位字母数字<br>- txnTime: 交易时间,格式:yyyyMMddHHmmss<br>- origTxnTime: 原始交易的交易时间,格式:YYYYMMDDhhmmss<br>- origOrderId: 原始交易的订单号,格式:8~40位字母数字<br>- origQryId: 原始交易的查询ID,格式:查询订单请求返回的若干位数字<br>- txnAmt: 待退货金额,单位元 | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br>- orderId: 支付订单ID |
| `query-unionpay-QrCode-trade` | 发起二维码类查询交易，支持申码支付交易查询、退款等查询类交易 | - orderId: 被查询交易订单号,原交易订单号,当使用原交易orderId和txnTime查询时必填<br>- txnTime: 被查询交易交易时间,原交易时间,当使用原交易orderId和txnTime查询时必填<br>- queryId: 被查询交易流水号,当使用原交易流水号查询时必填 | - code: 系统响应码 <br>- msg: 接口响应信息<br>- txnTime: 交易时间<br> |

此外对于平台商户和收单机构接入时候，需要额外在除query-unionpay-payment外的其他工具调用时候，上送如下字段:

| 名称 | 描述 | 参数 |
|-------|------|------|
merCatCode| 商户类别       |收单接入时必填
merName    |商户名称       |收单接入时必填
merAbbr    |商户简称       |收单接入时必填
subMerId   |二级商户代码   |平台商户接入时必填
subMerName| 二级商户名称   |平台商户接入时必填
subMerAbbr |二级商户简称   |平台商户接入时必填

## 5. 如何选择合适的支付方式  

在开发过程中，为了让 LLM 能更准确地选择合适的支付方式，建议在 Prompt 中清晰说明产品使用场景：
网页支付：适用于用户在电脑屏幕上看到支付界面的场景。如果智能体应用主要运行在桌面端（PC），可以在Prompt中说明："我的应用是桌面软件/PC网站，需要在电脑上展示支付二维码"。
手机支付：适用于用户在手机浏览器内发起支付的场景。如果应用是手机H5页面或移动端网站，可在Prompt中说明："我的页面是手机网页，需要直接在手机上发起在线付款"。
更多MCP支付工具正在研发中，敬请期待。

## 6. 注意事项  

* 最新使用指南请以银联开放平台-银联MCP智能支付服务解决方案为准（ <https://open.unionpay.com/tjweb/solution/detail?solId=613>  ）
* 银联MCP支付服务目前处于发布早期阶段，相关能力和配套设施正在持续完善中。在使用过程中，如有相关问题或建议，欢迎联系我们。
* 在开发任何使用 MCP Server的智能体服务，并提供给用户使用时，请了解必要的安全知识，防范AI应用特有的Prompt攻击、MCP Server任意命令执行等安全风险。
* 我们提供了MD5校验机制，参见'dist/checksums.md5'文件

## 7. 使用协议  

本工具是银联开放平台能力的组成部分。使用期间，请遵守中国银联开发者使用规范
（ <https://open.unionpay.com/tjweb/support/doc/online/3/122> ）、开放平台《中国银联服务协议》（ <https://user.95516.com/pages/misc/newAgree.html> ）和相关商业行为法规。
