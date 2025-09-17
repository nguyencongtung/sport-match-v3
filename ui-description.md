# SportMatch UI/UX Design - Consolidated Vision

## General UI/UX Principles:
1.  **Clarity & Simplicity:** All UI elements and flows must be intuitive and easy to understand at a glance. Reduce cognitive load for users.
2.  **Consistency is Key:** Maintain a consistent design language across all screens (colors, typography, iconography, button styles, interaction patterns). This builds familiarity and trust.
3.  **Vibrant & Energetic Design:** Leverage a dynamic, sports-appropriate color palette (e.g., blues, greens, oranges) that evokes movement and energy. Use bold typography for headings and key information.
4.  **Clear Visual Hierarchy:** Ensure important information and actionable elements (buttons) stand out. Use size, color, and placement to guide the user's eye effectively.
5.  **Purposeful Animations & Haptic Feedback:** Implement subtle, engaging animations for actions like swiping, button presses, and transitions. Haptic feedback should enhance the tactile experience, especially during swiping or successful actions, making the app feel responsive and engaging.
6.  **Accessibility:** Design with accessibility in mind, ensuring sufficient contrast, legible typography, and appropriate touch target sizes for all users.
7.  **Loading & Error States:** Implement clear loading indicators and user-friendly error messages with actionable guidance.
8.  **Empty States:** Provide helpful, encouraging messages and clear calls-to-action for screens with no content (e.g., "No matches yet—start swiping!", "No events nearby—create one!").

## Screen-Specific UI/UX Suggestions:

### 1. Onboarding & Authentication:

*   **Welcome Screen:**
    *   **Content:** App logo and name: "Welcome to SportMatch" with a brief tagline like "Find your perfect game partner" or "Match. Play. Connect!".
    *   **Action:** Automatically transitions to the Authentication screen after a brief moment (1-2 seconds) or via a "Get Started" button.
    *   **UI Suggestion:** Minimalist design with sports-inspired imagery, vibrant hero image, and a bold primary color for the button. Introduce subtle, sports-themed animations (e.g., a bouncing ball, a running figure) or a short, looping video background showcasing various sports activities to immediately convey the app's purpose and energy.
    *   **"Get Started" Button:** Make this button highly prominent with a strong primary color and clear, concise text. Consider a subtle animation on tap.
*   **Authentication Screen:**
    *   **Content:** Clean interface with primary calls-to-action.
    *   **UI Elements:**
        *   Button: `Sign Up`
        *   Button: `Log In`
        *   (Recommended): Social login options like `Continue with Google`, `Continue with Apple`, or `Continue with Facebook`.
    *   **Registration Flow (Initiated from `Sign Up`)**
        *   **Fields:** Email, Password (with confirmation).
        *   **Action:** `Create Account` button. On success, the user is logged in for the first time.
    *   **Login Flow (Initiated from `Log In`)**
        *   **Fields:** Email, Password.
        *   **Action:** `Log In` button.
    *   **Social Login Prominence:** Prioritize social login options by making them visually more prominent than email/password fields. Use official brand icons and colors for immediate recognition and trust.
    *   **Button Hierarchy:** Clearly differentiate `Sign Up` (primary action for new users) from `Log In` (secondary action) through color, size, or style.
    *   **UI:** Use clear separators (e.g., "OR") between social login options and traditional email/password fields.
*   **Community Standards Agreement (First-Time Login Only):**
    *   **Purpose:** To ensure a safe and respectful environment from the start.
    *   **Content:** Display a modal or full screen with a summary of key rules (e.g., "Be respectful," "Play safe," "Show up on time").
    *   **Action:** User must tap an `I Agree` button or check a box to proceed.
    *   **UI Suggestion:** Present as a clear, concise modal. Use checkmarks or simple icons next to each bullet point for quick readability. The "I Agree" button should be a prominent, affirmative button. Consider "Agree and Continue" to imply progression. Implement a checkbox next to "I have read and agree to the Community Standards" before enabling the button to ensure explicit consent.

### 2. User Profile Setup:

