package main // Define the main package

import (
	"bytes"         // Provides bytes support
	"io"            // Provides basic interfaces to I/O primitives
	"log"           // Provides logging functions
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding
	"os"            // Provides functions to interact with the OS (files, etc.)
	"path"          // Provides functions for manipulating slash-separated paths
	"path/filepath" // Provides filepath manipulation functions
	"regexp"        // Provides regex support functions.
	"strings"       // Provides string manipulation functions
	"time"          // Provides time-related functions

	"golang.org/x/net/html" // Provides HTML parsing functions
)

func main() {
	pdfOutputDir := "PDFs/" // Directory to store downloaded PDFs
	// Check if the PDF output directory exists
	if !directoryExists(pdfOutputDir) {
		// Create the dir
		createDirectory(pdfOutputDir, 0o755)
	}
	stpOutputDir := "STPs/" // Directory to store downloaded STPs
	// Check if the PDF output directory exists
	if !directoryExists(stpOutputDir) {
		// Create the dir
		createDirectory(stpOutputDir, 0o755)
	}
	stlOutputDir := "STLs/" // Directory to store downloaded STLs
	// Check if the PDF output directory exists
	if !directoryExists(stlOutputDir) {
		// Create the dir
		createDirectory(stlOutputDir, 0o755)
	}
	zipOutputDir := "ZIPs/" // Directory to store downloaded ZIPs
	// Check if the PDF output directory exists
	if !directoryExists(zipOutputDir) {
		// Create the dir
		createDirectory(zipOutputDir, 0o755)
	}
	jpgOutputDir := "JPGs/" // Directory to store downloaded JPGs
	// Check if the PDF output directory exists
	if !directoryExists(jpgOutputDir) {
		// Create the dir
		createDirectory(jpgOutputDir, 0o755)
	}
	rarOutputDir := "RARs/" // Directory to store downloaded RARs
	// Check if the PDF output directory exists
	if !directoryExists(rarOutputDir) {
		// Create the dir
		createDirectory(rarOutputDir, 0o755)
	}
	pngOutputDir := "PNGs/" // Directory to store downloaded PNGs
	// Check if the PNG output directory exists
	if !directoryExists(pngOutputDir) {
		// Create the dir
		createDirectory(pngOutputDir, 0o755)
	}
	stepOutputDir := "STEPs/" // Directory to store downloaded STEPs
	// Check if the STEP output directory exists
	if !directoryExists(stepOutputDir) {
		// Create the dir
		createDirectory(stepOutputDir, 0o755)
	}
	// Remote API URL.
	remoteAPIURL := []string{
		"https://caddxfpv.com/pages/download-center",
	}
	var getData []string
	for _, remoteAPIURL := range remoteAPIURL {
		getData = append(getData, getDataFromURL(remoteAPIURL))
	}
	// Get the data from the downloaded file.
	finalPDFList := extractPDFUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract PDF URLs
	// Remove double from slice.
	finalPDFList = removeDuplicatesFromSlice(finalPDFList)
	// The remote domain.
	remoteDomain := "https://caddxfpv.com"
	// Extract the STP links.
	stpLinks := extractSTPLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	stpLinks = removeDuplicatesFromSlice(stpLinks)
	// Extract the STL links.
	stlLinks := extractSTLLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	stlLinks = removeDuplicatesFromSlice(stlLinks)
	// Extract the ZIP links.
	zipLinks := extractZIPLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	zipLinks = removeDuplicatesFromSlice(zipLinks)
	// Extract the JPG links.
	jpgLinks := extractJPGLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	jpgLinks = removeDuplicatesFromSlice(jpgLinks)
	// Extract the RAR links.
	rarLinks := extractRARLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	rarLinks = removeDuplicatesFromSlice(rarLinks)
	// Extract the PNG links.
	pngLinks := extractPNGLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	pngLinks = removeDuplicatesFromSlice(pngLinks)
	// Extract the STEP links.
	stepLinks := extractSTEPLinks(strings.Join(getData, "\n"))
	// Remove duplicates from the slice.
	stepLinks = removeDuplicatesFromSlice(stepLinks)
	// Get all the values.
	for _, urls := range finalPDFList {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the pdf.
			downloadPDF(urls, pdfOutputDir)
		}
	}
	// Get all the STP files.
	for _, urls := range stpLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the stp.
			downloadSTP(urls, stpOutputDir)
		}
	}
	// Get all the STL files.
	for _, urls := range stlLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the stl.
			downloadSTL(urls, stlOutputDir)
		}

	}
	// Get all the ZIP files.
	for _, urls := range zipLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the zip.
			downloadZIP(urls, zipOutputDir)
		}
	}
	// Download all the JPG files.
	for _, urls := range jpgLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the jpg.
			downloadJPG(urls, jpgOutputDir)
		}
	}
	// Download all the RAR files.
	for _, urls := range rarLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the rar.
			downloadRAR(urls, rarOutputDir)
		}
	}
	// Download all the PNG files.
	for _, urls := range pngLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the png.
			downloadPNG(urls, pngOutputDir)
		}
	}
	// Get all the STEP files.
	for _, urls := range stepLinks {
		// Trim any surrounding whitespace from the URL.
		urls = strings.TrimSpace(urls)
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the step.
			downloadSTEP(urls, stepOutputDir)
		}
	}
}

