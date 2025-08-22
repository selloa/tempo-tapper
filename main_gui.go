package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Tempo Tapper</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            margin: 0;
            padding: 20px;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            justify-content: center;
        }
        .container {
            background: rgba(255, 255, 255, 0.1);
            border-radius: 15px;
            padding: 30px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
        }
        .bpm-display {
            font-size: 4em;
            font-weight: bold;
            margin: 20px 0;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
        }
        .tap-button {
            background: rgba(255, 255, 255, 0.2);
            border: 2px solid rgba(255, 255, 255, 0.3);
            border-radius: 50%;
            width: 120px;
            height: 120px;
            font-size: 1.2em;
            color: white;
            cursor: pointer;
            transition: all 0.3s ease;
            margin: 20px auto;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .tap-button:hover {
            background: rgba(255, 255, 255, 0.3);
            transform: scale(1.05);
        }
        .tap-button:active {
            transform: scale(0.95);
        }
        .instructions {
            margin-top: 20px;
            font-size: 0.9em;
            opacity: 0.8;
        }
        .status {
            margin-top: 10px;
            font-size: 0.8em;
            opacity: 0.7;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎵 Tempo Tapper</h1>
        <div class="bpm-display" id="bpm">--</div>
        <button class="tap-button" id="tapButton" onclick="tap()">
            TAP
        </button>
        <div class="instructions">
            Click the button or press SPACE to tap<br>
            Keep a steady rhythm to get accurate BPM
        </div>
        <div class="status" id="status">Ready to tap!</div>
    </div>

    <script>
        let taps = [];
        let lastTapTime = 0;
        
        function tap() {
            const now = Date.now();
            if (lastTapTime > 0) {
                const interval = now - lastTapTime;
                taps.push(interval);
                
                // Keep only last 8 taps
                if (taps.length > 8) {
                    taps = taps.slice(-8);
                }
                
                // Calculate BPM
                if (taps.length >= 2) {
                    const avgInterval = taps.reduce((a, b) => a + b, 0) / taps.length;
                    const bpm = Math.round(60000 / avgInterval);
                    document.getElementById('bpm').textContent = bpm;
                    document.getElementById('status').textContent = 
                        'Taps: ' + taps.length + ' | Avg interval: ' + Math.round(avgInterval) + 'ms';
                }
            }
            lastTapTime = now;
            
            // Visual feedback
            const button = document.getElementById('tapButton');
            button.style.background = 'rgba(255, 255, 255, 0.4)';
            setTimeout(() => {
                button.style.background = 'rgba(255, 255, 255, 0.2)';
            }, 100);
        }
        
        // Keyboard support
        document.addEventListener('keydown', function(event) {
            if (event.code === 'Space') {
                event.preventDefault();
                tap();
            }
        });
        
        // Reset after 3 seconds of no tapping
        setInterval(() => {
            const now = Date.now();
            if (now - lastTapTime > 3000 && taps.length > 0) {
                taps = [];
                lastTapTime = 0;
                document.getElementById('bpm').textContent = '--';
                document.getElementById('status').textContent = 'Ready to tap!';
            }
        }, 1000);
    </script>
</body>
</html>
`

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("tempo").Parse(htmlTemplate))
		tmpl.Execute(w, nil)
	})

	port := ":8080"
	url := "http://localhost" + port

	fmt.Printf("🎵 Tempo Tapper GUI starting...\n")
	fmt.Printf("Opening browser window...\n")
	fmt.Printf("If browser doesn't open automatically, go to: %s\n", url)
	fmt.Printf("Press Ctrl+C to quit\n\n")

	// Open browser after a short delay
	go func() {
		time.Sleep(1 * time.Second)
		openBrowser(url)
	}()

	log.Fatal(http.ListenAndServe(port, nil))
}
