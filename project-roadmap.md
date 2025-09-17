# SportMatch Project Roadmap (MVP Development Workflow)

This document outlines the phased development workflow for the SportMatch MVP, based on the `README.md` and `user-flow-final.txt` files. This roadmap provides a structured approach for faster initial release and user feedback collection.

## 1. Overall Development Strategy

*   **Lean MVP Focus:** Prioritize core features for initial release to gather user feedback quickly.
*   **Phased Implementation:** A structured approach to development, breaking down the project into manageable phases.
*   **Monorepo Advantage:** Leverage the monorepo structure for streamlined development, shared code, and simplified CI/CD for a small team.

## 2. Development Phases

### Phase 1: Onboarding & Profiles (4-6 weeks)
*   **Goal:** Establish core user authentication and profile management.
*   **Key Features:**
    *   Initialize monorepo with Yarn Workspaces (or npm workspaces).
    *   Set up `/frontend` with React Native, React Navigation, and AsyncStorage.
    *   Set up `/backend` with Gin Gonic, MongoDB Go driver, and JWT authentication.
    *   Implement Welcome Screen.
    *   Implement Authentication Screen (Sign Up, Log In, Social Login).
    *   Implement Community Standards Agreement.
    *   Implement User Profile Setup (Static Attributes, Dynamic Attributes, Photo Upload).
    *   Implement User Profile Viewing and Editing (for dynamic attributes).
    *   Implement Find Friend Screen (initial swipe interface).
*   **Relevant Files/Modules:**
    *   `/frontend/src/screens/WelcomeScreen.js`
    *   `/frontend/src/screens/LoginScreen.js`
    *   `/frontend/src/screens/RegisterScreen.js`
    *   `/frontend/src/screens/ProfileSetupScreen.js`
    *   `/frontend/src/screens/ProfileDynamicAttributesScreen.js`
    *   `/frontend/src/screens/ProfilePhotoUploadScreen.js`
    *   `/frontend/src/screens/ProfileScreen.js`
    *   `/frontend/src/screens/FindFriendScreen.js`
*   `/backend/internal/models/user.go`
*   `/backend/internal/routes/userRoutes.go`
*   `/backend/internal/controllers/userController.go`
*   `/backend/internal/middleware/auth.go`
*   `/backend/internal/middleware/upload.go` (for photos)
*   `/backend/internal/models/swipe.go`
*   `/backend/internal/routes/swipeRoutes.go`
*   `/backend/internal/controllers/swipeController.go`
    *   `/shared/validators/userValidator.js`

### Phase 2: Chat & Matching Core (3-4 weeks)
*   **Goal:** Enable real-time communication and basic event discovery.
*   **Key Features:**
*   Add Chat functionality (`/src/screens/ChatScreen.js`) with Gorilla WebSocket integration.
*   Implement real-time chat events (`/backend/internal/websocket/chatWebSocket.go`).
    *   Implement Matching Screen (`/src/screens/MatchingScreen.js`) for event discovery.
    *   Implement basic Event/Match Listing.
*   **Relevant Files/Modules:**
    *   `/frontend/src/screens/ChatScreen.js`
    *   `/frontend/src/screens/MatchingScreen.js`
    *   `/frontend/src/components/ChatBubble.js`
    *   `/frontend/src/components/MatchCard.js`
*   `/backend/internal/models/chat.go`
*   `/backend/internal/routes/chatRoutes.go`
*   `/backend/internal/controllers/chatController.go`
*   `/backend/internal/websocket/chatWebSocket.go`
*   `/backend/internal/models/match.go`
*   `/backend/internal/routes/matchRoutes.go`
*   `/backend/internal/controllers/matchController.go`

### Phase 3: Event Creation & Enhancements (2-3 weeks)
*   **Goal:** Allow users to create and manage sporting events, and integrate notifications.
*   **Key Features:**
    *   Build Create Match functionality (`/src/screens/CreateMatchScreen.js`) with form and map integration for location.
    *   Integrate AWS SNS for push notifications (for new likes, upcoming matches, chat messages) using Go SDK.
    *   Refine event filtering and joining logic.
*   **Relevant Files/Modules:**
    *   `/frontend/src/screens/CreateMatchScreen.js`
*   `/backend/internal/routes/matchRoutes.go` (for event creation endpoint)
*   `/backend/internal/controllers/matchController.go` (for event creation logic)
*   `/shared/constants/sports.go` (for sport selection)
*   (New Go modules for AWS SNS integration)

### Phase 4: Deployment, Testing & Launch (1-2 weeks)
*   **Goal:** Prepare the application for production and launch.
*   **Key Activities:**
    *   Deploy backend to AWS (EC2/Lambda) as a Go application.
    *   Deploy frontend to AWS Amplify (or equivalent for mobile app stores).
    *   Conduct end-to-end testing on iOS/Android devices.
    *   Perform security audits and performance optimizations.
    *   Prepare for app store submissions (iOS App Store, Google Play Store).
    *   Launch MVP.
*   **Relevant Files/Modules:**
    *   `/scripts/deploy.sh` (or similar deployment scripts)
    *   CI/CD configurations (e.g., GitHub Actions workflows)
    *   Testing frameworks (Jest, Supertest)

## 3. Additional Considerations During Development

*   **Geolocation:** Securely store Google Maps API keys and implement geolocation helpers.
*   **Security:** Implement robust JWT middleware (e.g., `gin-jwt`), input validation (e.g., `go-playground/validator`), and password hashing (e.g., `golang.org/x/crypto/bcrypt`).
*   **Testing:** Write comprehensive unit and integration tests for both frontend (React Native Testing Library) and backend (Go's built-in testing framework).
*   **Analytics:** Integrate analytics to track key user engagement metrics.
*   **Edge Cases:** Plan for and handle scenarios like no GPS, low match density, network errors.
