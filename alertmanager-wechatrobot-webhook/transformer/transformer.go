package transformer

import (
	"bytes"
	"fmt"

	"github.com/k8stech/alertmanager-wechatrobot-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	// buffer.WriteString(fmt.Sprintf("### 当前状态:%s \n", status))
	// buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n # 告警: <font color='warning'> %s </font>\n", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>【环境】 %s\n", labels["datacenter"]))
		buffer.WriteString(fmt.Sprintf("\n>【级别】 %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>【类型】 %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>【主机】 %s\n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n>【内容】 %s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>【当前状态】%s \n", status))
		buffer.WriteString(fmt.Sprintf("\n>【触发时间】 %s\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString("\n [跳转Grafana看板](http://1.1.1.1:3000/dashboards)")
		buffer.WriteString("\n @马乐乐 @彭勇 ")
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
