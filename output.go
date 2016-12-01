package main

import (
	"fmt"
	"log"

	"github.com/proxypoke/i3ipc"
)

func outputSwitch() {
	ipcsocket, err := i3ipc.GetIPCSocket()
	if err != nil {
		log.Fatalln(err)
	}

	workspaces, err := ipcsocket.GetWorkspaces()
	if err != nil {
		log.Fatalln(err)
	}

	var focusedWorkspacename string
	for _, w := range workspaces {
		if w.Focused {
			focusedWorkspacename = w.Name
			break
		}
	}

	outputs, err := ipcsocket.GetOutputs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, o := range outputs {
		if !o.Active {
			continue
		}

		if o.Current_Workspace == focusedWorkspacename {
			continue
		}

		cmd := fmt.Sprintf("workspace %s", o.Current_Workspace)
		_, err := ipcsocket.Command(cmd)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