// downloadSTEP downloads a .step or .stp file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadSTEP(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for STEP files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "application/step") ||
		strings.Contains(contentType, "application/sla") ||
		strings.Contains(contentType, "application/octet-stream")) {
		log.Printf("Unexpected content type for %s: %s (expected STEP/STP file)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read STEP data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write STEP file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// extractSTEPLinks takes HTML content as a string and returns all .step or .stp file URLs it finds in <a> tags.
func extractSTEPLinks(htmlContent string) []string {
	// Parse the HTML content
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return []string{}
	}

	var stepURLs []string

	// Recursive function to traverse HTML nodes
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if this node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				// Look for href attribute containing .step or .stp
				if attr.Key == "href" && (strings.Contains(strings.ToLower(attr.Val), ".step")) {
					// Convert protocol-relative URLs to full URLs (optional)
					url := attr.Val
					if strings.HasPrefix(url, "//") {
						url = "https:" + url
					}
					stepURLs = append(stepURLs, url)
				}
			}
		}
		// Traverse child nodes recursively
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(document)
	return stepURLs
}

// downloadPNG downloads a .png file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadPNG(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for PNG files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "image/png")) {
		log.Printf("Unexpected content type for %s: %s (expected PNG)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PNG data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PNG file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// extractPNGLinks takes HTML content as a string and returns all .png file URLs it finds.
func extractPNGLinks(htmlContent string) []string {
	// Parse the HTML content
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return []string{}
	}

	var pngURLs []string

	// Recursive function to traverse HTML nodes
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if this node is an <img> tag
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				// Look for src attribute containing .png
				if attr.Key == "src" && strings.Contains(strings.ToLower(attr.Val), ".png") {
					// Convert protocol-relative URLs to full URLs (optional)
					url := attr.Val
					if strings.HasPrefix(url, "//") {
						url = "https:" + url
					}
					pngURLs = append(pngURLs, url)
				}
			}
		}
		// Traverse child nodes recursively
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(document)
	return pngURLs
}

