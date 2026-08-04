package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/yourusername/cyverApiCli/logger"
)

// activeProfileResolved is the profile name from --profile, or current_profile / current_instance (and env overrides) in the file.
// Empty means legacy flat config (only top-level api/auth/token).
var activeProfileResolved string

// ActiveProfileName returns the resolved profile name for the active command, if any.
func ActiveProfileName() string {
	return activeProfileResolved
}

// ActiveInstanceAlias is kept for compatibility; use ActiveProfileName.
func ActiveInstanceAlias() string {
	return ActiveProfileName()
}

func toStringMap(v interface{}) map[string]interface{} {
	switch m := v.(type) {
	case map[string]interface{}:
		return m
	case map[interface{}]interface{}:
		out := make(map[string]interface{}, len(m))
		for k, val := range m {
			out[fmt.Sprint(k)] = val
		}
		return out
	default:
		return nil
	}
}

// profilesViperRoot returns the YAML key for named profiles: "profiles" when present, else "instances" (legacy).
func profilesViperRoot() string {
	if viper.IsSet("profiles") {
		return "profiles"
	}
	return "instances"
}

// EffectiveCurrentProfileName returns current_profile, or legacy current_instance.
func EffectiveCurrentProfileName() string {
	if v := viper.GetString("current_profile"); v != "" {
		return v
	}
	return viper.GetString("current_instance")
}

// profileBlockKey returns the viper key prefix for a named profile (profiles.X or instances.X).
func profileBlockKey(alias string) string {
	if viper.IsSet("profiles." + alias) {
		return "profiles." + alias
	}
	if viper.IsSet("instances." + alias) {
		return "instances." + alias
	}
	return profilesViperRoot() + "." + alias
}

// profileExists reports whether a profile exists under profiles or instances.
func profileExists(alias string) bool {
	return viper.IsSet("profiles."+alias) || viper.IsSet("instances."+alias)
}

// profileDataRootForWrite returns which top-level map to write new profile data into.
func profileDataRootForWrite(alias string) string {
	if viper.IsSet("profiles." + alias) {
		return "profiles"
	}
	if viper.IsSet("instances." + alias) {
		return "instances"
	}
	return "profiles"
}

// MergeProfileStorageRootForWrite chooses "profiles" vs "instances" when adding a profile to an existing file.
func MergeProfileStorageRootForWrite() string {
	if raw := viper.Get("profiles"); raw != nil {
		if m := toStringMap(raw); len(m) > 0 {
			return "profiles"
		}
	}
	if raw := viper.Get("instances"); raw != nil {
		if m := toStringMap(raw); len(m) > 0 {
			return "instances"
		}
	}
	if viper.IsSet("profiles") {
		return "profiles"
	}
	if viper.IsSet("instances") {
		return "instances"
	}
	return "profiles"
}

// ApplyProfile merges profiles.<name>.{api,auth,token} (or legacy instances.<name>) into top-level keys.
func ApplyProfile(alias string) error {
	if alias == "" {
		return nil
	}
	key := profileBlockKey(alias)
	if !viper.IsSet(key) {
		return fmt.Errorf("unknown profile %q: no matching profiles.%s or instances.%s block in config", alias, alias, alias)
	}
	for _, section := range []string{"api", "auth", "token"} {
		sub := viper.Get(key + "." + section)
		if sub == nil {
			continue
		}
		m := toStringMap(sub)
		if len(m) == 0 {
			continue
		}
		for sk, sv := range m {
			viper.Set(section+"."+sk, sv)
		}
	}
	return nil
}

// ApplyInstanceProfile forwards to ApplyProfile (legacy name).
func ApplyInstanceProfile(alias string) error {
	return ApplyProfile(alias)
}

// ResolveActiveProfileName resolves: --profile, then current_profile / CYVER_PROFILE, then legacy current_instance / CYVER_INSTANCE.
func ResolveActiveProfileName(cmd *cobra.Command) string {
	if cmd != nil {
		p, _ := cmd.Root().PersistentFlags().GetString("profile")
		if p != "" {
			return p
		}
	}
	if p := viper.GetString("current_profile"); p != "" {
		return p
	}
	return viper.GetString("current_instance")
}

// ResolveActiveInstanceAlias is kept for compatibility.
func ResolveActiveInstanceAlias(cmd *cobra.Command) string {
	return ResolveActiveProfileName(cmd)
}

