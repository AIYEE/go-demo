package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/AIYEE/go-demo/internal/chain"
	contract "github.com/AIYEE/go-demo/internal/chain/contract/confirm"
	confirm "github.com/AIYEE/go-demo/internal/confirmServer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

type program struct {
	start func()
	stop  func()
}

func (c *command) initStartCmd() error {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a demo service",
		Long:  "start a demo service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return cmd.Help()
			}
			v := strings.ToLower(c.config.GetString(optionNameVerbosity))
			logger, err := c.NewLogger(v)
			if err != nil {
				return fmt.Errorf("new logger: %v", err)
			}

			// new chain service
			rawurl := c.config.GetString(optionNameRPCEndPoint)
			actor, err := chain.New(rawurl)
			if err != nil {
				logger.Errorf("create service fail. Error: %v", err)
				return err
			}

			// new contract
			confirmContractAddress := common.HexToAddress(c.config.GetString(optionNameContractAddress))
			confirmContract, err := contract.NewConfirm(confirmContractAddress, actor.Client)
			if err != nil {
				logger.Errorf("create contract fail. Error: %v", err)
				return err
			}

			// new confirm service
			quit := make(chan struct{}, 1)
			wg := sync.WaitGroup{}
			dbFile := c.config.GetString(optionNameDbFile)
			ethSigner := c.config.GetString(optionNameEthSinger)
			needGas := c.config.GetBool(optionNameSendGas)
			confirmService, err := confirm.New(confirmContract, actor, logger, dbFile, quit, &wg, ethSigner, needGas)
			if err != nil {
				logger.Errorf("create confirmService fail. Error: %v", err)
				return err
			}
			go confirmService.Start()

			// Wait for termination or interrupt signals.
			// We want to clean up things at the end.
			interruptChannel := make(chan os.Signal, 1)
			signal.Notify(interruptChannel, syscall.SIGINT, syscall.SIGTERM)

			sig := <-interruptChannel
			logger.Infof("received signal: %v", sig)
			quit <- struct{}{}
			wg.Wait()
			logger.Info("quit program")

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return c.config.BindPFlags(cmd.Flags())
		},
	}
	c.setAllFlags(cmd)
	c.root.AddCommand(cmd)
	return nil
}
