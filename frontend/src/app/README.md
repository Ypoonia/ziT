Next.js App Router entrypoint.

Responsible for:

- Route definitions
- Nested layouts
- Server components
- Route-level loading/error boundaries
- Page composition

Rules:

- Keep business logic minimal
- Pages should orchestrate feature modules
- Avoid reusable logic/components here

Example:

```txt
app/
├── login/
├── dashboard/
├── board/[boardId]/
└── issue/[issueId]/
```
