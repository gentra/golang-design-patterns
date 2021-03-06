# Introduction
Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Personally, I'd say this is a leveled-up version of decorator pattern, so I'd suggest learning [decorator-pattern](https://github.com/gentra/golang-design-patterns/tree/master/decorator) first.

[![](https://mermaid.ink/img/eyJjb2RlIjoiY2xhc3NEaWFncmFtXG4gIGNsYXNzIG1haW57XG4gICAgbWFpbigpXG4gIH1cblxuICBjbGFzcyBDaGVja291dCB7XG4gICAgPDxpbnRlcmZhY2U-PlxuICAgICtFeGVjdXRlKCpDYXJ0KSBlcnJvclxuICAgICtTZXROZXh0KENoZWNrb3V0KVxuICB9XG5cbiAgQ2hlY2tvdXQ8fC4uRnJhdWRDaGVjazogaW1wbGVtZW50c1xuICAlJUZyYXVkQ2hlY2suLj5DaGVja291dDogY2FsbHMgbmV4dFxuICBDaGVja291dDx8Li5QYXltZW50UHJlcGFyYXRpb246IGltcGxlbWVudHNcbiAgJSVQYXltZW50UHJlcGFyYXRpb24uLj5DaGVja291dDogY2FsbHMgbmV4dFxuICBDaGVja291dDx8Li5Qcm9kdWN0VmFsaWRhdGlvbjogaW1wbGVtZW50c1xuICAlJVByb2R1Y3RWYWxpZGF0aW9uLi4-Q2hlY2tvdXQ6IGNhbGxzIG5leHRcbiAgQ2hlY2tvdXQ8fC4uUHJvbW9WYWxpZGF0aW9uOiBpbXBsZW1lbnRzXG4gICUlUHJvbW9WYWxpZGF0aW9uLi4-Q2hlY2tvdXQ6IGNhbGxzIG5leHRcblxuICBjbGFzcyBjaGVja291dENvbnN1bWVye1xuICAgIDw8aW5zdGFuY2U-PlxuICB9XG4gIFBheW1lbnRQcmVwYXJhdGlvbiAtLW8gY2hlY2tvdXRDb25zdW1lcjogcGFydCBvZlxuICBQcm9kdWN0VmFsaWRhdGlvbiAtLW8gY2hlY2tvdXRDb25zdW1lcjogcGFydCBvZlxuXG4gICAgY2xhc3MgY2hlY2tvdXRCaWdQcm9tb3tcbiAgICA8PGluc3RhbmNlPj5cbiAgfVxuICBQYXltZW50UHJlcGFyYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcbiAgUHJvZHVjdFZhbGlkYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcbiAgRnJhdWRDaGVjayAtLW8gY2hlY2tvdXRCaWdQcm9tbzogcGFydCBvZlxuICBQcm9tb1ZhbGlkYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcblxuICBjbGFzcyBwcm92aWRlQ2hlY2tvdXR7XG4gICAgPDxmdW5jdGlvbj4-XG4gIH1cblxuICBwcm92aWRlQ2hlY2tvdXQuLmNoZWNrb3V0Q29uc3VtZXJcbiAgcHJvdmlkZUNoZWNrb3V0Li5jaGVja291dEJpZ1Byb21vXG5cbiAgbWFpbi4ucHJvdmlkZUNoZWNrb3V0OiBjYWxsc1xuICBtYWluLi4-Q2hlY2tvdXQiLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCIsInRoZW1lVmFyaWFibGVzIjp7ImJhY2tncm91bmQiOiJ3aGl0ZSIsInByaW1hcnlDb2xvciI6IiNFQ0VDRkYiLCJzZWNvbmRhcnlDb2xvciI6IiNmZmZmZGUiLCJ0ZXJ0aWFyeUNvbG9yIjoiaHNsKDgwLCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJwcmltYXJ5Qm9yZGVyQ29sb3IiOiJoc2woMjQwLCA2MCUsIDg2LjI3NDUwOTgwMzklKSIsInNlY29uZGFyeUJvcmRlckNvbG9yIjoiaHNsKDYwLCA2MCUsIDgzLjUyOTQxMTc2NDclKSIsInRlcnRpYXJ5Qm9yZGVyQ29sb3IiOiJoc2woODAsIDYwJSwgODYuMjc0NTA5ODAzOSUpIiwicHJpbWFyeVRleHRDb2xvciI6IiMxMzEzMDAiLCJzZWNvbmRhcnlUZXh0Q29sb3IiOiIjMDAwMDIxIiwidGVydGlhcnlUZXh0Q29sb3IiOiJyZ2IoOS41MDAwMDAwMDAxLCA5LjUwMDAwMDAwMDEsIDkuNTAwMDAwMDAwMSkiLCJsaW5lQ29sb3IiOiIjMzMzMzMzIiwidGV4dENvbG9yIjoiIzMzMyIsIm1haW5Ca2ciOiIjRUNFQ0ZGIiwic2Vjb25kQmtnIjoiI2ZmZmZkZSIsImJvcmRlcjEiOiIjOTM3MERCIiwiYm9yZGVyMiI6IiNhYWFhMzMiLCJhcnJvd2hlYWRDb2xvciI6IiMzMzMzMzMiLCJmb250RmFtaWx5IjoiXCJ0cmVidWNoZXQgbXNcIiwgdmVyZGFuYSwgYXJpYWwiLCJmb250U2l6ZSI6IjE2cHgiLCJsYWJlbEJhY2tncm91bmQiOiIjZThlOGU4Iiwibm9kZUJrZyI6IiNFQ0VDRkYiLCJub2RlQm9yZGVyIjoiIzkzNzBEQiIsImNsdXN0ZXJCa2ciOiIjZmZmZmRlIiwiY2x1c3RlckJvcmRlciI6IiNhYWFhMzMiLCJkZWZhdWx0TGlua0NvbG9yIjoiIzMzMzMzMyIsInRpdGxlQ29sb3IiOiIjMzMzIiwiZWRnZUxhYmVsQmFja2dyb3VuZCI6IiNlOGU4ZTgiLCJhY3RvckJvcmRlciI6ImhzbCgyNTkuNjI2MTY4MjI0MywgNTkuNzc2NTM2MzEyOCUsIDg3LjkwMTk2MDc4NDMlKSIsImFjdG9yQmtnIjoiI0VDRUNGRiIsImFjdG9yVGV4dENvbG9yIjoiYmxhY2siLCJhY3RvckxpbmVDb2xvciI6ImdyZXkiLCJzaWduYWxDb2xvciI6IiMzMzMiLCJzaWduYWxUZXh0Q29sb3IiOiIjMzMzIiwibGFiZWxCb3hCa2dDb2xvciI6IiNFQ0VDRkYiLCJsYWJlbEJveEJvcmRlckNvbG9yIjoiaHNsKDI1OS42MjYxNjgyMjQzLCA1OS43NzY1MzYzMTI4JSwgODcuOTAxOTYwNzg0MyUpIiwibGFiZWxUZXh0Q29sb3IiOiJibGFjayIsImxvb3BUZXh0Q29sb3IiOiJibGFjayIsIm5vdGVCb3JkZXJDb2xvciI6IiNhYWFhMzMiLCJub3RlQmtnQ29sb3IiOiIjZmZmNWFkIiwibm90ZVRleHRDb2xvciI6ImJsYWNrIiwiYWN0aXZhdGlvbkJvcmRlckNvbG9yIjoiIzY2NiIsImFjdGl2YXRpb25Ca2dDb2xvciI6IiNmNGY0ZjQiLCJzZXF1ZW5jZU51bWJlckNvbG9yIjoid2hpdGUiLCJzZWN0aW9uQmtnQ29sb3IiOiJyZ2JhKDEwMiwgMTAyLCAyNTUsIDAuNDkpIiwiYWx0U2VjdGlvbkJrZ0NvbG9yIjoid2hpdGUiLCJzZWN0aW9uQmtnQ29sb3IyIjoiI2ZmZjQwMCIsInRhc2tCb3JkZXJDb2xvciI6IiM1MzRmYmMiLCJ0YXNrQmtnQ29sb3IiOiIjOGE5MGRkIiwidGFza1RleHRMaWdodENvbG9yIjoid2hpdGUiLCJ0YXNrVGV4dENvbG9yIjoid2hpdGUiLCJ0YXNrVGV4dERhcmtDb2xvciI6ImJsYWNrIiwidGFza1RleHRPdXRzaWRlQ29sb3IiOiJibGFjayIsInRhc2tUZXh0Q2xpY2thYmxlQ29sb3IiOiIjMDAzMTYzIiwiYWN0aXZlVGFza0JvcmRlckNvbG9yIjoiIzUzNGZiYyIsImFjdGl2ZVRhc2tCa2dDb2xvciI6IiNiZmM3ZmYiLCJncmlkQ29sb3IiOiJsaWdodGdyZXkiLCJkb25lVGFza0JrZ0NvbG9yIjoibGlnaHRncmV5IiwiZG9uZVRhc2tCb3JkZXJDb2xvciI6ImdyZXkiLCJjcml0Qm9yZGVyQ29sb3IiOiIjZmY4ODg4IiwiY3JpdEJrZ0NvbG9yIjoicmVkIiwidG9kYXlMaW5lQ29sb3IiOiJyZWQiLCJsYWJlbENvbG9yIjoiYmxhY2siLCJlcnJvckJrZ0NvbG9yIjoiIzU1MjIyMiIsImVycm9yVGV4dENvbG9yIjoiIzU1MjIyMiIsImNsYXNzVGV4dCI6IiMxMzEzMDAiLCJmaWxsVHlwZTAiOiIjRUNFQ0ZGIiwiZmlsbFR5cGUxIjoiI2ZmZmZkZSIsImZpbGxUeXBlMiI6ImhzbCgzMDQsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlMyI6ImhzbCgxMjQsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSIsImZpbGxUeXBlNCI6ImhzbCgxNzYsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlNSI6ImhzbCgtNCwgMTAwJSwgOTMuNTI5NDExNzY0NyUpIiwiZmlsbFR5cGU2IjoiaHNsKDgsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlNyI6ImhzbCgxODgsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSJ9fSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiY2xhc3NEaWFncmFtXG4gIGNsYXNzIG1haW57XG4gICAgbWFpbigpXG4gIH1cblxuICBjbGFzcyBDaGVja291dCB7XG4gICAgPDxpbnRlcmZhY2U-PlxuICAgICtFeGVjdXRlKCpDYXJ0KSBlcnJvclxuICAgICtTZXROZXh0KENoZWNrb3V0KVxuICB9XG5cbiAgQ2hlY2tvdXQ8fC4uRnJhdWRDaGVjazogaW1wbGVtZW50c1xuICAlJUZyYXVkQ2hlY2suLj5DaGVja291dDogY2FsbHMgbmV4dFxuICBDaGVja291dDx8Li5QYXltZW50UHJlcGFyYXRpb246IGltcGxlbWVudHNcbiAgJSVQYXltZW50UHJlcGFyYXRpb24uLj5DaGVja291dDogY2FsbHMgbmV4dFxuICBDaGVja291dDx8Li5Qcm9kdWN0VmFsaWRhdGlvbjogaW1wbGVtZW50c1xuICAlJVByb2R1Y3RWYWxpZGF0aW9uLi4-Q2hlY2tvdXQ6IGNhbGxzIG5leHRcbiAgQ2hlY2tvdXQ8fC4uUHJvbW9WYWxpZGF0aW9uOiBpbXBsZW1lbnRzXG4gICUlUHJvbW9WYWxpZGF0aW9uLi4-Q2hlY2tvdXQ6IGNhbGxzIG5leHRcblxuICBjbGFzcyBjaGVja291dENvbnN1bWVye1xuICAgIDw8aW5zdGFuY2U-PlxuICB9XG4gIFBheW1lbnRQcmVwYXJhdGlvbiAtLW8gY2hlY2tvdXRDb25zdW1lcjogcGFydCBvZlxuICBQcm9kdWN0VmFsaWRhdGlvbiAtLW8gY2hlY2tvdXRDb25zdW1lcjogcGFydCBvZlxuXG4gICAgY2xhc3MgY2hlY2tvdXRCaWdQcm9tb3tcbiAgICA8PGluc3RhbmNlPj5cbiAgfVxuICBQYXltZW50UHJlcGFyYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcbiAgUHJvZHVjdFZhbGlkYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcbiAgRnJhdWRDaGVjayAtLW8gY2hlY2tvdXRCaWdQcm9tbzogcGFydCBvZlxuICBQcm9tb1ZhbGlkYXRpb24gLS1vIGNoZWNrb3V0QmlnUHJvbW86IHBhcnQgb2ZcblxuICBjbGFzcyBwcm92aWRlQ2hlY2tvdXR7XG4gICAgPDxmdW5jdGlvbj4-XG4gIH1cblxuICBwcm92aWRlQ2hlY2tvdXQuLmNoZWNrb3V0Q29uc3VtZXJcbiAgcHJvdmlkZUNoZWNrb3V0Li5jaGVja291dEJpZ1Byb21vXG5cbiAgbWFpbi4ucHJvdmlkZUNoZWNrb3V0OiBjYWxsc1xuICBtYWluLi4-Q2hlY2tvdXQiLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCIsInRoZW1lVmFyaWFibGVzIjp7ImJhY2tncm91bmQiOiJ3aGl0ZSIsInByaW1hcnlDb2xvciI6IiNFQ0VDRkYiLCJzZWNvbmRhcnlDb2xvciI6IiNmZmZmZGUiLCJ0ZXJ0aWFyeUNvbG9yIjoiaHNsKDgwLCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJwcmltYXJ5Qm9yZGVyQ29sb3IiOiJoc2woMjQwLCA2MCUsIDg2LjI3NDUwOTgwMzklKSIsInNlY29uZGFyeUJvcmRlckNvbG9yIjoiaHNsKDYwLCA2MCUsIDgzLjUyOTQxMTc2NDclKSIsInRlcnRpYXJ5Qm9yZGVyQ29sb3IiOiJoc2woODAsIDYwJSwgODYuMjc0NTA5ODAzOSUpIiwicHJpbWFyeVRleHRDb2xvciI6IiMxMzEzMDAiLCJzZWNvbmRhcnlUZXh0Q29sb3IiOiIjMDAwMDIxIiwidGVydGlhcnlUZXh0Q29sb3IiOiJyZ2IoOS41MDAwMDAwMDAxLCA5LjUwMDAwMDAwMDEsIDkuNTAwMDAwMDAwMSkiLCJsaW5lQ29sb3IiOiIjMzMzMzMzIiwidGV4dENvbG9yIjoiIzMzMyIsIm1haW5Ca2ciOiIjRUNFQ0ZGIiwic2Vjb25kQmtnIjoiI2ZmZmZkZSIsImJvcmRlcjEiOiIjOTM3MERCIiwiYm9yZGVyMiI6IiNhYWFhMzMiLCJhcnJvd2hlYWRDb2xvciI6IiMzMzMzMzMiLCJmb250RmFtaWx5IjoiXCJ0cmVidWNoZXQgbXNcIiwgdmVyZGFuYSwgYXJpYWwiLCJmb250U2l6ZSI6IjE2cHgiLCJsYWJlbEJhY2tncm91bmQiOiIjZThlOGU4Iiwibm9kZUJrZyI6IiNFQ0VDRkYiLCJub2RlQm9yZGVyIjoiIzkzNzBEQiIsImNsdXN0ZXJCa2ciOiIjZmZmZmRlIiwiY2x1c3RlckJvcmRlciI6IiNhYWFhMzMiLCJkZWZhdWx0TGlua0NvbG9yIjoiIzMzMzMzMyIsInRpdGxlQ29sb3IiOiIjMzMzIiwiZWRnZUxhYmVsQmFja2dyb3VuZCI6IiNlOGU4ZTgiLCJhY3RvckJvcmRlciI6ImhzbCgyNTkuNjI2MTY4MjI0MywgNTkuNzc2NTM2MzEyOCUsIDg3LjkwMTk2MDc4NDMlKSIsImFjdG9yQmtnIjoiI0VDRUNGRiIsImFjdG9yVGV4dENvbG9yIjoiYmxhY2siLCJhY3RvckxpbmVDb2xvciI6ImdyZXkiLCJzaWduYWxDb2xvciI6IiMzMzMiLCJzaWduYWxUZXh0Q29sb3IiOiIjMzMzIiwibGFiZWxCb3hCa2dDb2xvciI6IiNFQ0VDRkYiLCJsYWJlbEJveEJvcmRlckNvbG9yIjoiaHNsKDI1OS42MjYxNjgyMjQzLCA1OS43NzY1MzYzMTI4JSwgODcuOTAxOTYwNzg0MyUpIiwibGFiZWxUZXh0Q29sb3IiOiJibGFjayIsImxvb3BUZXh0Q29sb3IiOiJibGFjayIsIm5vdGVCb3JkZXJDb2xvciI6IiNhYWFhMzMiLCJub3RlQmtnQ29sb3IiOiIjZmZmNWFkIiwibm90ZVRleHRDb2xvciI6ImJsYWNrIiwiYWN0aXZhdGlvbkJvcmRlckNvbG9yIjoiIzY2NiIsImFjdGl2YXRpb25Ca2dDb2xvciI6IiNmNGY0ZjQiLCJzZXF1ZW5jZU51bWJlckNvbG9yIjoid2hpdGUiLCJzZWN0aW9uQmtnQ29sb3IiOiJyZ2JhKDEwMiwgMTAyLCAyNTUsIDAuNDkpIiwiYWx0U2VjdGlvbkJrZ0NvbG9yIjoid2hpdGUiLCJzZWN0aW9uQmtnQ29sb3IyIjoiI2ZmZjQwMCIsInRhc2tCb3JkZXJDb2xvciI6IiM1MzRmYmMiLCJ0YXNrQmtnQ29sb3IiOiIjOGE5MGRkIiwidGFza1RleHRMaWdodENvbG9yIjoid2hpdGUiLCJ0YXNrVGV4dENvbG9yIjoid2hpdGUiLCJ0YXNrVGV4dERhcmtDb2xvciI6ImJsYWNrIiwidGFza1RleHRPdXRzaWRlQ29sb3IiOiJibGFjayIsInRhc2tUZXh0Q2xpY2thYmxlQ29sb3IiOiIjMDAzMTYzIiwiYWN0aXZlVGFza0JvcmRlckNvbG9yIjoiIzUzNGZiYyIsImFjdGl2ZVRhc2tCa2dDb2xvciI6IiNiZmM3ZmYiLCJncmlkQ29sb3IiOiJsaWdodGdyZXkiLCJkb25lVGFza0JrZ0NvbG9yIjoibGlnaHRncmV5IiwiZG9uZVRhc2tCb3JkZXJDb2xvciI6ImdyZXkiLCJjcml0Qm9yZGVyQ29sb3IiOiIjZmY4ODg4IiwiY3JpdEJrZ0NvbG9yIjoicmVkIiwidG9kYXlMaW5lQ29sb3IiOiJyZWQiLCJsYWJlbENvbG9yIjoiYmxhY2siLCJlcnJvckJrZ0NvbG9yIjoiIzU1MjIyMiIsImVycm9yVGV4dENvbG9yIjoiIzU1MjIyMiIsImNsYXNzVGV4dCI6IiMxMzEzMDAiLCJmaWxsVHlwZTAiOiIjRUNFQ0ZGIiwiZmlsbFR5cGUxIjoiI2ZmZmZkZSIsImZpbGxUeXBlMiI6ImhzbCgzMDQsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlMyI6ImhzbCgxMjQsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSIsImZpbGxUeXBlNCI6ImhzbCgxNzYsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlNSI6ImhzbCgtNCwgMTAwJSwgOTMuNTI5NDExNzY0NyUpIiwiZmlsbFR5cGU2IjoiaHNsKDgsIDEwMCUsIDk2LjI3NDUwOTgwMzklKSIsImZpbGxUeXBlNyI6ImhzbCgxODgsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSJ9fSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)

# Example
## Problem
Let's say our ecommerce website already has a basic checkout process which consists of these:
1. **Product Validation**
2. **Payment-process Preparation**

Then our BD come to us wanting to expand the business saying "We're going to also serve for Distributors too."

Besides, we have a big promo coming up. We need to prepare for any potential frauds and do additional validation for the promo.

So we have 2 new upcoming use-cases on our checkout process. After further analysis and break-downs of the requirements,
we formulate the steps of Distributors and Big-Promo use-cases.

Distributors Checkout Process:
1. Minimum quantity validation
2. **KYC validation**
3. Shop eligibility validation
4. Location validation
5. **Product validation**
6. **Payment-process preparation**

Big-Promo Checkout Process:
1. Promo validation
2. Fraud validation
3. **KYC validation**
4. **Product validation**
5. **Payment-process validation**

As you can see on the texts in bold, the three checkout sequences have some similar processes (KYC, Product, Payment). We,
of course, would like to minimize DRY (Do Not Repeat Yourself) so a design-choice to create 3 separate class won't fit
since it'll have some repetitions. Sure, we can add some if-statements on our checkout for different process-selection
for each use cases, but that will make the codes very long and complicated to maintain as time goes. It will also have
a low level of cohesion.

We need to have a design where each checkout-step is modular and reusable for each checkout-process.

## Solution
Enter chain-of-responsibility pattern. Here, each step can be made as modular and reusable. In this design-pattern, each
step will be set up as its own module which has a next step reference. The intention is to have a module with its own
responsibility which can call another similar module with different responsibility. So then we can later combine those
modules into a chain of steps, a chain of responsibility. With that intention in mind, we can set up this contract

```go
type Checkout interface {
	Execute(cart *entity.Cart) error
	SetNext(co Checkout)
}
```
Contract's set, now we'll implement each of the step on its own:
1. [Fraud validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/fraudcheck.go)
2. [KYC validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/kycvalidation.go)
3. [Location validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/locationvalidation.go)
4. [Minimum quantity validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/minimumqtyvalidation.go)
5. [Payment-process preparation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/paymentpreparation.go)
6. [Product validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/productvalidation.go)
7. [Promo validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/promovalidation.go)
8. [Shop eligibility validation](https://github.com/gentra/golang-design-patterns/blob/master/chainofresponsibility/internal/usecase/checkout/shopeligibility.go)

With the codes done, now we just need to combine and construct a checkout-process that we want

```go
func Checkout(coType constant.CheckoutType) internal.Checkout {
	switch coType {
	case constant.CheckoutConsumer:
		pay := checkout.NewPaymentPreparation() // sequence 2

		product := checkout.NewProductValidation() // sequence 1
		product.SetNext(pay)

		return product
	case constant.CheckoutDistributor:
		pay := checkout.NewPaymentPreparation() // sequence 6

		product := checkout.NewProductValidation() // sequence 5
		product.SetNext(pay)

		location := checkout.NewLocationValidation() // sequence 4
		location.SetNext(product)

		shopeligible := checkout.NewShopEligibility() // sequence 3
		shopeligible.SetNext(location)

		kyc := checkout.NewKYCValidation() // sequence 2
		kyc.SetNext(shopeligible)

		minqty := checkout.NewMinimumQtyValidation() // sequence 1
		minqty.SetNext(kyc)

		return minqty
	case constant.CheckoutBigPromo:
		pay := checkout.NewPaymentPreparation() // sequence 5

		product := checkout.NewProductValidation() // sequence 4
		product.SetNext(pay)

		fraudcheck := checkout.NewFraudCheck() // sequence 3
		fraudcheck.SetNext(product)

		kyc := checkout.NewKYCValidation() // sequence 2
		kyc.SetNext(kyc)

		promocheck := checkout.NewPromoValidation() // sequence 1
		promocheck.SetNext(fraudcheck)
		return promocheck
	default:
		return nil
	}
}
```

Then execute the instance
```go
	// Let's just pretend we get this checkout-type from user-type, config, or somewhere else
	checkoutType := constant.CheckoutDistributor
	co := provide.Checkout(checkoutType)

	err := co.Execute(&entity.Cart{})
	if err != nil {
		log.Println(err)
	}
```

In real world scenario, we might just want to make each checkout use-cases as a singleton. However, you might get the
idea now, with each step modularized, we can make any combinations we want. Now in real world case it just needs to be
integrated to the front-end.