# SportMatch Technical Requirements

This document outlines the technical stack, architectural considerations, and tooling for the SportMatch MVP, based on the `README.md` and `user-flow-final.txt` files.

## 1. Architecture Strategy

*   **Monorepo Structure:** A single repository housing frontend, backend, and shared utilities for simplified version control, CI/CD, and dependency management. This approach is chosen for its benefits in a lean MVP development cycle, allowing for faster iteration and easier collaboration.
    *   **Frontend:** `/frontend/` (React Native app for iOS and Android)
    *   **Backend:** `/backend/` (Golang backend)
    *   **Shared Utilities:** `/shared/` (Common code like constants, validators, types)

## 2. Technology Stack

### 2.1 Frontend
*   **Framework:** React Native
*   **Navigation:** React Navigation
*   **State Management:** (Implicit, but typically React Context API or Redux/Zustand for larger apps)
*   **UI Components:**
    *   `react-native-deck-swiper`: For the fluid card swiping interface on the Swipe Screen.
    *   `react-native-gifted-chat`: For real-time messaging bubbles in the Chat Screen.
*   **Local Storage:** AsyncStorage helpers (for user preferences, tokens, etc.).
*   **Environment Variables:** `react-native-config` (for managing API keys, etc.).
*   **Geolocation:** (Implicit, but will require native module integration for GPS access).

### 2.2 Backend
*   **Runtime:** Go (Golang 1.20+)
*   **Web Framework:** Gin Gonic (for building RESTful APIs)
*   **Database:** MongoDB (with official Go driver for schema definition and interaction).
*   **Real-time Communication:** Gorilla WebSocket (for real-time chat and match updates).
*   **Authentication:** JWT (JSON Web Tokens) with a suitable Go library for secure user authentication.
*   **Middleware:** Gin middleware for JWT authentication, validation, and file uploads.
*   **File Uploads:** (Implicit, but will require Go's `net/http` package for handling profile photo uploads).
*   **Environment Variables:** `.env` file or Go's `os` package (for `MONGODB_URI`, `JWT_SECRET`, etc.).

### 2.3 Shared Utilities
*   **Constants:** Shared definitions (e.g., sports list, skill levels).
*   **Validators:** Shared validation schemas (e.g., for user profiles, match creation) to ensure consistent data integrity across frontend and backend.
*   **Types:** (If using TypeScript, shared type definitions).

## 3. Deployment & Infrastructure

*   **Cloud Provider:** mac local
    *   **Backend Deployment:** mac local
    *   **Frontend Deployment:** mac local
    *   **Database Hosting:** self-hosted MongoDB on mac local
    *   **Storage:** mac local

## 4. Tooling and Development Workflow

*   **Package Management:** Yarn Workspaces or npm workspaces at the monorepo root `package.json` to manage dependencies for `frontend` and `backend`.
*   **Version Control:** Git with a branching strategy (e.g., `main`, `develop`, feature branches).
*   **CI/CD:** GitHub Actions or AWS CodePipeline for automated testing and deployment.
*   **Testing:** Go's built-in testing framework for backend unit/integration tests; React Native Testing Library for frontend tests.
*   **Environment Variables:** Managed via `.env` files and `react-native-config`.

## 5. Key Technical Considerations

*   **Geolocation:** Integration with Google Maps API for location selection in event creation and distance calculation for user profiles. API keys should be securely stored.
*   **Security:**
    *   JWT middleware for API authentication.
    *   Input validation using shared validators (e.g., `go-playground/validator` on backend).
    *   Password hashing (e.g., `golang.org/x/crypto/bcrypt`).
    *   Secure handling of environment variables.
*   **Real-time Communication:** Efficient management of Gorilla WebSocket connections for chat and live event updates.
*   **Scalability:** The chosen stack and monorepo structure are designed to be scalable for future features and increased user load, with options to split into microservices if needed post-MVP.
*   **Error Handling:** Robust error handling on both frontend and backend, with clear user feedback.
*   **Analytics:** Integration for tracking user engagement (swipes, matches, event joins).
