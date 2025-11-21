package config

// PruneConfig represents the configuration for pruning responses
type PruneConfig struct {
	// FuzzyKeys are prefixes for keys that should be removed
	FuzzyKeys []string `mapstructure:"fuzzy_keys"`

	// RemovePaths are exact paths or path suffixes that should be removed
	RemovePaths []string `mapstructure:"remove_paths"`
}

// DefaultPruneConfig returns the default prune configuration
func DefaultPruneConfig() PruneConfig {
	return PruneConfig{
		FuzzyKeys: []string{
			"customfield",
		},
		RemovePaths: []string{
			"emailAddress",
			"clone",
			"locked",
			"permittedOperations",
			"threadResolved",
			"avatarUrls",
			"timeZone",
			"thumbnail",
			"participants",
			"user.id",
			"user.links",
			//"user.name",
			"user.slug",
			"user.type",
			//"links",
			//"self",
			"scmId",
			"public",
			"author.id",
			//"author.name",
			"author.slug",
			"author.type",
			"author.key",
			"author.self",  //--
			"author.links", //--
			//"creator.name",
			"creator.key",
			"creator.self", //--
			//"reporter.name",
			"reporter.key",
			"reporter.self", //--
			"updateAuthor.name",
			"updateAuthor.key",
			"updateAuthor.self", //--
			"committer.id",
			//"committer.name",
			"committer.slug",
			"committer.type",
			"avatarId",
			"iconUrl",
			"statusCategory",
			"status.id",
			"status.self", //--
			"status.description",
			//"type",
			"fixVersions.id",
			"fixVersions.self",
			"issuetype.id",
			"issuetype.self",
			"priority.id",
			"priority.self",
			"lastViewed",
			"project.id",
			"projectCategory",
			"projectTypeKey",
			"resolution.id",
			"resolution.description",
			"resolution.self",
			"security",
			"versions.id",
			"versions.self", //--
			"votes",
			"watches",
			"displayId",
			"path.components", //?
			"path.extension",
			"path.name",
			"path.parent",
			"workratio", //?
			"type.id",
			"type.inward",
			"type.outward",
			"type.self", //
		},
	}
}
