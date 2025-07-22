# Met Museum Go Project

A modern web application built with Go, Gin, and Templ that interfaces with the Metropolitan Museum of Art's public API to explore museum departments and collections.

## 🏛️ Features

- **Home Page**: Beautiful landing page with museum branding and navigation
- **Department Browser**: View all Metropolitan Museum departments
- **Responsive Design**: Mobile-first responsive UI using Tailwind CSS
- **Modern Architecture**: Built with Go, Gin framework, and Templ templating
- **API Integration**: Real-time data from the Met Museum's public API

## 🛠️ Tech Stack

### Backend
- **Go 1.24.1** - Programming language
- **Gin** - HTTP web framework
- **Templ** - Type-safe HTML templating

### Frontend
- **Tailwind CSS 4.1.11** - Utility-first CSS framework
- **HTML5** - Semantic markup
- **JavaScript** - Interactive components

### External APIs
- **Met Museum API** - Collection and department data

## 📁 Project Structure

```
met-go-project/
├── assets/                 # Source CSS files
│   └── styles.css         # Tailwind CSS input file
├── components/            # Templ templates
│   ├── base.templ        # Base HTML layout
│   ├── home.templ        # Home page template
│   ├── departments.templ # Department list template
│   └── navigation.templ  # Navigation component
├── handlers/             # HTTP request handlers
│   ├── home.go          # Home page handler
│   └── get-departments.go # Department handlers
├── models/              # Data structures
│   └── departments.go   # Department models
├── public/              # Static assets
│   └── styles.css      # Compiled CSS
├── services/            # External API services
│   ├── met_departments.go    # Met API department service
│   └── met_collectionObject.go # Met API collection service
├── go.mod              # Go module dependencies
├── package.json        # Node.js dependencies (for Tailwind)
└── main.go            # Application entry point
```

## 🚀 Quick Start

### Prerequisites

- **Go 1.24.1 or higher**
- **Node.js** (for Tailwind CSS)
- **Git**

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd met-go-project
   ```

2. **Install Go dependencies**
   ```bash
   go mod download
   ```

3. **Install Node.js dependencies**
   ```bash
   npm install
   ```

4. **Build CSS assets**
   ```bash
   npm run build
   ```

5. **Generate Templ files**
   ```bash
   templ generate
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

7. **Visit the application**
   ```
   http://localhost:8085
   ```

## 🎯 Available Routes

| Route | Method | Description |
|-------|--------|-------------|
| `/` | GET | Home page with museum overview |
| `/met` | GET | Department list page |

## 🔧 Development

### CSS Development

For development with automatic CSS rebuilding:
```bash
npm run watch
```

This will watch for changes in `assets/styles.css` and automatically rebuild `public/styles.css`.

### Template Development

After modifying `.templ` files, regenerate the Go code:
```bash
templ generate
```

### Hot Reload Setup

For development with automatic server restart, you can use tools like:
- `air` - Live reload for Go apps
- `nodemon` - Monitor for changes

Install air:
```bash
go install github.com/air-verse/air@latest
```

Create `.air.toml` configuration and run:
```bash
air
```

## 🌐 API Integration

The application integrates with the Metropolitan Museum of Art's public API:

- **Base URL**: `https://collectionapi.metmuseum.org/public/collection/v1/`
- **Departments Endpoint**: `/departments`
- **No API key required**

### API Response Structure

```json
{
  "departments": [
    {
      "departmentId": 1,
      "displayName": "American Decorative Arts"
    }
  ]
}
```

## 🎨 Styling

The project uses **Tailwind CSS 4.1.11** for styling with a custom color scheme:

- **Primary**: Blue tones (`blue-600`, `blue-800`)
- **Secondary**: Gray tones (`gray-100`, `gray-900`)
- **Accents**: Green and purple for highlights

### Custom Styles

Main input file: `assets/styles.css`
Compiled output: `public/styles.css`

## 🏗️ Architecture

### Component Structure

- **Base Template**: Provides consistent layout, navigation, and styling
- **Page Templates**: Specific page content (Home, Departments)
- **Component Templates**: Reusable UI components (Navigation)

### Handler Pattern

Each route has a dedicated handler in the `handlers/` directory that:
1. Fetches data from services
2. Renders appropriate templates
3. Returns HTTP responses

### Service Layer

External API calls are abstracted into service functions:
- Clean separation of concerns
- Easy testing and mocking
- Consistent error handling

## 📱 Responsive Design

The application is fully responsive with breakpoints:
- **Mobile**: Default styles
- **Tablet**: `md:` prefix (768px+)
- **Desktop**: `lg:` prefix (1024px+)

### Navigation

- **Desktop**: Horizontal navigation bar
- **Mobile**: Hamburger menu with slide-down panel

## 🧪 Testing

To run tests (when implemented):
```bash
go test ./...
```

## 🚀 Deployment

### Build for Production

1. **Build CSS**
   ```bash
   npm run build
   ```

2. **Generate templates**
   ```bash
   templ generate
   ```

3. **Build Go binary**
   ```bash
   go build -o met-museum-app
   ```

4. **Run production server**
   ```bash
   ./met-museum-app
   ```

### Docker Deployment

Create a `Dockerfile`:
```dockerfile
FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o met-museum-app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/met-museum-app .
COPY --from=builder /app/public ./public
EXPOSE 8085
CMD ["./met-museum-app"]
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the ISC License.

## 🎯 Future Enhancements

- [ ] Individual artwork detail pages
- [ ] Search functionality
- [ ] User favorites system
- [ ] Collection filtering
- [ ] Image gallery views
- [ ] User authentication
- [ ] Admin dashboard
- [ ] API rate limiting
- [ ] Caching layer
- [ ] Unit and integration tests

## 📞 Support

If you encounter any issues or have questions, please open an issue in the repository.

## 🙏 Acknowledgments

- **Metropolitan Museum of Art** for providing the public API
- **Go Team** for the excellent programming language
- **Gin Framework** for the lightweight web framework
- **Templ** for type-safe HTML templating
- **Tailwind CSS** for the utility-first CSS framework