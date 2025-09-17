# SportMatch MVP – User Flow & Feature Breakdown

This document outlines the user flow and MVP features for the SportMatch application, detailing the step-by-step interactions users will have with the app.

## 1. Onboarding & Authentication

### 1.1 Welcome Screen
*   **Flow:** User opens the app.
*   **Step 1:** App displays "Welcome to SportMatch" with a tagline.
*   **Step 2:** (Automatic) App transitions to Authentication screen after 1-2 seconds, OR (User Action) User taps "Get Started" button.

### 1.2 Authentication Screen

#### Registration Flow
*   **Scenario:** New user registers an account.
*   **Preconditions:** User is on the Authentication screen.
*   **Steps:**
    1.  User taps `Sign Up` button.
    2.  App navigates to Registration form.
    3.  User enters `Email`, `Password`, and `Confirm Password`.
    4.  User taps `Create Account` button.
    5.  **System:** Validates input (email format, password strength, password match, unique email).
    6.  **If Valid:** System creates account, logs in user.
    7.  **If Valid:** App navigates to Community Standards Agreement (if first-time login) or Profile Setup.
    8.  **If Invalid:** App displays specific error messages (e.g., "Email already exists," "Passwords do not match").

#### Login Flow
*   **Scenario:** Existing user logs into their account.
*   **Preconditions:** User is on the Authentication screen.
*   **Steps:**
    1.  User taps `Log In` button.
    2.  App navigates to Login form.
    3.  User enters `Email` and `Password`.
    4.  User taps `Log In` button.
    5.  **System:** Validates credentials.
    6.  **If Valid:** System logs in user.
    7.  **If Valid:** App navigates to the default screen (Swipe Screen).
    8.  **If Invalid:** App displays error message (e.g., "Invalid email or password").

#### Social Login Flow (Recommended)
*   **Scenario:** User logs in/registers using a social media account.
*   **Preconditions:** User is on the Authentication screen.
*   **Steps:**
    1.  User taps `Continue with Google` (or Apple/Facebook) button.
    2.  **System:** Initiates social authentication flow (e.g., opens browser for Google login).
    3.  User completes social provider's authentication.
    4.  **System:** Receives authentication token.
    5.  **If Successful:** System logs in/registers user.
    6.  **If Successful:** App navigates to Community Standards Agreement (if first-time login) or Profile Setup/Swipe Screen.
    7.  **If Failed:** App displays error message.

### 1.3 Community Standards Agreement (First-Time Login Only)
*   **Scenario:** First-time user agrees to community standards.
*   **Preconditions:** User has just registered/logged in for the first time.
*   **Steps:**
    1.  App displays Community Standards modal/screen.
    2.  User reads the standards.
    3.  User taps `I Agree` button (or checks a box and then taps `Agree and Continue`).
    4.  **System:** Records user's agreement.
    5.  App navigates to Static Attributes Profile Setup.

## 2. User Profile Setup

### 2.1 Static Attributes (Permanent Details)
*   **Scenario:** User provides permanent profile details.
*   **Preconditions:** User has agreed to Community Standards.
*   **Steps:**
    1.  App displays "Let's Set Up Your Profile" screen.
    2.  User enters `Name`.
    3.  User selects `Date of Birth` using a date picker.
    4.  User selects `Gender` from options.
    5.  User taps `Continue` or `Next` button.
    6.  **System:** Validates input (Name required, Age 18+, Gender selected).
    7.  **If Valid:** System saves static attributes.
    8.  **If Valid:** App navigates to Dynamic Attributes Profile Setup.
    9.  **If Invalid:** App displays error messages.

### 2.2 Dynamic Attributes (Editable Later)
*   **Scenario:** User provides editable sports and matching preferences.
*   **Preconditions:** User has completed Static Attributes setup.
*   **Steps:**
    1.  App displays "What's Your Game?" screen.
    2.  User selects `My Sports` from a multi-select list/tags.
    3.  User selects `Looking For (Target Gender)` from options.
    4.  User adjusts `Age Range` using a dual-slider.
    5.  User selects `Skill Level` from options.
    6.  User taps `Continue` or `Next` button.
    7.  **System:** Validates input (at least one sport, all fields selected).
    8.  **If Valid:** System saves dynamic attributes.
    9.  **If Valid:** App navigates to Photo Upload Process.
    10. **If Invalid:** App displays error messages.

### 2.3 Photo Upload Process
*   **Scenario:** User uploads profile photos.
*   **Preconditions:** User has completed Dynamic Attributes setup.
*   **Steps:**
    1.  App displays "Add Your Photos" screen with 6 photo slots.
    2.  User taps a `+` icon in a slot.
    3.  User selects a photo from gallery or takes a new one with camera.
    4.  **System:** Uploads photo, displays it in the slot. (Repeat for up to 6 photos).
    5.  (Optional) User drags and drops photos to reorder them.
    6.  (Optional) User taps `X` icon to remove a photo.
    7.  User taps `Finish Setup` or `Let's Go!` button.
    8.  **System:** Designates the first photo as primary.
    9.  **System:** Displays "Profile Complete!" animation.
    10. App redirects to the Swipe Screen.

## 3. Main Navigation & Screens

### 3.1 Bottom Navigation Bar
*   **Scenario:** User navigates between main sections of the app.
*   **Preconditions:** User is logged in and has completed profile setup.
*   **Steps:**
    1.  User taps one of the four icons in the bottom navigation bar: `Swipe`, `Events/Matches`, `Chat`, `Profile`.
    2.  **System:** Highlights the active tab.
    3.  **System:** Navigates to the corresponding screen.
    4.  **System:** Displays notification badges on `Chat` or `Events/Matches` icons if there are unread messages or new event updates.

