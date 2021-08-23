package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Preparing command")
	apiKey, _ := os.LookupEnv("TWITCH_API_KEY")
	args := `raspivid -rot 90 -n -t 0 -w $WIDTH -h $HEIGHT -fps $FRAMERATE -b $BITRATE -g $KEYFRAME -o - | ffmpeg -f lavfi -i anullsrc -c:a aac -r $FRAMERATE -i - -g $KEYFRAME -strict experimental -threads 4 -vcodec copy -map 0:a -map 1:v -b:v $BITRATE -preset ultrafast -f flv out.flv`
	c := exec.Command("bash", "-c", args)
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	c.Env = []string{
		"WIDTH=1920",
		"HEIGHT=1080",
		"FRAMERATE=30",
		"KEYFRAME=60",
		"BITRATE=3500000",
		"URL=rtmp://fra02.contribute.live-video.net/app",
		fmt.Sprintf("KEY=%s", apiKey),
	}
	if err := c.Start(); err != nil {
		fmt.Printf("Failed to start: %v\n", err)
		return
	}

	fmt.Println("Spawning monitor...")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Printf("PID: %v\n", c.Process.Pid)
		}
	}()

	wg.Wait()
	fmt.Println("Killing cmd...")
	pgid, err := syscall.Getpgid(c.Process.Pid)
	if err != nil {
		panic("failed to get PGID")
	}

	syscall.Kill(-pgid, 15)

	fmt.Println("Waiting for termination...")
	c.Wait()

	fmt.Printf("Killed. Status: %v, Process: %v\n", c.ProcessState, c.Process)
	fmt.Println("Done")
}
