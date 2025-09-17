# SportMatch Data Model

This document defines the core data entities, their attributes, and relationships within the SportMatch application. This model is based on the information provided in `README.md` and `user-flow-final.txt`.

## 1. Entities

### 1.1 User
Represents an individual user of the SportMatch application.

*   **Attributes:**
    *   `_id` (ObjectId, Primary Key): Unique identifier for the user.
    *   `name` (String, Required): User's full name.
    *   `email` (String, Required, Unique): User's email address, used for login.
    *   `passwordHash` (String, Required): Hashed password for security.
    *   `dateOfBirth` (Date, Required): User's date of birth, used to calculate age.
    *   `gender` (String, Required): User's self-identified gender (e.g., "Male", "Female", "Other/Prefer not to say").
    *   `profilePictureUrls` (Array of Strings, Optional): URLs to user's uploaded profile photos (up to 6). The first URL in the array is typically the primary photo.
    *   `sports` (Array of Strings, Required): List of sports the user plays.
    *   `lookingForGender` (String, Required): Gender preference for connections/matches (e.g., "Men", "Women", "Everyone/Any").
    *   `ageRange` (Object, Required): Preferred age range for connections/matches.
        *   `min` (Number)
        *   `max` (Number)
    *   `skillLevel` (String, Required): Self-assessed skill level (e.g., "Beginner", "Intermediate", "Advanced", "All Levels").
    *   `createdAt` (Date, Required): Timestamp of user account creation.
    *   `updatedAt` (Date, Required): Timestamp of last profile update.

### 1.2 Swipe
Records a user's swipe action on another user's profile.

*   **Attributes:**
    *   `_id` (ObjectId, Primary Key): Unique identifier for the swipe record.
    *   `swiper` (ObjectId, Required, Reference to User): The user who performed the swipe.
    *   `swiped` (ObjectId, Required, Reference to User): The user whose profile was swiped on.
    *   `direction` (String, Required): The direction of the swipe ("right" for like, "left" for pass).
    *   `timestamp` (Date, Required): When the swipe occurred.

### 1.3 Connection
Represents a mutual like between two users, allowing them to chat. This is an implicit entity formed by two `Swipe` records.

*   **Attributes:** (No explicit model, but conceptually links two users)
    *   `user1` (ObjectId, Reference to User): One of the connected users.
    *   `user2` (ObjectId, Reference to User): The other connected user.
    *   `createdAt` (Date, Required): Timestamp when the connection was established.

### 1.4 Chat
Represents a single message within a conversation between two connected users.

*   **Attributes:**
    *   `_id` (ObjectId, Primary Key): Unique identifier for the chat message.
    *   `sender` (ObjectId, Required, Reference to User): The user who sent the message.
    *   `receiver` (ObjectId, Required, Reference to User): The user who received the message.
    *   `message` (String, Required): The content of the message.
    *   `timestamp` (Date, Required): When the message was sent.
    *   `read` (Boolean, Default: false): Indicates if the message has been read by the receiver.

### 1.5 Match (Event)
Represents a scheduled sporting activity that users can create and join.

*   **Attributes:**
    *   `_id` (ObjectId, Primary Key): Unique identifier for the match/event.
    *   `creator` (ObjectId, Required, Reference to User): The user who created the match.
    *   `sport` (String, Required): The name of the sport (e.g., "Badminton", "Soccer").
    *   `startTime` (Date, Required): The date and time the match is scheduled to start.
    *   `maxPeople` (Number, Required): The maximum number of participants allowed for the match.
    *   `currentParticipants` (Number, Default: 1): The current number of participants.
    *   `fee` (Number, Optional, Default: 0): The cost to join the match, if any.
    *   `location` (Object, Required): Details about the match location.
        *   `name` (String, Required): Name of the venue/location.
        *   `address` (String, Optional): Full address.
        *   `coordinates` (Array of Numbers [longitude, latitude], Optional): Geo-coordinates for map integration.
    *   `level` (String, Required): The recommended skill level for the match (e.g., "Beginner", "Intermediate").
    *   `description` (String, Optional): Additional details or notes about the match.
    *   `participants` (Array of ObjectId, Reference to User): List of users who have joined the match.
    *   `createdAt` (Date, Required): Timestamp of match creation.
    *   `updatedAt` (Date, Required): Timestamp of last match update.

## 2. Relationships

*   **User to Swipe:** One-to-Many (A User can perform many Swipes, and be swiped on many times).
*   **User to Connection:** Many-to-Many (Users form Connections with each other).
*   **User to Chat:** One-to-Many (A User can send/receive many Chat messages).
*   **User to Match:**
    *   One-to-Many (A User can `create` many Matches).
    *   Many-to-Many (A User can `participate` in many Matches, and a Match can have many Participants).

## 3. Validation Rules (Examples)

*   **User:**
    *   `email` must be unique and valid format.
    *   `password` must meet complexity requirements.
    *   `dateOfBirth` must result in an age of 18 or older.
    *   `sports` array must not be empty.
*   **Match:**
    *   `startTime` must be in the future.
    *   `maxPeople` must be a positive integer.
    *   `currentParticipants` must not exceed `maxPeople`.
    *   `sport`, `startTime`, `maxPeople`, `location`, `level` are required.
