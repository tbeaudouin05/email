package emailconf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// EmailConf struct defines how config.yaml should be built
type EmailConf struct {
	EmailTitle      string `yaml:"title"`
	EmailBody       string `yaml:"body"`
	EmailAttachPath string `yaml:"attachPath"`
	EmailRecipient  string `yaml:"recipient"`
	SenderEmail     string `yaml:"senderEmail"`
	SenderName      string `yaml:"senderName"`
	SenderPw        string `yaml:"senderPw"`
	SMTPAddr        string `yaml:"smtpAddr"`
	SMTPPort        string `yaml:"smtpPort"`
}

// ReadYamlEmailConf reads the configuration of the email to send from config.yaml
// config.yaml should be in the same folder as sendemailmain.exe
// config.yaml should contain title:, body:, attachPath: and recipient
// if no attachment, attachPath: NA
func (c *EmailConf) ReadYamlEmailConf() *EmailConf {

	yamlFile, err := ioutil.ReadFile("email_config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
