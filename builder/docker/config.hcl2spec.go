// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package docker

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	Author                    *string           `mapstructure:"author" cty:"author"`
	Changes                   []string          `mapstructure:"changes" cty:"changes"`
	Commit                    *bool             `mapstructure:"commit" required:"true" cty:"commit"`
	ContainerDir              *string           `mapstructure:"container_dir" required:"false" cty:"container_dir"`
	Discard                   *bool             `mapstructure:"discard" required:"true" cty:"discard"`
	ExecUser                  *string           `mapstructure:"exec_user" required:"false" cty:"exec_user"`
	ExportPath                *string           `mapstructure:"export_path" required:"true" cty:"export_path"`
	Image                     *string           `mapstructure:"image" required:"true" cty:"image"`
	Message                   *string           `mapstructure:"message" required:"true" cty:"message"`
	Privileged                *bool             `mapstructure:"privileged" required:"false" cty:"privileged"`
	Pty                       *bool             `cty:"pty"`
	Pull                      *bool             `mapstructure:"pull" required:"false" cty:"pull"`
	RunCommand                []string          `mapstructure:"run_command" required:"false" cty:"run_command"`
	Volumes                   map[string]string `mapstructure:"volumes" required:"false" cty:"volumes"`
	FixUploadOwner            *bool             `mapstructure:"fix_upload_owner" required:"false" cty:"fix_upload_owner"`
	WindowsContainer          *bool             `mapstructure:"windows_container" required:"false" cty:"windows_container"`
	Login                     *bool             `mapstructure:"login" required:"false" cty:"login"`
	LoginPassword             *string           `mapstructure:"login_password" required:"false" cty:"login_password"`
	LoginServer               *string           `mapstructure:"login_server" required:"false" cty:"login_server"`
	LoginUsername             *string           `mapstructure:"login_username" required:"false" cty:"login_username"`
	EcrLogin                  *bool             `mapstructure:"ecr_login" required:"false" cty:"ecr_login"`
	AccessKey                 *string           `mapstructure:"aws_access_key" required:"false" cty:"aws_access_key"`
	SecretKey                 *string           `mapstructure:"aws_secret_key" required:"false" cty:"aws_secret_key"`
	Token                     *string           `mapstructure:"aws_token" required:"false" cty:"aws_token"`
	Profile                   *string           `mapstructure:"aws_profile" required:"false" cty:"aws_profile"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"author":                       &hcldec.AttrSpec{Name: "author", Type: cty.String, Required: false},
		"changes":                      &hcldec.AttrSpec{Name: "changes", Type: cty.List(cty.String), Required: false},
		"commit":                       &hcldec.AttrSpec{Name: "commit", Type: cty.Bool, Required: false},
		"container_dir":                &hcldec.AttrSpec{Name: "container_dir", Type: cty.String, Required: false},
		"discard":                      &hcldec.AttrSpec{Name: "discard", Type: cty.Bool, Required: false},
		"exec_user":                    &hcldec.AttrSpec{Name: "exec_user", Type: cty.String, Required: false},
		"export_path":                  &hcldec.AttrSpec{Name: "export_path", Type: cty.String, Required: false},
		"image":                        &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"message":                      &hcldec.AttrSpec{Name: "message", Type: cty.String, Required: false},
		"privileged":                   &hcldec.AttrSpec{Name: "privileged", Type: cty.Bool, Required: false},
		"pty":                          &hcldec.AttrSpec{Name: "pty", Type: cty.Bool, Required: false},
		"pull":                         &hcldec.AttrSpec{Name: "pull", Type: cty.Bool, Required: false},
		"run_command":                  &hcldec.AttrSpec{Name: "run_command", Type: cty.List(cty.String), Required: false},
		"volumes":                      &hcldec.BlockAttrsSpec{TypeName: "volumes", ElementType: cty.String, Required: false},
		"fix_upload_owner":             &hcldec.AttrSpec{Name: "fix_upload_owner", Type: cty.Bool, Required: false},
		"windows_container":            &hcldec.AttrSpec{Name: "windows_container", Type: cty.Bool, Required: false},
		"login":                        &hcldec.AttrSpec{Name: "login", Type: cty.Bool, Required: false},
		"login_password":               &hcldec.AttrSpec{Name: "login_password", Type: cty.String, Required: false},
		"login_server":                 &hcldec.AttrSpec{Name: "login_server", Type: cty.String, Required: false},
		"login_username":               &hcldec.AttrSpec{Name: "login_username", Type: cty.String, Required: false},
		"ecr_login":                    &hcldec.AttrSpec{Name: "ecr_login", Type: cty.Bool, Required: false},
		"aws_access_key":               &hcldec.AttrSpec{Name: "aws_access_key", Type: cty.String, Required: false},
		"aws_secret_key":               &hcldec.AttrSpec{Name: "aws_secret_key", Type: cty.String, Required: false},
		"aws_token":                    &hcldec.AttrSpec{Name: "aws_token", Type: cty.String, Required: false},
		"aws_profile":                  &hcldec.AttrSpec{Name: "aws_profile", Type: cty.String, Required: false},
	}
	return s
}