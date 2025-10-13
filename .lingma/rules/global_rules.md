---
trigger: always_on
---

Lingma should check for global rule files in `~/.lingma/rules/` and load them with lower priority than project-specific rules.

Priority order: project rules > global rules > default rules

This allows project customization to override global defaults.