*   **Navigation/Progress:** Use a clear, persistent visual progress bar or "Step X of Y" indicator at the top of the screen to manage user expectations and show progress. Ideally, tapping it could show a summary of steps.
*   **"Continue" / "Next" Buttons:** These should be consistently placed (e.g., bottom right) and styled as primary action buttons. They should only become active once all required fields on the current step are valid.
*   **Static Attributes (Permanent Details):**
    *   **Screen Title:** "Let's Set Up Your Profile"
    *   **Helper Text:** "These details help others find you and cannot be changed later."
    *   **Input Fields:** `Name` (Text field), `Date of Birth` (Date picker, to calculate Age, with minimum age validation 18+), `Gender` (Selectable options: Male, Female, Other/Prefer not to say).
    *   **Action:** `Continue` or `Next` button.
    *   **UI Suggestion:** Progressive disclosure with clear progress indicators (e.g., "Step 1 of 3"). Clearly indicate why these fields are permanent (e.g., a small "i" icon with a tooltip explaining immutability).
*   **Dynamic Attributes (Editable Later):**
    *   **Screen Title:** "What's Your Game?"
    *   **Helper Text:** "Tell us about your sports and what you're looking for. You can change this anytime."
    *   **Input Fields / Components:**
        *   **My Sports:** A multi-select list or tag system where users can pick the sports they play.
            *   **UI Suggestion:** Implement a highly visual and interactive multi-select. Use a grid of large, tappable sport icons/cards. When selected, they should highlight or animate. Include a search bar with pre-set popular sports tags and recognizable icons. Tapping a sport tag should toggle its selection, with clear visual feedback (e.g., a checkmark, highlight, or moving it to a "Selected Sports" section).
        *   **Looking For (Target Gender):** Selectable options (Men, Women, Everyone/Any). Use clear, distinct buttons or segmented controls (e.g., "Men," "Women," "Everyone") for easy selection.
        *   **Age Range:** A dual-slider to select the desired age range of other users (e.g., 18-30). Ensure the current selected range is clearly displayed numerically above or below the slider.
        *   **Skill Level:** Self-assessment using selectable options (Beginner, Intermediate, Advanced, All Levels). Use a segmented control or a set of distinct buttons for easy selection. Consider a more descriptive scale (e.g., a 1-5 star rating, or a slider with labels like "Casual Player" to "Competitive Athlete") for nuanced self-assessment.
    *   **Action:** `Continue` or `Next` button.
*   **Photo Upload Process:**
    *   **Screen Title:** "Add Your Photos"
    *   **Helper Text:** "Upload up to 6 photos to stand out!"
    *   **Input Fields / Components:** A grid of 6 slots for pictures. The first slot is the primary profile picture.
    *   **UI Suggestion:** Use a `+` icon in each slot. Tapping it opens the phone's gallery or camera. Allow users to reorder photos by dragging and dropping. Provide clear guidance on photo requirements. Clearly indicate drag-and-drop functionality (e.g., subtle drag handles or a brief tutorial on first use). Add a small, easily discoverable "X" icon on each photo slot to allow users to remove photos.
    *   **Action:** `Finish Setup` or `Let's Go!` button.
    *   **Completion:** "Profile Complete!" animation (confetti or sports ball bounce), then redirect to default page (Swipe-Card). The "Finish Setup" / "Let's Go!" Button should be the final, celebratory call to action, perhaps with a slightly different, more encouraging tone than "Continue."

### 3. Main Navigation & Bottom Navigation Bar:

*   **UI:** A bar at the bottom of the screen with four clearly labeled icons:
    1.  **Swipe** (Default): For discovering other users. Icon: 🃏
    2.  **Events/Matches**: For viewing and creating sporting events. Icon: 📅
    3.  **Chat**: For messaging with connections. Icon: 💬
    4.  **Profile**: For viewing and managing your own profile. Icon: 👤
*   **Icons & Labels:** Use simple, universally recognizable icons. Add concise text labels below each icon for absolute clarity, especially for new users.
*   **Active State:** Clearly highlight the currently active tab through a distinct color change, subtle animation, a bold text label, or a subtle underline.
*   **Order:** The order (Swipe, Events, Chat, Profile) is logical for a dating/social app.
*   **Notification Badges:** Implement small, clear notification badges on the `Chat` icon for unread messages and potentially on the `Events/Matches` icon for new event invitations or updates to joined events.

### 4. Detailed Feature Screens:

