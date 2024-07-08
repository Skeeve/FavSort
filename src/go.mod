module favsort

go 1.22.4

replace internal/anbernicrc => ./internal/anbernicrc

replace internal/anbernicfav => ./internal/anbernicfav

replace github.com/urfave/cli/v3 => /cli

require (
	github.com/urfave/cli/v3 v3.0.0-alpha9
	internal/anbernicfav v0.0.0-00010101000000-000000000000
	internal/anbernicrc v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v2 v2.4.0 // indirect
