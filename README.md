# Tempo Tapper - Enhanced Edition

> A BPM (Beats Per Minute) calculator that detects tempo from keyboard input, mouse clicks, or touch taps. This is an enhanced version of the original [tempo-tapper](https://github.com/nobuyo/tempo-tapper) by [nobuyo](https://github.com/nobuyo).

**Topics:** `bpm-calculator` `tempo-detector` `music-tools` `rhythm-tapper` `web-app` `go` `html` `javascript` `cross-platform` `mobile-friendly`

## 🎵 Features

- **Multiple Input Methods**: Tap with keyboard, mouse clicks, or touch
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Web Version**: Standalone HTML file that works in any browser
- **Mobile Friendly**: Perfect for phones - add to home screen for quick access
- **Real-time BPM Detection**: Instant tempo calculation as you tap
- **Beautiful UI**: Modern, responsive design with gradient backgrounds

## 📱 Mobile Usage

The HTML version is perfect for mobile devices:

1. Open `tempo_tapper.html` in your phone's browser
2. Tap the screen or use the spacebar to set the tempo
3. Add to your home screen for quick access:
   - **iOS**: Tap the share button → "Add to Home Screen"
   - **Android**: Tap the menu → "Add to Home Screen"

## 🚀 Quick Start

### Web Version (Recommended)
Simply open `tempo_tapper.html` in any web browser - no installation required!

### Desktop Applications
Choose the appropriate executable for your platform:

- **Windows**: `tempo_windows.exe` or `tempo_gui.exe`
- **macOS/Linux**: `tempo` or `tempo_simple`
- **All Platforms**: `tempo_gui.exe` (GUI version)

## 🛠️ Building from Source

### Prerequisites
- Go 1.16 or later
- `stty` command (for Unix-like systems)

### Build Commands

```bash
# Original CLI version
go build -o tempo main.go

# Simple version
go build -o tempo_simple main_simple.go

# Windows-specific version
go build -o tempo_windows main_windows.go

# GUI version (serves HTML interface)
go build -o tempo_gui main_gui.go
```

## 📖 Usage

### Web Version
1. Open `tempo_tapper.html` in your browser
2. Tap the "TAP" button or press the spacebar to the rhythm
3. Keep a steady rhythm for accurate BPM detection
4. The BPM will be displayed in real-time

### Desktop Version
1. Run the appropriate executable for your platform
2. Tap any key to the rhythm
3. Press `Ctrl-c` to quit

## 🎯 How It Works

The tempo tapper calculates BPM by measuring the intervals between your taps. It uses a rolling window of recent taps to provide accurate, real-time tempo detection. The more consistent your rhythm, the more accurate the BPM reading.

## 📁 Project Structure

- `tempo_tapper.html` - Standalone web version
- `main.go` - Original CLI version
- `main_simple.go` - Simplified CLI version
- `main_windows.go` - Windows-optimized version
- `main_gui.go` - GUI version that serves the web interface
- Pre-built executables for all platforms

## 🤝 Credits

This project is based on the original [tempo-tapper](https://github.com/nobuyo/tempo-tapper) by [nobuyo](https://github.com/nobuyo). The original project provided the core BPM calculation logic, which has been enhanced with:

- Web interface for cross-platform compatibility
- Mobile-optimized design
- Multiple input methods (keyboard, mouse, touch)
- Pre-built executables for easy distribution
- Enhanced UI/UX

## 📄 License

MIT License - same as the original project.

## 🌟 Enhancements Made

- **Web Interface**: Created a standalone HTML file that works in any browser
- **Mobile Support**: Optimized for touch devices and mobile browsers
- **Multiple Builds**: Added Windows-specific and GUI versions
- **Pre-built Binaries**: Included executables for all major platforms
- **Enhanced UI**: Modern design with gradients and responsive layout
- **Input Flexibility**: Support for keyboard, mouse, and touch input
- **Home Screen Integration**: Easy to add to mobile home screens

---

**Original Project**: [nobuyo/tempo-tapper](https://github.com/nobuyo/tempo-tapper)  
**Original Author**: [nobuyo](https://github.com/nobuyo)
