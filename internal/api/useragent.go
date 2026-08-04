package api

// ChromeUserAgent is a typical Chrome-on-Windows desktop User-Agent plus a CyverCliTool product
// token so traffic can be distinguished from real browsers while keeping a browser-like prefix.
// The Chrome portion is bumped periodically to stay in line with current stable Chrome major lines.
const ChromeUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36 CyverCliTool/1.0"
