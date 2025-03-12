# **Online Art Gallery - Technical Documentation**

## **Tech Stack**

- **Backend:** FastAPI (Render.com)
- **Frontend:** Next.js (Vercel)
- **Database:** PostgreSQL (Render.com)
- **Image Storage:** Cloudinary
- **Future Migration:** AWS (S3, RDS, EC2, Amplify)

---

## **Database Schema (PostgreSQL)**

### **1. Users Table** (Only one master user)

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **2. Art Posts Table**

```sql
CREATE TABLE art_posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT NOT NULL,
    category VARCHAR(50) NOT NULL CHECK (category IN ('photography', 'painting')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **3. Artist Info Table**

```sql
CREATE TABLE artist_info (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio TEXT NOT NULL,  -- Markdown formatted text
    email VARCHAR(255),
    social_links JSONB,
    profile_image_url TEXT
);
```

---

## **API Endpoints (FastAPI)**

### **Authentication**

- `POST /auth/login` → Login as master user
- `POST /auth/logout` → Logout

### **Art Posts Management**

- `GET /posts` → Get all art posts
- `GET /posts/{id}` → Get a single post
- `GET /posts/category/{category}` → Get art posts by category (photography/painting)
- `POST /posts` → Create a new art post (Admin only)
- `PUT /posts/{id}` → Update an art post (Admin only)
- `DELETE /posts/{id}` → Delete an art post (Admin only)

### **Artist Info Management**

- `GET /artist` → Get artist info (Bio in Markdown format)
- `PUT /artist` → Update artist bio, contact details, and social links (Admin only)

### **Image Upload (Cloudinary)**

- `POST /upload` → Upload an image to Cloudinary and return the URL

---

## **Frontend UI Elements (Next.js)**

### **1. Main Page (Gallery Display)**

- Header with **Gallery Name & Owner Info**
- Navigation:
  - **Photography** Page → Displays photography art posts
  - **Paintings** Page → Displays painting art posts
  - **About the Artist** Page → Displays artist's bio and contact info
- Grid/List view of all art posts (title, image, short description)
- Clickable posts to view details

### **2. Art Post Detail Page**

- Full-size artwork image
- Title & description
- Created date
- Back to respective gallery page button

### **3. Admin Panel (Only for Master User)**

- **Login Page** → Username & password authentication
- **Dashboard** → List all posts with edit/delete buttons
- **Create New Post** → Upload image (Cloudinary), enter title, category & description
- **Edit Post Page** → Modify existing post details
- **Delete Confirmation** → Prompt before deletion
- **Edit Artist Info** → Modify artist bio (supports Markdown), email, social links, and profile image

### **4. About the Artist Page**

- Displays **artist's biography** from the database, rendered from Markdown
- Shows **contact details** (Email, social media links)
- Displays **profile image**
- Personal artwork showcase (if applicable)

---

## **Deployment Plan**

1. **Backend** → Deploy FastAPI on Render.com
2. **Frontend** → Deploy Next.js on Vercel
3. **Database** → Use PostgreSQL on Render
4. **Storage** → Cloudinary for image uploads
5. **Future AWS Migration** → Transition to AWS RDS, S3, EC2, Amplify

This ensures an easy start with minimal cost while keeping migration to AWS smooth when needed.
