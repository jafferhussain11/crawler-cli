# Web Crawler CLI app in Go - mvp

A high-performance, concurrent web crawler built in Go using Test-Driven Development (TDD) principles (as much as possible). This crawler efficiently extracts structured data from websites and generates a CSV report.

## 🚀 Features

### Core Functionality
- **Concurrent Web Crawling**: Multi-threaded crawling with configurable concurrency limits
- **Domain Restriction**: Automatically restricts crawling to the same domain as the starting URL
- **URL Normalization**: Intelligent URL processing and deduplication
- **Structured Data Extraction**: Extracts H1 tags, first paragraphs, outgoing links, and image URLs
- **CSV Report Generation**: Comprehensive data export in CSV format
- **Real-time Progress Tracking**: Live console output showing crawling progress

### Technical Features
- **Thread-Safe Operations**: Uses `sync.Map` and atomic operations for concurrent access
- **Memory Efficient**: Streams HTML content and processes data incrementally
- **Error Handling**: Robust error handling for network failures and parsing errors
- **Content Type Validation**: Ensures only HTML content is processed
- **User-Agent Headers**: Proper HTTP headers for respectful crawling

## 📋 Usage

```bash
go run main.go <base_url> <max_pages> <max_concurrency>
```

### Example
```bash
go run main.go https://example.com 100 10
```

**Parameters:**
- `base_url`: Starting URL for the crawl
- `max_pages`: Maximum number of pages to crawl
- `max_concurrency`: Maximum concurrent goroutines (recommended: 5-15)

## 🏗️ Architecture

### Core Components

1. **Main Crawler (`crawl-page.go`)**
   - Manages crawling configuration and concurrency
   - Handles domain validation and URL normalization
   - Coordinates concurrent page processing

2. **HTML Processing (`extract_html.go`, `extract_links_images.go`)**
   - Extracts H1 tags and first paragraphs using goquery
   - Processes outgoing links and image URLs
   - Handles relative URL resolution

3. **Data Structures (`structured_page_data.go`)**
   - Defines `PageData` struct for organized data storage
   - Manages structured data extraction pipeline

4. **URL Management (`normalize_url.go`)**
   - URL normalization and deduplication
   - Domain validation and path processing

5. **HTTP Client (`getHTML.go`)**
   - Robust HTTP client with proper headers
   - Content type validation and error handling

6. **Report Generation (`csv_report.go`)**
   - CSV export functionality
   - Thread-safe data aggregation

## 📊 Output Format

The crawler generates a CSV report (`report.csv`) with the following columns:
- `page_url`: The crawled page URL
- `h1`: Main heading text
- `first_paragraph`: First paragraph content
- `outgoing_link_urls`: Semicolon-separated list of outgoing links
- `image_urls`: Semicolon-separated list of image URLs

## 🧪 Testing

The project includes a fair number of tests covering:
- URL normalization tests
- HTML extraction tests
- Link and image extraction tests
- Structured data processing tests

Run tests with:
```bash
go test ./...
```

## 🔮 Future Enhancements & Roadmap

### Phase 1: Enhanced Error Handling & Reliability
- **Better Error Recovery**: Implement retry mechanisms for failed URLs
- **Rate Limiting**: Add configurable delays between requests to be respectful to target servers
- **Timeout Management**: Implement request timeouts and connection pooling
- **Robust Error Logging**: Detailed error tracking and reporting system

### Phase 2: Server Deployment & API
- **REST API**: HTTP API endpoints for triggering crawls remotely
- **Web Dashboard**: Real-time monitoring interface for crawl progress
- **Queue System**: Background job processing with Redis/RabbitMQ
- **Authentication**: Secure API access with JWT tokens

### Phase 3: Advanced Features
- **Email Notifications**: Send crawl results via email when complete
- **Metrics & Analytics**: Detailed performance metrics and crawl statistics
- **Scheduled Crawling**: Cron-based automatic crawling

### Phase 4: Enterprise Features
- **Multi-tenant Support**: Isolated crawling environments
- **Custom Extractors**: Plugin system for custom data extraction
- **Crawl Depth Control**: Configurable crawling depth limits
- **Content Filtering**: Advanced content type and URL pattern filtering
- **Performance Monitoring**: APM integration and health checks

### Phase 5: AI
- **Content Classification**: Automatic content categorization
- **Sentiment Analysis**: Content sentiment scoring
- **SEO Analysis**: Automated SEO metrics extraction

## 🛠️ Technical Stack

- **Language**: Go 1.25+
- **Dependencies**: 
  - `github.com/PuerkitoBio/goquery` - HTML parsing
- **Concurrency**: Goroutines with manual mutex use to lock map along with an alternate implementation using
sync.Map and atomic operations.
- **Testing**: Built-in Go testing framework with TDD approach

## 📈 Performance Considerations

- **Memory Usage**: Optimized for large-scale crawling with minimal memory footprint
- **Concurrency**: Configurable concurrency to balance speed and server load
- **Network Efficiency**: Connection reuse and proper HTTP client configuration
- **Data Processing**: Stream-based HTML processing to handle large pages



## 📄 License

This project is part of a learning exercise in Go web crawling and TDD practices.
