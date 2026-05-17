# zit

A lightweight task management tool built around a visual kanban board. Create projects, break work into tasks, assign owners, set priorities and deadlines, and track progress as cards move across customizable columns. Designed to keep teams aligned, surface blockers early, and turn scattered to-dos into shippable work.

<details>
<summary><strong>🚀 Repository Standards: Commits, Branches, and Pull Requests</strong></summary>

<br />

Maintaining consistent naming conventions across the repository is crucial for project maintainability, automated changelog generation, and streamlined code reviews. Please adhere to the following guidelines when contributing. 🛠️

---

## 1. 📝 Commit Message Convention

We follow a structured commit format tailored for a multi-layered architecture, combining the affected layer with standard Conventional Commit types.

**Format:**

```txt
<layer>/<type>: <short-description>
```

**General Rules 📏:**

- **Imperative Mood:** Write the description in the imperative, present tense (e.g., use `add` instead of `added` or `adds`).
- **Formatting:** The description should be all lowercase and should _not_ end with a period.
- **Length Limit:** Keep the subject line under 50 characters. If more context is needed, leave a blank line and add a detailed commit body wrapping at 72 characters.

### 🏗️ Layers

| Layer      | Emoji | Description                                 |
| ---------- | ----- | ------------------------------------------- |
| `frontend` | 🎨    | Next.js client-side application changes     |
| `backend`  | ⚙️    | Go backend APIs, services, and routing      |
| `infra`    | ☁️    | Docker, CI/CD pipelines, Terraform, configs |
| `docs`     | 📚    | Documentation, READMEs, and wiki updates    |
| `shared`   | 🔗    | Shared contracts, types, and interfaces     |

### 🧩 Types

| Type       | Emoji | Purpose                                                      |
| ---------- | ----- | ------------------------------------------------------------ |
| `feat`     | ✨    | Introduces a new feature to the codebase                     |
| `fix`      | 🐛    | Patches a bug in the codebase                                |
| `refactor` | ♻️    | Code changes that neither fix a bug nor add a feature        |
| `chore`    | 🧹    | Tooling, config, or minor dependency changes                 |
| `test`     | 🧪    | Adding missing tests or correcting existing ones             |
| `perf`     | ⚡    | Code changes that improve performance                        |
| `style`    | 💅    | Formatting, missing semi-colons, etc. (no code logic change) |

### 💡 Commit Examples

**Frontend 🎨**

```txt
frontend/feat: implement kanban drag and drop
frontend/fix: resolve sidebar hydration mismatch
frontend/style: format component files with prettier
```

**Backend ⚙️**

```txt
backend/feat: add websocket issue events
backend/fix: prevent duplicate sprint creation
backend/perf: optimize database query for user retrieval
```

---

## 2. 🌿 Branch Naming Convention

Branch names should be descriptive, lowercase, and use hyphens for separation. Including an issue or ticket number is highly encouraged for traceability. 🔍

**Format:**

```txt
<layer>/[issue-number]-<short-description>
```

**Examples:**

```txt
frontend/102-kanban-board
backend/45-websocket-engine
infra/docker-setup
```

_(Note: Omit the issue number only if the branch is a minor chore or undocumented hotfix)._

---

## 3. 📥 Pull Request Guidelines

Pull Request titles serve as the primary entry for release notes and history tracking. They should be clear and concise. 📖

**Format:**

```txt
[<Layer>] <Summary changes of>
```

**Examples:**

```txt
[Frontend] Add drag-and-drop board interactions
[Backend] Implement JWT authentication middleware
[Infra] Migrate CI pipeline to GitHub Actions
```

**PR Best Practices ✅:**

- **Link Issues:** Always link the relevant ticket/issue in the PR description (e.g., `Closes #102`).
- **Draft Status:** Use Draft PRs 🚧 for work-in-progress features to signal that the code is not yet ready for review.
- **Description:** Provide a brief summary of _why_ the change is being made and _how_ it was implemented, highlighting any breaking changes ⚠️ or required database migrations.

</details>
</Summary></Layer>
