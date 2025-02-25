# daily-learning-feed
The Daily Learning Feed will be a web app that aggregates and displays learning resources (articles, videos, blog posts) from various sources. Users will be able to view, save, and categorize learning materials, either by manually adding them or fetching from external APIs (optional).

## Core Goals:

- Allow users to add learning resources (title, URL, category).
- Automatically fetch daily learning materials from external sources (optional).
- Enable users to categorize and filter their saved resources.
- Deploy the app online using Railway or Heroku.

## Project Breakdown by Day
### Define Scope & Set Up Project

✅ Goals:

    Plan out the features and database schema.
    Initialize the project with necessary libraries (gin, gorm, etc.).

✅ Tasks:

    Define database schema (Resources, Categories).
    Sketch a rough UI wireframe (learning feed, saved resources, categories).
    Set up main.go and install dependencies.

📌 Deliverable: Project is initialized with a clear scope & structure.

### Set Up Database & Models

✅ Goals:

    Set up the PostgreSQL database using gorm.

✅ Tasks:

    Create models for:
        Resources (ID, Title, URL, Category, Source, DateAdded).
        Categories (ID, Name, Description).
    Write migration functions to initialize the database.

📌 Deliverable: Database schema created & connected using GORM.

### Implement Backend CRUD API for Learning Resources

✅ Goals:

    Build API endpoints for adding, retrieving, updating, and deleting resources.

✅ Tasks:

    Create API routes using gin:
        POST /resources → Add a new learning resource
        GET /resources → Fetch all resources
        GET /resources/:id → Fetch a single resource
        PUT /resources/:id → Edit a resource
        DELETE /resources/:id → Remove a resource
    Implement basic validation (e.g., check if the URL is valid).

📌 Deliverable: Fully functional API for managing learning resources.

### Fetch External Learning Resources (Optional Feature)

✅ Goals:

    Automatically fetch new learning materials from external sources (API or RSS feeds).

✅ Tasks:

    Implement RSS feed parsing using gofeed (if using blog sources).
    OR Use an API like Dev.to, Medium, or YouTube to fetch learning resources.
    Store fetched resources in the database.

📌 Deliverable: The app can automatically pull new learning resources.

### Build the UI & Connect Frontend to Backend

✅ Goals:

    Create a simple frontend to display learning resources.

✅ Tasks:

    Build HTML templates using Go’s html/template.
    Display learning materials in a clean table layout.
    Add category-based filtering.
    Use AJAX to interact with the API dynamically.

📌 Deliverable: Users can view and filter learning resources in the UI.

### Deployment & Documentation

✅ Goals:

    Deploy the Daily Learning Feed online.

✅ Tasks:

    Choose a deployment platform (Railway, Heroku, Fly.io).
    Set up environment variables (.env using godotenv).
    Test API access and database connections post-deployment.

📌 Deliverable: The app is live and accessible online.

### Final Testing & Project Wrap-Up

✅ Goals:

    Perform final optimizations.

✅ Tasks:

    Fix any UI/UX issues and test API responses.
    Optimize database queries.
    Record a recap video showcasing how the Daily Learning Feed works.

📌 Deliverable: Final working app is tested, polished.
