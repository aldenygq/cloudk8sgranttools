/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"fmt"
	"cloudk8sgranttools/server"

	"github.com/spf13/cobra"
)

// ackgrantCmd represents the ackgrant command
var ackgrantCmd = &cobra.Command{
	Use:   "ackgrant",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		uid,_ := cmd.Flags().GetString("uid")
		role_name, _ := cmd.Flags().GetString("role_name")
		role_type, _ := cmd.Flags().GetString("role_type")
		is_custom, _ := cmd.Flags().GetBool("is_custom")
		cluster, _ := cmd.Flags().GetString("cluster")
		server.StartAckGrant(uid,role_name,role_type,cluster,is_custom)	
	},
}

func init() {
	rootCmd.AddCommand(ackgrantCmd)
	var (
		uid       string
		cluster   string
		is_custom bool
		role_name string
		role_type string
	)
	ackgrantCmd.Flags().StringVarP(&uid, "uid", "", "", "aliyun ram role id")
	_ = ackgrantCmd.MarkFlagRequired("uid")
	ackgrantCmd.Flags().BoolVarP(&is_custom, "is_custom", "", false, "is it a custom permission authorization")
	_ = ackgrantCmd.MarkFlagRequired("is_custom")
	ackgrantCmd.Flags().StringVarP(&role_name, "role_name", "","", "k8s role name")
	_ = ackgrantCmd.MarkFlagRequired("role_name")
	ackgrantCmd.Flags().StringVarP(&role_type, "role_type", "", "", "role_type")
	_ = ackgrantCmd.MarkFlagRequired("role_type")
	ackgrantCmd.Flags().StringVarP(&cluster, "cluster", "", "", "aliyun ack cluster id")
	_ = ackgrantCmd.MarkFlagRequired("cluster")
}