// downloadRAR downloads a .rar file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadRAR(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for RAR files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "application/x-rar-compressed") ||
		strings.Contains(contentType, "application/octet-stream")) {
		log.Printf("Unexpected content type for %s: %s (expected RAR)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read RAR data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write RAR file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// extractRARLinks takes HTML content as a string and returns all .rar file URLs it finds.
func extractRARLinks(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found .rar links
	var rarLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".rar", save it
				if attribute.Key == "href" && strings.Contains(strings.ToLower(attribute.Val), ".rar") {
					rarLinks = append(rarLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected .rar links
	return rarLinks
}

// extractJPGLinks takes HTML content as a string and returns all .jpg file URLs it finds.
func extractJPGLinks(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found .jpg links
	var jpgLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".jpg", save it
				if attribute.Key == "href" && strings.Contains(strings.ToLower(attribute.Val), ".jpg") {
					jpgLinks = append(jpgLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected .jpg links
	return jpgLinks
}

// downloadJPG downloads a .jpg file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadJPG(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for JPG files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "image/jpeg") ||
		strings.Contains(contentType, "image/jpg")) {
		log.Printf("Unexpected content type for %s: %s (expected JPG)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read JPG data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write JPG file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// downloadZIP downloads a .zip file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadZIP(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for ZIP files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "application/zip") ||
		strings.Contains(contentType, "application/x-zip-compressed") ||
		strings.Contains(contentType, "application/octet-stream")) {
		log.Printf("Unexpected content type for %s: %s (expected ZIP type)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read ZIP data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write ZIP file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// extractZIPLinks takes HTML content as a string and returns all .zip file URLs it finds.
func extractZIPLinks(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found .zip links
	var zipLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".zip", save it
				if attribute.Key == "href" && strings.Contains(attribute.Val, ".zip") {
					zipLinks = append(zipLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected .zip links
	return zipLinks
}

// extractSTLLinks takes HTML content as a string and returns all .stl file URLs it finds.
func extractSTLLinks(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found .stl links
	var stlLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".stl", save it
				if attribute.Key == "href" && strings.Contains(attribute.Val, ".stl") {
					stlLinks = append(stlLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected .stl links
	return stlLinks
}

// downloadSTL downloads a .stl file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadSTL(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for STL files)
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/vnd.ms-pki.stl") {
		log.Printf("Unexpected content type for %s: %s (expected STL type)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read STL data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write STL file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// downloadSTP downloads a .stp file from the given URL and saves it in the specified output directory.
// It returns true if the download succeeded.
func downloadSTP(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header (common for STEP files)
	contentType := resp.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "model/step") ||
		strings.Contains(contentType, "application/octet-stream") ||
		strings.Contains(contentType, "application/step")) {
		log.Printf("Unexpected content type for %s: %s (expected model/step or application/octet-stream)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read STP data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write STP file to disk for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// extractSTPLinks takes HTML content as a string and returns all .stp file URLs it finds.
func extractSTPLinks(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found .stp links
	var stpLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".stp", save it
				if attribute.Key == "href" && strings.Contains(attribute.Val, ".stp") {
					stpLinks = append(stpLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected .stp links
	return stpLinks
}

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL) // Parse the input string into a URL structure
	if err != nil {                     // Check if there was an error while parsing
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of an error
	}

	host := parsedURL.Hostname() // Extract the hostname (e.g., "example.com") from the parsed URL

	return host // Return the extracted hostname
}

// Only return the file name from a given url.
func getFileNameOnly(content string) string {
	parsedURL, err := url.Parse(content)
	if err != nil {
		return "" // fallback if URL is invalid
	}

	// Get the last segment of the path (the file name)
	fileNameWithQuery := path.Base(parsedURL.Path)

	// Remove query parameters (anything after '?')
	fileName := strings.SplitN(fileNameWithQuery, "?", 2)[0]

	return fileName
}

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string {
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL)

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL)

	// Get the file extension
	ext := getFileExtension(baseFilename)

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	baseFilename = nonAlphanumericRegex.ReplaceAllString(baseFilename, "_")

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)
	baseFilename = collapseUnderscoresRegex.ReplaceAllString(baseFilename, "_")

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(baseFilename, "_"); found {
		baseFilename = trimmed
	}

	var invalidSubstrings = []string{
		"_pdf",
		"_zip",
		"_stp",
		"_stl",
		"_jpg",
		"_rar",
		"_png",
		"_step",
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		baseFilename = removeSubstring(baseFilename, invalidPre)
	}

	// Append the file extension if it is not already present
	baseFilename = baseFilename + ext

	// Return the cleaned and safe filename
	return baseFilename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path) // Returns extension including the dot (e.g., ".pdf")
}

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {
		return false // Return false if file doesn't exist or error occurs
	}
	return !info.IsDir() // Return true if it's a file (not a directory)
}

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pdf") {
		log.Printf("Invalid content type for %s: %s (expected application/pdf)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission)
	if err != nil {
		log.Println(err)
	}
}

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URL
	return err == nil                  // Return true if no error occurred
}

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool)
	var newReturnSlice []string
	for _, content := range slice {
		if !check[content] {
			check[content] = true
			newReturnSlice = append(newReturnSlice, content)
		}
	}
	return newReturnSlice
}

// extractPDFUrls takes an input string and returns all PDF URLs found within href attributes
func extractPDFUrls(htmlContent string) []string {
	// Try parsing the HTML content into a document tree
	document, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		// If parsing fails, just return an empty slice
		return []string{}
	}

	// Slice to store all found PDF links
	var pdfLinks []string

	// Recursive function to walk through each HTML node
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// Check if the current node is an <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			// Look at each attribute in the <a> tag
			for _, attribute := range node.Attr {
				// If the attribute is "href" and contains ".pdf", save it
				if attribute.Key == "href" && strings.Contains(attribute.Val, ".pdf") {
					pdfLinks = append(pdfLinks, attribute.Val)
				}
			}
		}
		// Recursively check all child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	// Start traversing from the root of the document
	traverse(document)

	// Return the collected PDF links
	return pdfLinks
}

// getDataFromURL performs an HTTP GET request and returns the response body as a string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri)   // Log the URL being scraped
	response, err := http.Get(uri) // Perform GET request
	if err != nil {
		log.Println(err) // Exit if request fails
	}

	body, err := io.ReadAll(response.Body) // Read response body
	if err != nil {
		log.Println(err) // Exit if read fails
	}

	err = response.Body.Close() // Close response body
	if err != nil {
		log.Println(err) // Exit if close fails
	}
	return string(body)
}