#### 4.1 Swipe Screen (Default View)
*   **Purpose:** The primary user discovery engine, mimicking Tinder-style card swiping.
*   **UI:**
    *   A stack of **Swipe Cards** dominates the screen.
    *   Each card displays another user's information: Primary Photo (takes up most of the card, with a swipeable carousel for up to 6 photos), Name, Age, Sports in common (highlighted with icons or colored tags), Distance away (e.g., "5 km away"), Short bio or preferred play level/skill level indicator.
    *   Action buttons (optional, as gestures are primary): `Pass (X)`, `Like (❤️)`.
    *   **Swipe Cards - Information Hierarchy:** Prioritize key information (Name, Age, Sports in common, Distance) at a glance. Use clear, readable typography and icons.
    *   **Swipe Cards - Photo Carousel:** Indicate multiple photos with subtle dots or a small counter (e.g., "1/6") at the bottom of the card. Make it intuitive to swipe through photos within the card.
    *   **Swipe Cards - Gestures:** Emphasize the swipe gestures. Ensure smooth transitions and responsive feedback.
    *   **Swipe Cards - Explicit Action Buttons:** While gestures are primary, prominent `Pass (X)` and `Like (❤️)` buttons are crucial for accessibility and clarity. They should be large, easily tappable, and visually distinct (e.g., red for Pass, green for Like). Consider a subtle animation for these buttons when a swipe gesture is performed.
    *   **Swipe Cards - "Sports in Common" Highlighting:** This is a key differentiator. Make common sports visually distinct (e.g., bold text, a different accent color, or a small, shared icon).
    *   **Swipe Cards - "Distance Away":** Ensure this information is clearly visible and easy to locate on the card.
    *   **Connection Notification:** A full-screen, visually rewarding animation is excellent for celebrating a mutual connection. The `Start Chatting` button should be the primary call to action, with `Keep Swiping` as a secondary, less prominent option.
    *   **Empty Stack:** "No more matches nearby—adjust preferences?" with a link to the Profile screen. The message should be encouraging, with a clear, actionable button or tappable link that directly navigates to the Profile screen for preference adjustments.
*   **UI Inspiration:** Tinder's card stack interface with `react-native-deck-swiper` for fluid gestures, engaging, addictive swiping with subtle animations (card tilt on drag).

#### 4.2 Events/Matches Screen
*   **Purpose:** To find and organize actual sporting activities.
*   **UI:**
    *   **Header:** "Upcoming Events" or "Upcoming Matches". "Upcoming Events" is clearer than "Upcoming Matches" as it encompasses a broader range of activities.
    *   **Primary Action:** A `+` icon on the top right to trigger the "Create Event/Match" flow. Place this prominently in the top right corner, perhaps within a circle, to signify creation. Alternatively, consider a floating action button (FAB) at the bottom right for greater discoverability.
    *   **Main Content:** A vertical, scrollable list of "Event/Match Cards." Events are sorted by `Start Time`, with today's events first.
    *   **Event/Match Card Information:** Each card clearly displays: Sport Name & Level (e.g., "Badminton - Intermediate" with recognizable icon), Start Time (e.g., "Today, 6:00 PM" or "Sep 18, 2025 at 2:00 PM" with relative time indicator), Participants (`👤 5/8` (Current / Max) with visual people icons), Location (`📍 Binh Duong Olympic Stadium` with map pin icon), Fee (`💲 50,000 VND` or `Free`).
    *   **Event/Match Card Information - Visual Clarity:** Use a card-based layout with clear separation of information.
    *   **Event/Match Card Information - Sport Name & Level:** Combine these visually (e.g., "Badminton - Intermediate" with a small badminton icon).
    *   **Event/Match Card Information - Time & Date:** Use relative time (e.g., "Today, 6:00 PM") for immediate events, and full date/time for future ones.
    *   **Event/Match Card Information - Participants:** The `👤 5/8` format is excellent. Consider a small progress bar or visual representation of filled slots.
    *   **Event/Match Card Information - Location:** A small map pin icon is effective.
    *   **Event/Match Card Information - Fee:** Clearly display currency.
    *   **Filters:** Top bar dropdowns for sport, date range, level—auto-refresh list. Implement a persistent filter bar at the top, allowing users to quickly refine results without navigating away. Use interactive filter components (e.g., tapping a filter opens a modal or an overlay for selection). Include a "Clear Filters" option. Consider horizontal scrollable "tag" filters for common sports/levels.
    *   **Pagination:** Infinite scroll or "Load More" button, showing a maximum of 20 events at a time. Infinite scroll is preferred for a seamless browsing experience.
*   **Functionality:** Tapping an Event/Match Card opens a detailed view with more information and a "Join" button.
*   **UI Inspiration:** Meetup or Eventbrite's event listings—card-based for visual appeal, with RSVP buttons.

