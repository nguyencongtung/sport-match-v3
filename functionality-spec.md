# SportMatch Functionality Specification

This document details the core functionalities of the SportMatch application, outlining user actions, system responses, business logic, and error handling for each feature.

## 1. Onboarding & Authentication

### 1.1 Welcome Screen
*   **Feature Name:** Welcome Screen Display
*   **Description:** Displays the app logo, name, and a tagline, then transitions to the Authentication screen.
*   **User Actions/Inputs:** None (automatic transition) or "Get Started" button tap.
*   **System Responses/Outputs:** Displays welcome message, then navigates to Authentication screen.
*   **Business Logic/Rules:** Transitions after 1-2 seconds or user interaction.

### 1.2 Authentication Screen

#### Registration Flow
*   **Feature Name:** User Registration
*   **Description:** Allows new users to create an account using email and password.
*   **User Actions/Inputs:** User enters `Email`, `Password`, `Confirm Password`. Taps `Create Account` button.
*   **System Responses/Outputs:**
    *   On success: Logs in the user, redirects to Profile Setup.
    *   On failure: Displays error messages (e.g., "Email already exists," "Passwords do not match," "Invalid email format," "Password too short").
*   **Business Logic/Rules:**
    *   Email must be a valid format.
    *   Password must meet complexity requirements (e.g., minimum length, special characters).
    *   Password and Confirm Password must match.
    *   Email must be unique.

#### Login Flow
*   **Feature Name:** User Login
*   **Description:** Allows existing users to log into their account.
*   **User Actions/Inputs:** User enters `Email`, `Password`. Taps `Log In` button.
*   **System Responses/Outputs:**
    *   On success: Logs in the user, redirects to the default screen (Swipe Screen).
    *   On failure: Displays error message (e.g., "Invalid email or password").
*   **Business Logic/Rules:** Validates credentials against stored user data.

#### Social Login (Recommended)
*   **Feature Name:** Social Login (Google, Apple, Facebook)
*   **Description:** Allows users to register or log in using their social media accounts.
*   **User Actions/Inputs:** User taps `Continue with [Social Provider]` button.
*   **System Responses/Outputs:** Integrates with social provider for authentication. On success, logs in/registers user and redirects.
*   **Business Logic/Rules:** Handles OAuth flow with selected social provider.

### 1.3 Community Standards Agreement
*   **Feature Name:** Community Standards Agreement
*   **Description:** Presents community guidelines to first-time users, requiring agreement to proceed.
*   **User Actions/Inputs:** User reads standards, checks "I Agree" checkbox (if present), taps `I Agree` or `Agree and Continue` button.
*   **System Responses/Outputs:**
    *   On agreement: Proceeds to Profile Setup.
    *   If not agreed: Prevents progression.
*   **Business Logic/Rules:** Must be agreed upon by first-time users before accessing core app features.

## 2. User Profile Setup

### 2.1 Static Attributes Setup
*   **Feature Name:** Static Profile Setup
*   **Description:** Collects permanent user details like name, date of birth, and gender.
*   **User Actions/Inputs:** User enters `Name`, selects `Date of Birth` (via date picker), selects `Gender`. Taps `Continue` or `Next`.
*   **System Responses/Outputs:** Stores static profile data. Navigates to Dynamic Attributes Setup.
*   **Business Logic/Rules:**
    *   Name is required.
    *   Date of Birth is required; calculates age; minimum age validation (18+).
    *   Gender selection is required.
    *   These fields are unchangeable after initial setup.

### 2.2 Dynamic Attributes Setup
*   **Feature Name:** Dynamic Profile Setup
*   **Description:** Collects editable user preferences like sports, target gender, age range, and skill level.
*   **User Actions/Inputs:** User selects `My Sports` (multi-select), `Looking For (Target Gender)`, `Age Range` (dual-slider), `Skill Level`. Taps `Continue` or `Next`.
*   **System Responses/Outputs:** Stores dynamic profile data. Navigates to Photo Upload Process.
*   **Business Logic/Rules:**
    *   At least one sport must be selected.
    *   Target Gender, Age Range, and Skill Level are required.
    *   These fields are editable later via the Profile Screen.

### 2.3 Photo Upload Process
*   **Feature Name:** Profile Photo Upload
*   **Description:** Allows users to upload up to 6 photos, with the first being the primary profile picture.
*   **User Actions/Inputs:** User taps `+` icon in photo slots, selects photos from gallery or takes new ones with camera. Drags and drops to reorder. Taps `X` to remove. Taps `Finish Setup` or `Let's Go!`.
*   **System Responses/Outputs:** Uploads and stores photos. Sets primary photo. Displays "Profile Complete!" animation. Redirects to Swipe Screen.
*   **Business Logic/Rules:**
    *   Allows up to 6 photos.
    *   First uploaded photo (or first in order) is designated as primary.
    *   Photos can be reordered.
    *   Photos can be removed.

## 3. Main Navigation & Screens

