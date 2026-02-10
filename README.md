# Learn System Design

This repository contains a collection of mini-projects and isolated concepts designed to practice System Design principles and explore new technology stacks.

## 📂 Repository Structure

### 🔹 [Concepts](./concepts/)
Small, focused scripts or demos illustrating specific algorithms or system design building blocks.
*Used for: Learning a single piece of the puzzle (e.g., "How does a Bloom Filter work?").*

### 🔹 [Projects](./projects/)
Complete, standalone mini-applications solving broader system design problems.
*Used for: Integrating multiple concepts into a working solution (e.g., "Design a Chat System").*

**Structure per Project:**
- `README.md`: The shared problem statement.
- `challenge/`: A clean starting point for students (teaching mode).
- `solution/`: Your reference implementation and full design notes (learning mode).

### 🔹 [Templates](./templates/)
Starter boilerplate for creating new projects with a consistent structure.

---

## 🚀 Getting Started

### Creating a New Mini-Project
1. **Choose a name** for your project (e.g., `url-shortener`).
2. **Copy the template**:
   ```bash
   cp -r templates/project-starter projects/url-shortener
   ```
3. **Rename the language folder**:
   ```bash
   mv projects/url-shortener/lang projects/url-shortener/go
   ```
4. **Initialize the project** in your desired language within `projects/url-shortener/go/solution/`.
5. **Document your design** in `projects/url-shortener/design.md`.

## 📚 Syllabus / Progress Log

| Project / Concept | Type | Tech Stack | Status | Key Learnings |
|-------------------|------|------------|--------|---------------|
| LRU Cache | Project| Go | ✅ | Thread Safety, Map+LinkedList |
| LFU Cache | Project | Go | 🚧 | Frequency Tracking, O(1) Complexity |
| Kth Largest | Project | Go | 🚧 | Min-Heaps, Streaming Data |

---
*Created for teaching & continuous learning.*