#### Sub-Flow: Create Event/Match
*   **Trigger:** User taps the `+` icon on the Events/Matches screen.
*   **UI:** A full-screen form or modal with the following fields: Sport Name (A dropdown or searchable list (same as in user profile setup, with icon-based picker)), Start Time (A date and time picker (calendar + clock)), Max People (A number input field/slider), Fee (A number input field (with currency symbol). A toggle for "Free" can be included), Location (A text field, ideally integrated with a map API (like Google Maps) to select a venue (searchable map picker)), Level (Selectable options (e.g., Beginner, Intermediate, Advanced)), Description (Optional: A text area for extra details).
    *   **Sport Name:** Use the same visual, icon-based picker as in the profile setup for consistency.
    *   **Location:** Integrate a map view directly into the form, allowing users to drop a pin or search for venues. Provide a clear "Select Location on Map" button.
    *   **Max People / Fee / Level:** Use intuitive input types (e.g., steppers for numbers, segmented controls for level).
    *   **Description:** A multi-line text area.
*   **Action:** A `Create Event` button at the bottom. Prominent at the bottom of the form, enabled only when all required fields are filled.
*   **Success:** Toast notification "Match Created!". A clear, non-intrusive toast notification.

#### 4.3 Chat Screen
*   **Purpose:** To facilitate communication between Connections.
*   **UI:**
    *   **Top Section (New Connections/Friends List):** A horizontally scrolling list of new Connections who you haven't chatted with yet, or frequent contacts. Each item shows the user's profile picture/avatar. "New Connections" is a good label. Use larger, circular profile pictures for this horizontal scroll. A small "New" badge could indicate unread messages from new connections.
    *   **Bottom Section (Conversations):** A standard, vertical list of all ongoing chats. Each list item shows: User's Profile Picture, User's Name, The last message preview (e.g., "Sounds good, see you there!"), Timestamp of the last message, Unread message indicators.
    *   **Bottom Section (Conversations) - Unread Indicators:** Clearly distinguish unread messages (e.g., bold text for the last message, a small colored dot or number badge on the profile picture).
    *   **Bottom Section (Conversations) - Last Message Preview:** Truncate long messages and display the timestamp clearly.
    *   **Chat Window:** Standard chat UI with a clear text input field and a distinct "Send" button (e.g., paper airplane icon).
    *   **Empty State:** "No chats yet—start swiping to make friends!" with a clear clickable link/button that navigates directly to the Swipe screen.
*   **Functionality:** Tapping a new connection or an existing conversation opens the one-on-one chat window. Standard chat functionality: text input, send button, real-time message display. **MVP:** Text only. Future features: media sharing, voice notes, typing indicators, read receipts. **Safety:** Reporting and blocking functionality accessible through each conversation's context menu.
*   **UI Inspiration:** WhatsApp or Messenger's clean conversation views, `react-native-gifted-chat` for real-time messaging bubbles.

#### 4.4 Profile Screen
*   **Purpose:** Allows the user to view their own profile as others see it and to edit their information.
*   **UI:**
    *   Visually displays the user's 6 photos in a grid (editable order/upload). Display photos in an appealing grid. Allow users to tap on photos to view them in full screen and manage their order/primary status.
    *   Shows all their static and dynamic information (Name, Age, Gender, Sports, Looking For, Age Range, Skill Level). Present static and dynamic information clearly. Clearly differentiate static (uneditable) fields from dynamic (editable) fields visually (e.g., greyed-out text for static fields, clear input fields for dynamic ones).
    *   **Primary Action Buttons:**
        *   `Edit Profile`: Takes the user to a screen where they can change their **dynamic attributes** (Photos, My Sports, Looking For, Age Range, Skill Level). Make this a prominent, easily discoverable button, perhaps a primary button or a clear icon (e.g., a pencil icon) next to the profile details. The "Edit Profile" flow should be a dedicated screen.
        *   `Settings` (Gear Icon ⚙️): For app-level settings like notifications, privacy, and logout. Place this in a standard location, like the top right corner, for quick access to app-level settings.
*   **Functionality:**
    *   **Editing:** The "Edit Profile" flow should make it clear which fields (dynamic) can be changed. Static fields (Name, Age, Gender) should be displayed but greyed out/uneditable.
    *   **Settings:** Logout, Community Guidelines, Help.
*   **UI Inspiration:** Hinge's detailed but editable profiles, Tinder's edit view for quick tweaks.