### 3.1 Bottom Navigation Bar
*   **Feature Name:** Persistent Bottom Navigation
*   **Description:** Provides quick access to main app sections: Swipe, Events/Matches, Chat, Profile.
*   **User Actions/Inputs:** User taps on navigation icons.
*   **System Responses/Outputs:** Navigates to the corresponding screen. Displays active tab state. Displays notification badges.
*   **Business Logic/Rules:**
    *   Swipe screen is the default view after onboarding.
    *   Notification badges appear for unread messages (Chat) or new event updates (Events/Matches).

## 4. Detailed Feature Screens

### 4.1 Swipe Screen
*   **Feature Name:** User Discovery (Swiping)
*   **Description:** Presents user profiles as swipeable cards for liking or passing.
*   **User Actions/Inputs:** User swipes right (Like), swipes left (Pass), or taps `Like (❤️)` / `Pass (X)` buttons.
*   **System Responses/Outputs:**
    *   Records user's preference (like/pass).
    *   If mutual like: Triggers "Connection" notification/modal. Creates a new Connection.
    *   If no more profiles: Displays "Empty Stack" message with prompt to adjust preferences.
*   **Business Logic/Rules:**
    *   Only shows users matching the current user's dynamic preferences (age range, target gender, sports).
    *   Prevents showing already-swiped profiles.
    *   "Connection" is formed only on mutual right swipes.

### 4.2 Events/Matches Screen

#### Event/Match Listing
*   **Feature Name:** Event/Match Discovery
*   **Description:** Displays a scrollable list of upcoming sporting events/matches.
*   **User Actions/Inputs:** User scrolls through the list. Taps on an Event/Match Card. Uses filter options (sport, date range, level).
*   **System Responses/Outputs:**
    *   Displays Event/Match Cards sorted by start time.
    *   Opens detailed event view on card tap.
    *   Refreshes list based on applied filters.
*   **Business Logic/Rules:**
    *   Events are sorted by `Start Time`.
    *   Filters apply dynamically to the displayed list.
    *   Pagination/infinite scroll loads more events as user scrolls.

#### Create Event/Match
*   **Feature Name:** Event/Match Creation
*   **Description:** Allows users to create new sporting events.
*   **User Actions/Inputs:** User taps `+` icon. Fills out form: `Sport Name`, `Start Time`, `Max People`, `Fee`, `Location`, `Level`, `Description` (optional). Taps `Create Event` button.
*   **System Responses/Outputs:**
    *   On success: Creates new event, displays "Match Created!" toast notification. Adds event to the list.
    *   On failure: Displays error messages for invalid input (e.g., "Max people must be greater than 0").
*   **Business Logic/Rules:**
    *   All fields except Description are required.
    *   `Max People` must be a positive integer.
    *   `Start Time` must be in the future.
    *   Location can be selected via map integration.

### 4.3 Chat Screen

#### Conversations List
*   **Feature Name:** Chat List Display
*   **Description:** Shows a list of new connections and ongoing chat conversations.
*   **User Actions/Inputs:** User scrolls through lists. Taps on a connection/conversation.
*   **System Responses/Outputs:**
    *   Displays new connections horizontally.
    *   Displays ongoing conversations vertically with last message preview and timestamp.
    *   Opens one-on-one chat window on tap.
    *   Displays unread message indicators.
*   **Business Logic/Rules:**
    *   New connections are those with mutual likes but no prior chat messages.
    *   Conversations are sorted by last message timestamp.

#### One-on-One Chat
*   **Feature Name:** Real-time One-on-One Chat
*   **Description:** Enables real-time text communication between connected users.
*   **User Actions/Inputs:** User types message into input field, taps `Send` button.
*   **System Responses/Outputs:** Sends message, displays message in real-time in chat window.
*   **Business Logic/Rules:**
    *   Only text messages are supported in MVP.
    *   Messages are exchanged in real-time via Socket.IO.
    *   Provides options to report or block users.

### 4.4 Profile Screen

#### Profile Viewing
*   **Feature Name:** User Profile Viewing
*   **Description:** Displays the user's own profile information and photos.
*   **User Actions/Inputs:** User views their profile. Taps on photos to view full screen.
*   **System Responses/Outputs:** Displays user's photos, static attributes (Name, Age, Gender), and dynamic attributes (Sports, Looking For, Age Range, Skill Level).
*   **Business Logic/Rules:** Static attributes are displayed as uneditable.

#### Profile Editing
*   **Feature Name:** Profile Editing (Dynamic Attributes)
*   **Description:** Allows users to modify their dynamic profile attributes and manage photos.
*   **User Actions/Inputs:** User taps `Edit Profile` button. Modifies `Photos`, `My Sports`, `Looking For`, `Age Range`, `Skill Level`. Saves changes.
*   **System Responses/Outputs:** Updates dynamic profile data. Reflects changes on the profile screen.
*   **Business Logic/Rules:** Only dynamic attributes are editable.

#### Settings Access
*   **Feature Name:** App Settings Access
*   **Description:** Provides access to app-level settings like notifications, privacy, and logout.
*   **User Actions/Inputs:** User taps `Settings` (gear icon). Selects options like `Logout`, `Community Guidelines`, `Help`.
*   **System Responses/Outputs:** Navigates to selected setting screen or performs action (e.g., logs out user).
*   **Business Logic/Rules:** Standard app settings functionality.