## 4. Detailed Feature Screens

### 4.1 Swipe Screen (Default View)
*   **Scenario:** User discovers and interacts with other profiles.
*   **Preconditions:** User is on the Swipe Screen.
*   **Steps:**
    1.  App displays a stack of **Swipe Cards**, each showing another user's profile.
    2.  User views the profile details (photos, name, age, common sports, distance, bio).
    3.  User performs one of the following actions:
        *   **Swipe Right** (or taps `Like ❤️` button): Likes the profile.
        *   **Swipe Left** (or taps `Pass X` button): Passes on the profile.
    4.  **System:** Processes the swipe/tap.
    5.  **If Mutual Like:**
        *   **System:** Triggers a "Connection!" notification/modal.
        *   User sees options: `Start Chatting` or `Keep Swiping`.
        *   User taps `Start Chatting`: App navigates to the chat window with the new connection.
        *   User taps `Keep Swiping`: App dismisses modal and continues to the next profile.
    6.  **If No More Profiles:** App displays "No more matches nearby—adjust preferences?" with a link to the Profile screen.
    7.  User taps link: App navigates to Profile screen.

### 4.2 Events/Matches Screen

#### Event/Match Listing Flow
*   **Scenario:** User browses upcoming sporting events.
*   **Preconditions:** User is on the Events/Matches Screen.
*   **Steps:**
    1.  App displays a vertical, scrollable list of "Event/Match Cards," sorted by start time.
    2.  User scrolls down to view more events (infinite scroll).
    3.  (Optional) User applies filters (sport, date range, level) using the filter bar.
    4.  **System:** Refreshes the list of events based on filters.
    5.  User taps an Event/Match Card.
    6.  App opens a detailed view of the selected event.
    7.  User taps `Join` button (on detailed view).
    8.  **System:** Adds user to event participants. Displays confirmation.

#### Create Event/Match Flow
*   **Scenario:** User creates a new sporting event.
*   **Preconditions:** User is on the Events/Matches Screen.
*   **Steps:**
    1.  User taps the `+` icon (or FAB) to create a new event.
    2.  App navigates to the "Create Event/Match" form.
    3.  User fills out the form:
        *   Selects `Sport Name` (from dropdown/picker).
        *   Selects `Start Time` (date and time picker).
        *   Enters `Max People`.
        *   Enters `Fee` (or toggles "Free").
        *   Selects `Location` (via text input, ideally with map integration).
        *   Selects `Level`.
        *   (Optional) Enters `Description`.
    4.  User taps `Create Event` button.
    5.  **System:** Validates form input.
    6.  **If Valid:** System creates the event.
    7.  **If Valid:** App displays "Match Created!" toast notification.
    8.  **If Valid:** App returns to the Events/Matches list, with the new event visible.
    9.  **If Invalid:** App displays error messages for invalid fields.

### 4.3 Chat Screen

#### Conversations List Flow
*   **Scenario:** User views their chat connections and conversations.
*   **Preconditions:** User is on the Chat Screen.
*   **Steps:**
    1.  App displays a horizontally scrolling list of "New Connections" (mutual likes without prior chat).
    2.  App displays a vertical list of "Conversations" (ongoing chats) with last message preview and unread indicators.
    3.  User taps on a "New Connection" or an existing "Conversation" item.
    4.  App navigates to the one-on-one chat window for that connection.
    5.  **If No Chats:** App displays "No chats yet—start swiping to make friends!" with a clickable link to the Swipe screen.

#### One-on-One Chat Flow
*   **Scenario:** User sends and receives messages in a chat.
*   **Preconditions:** User is in a one-on-one chat window.
*   **Steps:**
    1.  User types a message into the text input field.
    2.  User taps the `Send` button.
    3.  **System:** Sends the message.
    4.  **System:** Displays the sent message in the chat window in real-time.
    5.  **System:** Displays incoming messages from the other user in real-time.
    6.  (Optional) User accesses context menu to `Report` or `Block` the other user.

### 4.4 Profile Screen

#### Profile Viewing Flow
*   **Scenario:** User views their own profile.
*   **Preconditions:** User is on the Profile Screen.
*   **Steps:**
    1.  App displays user's profile photos in a grid.
    2.  App displays user's static attributes (Name, Age, Gender) and dynamic attributes (Sports, Looking For, Age Range, Skill Level).
    3.  (Optional) User taps on a photo to view it full screen.

#### Profile Editing Flow
*   **Scenario:** User edits their dynamic profile information.
*   **Preconditions:** User is on the Profile Screen.
*   **Steps:**
    1.  User taps `Edit Profile` button.
    2.  App navigates to the Profile Editing screen.
    3.  User modifies dynamic attributes (e.g., adds/removes photos, changes selected sports, adjusts age range, updates skill level).
    4.  User taps `Save Changes` or `Done` button.
    5.  **System:** Validates changes.
    6.  **If Valid:** System updates profile data.
    7.  **If Valid:** App returns to the Profile Screen, displaying updated information.
    8.  **If Invalid:** App displays error messages.

#### Settings Access Flow
*   **Scenario:** User accesses app settings.
*   **Preconditions:** User is on the Profile Screen.
*   **Steps:**
    1.  User taps `Settings` (gear icon).
    2.  App navigates to the Settings screen.
    3.  User selects an option (e.g., `Logout`, `Community Guidelines`, `Help`).
    4.  **If `Logout`:** System logs out the user. App navigates to the Authentication screen.
    5.  **If `Community Guidelines` / `Help`:** App displays the relevant information.
