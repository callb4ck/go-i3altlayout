package main

import (
    "log"

    "github.com/mdirkse/i3ipc"
)

func onWindowFocus(i3 *i3ipc.IPCSocket, event i3ipc.Event) {
    window, err := i3.GetTree()

    if err != nil {
        panic("Error while fetching layout tree:\n" + err.Error())
    }

    width := window.FindFocused().Rect.Width
    height := window.FindFocused().Rect.Height

    log.Printf("Width: %v | Height: %v", width, height)

    if height > width {
        log.Println("Setting vertical split")
        i3.Command("split vertical")
    } else {
        log.Println("Setting horizontal split")
        i3.Command("split horizontal")
    }
}

func main() {
    i3, err := i3ipc.GetIPCSocket()

    if err != nil {
        panic("Error while creating new IPC socket:\n" + err.Error())
    }

    version, err := i3.GetVersion()

    if err != nil {
        panic("Error while fetching version:\n" + err.Error())
    }

    log.Println(version)


    i3ipc.StartEventListener()
    ws_events, err := i3ipc.Subscribe(i3ipc.I3WindowEvent)
    for {
        event := <- ws_events
        onWindowFocus(i3, event)
        log.Printf("Received an event: %v\n", event)
    }
}