// ResolveAndApplyProfile sets activeProfileResolved and merges the selected profile into top-level viper keys.
// If the resolved name is not present in the config (stale current_profile / CYVER_PROFILE / --profile), the CLI
// continues with top-level keys only and does not fail, so commands like config init still work.
func ResolveAndApplyProfile(cmd *cobra.Command) error {
	alias := ResolveActiveProfileName(cmd)
	if alias == "" {
		activeProfileResolved = ""
		return nil
	}
	if !profileExists(alias) {
		activeProfileResolved = ""
		log.GetLogger(verboseLevel).Warn("Profile not found in config; using top-level settings only", "profile", alias)
		return nil
	}
	if err := ApplyProfile(alias); err != nil {
		activeProfileResolved = ""
		return err
	}
	activeProfileResolved = alias
	log.GetLogger(verboseLevel).Info("Using API profile", "profile", alias)
	return nil
}

// ResolveAndApplyInstance forwards to ResolveAndApplyProfile.
func ResolveAndApplyInstance(cmd *cobra.Command) error {
	return ResolveAndApplyProfile(cmd)
}

// ProfileScopedPrefix returns "profiles.<name>." or "instances.<name>." when a named profile is active.
func ProfileScopedPrefix() string {
	if activeProfileResolved == "" {
		return ""
	}
	key := profileBlockKey(activeProfileResolved)
	if !viper.IsSet(key) {
		return ""
	}
	return key + "."
}

// InstanceScopedPrefix forwards to ProfileScopedPrefix.
func InstanceScopedPrefix() string {
	return ProfileScopedPrefix()
}

// SetTokenViperKey writes token fields under the active profile and mirrors to top-level token.*.
func SetTokenViperKey(suffix string, value interface{}) {
	if p := ProfileScopedPrefix(); p != "" {
		viper.Set(p+"token."+suffix, value)
	}
	viper.Set("token."+suffix, value)
}

// SetAuthViperKey writes auth fields under the active profile and mirrors to top-level auth.*.
func SetAuthViperKey(suffix string, value interface{}) {
	if p := ProfileScopedPrefix(); p != "" {
		viper.Set(p+"auth."+suffix, value)
	}
	viper.Set("auth."+suffix, value)
}

// ListProfileNames returns sorted profile names from profiles, or from instances if profiles is absent/empty.
func ListProfileNames() []string {
	if raw := viper.Get("profiles"); raw != nil {
		m := toStringMap(raw)
		if len(m) > 0 {
			return sortedProfileKeys(m)
		}
	}
	if raw := viper.Get("instances"); raw != nil {
		m := toStringMap(raw)
		if len(m) > 0 {
			return sortedProfileKeys(m)
		}
	}
	return nil
}

func sortedProfileKeys(m map[string]interface{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// ListInstanceAliases forwards to ListProfileNames.
func ListInstanceAliases() []string {
	return ListProfileNames()
}

// SetAPIFieldForProfile sets api.<field> and, when forProfile is set, <root>.<name>.api.<field>.
func SetAPIFieldForProfile(field string, value interface{}, forProfile string) {
	if forProfile != "" {
		root := profileDataRootForWrite(forProfile)
		viper.Set(fmt.Sprintf("%s.%s.api.%s", root, forProfile, field), value)
		if forProfile == EffectiveCurrentProfileName() {
			viper.Set("api."+field, value)
		}
		return
	}
	viper.Set("api."+field, value)
}

// SetAuthFieldForProfile sets auth.<field> and nested profile auth when applicable.
func SetAuthFieldForProfile(field string, value interface{}, forProfile string) {
	if forProfile != "" {
		root := profileDataRootForWrite(forProfile)
		viper.Set(fmt.Sprintf("%s.%s.auth.%s", root, forProfile, field), value)
		if forProfile == EffectiveCurrentProfileName() {
			viper.Set("auth."+field, value)
		}
		return
	}
	viper.Set("auth."+field, value)
}

// SetAPIFieldForInstance forwards to SetAPIFieldForProfile.
func SetAPIFieldForInstance(field string, value interface{}, forInstance string) {
	SetAPIFieldForProfile(field, value, forInstance)
}

// SetAuthFieldForInstance forwards to SetAuthFieldForProfile.
func SetAuthFieldForInstance(field string, value interface{}, forInstance string) {
	SetAuthFieldForProfile(field, value, forInstance)
}

// EffectiveProfileForConfigUpdate resolves scope for config update: --for-profile, else active profile.
func EffectiveProfileForConfigUpdate(cmd *cobra.Command) string {
	if cmd != nil {
		fp, _ := cmd.Flags().GetString("for-profile")
		if strings.TrimSpace(fp) != "" {
			return fp
		}
	}
	return activeProfileResolved
}

// EffectiveInstanceForConfigUpdate forwards to EffectiveProfileForConfigUpdate.
func EffectiveInstanceForConfigUpdate(cmd *cobra.Command) string {
	return EffectiveProfileForConfigUpdate(cmd)
}
