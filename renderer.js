const { ipcRenderer } = require('electron');

// Make window draggable
let isDragging = false;
let dragOffset = { x: 0, y: 0 };

// Window control elements
const titleBar = document.getElementById('title-bar');
const minimizeBtn = document.getElementById('minimize-btn');
const maximizeBtn = document.getElementById('maximize-btn');
const closeBtn = document.getElementById('close-btn');

// Window control event listeners
if (minimizeBtn) {
  minimizeBtn.addEventListener('click', () => {
    ipcRenderer.send('minimize-window');
  });
}

if (maximizeBtn) {
  maximizeBtn.addEventListener('click', () => {
    ipcRenderer.send('maximize-window');
  });
}

if (closeBtn) {
  closeBtn.addEventListener('click', () => {
    ipcRenderer.send('close-window');
  });
}

// Make title bar draggable
if (titleBar) {
  titleBar.addEventListener('mousedown', (e) => {
    isDragging = true;
    dragOffset = {
      x: e.clientX,
      y: e.clientY
    };
  });
}

// Handle mouse move for dragging
document.addEventListener('mousemove', (e) => {
  if (isDragging) {
    // Use IPC to communicate with main process for window movement
    ipcRenderer.send('move-window', {
      deltaX: e.clientX - dragOffset.x,
      deltaY: e.clientY - dragOffset.y
    });
    
    dragOffset = {
      x: e.clientX,
      y: e.clientY
    };
  }
});

// Stop dragging on mouse up
document.addEventListener('mouseup', () => {
  isDragging = false;
});

// Global keyboard shortcuts
document.addEventListener('keydown', (e) => {
  // R key to restart/reset
  if (e.key === 'r' || e.key === 'R') {
    e.preventDefault();
    ipcRenderer.send('restart-app');
  }
  
  // Escape key to close
  if (e.key === 'Escape') {
    e.preventDefault();
    ipcRenderer.send('close-window');
  }
  
  // F11 for fullscreen toggle
  if (e.key === 'F11') {
    e.preventDefault();
    ipcRenderer.send('toggle-fullscreen');
  }
});

// Prevent context menu
document.addEventListener('contextmenu', (e) => {
  e.preventDefault();
});

// Focus management
window.addEventListener('focus', () => {
  document.body.classList.add('focused');
});

window.addEventListener('blur', () => {
  document.body.classList.remove('focused');
});
